package handler

import (
	"context"
	"net/http"
	"strings"
	"test_backend/features/notes"
	"test_backend/helper"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type NoteHandler struct {
	noteHandler notes.NotesServiceInterface
	log         *logrus.Entry
}

// CreateNotes godoc
//
//	@Summary		Create Note
//	@Description	create new Note with the provided details
//	@Tags			Notes
//	@Accept			json
//	@Produce		json
//	@Param			request	body		NotesRequest	true	"New Note"
//	@Success		201		{object}	helper.Response
//	@Failure		400		{object}	helper.Response
//	@Failure		500		{object}	helper.Response
//	@Router			/notes [post]
func (handler *NoteHandler) AddNotes(c *gin.Context) {
	var req NotesRequest
	errBind := c.ShouldBindJSON(&req)
	if errBind != nil {
		handler.log.Errorf("binding failed: %v", errBind)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "failed binding note",
		})
		return
	}

	entity := RequestToEntity(req)
	id, err := handler.noteHandler.Add(entity, context.Background())
	if err != nil {
		if strings.Contains(err.Error(), "error validate") {
			handler.log.Errorf("create failed: %v", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "failed insert data, title/content required",
			})
			return
		} else {
			handler.log.Errorf("create failed: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status":  "fail",
				"message": "failed insert data note",
			})
			return
		}

	}
	c.JSON(http.StatusCreated, helper.Response{
		Id:      id,
		Status:  "success",
		Message: "Success create note",
	})
}

// ListNotes godoc
//
//	@Summary		list Notes
//	@Description	list Notes
//	@Tags			Notes
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		helper.Response
//	@Failure		500	{object}	helper.Response
//	@Router			/notes [get]
func (handler *NoteHandler) GetAll(c *gin.Context) {
	data, err := handler.noteHandler.GetAll(context.Background())
	if err != nil {
		handler.log.Errorf("get list failed: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "fail",
			"message": "failed get list note",
		})
		return
	}
	resp := ListEntityToResponse(data)
	c.JSON(http.StatusOK, helper.Response{
		Data:    resp,
		Status:  "success",
		Message: "Success get list note",
	})
}

// DetailNote godoc
//
//	@Summary		Note By Id
//	@Description	get note by id  with id note
//	@Tags			Notes
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Note ID"
//	@Success		200	{object}	helper.Response
//	@Failure		500	{object}	helper.Response
//	@Router			/notes/{id} [get]
func (handler *NoteHandler) GetById(c *gin.Context) {
	id := c.Param("id")

	data, err := handler.noteHandler.GetById(context.Background(), id)
	if err != nil {
		handler.log.Errorf("get by id failed: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "fail",
			"message": "failed get note by id",
		})
		return
	}
	resp := EntityToResponse(data)
	c.JSON(http.StatusOK, helper.Response{
		Data:    resp,
		Status:  "success",
		Message: "Success get note by id",
	})
}

// UpdateNotes godoc
//
//	@Summary		Update Notes
//	@Description	update Note with the provided details
//	@Tags			Notes
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string			true	"Note ID"
//	@Param			request	body		NotesRequest	true	"Updates Note"
//	@Success		200		{object}	helper.Response
//	@Failure		400		{object}	helper.Response
//	@Failure		500		{object}	helper.Response
//	@Router			/notes/{id} [put]
func (handler *NoteHandler) UpdateById(c *gin.Context) {
	var req NotesRequest
	errBind := c.ShouldBindJSON(&req)
	if errBind != nil {
		handler.log.Errorf("binding failed: %v", errBind)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "failed binding data note",
		})
		return
	}
	entity := RequestToEntity(req)
	id := c.Param("id")
	err := handler.noteHandler.EditById(context.Background(), id, entity)
	if err != nil {
		handler.log.Errorf("update failed: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "fail",
			"message": "failed update note",
		})
		return
	}
	c.JSON(http.StatusOK, helper.Response{
		Status:  "success",
		Message: "Success update note",
	})
}

// DeleteNote godoc
//
//	@Summary		Delete Note
//	@Description	Delete a Note by its Id
//	@Tags			Notes
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Note ID"
//	@Success		200	{object}	helper.Response
//	@Failure		500	{object}	helper.Response
//	@Router			/notes/{id} [delete]
func (handler *NoteHandler) DeleteById(c *gin.Context) {
	id := c.Param("id")
	err := handler.noteHandler.DeleteById(context.Background(), id)
	if err != nil {
		handler.log.Errorf("delete failed: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "fail",
			"message": "failed delete note by id",
		})
		return
	}
	c.JSON(http.StatusOK, helper.Response{
		Status:  "success",
		Message: "Success delete note",
	})
}

func New(handler notes.NotesServiceInterface, log *logrus.Entry) *NoteHandler {
	return &NoteHandler{
		noteHandler: handler,
		log:         log,
	}
}
