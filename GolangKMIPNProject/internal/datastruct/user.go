package datastruct

const UserTableName = "user"

type User struct {
	UserId   int64  `db:"UserID"`
	Fullname string `db:"FullName"`
	Password string `db:"Password"`
	Email    string `db:"Email"`
	Nik      string `db:"NIK"`
	Telenum  string `db:"TeleNum"`
	Kota     string `db:"KotaAsal"`
	Minat    string `db:"Minat"`
}
