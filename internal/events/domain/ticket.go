package domain

type TicketType string

const (
	TicketTypeHalf TicketType = "half" //meia
	TicketTypeFull TicketType = "full" //inteira
)

type Ticket struct {
	ID         string
	EventID    string
	Spot       *Spot
	TicketType TicketType
	Price      float64
}
