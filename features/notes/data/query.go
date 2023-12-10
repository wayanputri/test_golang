package data

import (
	"context"
	"errors"
	"fmt"
	"log"
	"test_backend/features/notes"
	"test_backend/helper"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type NotesData struct {
	db   *mongo.Client
	coll *mongo.Collection
}

// DeleteById implements notes.NotesDataInterface.
func (repo *NotesData) DeleteById(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	objectID, errObject := primitive.ObjectIDFromHex(id)
	if errObject != nil {
		fmt.Println("error id", errObject)
		return helper.ErrNoteNotFound
	}
	filter := bson.M{"_id": objectID, "deleted": false}
	update := bson.M{
		"$set": bson.M{
			"deleted":    true,
			"deleted_at": time.Now().UTC(),
		},
	}

	_, err := repo.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil

}

// UpdateById implements notes.NotesDataInterface.
func (repo *NotesData) UpdateById(ctx context.Context, id string, note notes.NotesEntity) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	objectID, errObject := primitive.ObjectIDFromHex(id)
	if errObject != nil {
		fmt.Println("error id", errObject)
		return helper.ErrNoteNotFound
	}
	filter := bson.M{"_id": objectID, "deleted": false}
	update := bson.M{
		"$set": bson.M{
			"title":      note.Title,
			"content":    note.Content,
			"updated_at": time.Now(),
		},
	}
	_, err := repo.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

// SelectById implements notes.NotesDataInterface.
func (repo *NotesData) SelectById(ctx context.Context, id string) (notes.NotesEntity, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	objectID, errObject := primitive.ObjectIDFromHex(id)
	if errObject != nil {
		fmt.Println("error id", errObject)
		return notes.NotesEntity{}, helper.ErrNoteNotFound
	}
	filter := bson.M{"_id": objectID, "deleted": false}
	var result Notes
	err := repo.coll.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return notes.NotesEntity{}, err
		}
		panic(err)
	}

	resp := ModelToEntity(result)
	log.Println("Entity response:", resp)
	return resp, nil
}

// SelectAll implements notes.NotesDataInterface.
func (repo *NotesData) SelectAll(ctx context.Context) ([]notes.NotesEntity, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"deleted": false}
	result, err := repo.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer result.Close(ctx)
	var notesList []Notes
	if err := result.All(ctx, &notesList); err != nil {
		return nil, err
	}

	response := ListModelToEntity(notesList)
	return response, nil
}

// Insert implements notes.NotesDataInterface.
func (repo *NotesData) Insert(note notes.NotesEntity, ctx context.Context) (string, error) {
	model := EntityToModel(note)
	model.Deleted = false
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	result, err := repo.coll.InsertOne(ctx, model)
	if err != nil {
		fmt.Println("error insert to mongodb : ", err)
		return "", errors.New("failed create notes")
	}
	insertedID := result.InsertedID.(primitive.ObjectID).Hex()
	return insertedID, nil
}

func New(db *mongo.Client, coll *mongo.Collection) notes.NotesDataInterface {
	return &NotesData{
		db:   db,
		coll: coll,
	}
}
