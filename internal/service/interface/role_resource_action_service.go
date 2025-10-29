package service

import (
	"authsome-rbac/api/dto"
	"context"

	"github.com/google/uuid"
)

type IRoleResourceActionService interface {
	LinkRoleResourceAction(ctx context.Context, req dto.LinkRoleResourceActionRequestDTO) (*dto.RoleResourceActionDTO, error)
	GetRoleResourceActionByID(ctx context.Context, id uuid.UUID) (*dto.RoleResourceActionDTO, error)
	GetRoleResourceActions(ctx context.Context, limit int, cursor string) (*dto.GetRoleResourceActionsResponseDTO, error)
	UnlinkRoleResourceAction(ctx context.Context, id uuid.UUID) error
}
