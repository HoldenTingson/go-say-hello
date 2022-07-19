package datastruct

const EoTableName = "eo"

type Eo struct {
	EoId       int64  `db:"EOID"`
	EoName     string `db:"EOName"`
	EoPassword string `db:"EOPassword"`
	EoEmail    string `db:"EOEmail"`
}
