package datastruct

const EventTableName = "event"

type Event struct {
	EventId         int64  `db:"EventID"`
	EventName       string `db:"EventName"`
	EventCode       string `db:"EventCode"`
	EventDate       string `db:"EventDate"`
	EventGenre      string `db:"EventGenre"`
	EventDesc       string `db:"EventDesc"`
	EventTier       string `db:"EventTier"`
	EventTierDesc   string `db:"EventTierDesc"`
	EventTierPrice  int64  `db:"EventTierPrice"`
	EventPicture    string `db:"EventPicture"`
	EventPlace      string `db:"EventPlace"`
	EventTierAmount int64  `db:"EventTierAmount"`
	EventTime       string `db:"EventTime"`
}
