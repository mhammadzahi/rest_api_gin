// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package database

import (
	"database/sql"
)

type User struct {
	ID                int32
	Name              sql.NullString
	PhoneNumber       sql.NullString
	Otp               sql.NullString
	OtpExpirationTime sql.NullTime
}
