package dto

import "github.com/google/uuid"

type CreateResourceRequestDTO struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

type ResourceDTO struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
	Cursor      string    `json:"cursor,omitempty"`
}

// This struct is now correct:
type GetResourcesResponseDTO struct {
	Resources []ResourceDTO `json:"resources"`
}
