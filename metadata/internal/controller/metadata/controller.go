package metadata

import (
	"context"
	"errors"
	"github.com/MSVelan/movieapp/metadata/internal/repository"
	"github.com/MSVelan/movieapp/metadata/pkg/model"
)

// ErrNotFound returned when requested record is not found
var ErrNotFound = errors.New("not found")

type metadataRepository interface {
	Get(ctx context.Context, id string) (*model.Metadata, error)
	Put(ctx context.Context, id string, metadata *model.Metadata) error
}

type Controller struct {
	repo metadataRepository
}

// creates new metadata service controller.
func New(repo metadataRepository) *Controller {
	return &Controller{repo}
}

func (c *Controller) Get(ctx context.Context, id string) (*model.Metadata, error) {
	res, err := c.repo.Get(ctx, id)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		return nil, ErrNotFound
	}
	return res, nil
}

func (c *Controller) Put(ctx context.Context, id string, metadata *model.Metadata) error {
	return c.repo.Put(ctx, id, metadata)
}
