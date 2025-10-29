package service

import (
	"authsome-rbac/api/dto"
	"context"

	"github.com/google/uuid"
)

type IResourceActionService interface {
	LinkResourceAction(ctx context.Context, req dto.LinkResourceActionRequestDTO) (*dto.ResourceActionDTO, error)
	GetResourceActionByID(ctx context.Context, id uuid.UUID) (*dto.ResourceActionDTO, error)
	GetResourceActions(ctx context.Context, limit int, cursor string) (*dto.GetResourceActionsResponseDTO, error)
	UnlinkResourceAction(ctx context.Context, id uuid.UUID) error
}
