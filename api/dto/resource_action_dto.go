package dto

import "github.com/google/uuid"

type LinkResourceActionRequestDTO struct {
	ResourceID uuid.UUID `json:"resourceId"`
	ActionID   uuid.UUID `json:"actionId"`
}

type ResourceActionDTO struct {
	ID         uuid.UUID `json:"id"`
	ResourceID uuid.UUID `json:"resourceId"`
	ActionID   uuid.UUID `json:"actionId"`
	Cursor     string    `json:"cursor,omitempty"`
}
