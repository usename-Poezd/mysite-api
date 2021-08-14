package domain

import "time"

type User struct {
	Id			string		`json:"id"`
	Name 		string 		`json:"name"`
	Email		string		`json:"email"`
	Password	string		`json:"-"`
	UpdatedAt	time.Time	`json:"updatedAt" db:"updated_at"`
	CreatedAt	time.Time	`json:"createdAt" db:"created_at"`
}
