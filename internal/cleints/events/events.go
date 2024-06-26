package events

type Fetcher interface {
	Fetch(limit int) ([]Event, error)
}

type Processor interface {
	Process(e Event) error
}

type Type int

const (
	Unknow Type = iota
	message
)

type Event struct {
	Type Type
	Text string
}
