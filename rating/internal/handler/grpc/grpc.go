package grpc

import (
	"context"
	"errors"

	"github.com/MSVelan/movieapp/gen"
	"github.com/MSVelan/movieapp/rating/internal/controller/rating"
	"github.com/MSVelan/movieapp/rating/pkg/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Handler defines gRPC API handler.
type Handler struct {
	gen.UnimplementedRatingServiceServer
	ctrl *rating.Controller
}

// New creates a new rating gRPC handler.
func New(ctrl *rating.Controller) *Handler {
	return &Handler{ctrl: ctrl}
}

// GetAggregatedRating returns the aggregated rating for a record.
func (h *Handler) GetAggregatedRating(ctx context.Context, req *gen.GetAggregatedRatingRequest) (*gen.GetAggregatedRatingResponse, error) {
	if req == nil || req.RecordId == "" || req.RecordType == "" {
		return nil, status.Errorf(codes.InvalidArgument, "nil request or empty id or empty record type")
	}
	v, err := h.ctrl.GetAggregatedRating(ctx, model.RecordID(req.RecordId), model.RecordType(req.RecordType))
	if err != nil && errors.Is(err, rating.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &gen.GetAggregatedRatingResponse{
		RatingValue: v,
	}, nil
}

// PutRating writes a rating for a given record.
func (h *Handler) PutRating(ctx context.Context, req *gen.PutRatingRequest) (*gen.PutRatingResponse, error) {
	if req == nil || req.RecordId == "" || req.UserId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "nil request or empty record id or empty user id")
	}
	err := h.ctrl.Put(ctx, model.RecordID(req.RecordId), model.RecordType(req.RecordType), &model.Rating{
		UserID: model.UserID(req.UserId),
		Value:  model.RatingValue(req.RatingValue),
	})
	if err != nil {
		return nil, err
	}
	return &gen.PutRatingResponse{}, nil
}
