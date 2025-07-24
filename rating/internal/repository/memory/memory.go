package memory

import (
	"context"

	"github.com/MSVelan/movieapp/rating/internal/repository"
	"github.com/MSVelan/movieapp/rating/pkg/model"
)

// Repository defines rating repository.
type Repository struct {
	data map[model.RecordType]map[model.RecordID][]model.Rating // data[rType][rID] = array of ratingStructs
}

// New creates a new repository.
func New() *Repository {
	return &Repository{
		map[model.RecordType]map[model.RecordID][]model.Rating{},
	}
}

// Get retrieves all ratings for a given record.
func (r *Repository) Get(ctx context.Context, recordID model.RecordID, recordType model.RecordType) ([]model.Rating, error) {
	if _, ok := r.data[recordType]; !ok {
		return nil, repository.ErrNotFound
	}
	if ratings, ok := r.data[recordType][recordID]; !ok || len(ratings) == 0 {
		return nil, repository.ErrNotFound
	}
	return r.data[recordType][recordID], nil
}

// Put adds a rating for given record.
func (r *Repository) Put(ctx context.Context, recordID model.RecordID, recordType model.RecordType, rating *model.Rating) error {
	if _, ok := r.data[recordType]; !ok {
		r.data[recordType] = map[model.RecordID][]model.Rating{}
	}
	r.data[recordType][recordID] = append(r.data[recordType][recordID], *rating)
	return nil
}
