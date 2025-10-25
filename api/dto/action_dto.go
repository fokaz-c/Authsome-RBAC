package dto

import "github.com/google/uuid"

type CreateActionRequestDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ActionDTO struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Cursor      string    `json:"cursor,omitempty"`
}
