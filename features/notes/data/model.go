package data

import (
	"test_backend/features/notes"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Notes struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Title     string             `bson:"title"`
	Content   string             `bson:"content"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	Deleted   bool               `bson:"deleted"`
	DeletedAt time.Time
}

func ModelToEntity(note Notes) notes.NotesEntity {
	return notes.NotesEntity{
		ID:        note.ID.Hex(),
		Title:     note.Title,
		Content:   note.Content,
		CreatedAt: note.CreatedAt,
		UpdatedAt: note.UpdatedAt,
	}
}

func ListModelToEntity(note []Notes) []notes.NotesEntity {
	var listNotes []notes.NotesEntity
	for _, value := range note {
		listNotes = append(listNotes, ModelToEntity(value))
	}
	return listNotes
}

func EntityToModel(note notes.NotesEntity) Notes {
	return Notes{
		Title:     note.Title,
		Content:   note.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
