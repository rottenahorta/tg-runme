package events

type Fetcher interface {
	Fetch(l int) ([]Event, error)
}

type Processor interface {
	Process(e Event) error
	SetWH(u string) error
	CheckWH(u string) ([]byte, error)
	ChangeHost(h string)
}

type Type int

const (
	Unknown Type = iota
	Message
)

type Event struct {
	Type Type
	Text string
	Meta interface{}
}
