package models

import (
	"time"

	"github.com/uptrace/bun"
)

type SafeUser struct {
	ID        int64     `bun:"id,pk,autoincrement" json:"id"`
	Username  string    `bun:"username,unique,notnull" json:"username"`
	Email     string    `bun:"email,unique,notnull" json:"email"`
	FirstName string    `bun:"first_name,notnull" json:"first_name"`
	LastName  string    `bun:"last_name,notnull" json:"last_name"`
	Orders    []*Order  `bun:"rel:has-many,join:id=user_id" json:"orders"`
	CreatedAt time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp" json:"updated_at"`
	Active    bool      `bun:"active,notnull" json:"active"`
}

type User struct {
	SafeUser
	bun.BaseModel `bun:"table:users"`

	Password string `bun:"password,notnull"`
	Salt     string `bun:"salt,notnull"`
}
