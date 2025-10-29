package dto

import "github.com/google/uuid"

type LinkRoleResourceActionRequestDTO struct {
	RoleID           uuid.UUID `json:"roleId"`
	ResourceActionID uuid.UUID `json:"resourceActionId"`
}

type RoleResourceActionDTO struct {
	ID               uuid.UUID `json:"id"`
	RoleID           uuid.UUID `json:"roleId"`
	ResourceActionID uuid.UUID `json:"resourceActionId"`
	Cursor           string    `json:"cursor,omitempty"`
}

type GetRoleResourceActionsResponseDTO struct {
	RoleResourceActions []RoleResourceActionDTO `json:"roleResourceActions"`
}
