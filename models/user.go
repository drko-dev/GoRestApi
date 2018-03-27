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

// Usuario db
type Usuario struct {
	ID                       int
	Apellido                 sql.NullString
	Celular                  sql.NullString
	Comentario               sql.NullString
	CredencialesMp           sql.NullString
	Cuit                     sql.NullString
	DistanciaDeBusqueda      sql.NullInt64
	Domicilio                sql.NullString
	DominioID                int
	Email                    string
	EstadoID                 int
	FacebookID               sql.NullString
	FechaHabilitacion        mysql.NullTime
	FechaUltimoLogin         mysql.NullTime
	HabilitacionMercadoEnvio int
	HabilitacionMp           int
	Horario                  sql.NullString
	Iva                      sql.NullString
	LocalidadBusquedaID      sql.NullInt64
	LocalidadID              sql.NullInt64
	MediosDePago             sql.NullString
	MotivoEstadoID           sql.NullInt64
	Nombre                   sql.NullString
	NombreUsuario            sql.NullString
	Password                 string
	RazonSocial              sql.NullString
	Role                     sql.NullString
	Rubro                    sql.NullString
	Salt                     string
	Slug                     string
	Telefono                 sql.NullString
	Web                      sql.NullString
}

//https://github.com/variadico/scaneo
//https://github.com/Shelnutt2/db2struct
//$ db2struct --host localhost -d symfony -t user --package models --struct user -p monaFu69ma --user user_proyect
