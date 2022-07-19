package datastruct

const TicketTableName = "ticket"

type Ticket struct {
	TicketId  int64 `db:"TicketID"`
	UserId    int64 `db:"UserID"`
	EventId   int64 `db:"EventID"`
	OrderDate int64 `db:"OrderDate"`
	EventTier int64 `db:"EventTier"`
	EoId      int64 `db:"EOID"`
	Barcode   int64 `db:"Barcode"`
}
