package errors

type Bubbles struct {
	Internal     []*Internal
	Failure      []*Failure
	InvalidValue []*InvalidValue
	AlreadyExist []*AlreadyExist
	NotExist     []*NotExist
	Unknown      []error
	Amount       int
}

type (
	single = interface{ Unwrap() error }
	joined = interface{ Unwrap() []error }
)

func ExtractBubbles(err error) []error {
	var (
		errs    []error
		extract func(error)
	)

	extract = func(err error) {
		switch wrapped := err.(type) {
		case nil:
			return
		case single:
			extract(wrapped.Unwrap())
		case joined:
			for _, unwrapped := range wrapped.Unwrap() {
				extract(unwrapped)
			}
		default:
			errs = append(errs, err)
		}
	}

	extract(err)

	return errs
}

func FilterBubbles(errs []error, bubbles *Bubbles) {
	if bubbles == nil {
		Panic(Standard("Cannot filter if \"Bubbles\" are not defined"))
	}

	for _, err := range errs {
		switch bubble := err.(type) {
		case *Internal:
			bubbles.Internal = append(bubbles.Internal, bubble)
		case *Failure:
			bubbles.Failure = append(bubbles.Failure, bubble)
		case *InvalidValue:
			bubbles.InvalidValue = append(bubbles.InvalidValue, bubble)
		case *AlreadyExist:
			bubbles.AlreadyExist = append(bubbles.AlreadyExist, bubble)
		case *NotExist:
			bubbles.NotExist = append(bubbles.NotExist, bubble)
		default:
			bubbles.Unknown = append(bubbles.Unknown, bubble)
		}
	}

	bubbles.Amount = len(errs)
}

func Unwrap(err error, bubbles *Bubbles) {
	FilterBubbles(ExtractBubbles(err), bubbles)
}
