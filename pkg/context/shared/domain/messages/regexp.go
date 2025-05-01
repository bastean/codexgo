package messages

const (
	RExOrganization = `([a-z0-9]{1,30})`
	RExService      = `([a-z0-9]{1,30})`
	RExVersion      = `(\d+)`
	RExType         = `(event|command|query|response)`
	RExEntity       = `([a-z]{1,30})`
	RExTrigger      = `([a-z_]{1,30})`
	RExAction       = `([a-z]{1,30})`
	RExStatus       = `(queued|succeeded|failed|done)`
)

var Type = struct {
	Event, Command, Query, Response string
}{
	Event:    "event",
	Command:  "command",
	Query:    "query",
	Response: "response",
}

var Status = struct {
	Queued, Succeeded, Failed, Done string
}{
	Queued:    "queued",
	Succeeded: "succeeded",
	Failed:    "failed",
	Done:      "done",
}
