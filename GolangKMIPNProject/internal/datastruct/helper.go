package datastruct

const HelperTableName = "helper"

type Helper struct {
	HelperId  int64  `db:"HelperID"`
	UserId    int64  `db:"UserID"`
	EventId   int64  `db:"EventID"`
	EoId      int64  `db:"EOID"`
	EventCode string `db:"EventCode"`
}
