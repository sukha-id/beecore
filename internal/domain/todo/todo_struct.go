package domain

import "time"

type Todo struct {
	ID        int       `json:"ID,omitempty" db:"id"`
	Task      string    `json:"task,omitempty" db:"task"`
	CreatedAt string    `json:"createdAt,omitempty" db:"create_at"`
	UpdatedAt string    `json:"updatedAt,omitempty" db:"create_at"`
	IsDeleted time.Time `json:"-" db:"is_deleted"`
	DeletedAt time.Time `json:"-" db:"deleted_at"`
}
