package table

import (
	"gorm.io/gorm"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/persistences/sqlite"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/role"
)

type User struct {
	*gorm.Model
	Created, Updated string
	Verify, Reset    string
	ID               string `gorm:"uniqueIndex"`
	Email            string `gorm:"uniqueIndex"`
	Username         string `gorm:"uniqueIndex"`
	Password         string
	Verified         bool
}

type Table struct {
	*gorm.DB
}

func (t *Table) Create(user *user.User) error {
	err := user.CreationStamp()

	if err != nil {
		return errors.BubbleUp(err, "Create")
	}

	aggregate := user.ToPrimitive()

	err = t.DB.Create(&User{
		Created:  aggregate.Created,
		Updated:  aggregate.Updated,
		Verify:   aggregate.Verify,
		Reset:    aggregate.Reset,
		ID:       aggregate.ID,
		Email:    aggregate.Email,
		Username: aggregate.Username,
		Password: aggregate.Password,
		Verified: aggregate.Verified,
	}).Error

	switch {
	case sqlite.IsErrDuplicateValue(err):
		return errors.BubbleUp(sqlite.HandleErrDuplicateValue(err), "Create")
	case err != nil:
		return errors.New[errors.Internal](&errors.Bubble{
			Where: "Create",
			What:  "Failure to create a User",
			Why: errors.Meta{
				"ID": user.ID.Value,
			},
			Who: err,
		})
	}

	return nil
}

func (t *Table) Update(user *user.User) error {
	err := user.UpdatedStamp()

	if err != nil {
		return errors.BubbleUp(err, "Update")
	}

	err = t.DB.Where(&User{ID: user.ID.Value}).Updates(user.ToPrimitive()).Error

	switch {
	case sqlite.IsErrDuplicateValue(err):
		return errors.BubbleUp(sqlite.HandleErrDuplicateValue(err), "Update")
	case err != nil:
		return errors.New[errors.Internal](&errors.Bubble{
			Where: "Update",
			What:  "Failure to update a User",
			Why: errors.Meta{
				"ID": user.ID.Value,
			},
			Who: err,
		})
	}

	return nil
}

func (t *Table) Delete(id *user.ID) error {
	err := t.DB.Where(&User{ID: id.Value}).Unscoped().Delete(new(User)).Error

	if err != nil {
		return errors.New[errors.Internal](&errors.Bubble{
			Where: "Delete",
			What:  "Failure to delete a User",
			Why: errors.Meta{
				"ID": id.Value,
			},
			Who: err,
		})
	}

	return nil
}

func (t *Table) Search(criteria *user.Criteria) (*user.User, error) {
	var (
		filter *User
		index  string
	)

	switch {
	case criteria.ID != nil:
		filter = &User{ID: criteria.ID.Value}
		index = criteria.ID.Value
	case criteria.Email != nil:
		filter = &User{Email: criteria.Email.Value}
		index = criteria.Email.Value
	case criteria.Username != nil:
		filter = &User{Username: criteria.Username.Value}
		index = criteria.Username.Value
	default:
		return nil, errors.New[errors.Internal](&errors.Bubble{
			Where: "Search",
			What:  "Criteria not defined",
		})
	}

	primitive := new(user.Primitive)

	err := t.DB.Where(filter).First(&primitive).Error

	switch {
	case sqlite.IsErrNotFound(err):
		return nil, errors.BubbleUp(sqlite.HandleErrNotFound(err, index), "Search")
	case err != nil:
		return nil, errors.New[errors.Internal](&errors.Bubble{
			Where: "Search",
			What:  "Failure to search a User",
			Why: errors.Meta{
				"Index": index,
			},
			Who: err,
		})
	}

	user, err := user.FromPrimitive(primitive)

	if err != nil {
		return nil, errors.New[errors.Internal](&errors.Bubble{
			Where: "Search",
			What:  "Failure to create a User from a Primitive",
			Why: errors.Meta{
				"Index":     index,
				"Primitive": primitive,
			},
			Who: err,
		})
	}

	return user, nil
}

func Open(database *sqlite.Database) (role.Repository, error) {
	if err := database.Session.AutoMigrate(new(User)); err != nil {
		return nil, errors.New[errors.Internal](&errors.Bubble{
			Where: "Open",
			What:  "Failure to run auto migration for User model",
			Who:   err,
		})
	}

	return &Table{
		DB: database.Session.Model(new(User)).Session(new(gorm.Session)),
	}, nil
}
