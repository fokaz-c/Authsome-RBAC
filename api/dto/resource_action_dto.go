package dto

import "github.com/google/uuid"

type LinkResourceActionRequestDTO struct {
	ResourceID uuid.UUID `json:"resourceId"`
	ActionID   uuid.UUID `json:"actionId"`
}

type ResourceActionDTO struct {
	ID         uuid.UUID `json:"Id"`
	ResourceID uuid.UUID `json:"resourceId"`
	ActionID   uuid.UUID `json:"actionId"`
}
