// in-memory implementation of the data model.
package memory

import (
	"context"
	"github.com/MSVelan/movieapp/metadata/internal/repository"
	"github.com/MSVelan/movieapp/metadata/pkg/model"
	"sync"
)

type Repository struct {
	sync.RWMutex
	data map[string]*model.Metadata
}

// Create new repository.
func New() *Repository {
	return &Repository{
		data: map[string]*model.Metadata{},
	}
}

// Retrieve movie metadata by movie id.
func (r *Repository) Get(_ context.Context, id string) (*model.Metadata, error) {
	r.RLock()
	defer r.RUnlock()
	m, ok := r.data[id]
	if !ok {
		return nil, repository.ErrNotFound
	}
	return m, nil
}

// Add movie metadata for given movie id.
func (r *Repository) Put(_ context.Context, id string, metadata *model.Metadata) error {
	r.Lock()
	defer r.Unlock()
	r.data[id] = metadata
	return nil
}
