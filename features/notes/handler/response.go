package handler

import (
	"test_backend/features/notes"
	"time"
)

type NotesResponse struct {
	ID        string    `json:"_id,omitempty"`
	Title     string    `json:"title,omitempty"`
	Content   string    `json:"content,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func EntityToResponse(note notes.NotesEntity) NotesResponse {
	return NotesResponse{
		ID:        note.ID,
		Title:     note.Title,
		Content:   note.Content,
		CreatedAt: note.CreatedAt,
		UpdatedAt: note.UpdatedAt,
	}
}

func ListEntityToResponse(note []notes.NotesEntity) []NotesResponse {
	var listNotes []NotesResponse
	for _, value := range note {
		listNotes = append(listNotes, EntityToResponse(value))
	}
	return listNotes
}
