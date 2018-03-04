//https://github.com/variadico/scaneo
//https://github.com/Shelnutt2/db2struct
//$ db2struct --host localhost -d symfony -t user --package models --struct user -p monaFu69ma --user user_proyect
package models

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

// User estructura que representa la bd user
type User struct {
	ID                  int
	Username            string
	UsernameCanonical   string
	Email               string
	EmailCanonical      string
	Enabled             int
	Salt                sql.NullString
	Password            string
	LastLogin           mysql.NullTime
	ConfirmationToken   sql.NullString
	PasswordRequestedAt mysql.NullTime
	Roles               string
}
