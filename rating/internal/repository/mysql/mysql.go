package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/MSVelan/movieapp/rating/internal/repository"
	"github.com/MSVelan/movieapp/rating/pkg/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// Repository defines MySQL based rating repository.
type Repository struct {
	db *sql.DB
}

// New creates a MySQL based rating repository.
func New() (*Repository, error) {
	err := godotenv.Load("pkg/mysql/.env")
	if err != nil {
		log.Fatalf("error loading .env file: %v", err)
		return nil, err
	}

	password := os.Getenv("PASSWORD")
	db, err := sql.Open("mysql", fmt.Sprintf("root:%s@/movieapp", password))
	if err != nil {
		log.Fatalf("mysql error: %v", err)
		return nil, err
	}

	return &Repository{db}, nil
}

// Get retrieves all ratings for a given record(record_id, record_type).
func (r *Repository) Get(ctx context.Context, recordID model.RecordID, recordType model.RecordType) ([]model.Rating, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT user_id, value FROM ratings WHERE record_id = ? AND record_type = ?", recordID, recordType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []model.Rating
	for rows.Next() {
		var userID string
		var value int32
		if err := rows.Scan(&userID, &value); err != nil {
			return nil, err
		}

		res = append(res, model.Rating{
			UserID: model.UserID(userID),
			Value:  model.RatingValue(value),
		})
	}

	if len(res) == 0 {
		return nil, repository.ErrNotFound
	}
	return res, nil
}

// Put adds a rating for a given record(record_id, record_type).
func (r *Repository) Put(ctx context.Context, recordID model.RecordID, recordType model.RecordType, rating *model.Rating) error {
	_, err := r.db.ExecContext(ctx,
		"INSERT into ratings (record_id, record_type, user_id, value) VALUES (?, ?, ?, ?)",
		recordID, recordType, rating.UserID, rating.Value)

	return err
}
