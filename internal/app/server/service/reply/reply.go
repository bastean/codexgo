package reply

type (
	Payload = map[string]any
)

type JSON struct {
	Success bool
	Message string
	Data    Payload
}
