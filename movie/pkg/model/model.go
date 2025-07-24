package model

import (
	"github.com/MSVelan/movieapp/metadata/pkg/model"
)

// MovieDetails include movie metadata and its aggregated rating.
type MovieDetails struct {
	Rating   *float64       `json:"rating,omitempty"`
	Metadata model.Metadata `json:"metadata"`
}
