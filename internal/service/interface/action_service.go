package service

import (
	"authsome-rbac/api/dto"
	"context"

	"github.com/google/uuid"
)

type IActionService interface {
	CreateAction(ctx context.Context, req dto.CreateActionRequestDTO) (*dto.ActionDTO, error)
	GetActionByID(ctx context.Context, id uuid.UUID) (*dto.ActionDTO, error)
	GetActions(ctx context.Context, limit int, cursor string) (*dto.GetActionsResponseDTO, error)
	UpdateAction(ctx context.Context, id uuid.UUID, req dto.CreateActionRequestDTO) (*dto.ActionDTO, error)
	DeleteAction(ctx context.Context, id uuid.UUID) error
}
