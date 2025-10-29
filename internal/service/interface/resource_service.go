package service

import (
	"authsome-rbac/api/dto"
	"context"

	"github.com/google/uuid"
)

type IResourceService interface {
	CreateResource(ctx context.Context, req dto.CreateResourceRequestDTO) (*dto.ResourceDTO, error)
	GetResourceByID(ctx context.Context, id uuid.UUID) (*dto.ResourceDTO, error)
	GetResources(ctx context.Context, limit int, cursor string) (*dto.GetResourcesResponseDTO, error)
	UpdateResource(ctx context.Context, id uuid.UUID, req dto.CreateResourceRequestDTO) (*dto.ResourceDTO, error)
	DeleteResource(ctx context.Context, id uuid.UUID) error
}
