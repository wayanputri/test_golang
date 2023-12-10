package notes

import (
	"context"
	"time"
)

type NotesEntity struct {
	ID        string
	Title     string `validate:"required"`
	Content   string `validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type NotesDataInterface interface {
	Insert(notes NotesEntity, ctx context.Context) (string, error)
	SelectAll(ctx context.Context) ([]NotesEntity, error)
	SelectById(ctx context.Context, id string) (NotesEntity, error)
	UpdateById(ctx context.Context, id string, note NotesEntity) error
	DeleteById(ctx context.Context, id string) error
}
type NotesServiceInterface interface {
	Add(notes NotesEntity, ctx context.Context) (string, error)
	GetAll(ctx context.Context) ([]NotesEntity, error)
	GetById(ctx context.Context, id string) (NotesEntity, error)
	EditById(ctx context.Context, id string, note NotesEntity) error
	DeleteById(ctx context.Context, id string) error
}
