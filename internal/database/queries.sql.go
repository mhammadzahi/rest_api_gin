// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: queries.sql

package database

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (name, phone_number) VALUES ($1, $2) RETURNING id, name, phone_number, otp, otp_expiration_time
`

type CreateUserParams struct {
	Name        sql.NullString
	PhoneNumber sql.NullString
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Name, arg.PhoneNumber)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.PhoneNumber,
		&i.Otp,
		&i.OtpExpirationTime,
	)
	return i, err
}

const generateOTP = `-- name: GenerateOTP :one
UPDATE users SET otp = $1, otp_expiration_time = $2 WHERE phone_number = $3 RETURNING id, name, phone_number, otp, otp_expiration_time
`

type GenerateOTPParams struct {
	Otp               sql.NullString
	OtpExpirationTime sql.NullTime
	PhoneNumber       sql.NullString
}

func (q *Queries) GenerateOTP(ctx context.Context, arg GenerateOTPParams) (User, error) {
	row := q.db.QueryRowContext(ctx, generateOTP, arg.Otp, arg.OtpExpirationTime, arg.PhoneNumber)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.PhoneNumber,
		&i.Otp,
		&i.OtpExpirationTime,
	)
	return i, err
}

const verifyOTP = `-- name: VerifyOTP :one
SELECT id, name, phone_number, otp, otp_expiration_time FROM users WHERE phone_number = $1 AND otp = $2 AND otp_expiration_time > NOW()
`

type VerifyOTPParams struct {
	PhoneNumber sql.NullString
	Otp         sql.NullString
}

func (q *Queries) VerifyOTP(ctx context.Context, arg VerifyOTPParams) (User, error) {
	row := q.db.QueryRowContext(ctx, verifyOTP, arg.PhoneNumber, arg.Otp)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.PhoneNumber,
		&i.Otp,
		&i.OtpExpirationTime,
	)
	return i, err
}
