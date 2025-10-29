package service

import (
	"authsome-rbac/api/dto"
	"context"

	"github.com/google/uuid"
)

type IRoleService interface {
	CreateRole(ctx context.Context, req dto.CreateRoleRequestDTO) (*dto.RoleDTO, error)
	GetRoleByID(ctx context.Context, id uuid.UUID) (*dto.RoleDTO, error)
	GetRoles(ctx context.Context, limit int, cursor string) (*dto.GetRolesResponseDTO, error)
	UpdateRole(ctx context.Context, id uuid.UUID, req dto.CreateRoleRequestDTO) (*dto.RoleDTO, error)
	DeleteRole(ctx context.Context, id uuid.UUID) error
}
