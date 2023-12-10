package service

import (
	"context"
	"errors"
	"test_backend/features/notes"
	"test_backend/helper"
	"test_backend/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	ctx  = context.Background()
	repo = new(mocks.NoteData)
	srv  = New(repo)
)

func TestCreate(t *testing.T) {
	inputNote := notes.NotesEntity{
		Title: "catatan 1", Content: "Lorem Ipsum is dummy text of the printing and typesetting industry, derived from a Latin passage by Cicero.",
	}

	t.Run("success create notes", func(t *testing.T) {
		repo.On("Insert", inputNote, ctx).Return("65744db7fd950a80193d89bb", nil).Once()
		id, err := srv.Add(inputNote, ctx)
		assert.NotEmpty(t, id)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("failed create notes", func(t *testing.T) {
		repo.On("Insert", inputNote, ctx).Return("", errors.New("fail create notes")).Once()
		id, err := srv.Add(inputNote, ctx)
		assert.Empty(t, id)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("failed validation struct", func(t *testing.T) {
		inputNote := notes.NotesEntity{
			Title: "catatan 1",
		}
		id, err := srv.Add(inputNote, ctx)
		assert.Empty(t, id)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestGetAll(t *testing.T) {

	returnNote := []notes.NotesEntity{
		{ID: "65744db7fd950a80193d89bb", Title: "catatan 1", Content: "Lorem Ipsum is dummy text of the printing and typesetting industry, derived from a Latin passage by Cicero.", CreatedAt: helper.ParseTime("2023-12-09T11:21:27.887+00:00"), UpdatedAt: helper.ParseTime("2023-12-09T11:21:27.887+00:00")},
	}

	t.Run("success get all notes", func(t *testing.T) {
		repo.On("SelectAll", ctx).Return(returnNote, nil).Once()
		resp, err := srv.GetAll(ctx)
		assert.Equal(t, returnNote, resp)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("failed get all notes", func(t *testing.T) {
		repo.On("SelectAll", ctx).Return(nil, errors.New("fail get all notes")).Once()
		resp, err := srv.GetAll(ctx)
		assert.Nil(t, resp)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}
func TestGetById(t *testing.T) {

	returnNote := notes.NotesEntity{
		ID: "65744db7fd950a80193d89bb", Title: "catatan 1", Content: "Lorem Ipsum is dummy text of the printing and typesetting industry, derived from a Latin passage by Cicero.", CreatedAt: helper.ParseTime("2023-12-09T11:21:27.887+00:00"), UpdatedAt: helper.ParseTime("2023-12-09T11:21:27.887+00:00"),
	}

	t.Run("success get notes by id ", func(t *testing.T) {
		repo.On("SelectById", ctx, "65744db7fd950a80193d89bb").Return(returnNote, nil).Once()
		resp, err := srv.GetById(ctx, "65744db7fd950a80193d89bb")
		assert.Equal(t, returnNote, resp)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("failed get notes by id", func(t *testing.T) {
		repo.On("SelectById", ctx, "65744db7fd950a80193d89bb").Return(notes.NotesEntity{}, errors.New("fail get notes by id")).Once()
		resp, err := srv.GetById(ctx, "65744db7fd950a80193d89bb")
		assert.NotEqual(t, returnNote, resp)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	inputNote := notes.NotesEntity{
		Title: "catatan 1", Content: "Lorem Ipsum is dummy text of the printing and typesetting industry, derived from a Latin passage by Cicero.",
	}
	returnNote := notes.NotesEntity{
		ID: "65744db7fd950a80193d89bb", Title: "catatan 1", Content: "Lorem Ipsum is dummy text of the printing and typesetting industry, derived from a Latin passage by Cicero.", CreatedAt: helper.ParseTime("2023-12-09T11:21:27.887+00:00"), UpdatedAt: helper.ParseTime("2023-12-09T11:21:27.887+00:00"),
	}

	t.Run("success update notes", func(t *testing.T) {
		repo.On("SelectById", ctx, "65744db7fd950a80193d89bb").Return(returnNote, nil).Once()
		repo.On("UpdateById", ctx, "65744db7fd950a80193d89bb", inputNote).Return(nil).Once()
		err := srv.EditById(ctx, "65744db7fd950a80193d89bb", inputNote)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("fail get notes by id", func(t *testing.T) {
		repo.On("SelectById", ctx, "65744db7fd950a80193d89bb").Return(notes.NotesEntity{}, errors.New("error get notes by id")).Once()
		err := srv.EditById(ctx, "65744db7fd950a80193d89bb", inputNote)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("fail update notes by id", func(t *testing.T) {
		repo.On("SelectById", ctx, "65744db7fd950a80193d89bb").Return(returnNote, nil).Once()
		repo.On("UpdateById", ctx, "65744db7fd950a80193d89bb", inputNote).Return(errors.New("error update notes by id")).Once()
		err := srv.EditById(ctx, "65744db7fd950a80193d89bb", inputNote)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	t.Run("success delete note by id", func(t *testing.T) {
		repo.On("DeleteById", ctx, "65744db7fd950a80193d89bb").Return(nil).Once()
		err := srv.DeleteById(ctx, "65744db7fd950a80193d89bb")
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("failed delete note by id", func(t *testing.T) {
		repo.On("DeleteById", ctx, "65744db7fd950a80193d89bb").Return(errors.New("error delete notes by id")).Once()
		err := srv.DeleteById(ctx, "65744db7fd950a80193d89bb")
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}
