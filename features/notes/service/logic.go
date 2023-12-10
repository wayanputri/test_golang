package service

import (
	"context"
	"errors"
	"test_backend/features/notes"

	"github.com/go-playground/validator/v10"
)

type NotesService struct {
	notesService notes.NotesDataInterface
	validate     *validator.Validate
}

// DeleteById implements notes.NotesServiceInterface.
func (service *NotesService) DeleteById(ctx context.Context, id string) error {
	err := service.notesService.DeleteById(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

// EditById implements notes.NotesServiceInterface.
func (service *NotesService) EditById(ctx context.Context, id string, note notes.NotesEntity) error {
	data, errGet := service.notesService.SelectById(ctx, id)
	if errGet != nil {
		return errGet
	}
	if note.Content == "" {
		note.Content = data.Content
	}
	if note.Title == "" {
		note.Title = data.Title
	}
	err := service.notesService.UpdateById(ctx, id, note)
	if err != nil {
		return err
	}
	return nil
}

// GetById implements notes.NotesServiceInterface.
func (service *NotesService) GetById(ctx context.Context, id string) (notes.NotesEntity, error) {
	data, err := service.notesService.SelectById(ctx, id)
	if err != nil {
		return notes.NotesEntity{}, err
	}

	return data, nil
}

// GetAll implements notes.NotesServiceInterface.
func (service *NotesService) GetAll(ctx context.Context) ([]notes.NotesEntity, error) {
	data, err := service.notesService.SelectAll(ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Add implements notes.NotesServiceInterface.
func (service *NotesService) Add(notes notes.NotesEntity, ctx context.Context) (string, error) {
	errValidate := service.validate.Struct(notes)
	if errValidate != nil {
		return "", errors.New("error validate, title/content required")
	}
	id, err := service.notesService.Insert(notes, ctx)
	if err != nil {
		return "", err
	}
	return id, nil
}

func New(service notes.NotesDataInterface) notes.NotesServiceInterface {
	return &NotesService{
		notesService: service,
		validate:     validator.New(),
	}
}
