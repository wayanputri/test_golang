package handler

import "test_backend/features/notes"

type NotesRequest struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
}

func RequestToEntity(note NotesRequest) notes.NotesEntity {
	return notes.NotesEntity{
		Title:   note.Title,
		Content: note.Content,
	}
}
