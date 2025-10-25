package dto

import "github.com/google/uuid"

type LinkRoleResourceActionRequestDTO struct {
	RoleID           uuid.UUID `json:"roleId"`
	ResourceActionID uuid.UUID `json:"resourceActionId"`
}

type RoleResourceActionDTO struct {
	ID uuid.UUID `json:"id"`
}
