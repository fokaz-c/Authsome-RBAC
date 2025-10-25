package dto

import "github.com/google/uuid"

type CreateRoleRequestDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type RoleDTO struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Cursor      string    `json:"cursor,omitempty"`
}
