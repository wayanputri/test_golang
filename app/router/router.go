package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"

	dataNote "test_backend/features/notes/data"
	handlerNote "test_backend/features/notes/handler"
	serviceNote "test_backend/features/notes/service"
)

func InitRouter(db *mongo.Client, r *gin.Engine, log *logrus.Entry, coll *mongo.Collection) {
	DNote := dataNote.New(db, coll)
	SNote := serviceNote.New(DNote)
	HNote := handlerNote.New(SNote, log)
	r.POST("/notes", HNote.AddNotes)
	r.GET("/notes", HNote.GetAll)
	r.GET("/notes/:id", HNote.GetById)
	r.PUT("/notes/:id", HNote.UpdateById)
	r.DELETE("/notes/:id", HNote.DeleteById)
}
