package dto

import "github.com/google/uuid"

type UserFromToken struct {
	ID   uuid.UUID
	Role string
}
