package kv

import (
	"bytes"
	"encoding/json"
	"slices"

	"github.com/dgraph-io/badger/v4"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/persistences/badgerdb"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/role"
)

const (
	TotalCriteriaFields = 3
	IxID                = 0
	IxEmail             = 1
	IxUsername          = 2
)

type KV struct {
	*badger.DB
}

func (kv *KV) keyIterator(parser func(keyValues []string, item *badger.Item) (bool, error)) error {
	return kv.DB.View(func(txn *badger.Txn) error {
		options := badger.DefaultIteratorOptions

		options.PrefetchValues = false

		iterator := txn.NewIterator(options)

		defer iterator.Close()

		for iterator.Rewind(); iterator.Valid(); iterator.Next() {
			item := iterator.Item()

			values := badgerdb.ParseKey(item.Key())

			if len(values) != TotalCriteriaFields {
				return errors.New[errors.Internal](&errors.Bubble{
					What: "Invalid Key length",
				})
			}

			shouldBreak, err := parser(values, item)

			if shouldBreak {
				if err != nil {
					return errors.BubbleUp(err)
				}

				return nil
			}
		}

		return nil
	})
}

func (kv *KV) isDuplicate(criteria *user.Criteria, isItselfExcluded bool) error {
	var indexes []string

	if criteria.ID == nil {
		return errors.New[errors.Internal](&errors.Bubble{
			What: "ID criteria not defined",
		})
	}

	err := kv.keyIterator(func(keyValues []string, item *badger.Item) (bool, error) {
		switch {
		case isItselfExcluded && criteria.ID.Value() == keyValues[IxID]:
			return false, nil
		case !isItselfExcluded && criteria.ID.Value() == keyValues[IxID]:
			indexes = append(indexes, "ID")
		}

		if criteria.Email != nil && criteria.Email.Value() == keyValues[IxEmail] {
			indexes = append(indexes, "Email")
		}

		if criteria.Username != nil && criteria.Username.Value() == keyValues[IxUsername] {
			indexes = append(indexes, "Username")
		}

		if len(indexes) > 0 {
			return true, errors.BubbleUp(badgerdb.HandleErrDuplicateValue(indexes[0]))
		}

		return false, nil
	})

	if err != nil {
		return errors.BubbleUp(err)
	}

	return nil
}

func (kv *KV) actualKey(id *user.ID) ([]byte, error) {
	if id == nil {
		return nil, errors.New[errors.Internal](&errors.Bubble{
			What: "ID not defined",
		})
	}

	var key []byte

	err := kv.keyIterator(func(keyValues []string, item *badger.Item) (bool, error) {
		if slices.Contains(keyValues, id.Value()) {
			key = append(key, item.Key()...)
			return true, nil
		}

		return false, nil
	})

	if err != nil {
		return nil, errors.BubbleUp(err)
	}

	return key, nil
}

func (kv *KV) Create(aggregate *user.User) error {
	err := kv.isDuplicate(&user.Criteria{
		ID:       aggregate.ID,
		Email:    aggregate.Email,
		Username: aggregate.Username,
	}, false)

	if err != nil {
		return errors.BubbleUp(err)
	}

	err = aggregate.CreationStamp()

	if err != nil {
		return errors.BubbleUp(err)
	}

	value, err := json.Marshal(aggregate.ToPrimitive())

	if err != nil {
		return errors.New[errors.Internal](&errors.Bubble{
			What: "Failure to encode a User",
			Why: errors.Meta{
				"ID": aggregate.ID.Value(),
			},
			Who: err,
		})
	}

	key, err := badgerdb.NewKey(
		aggregate.ID.Value(),
		aggregate.Email.Value(),
		aggregate.Username.Value(),
	)

	if err != nil {
		return errors.BubbleUp(err)
	}

	err = kv.DB.Update(func(txn *badger.Txn) error {
		return txn.Set(key, value)
	})

	if err != nil {
		return errors.New[errors.Internal](&errors.Bubble{
			What: "Failure to create a User",
			Why: errors.Meta{
				"ID": aggregate.ID.Value(),
			},
			Who: err,
		})
	}

	return nil
}

func (kv *KV) Update(aggregate *user.User) error {
	err := kv.isDuplicate(&user.Criteria{
		ID:       aggregate.ID,
		Email:    aggregate.Email,
		Username: aggregate.Username,
	}, true)

	if err != nil {
		return errors.BubbleUp(err)
	}

	err = aggregate.UpdatedStamp()

	if err != nil {
		return errors.BubbleUp(err)
	}

	value, err := json.Marshal(aggregate.ToPrimitive())

	if err != nil {
		return errors.New[errors.Internal](&errors.Bubble{
			What: "Failure to encode a User",
			Why: errors.Meta{
				"ID": aggregate.ID.Value(),
			},
			Who: err,
		})
	}

	actualKey, err := kv.actualKey(aggregate.ID)

	if err != nil {
		return errors.BubbleUp(err)
	}

	newKey, err := badgerdb.NewKey(
		aggregate.ID.Value(),
		aggregate.Email.Value(),
		aggregate.Username.Value(),
	)

	if err != nil {
		return errors.BubbleUp(err)
	}

	err = kv.DB.Update(func(txn *badger.Txn) error {
		if bytes.Equal(actualKey, newKey) {
			return txn.Set(actualKey, value)
		}

		errTxn := txn.Delete(actualKey)

		if errTxn != nil {
			return errTxn
		}

		return txn.Set(newKey, value)
	})

	if err != nil {
		return errors.New[errors.Internal](&errors.Bubble{
			What: "Failure to update a User",
			Why: errors.Meta{
				"ID": aggregate.ID.Value(),
			},
			Who: err,
		})
	}

	return nil
}

func (kv *KV) Delete(id *user.ID) error {
	key, err := kv.actualKey(id)

	if err != nil {
		return errors.BubbleUp(err)
	}

	err = kv.DB.Update(func(txn *badger.Txn) error {
		return txn.Delete(key)
	})

	if err != nil {
		return errors.New[errors.Internal](&errors.Bubble{
			What: "Failure to delete a User",
			Why: errors.Meta{
				"ID": id.Value(),
			},
			Who: err,
		})
	}

	return nil
}

func (kv *KV) Search(criteria *user.Criteria) (*user.User, error) {
	var index string

	switch {
	case criteria.ID != nil:
		index = criteria.ID.Value()
	case criteria.Email != nil:
		index = criteria.Email.Value()
	case criteria.Username != nil:
		index = criteria.Username.Value()
	default:
		return nil, errors.New[errors.Internal](&errors.Bubble{
			What: "Criteria not defined",
		})
	}

	var rawPrimitive []byte

	err := kv.keyIterator(func(keyValues []string, item *badger.Item) (bool, error) {
		if slices.Contains(keyValues, index) {
			return true, item.Value(func(value []byte) error {
				rawPrimitive = append(rawPrimitive, value...)
				return nil
			})
		}

		return false, nil
	})

	switch {
	case err != nil:
		return nil, errors.BubbleUp(err)
	case rawPrimitive == nil:
		return nil, errors.BubbleUp(badgerdb.HandleErrNotFound(index))
	}

	primitive := new(user.Primitive)

	err = json.Unmarshal(rawPrimitive, primitive)

	if err != nil {
		return nil, errors.New[errors.Internal](&errors.Bubble{
			What: "Failure to decode a result",
			Why: errors.Meta{
				"Index": index,
			},
			Who: err,
		})
	}

	aggregate, err := user.FromPrimitive(primitive)

	if err != nil {
		return nil, errors.New[errors.Internal](&errors.Bubble{
			What: "Failure to create a User from a Primitive",
			Why: errors.Meta{
				"Index":     index,
				"Primitive": primitive,
			},
			Who: err,
		})
	}

	return aggregate, nil
}

func Open(database *badgerdb.Database) (role.Repository, error) {
	return &KV{
		DB: database.Session,
	}, nil
}
