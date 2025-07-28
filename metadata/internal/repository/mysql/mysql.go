package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/MSVelan/movieapp/metadata/internal/repository"
	"github.com/MSVelan/movieapp/metadata/pkg/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// Repository defines a MySQL based movie metadata repository
type Repository struct {
	db *sql.DB
}

// New creates a new MySQL based repository
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

// Get retrieves movie metadata by movie id.
func (r *Repository) Get(ctx context.Context, id string) (*model.Metadata, error) {
	var title, description, director string
	row := r.db.QueryRowContext(ctx, "SELECT title, description, director FROM movies WHERE id = ?", id)
	if err := row.Scan(&title, &description, &director); err != nil {
		if err == sql.ErrNoRows {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}

	return &model.Metadata{
		ID:          id,
		Title:       title,
		Description: description,
		Director:    director,
	}, nil
}

// Put adds movie metadata for a given movie id.
func (r *Repository) Put(ctx context.Context, id string, metadata *model.Metadata) error {
	_, err := r.db.ExecContext(ctx,
		"INSERT into movies (id, title, description, director) VALUES (?, ?, ?, ?)",
		id, metadata.Title, metadata.Description, metadata.Director)

	return err
}
