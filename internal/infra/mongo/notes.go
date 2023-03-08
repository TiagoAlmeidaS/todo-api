package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
	"time"
	"todo_project.com/internal/app/repository"
	"todo_project.com/internal/domain/notes"
)

type NotesRepository struct {
	Collection *mongo.Collection
}

type Notes struct {
	ID          primitive.ObjectID  `bson:"_id"`
	IDUser      string              `bson:"id_user"`
	Title       string              `bson:"title"`
	Description string              `bson:"description"`
	DateCreated *primitive.DateTime `bson:"date_created"`
	DateUpdate  *primitive.DateTime `bson:"date_update"`
}

func domainToNotes(notesDomain notes.Notes) *Notes {
	objectId, err := primitive.ObjectIDFromHex(notesDomain.ID)
	if err != nil {
		objectId = primitive.NewObjectID()
	}

	notesModel := Notes{
		ID:          objectId,
		Title:       notesDomain.Title,
		Description: notesDomain.Description,
		IDUser:      notesDomain.IDUser,
	}

	if notesDomain.DateCreated != nil {
		dateCreated := primitive.NewDateTimeFromTime(*notesDomain.DateCreated)
		notesModel.DateCreated = &dateCreated
	}

	if notesDomain.DateUpdate != nil {
		dateUpdate := primitive.NewDateTimeFromTime(*notesDomain.DateUpdate)
		notesModel.DateUpdate = &dateUpdate
	}

	return &notesModel
}

func (t Notes) toDomain() *notes.Notes {
	notesDomain := notes.Notes{
		ID:          t.ID.Hex(),
		Title:       t.Title,
		Description: t.Description,
		IDUser:      t.IDUser,
	}

	if t.DateCreated != nil {
		dateCreated := t.DateCreated.Time()
		notesDomain.DateCreated = &dateCreated
	}

	if t.DateUpdate != nil {
		DateUpdate := t.DateUpdate.Time()
		notesDomain.DateUpdate = &DateUpdate
	}

	return &notesDomain
}

func (t NotesRepository) GetById(id string) (*notes.Notes, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, repository.ErrNotesNotFound
	}

	notesModel := Notes{}
	err = t.Collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&notesModel)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, repository.ErrNotesNotFound
		}
		return nil, err
	}

	return notesModel.toDomain(), nil
}

func (t NotesRepository) Insert(notesDomain notes.Notes) (*notes.Notes, error) {
	notesModel := domainToNotes(notesDomain)
	_, err := t.Collection.InsertOne(context.Background(), notesModel)
	if err != nil {
		return nil, err
	}

	return notesModel.toDomain(), nil
}

func (t NotesRepository) Update(notesDomain notes.Notes) (*notes.Notes, error) {
	notesModel := domainToNotes(notesDomain)
	_, err := t.Collection.UpdateOne(context.Background(), bson.M{"_id": notesModel.ID}, bson.M{"$set": notesModel})
	if err != nil {
		return nil, repository.ErrNotesUpdateNotFound
	}

	return notesModel.toDomain(), nil
}

func (t NotesRepository) GetAllByClientId(clientId string) (*[]notes.Notes, error) {
	var notesList []notes.Notes

	cur, err := t.Collection.Find(context.Background(), bson.M{"id_user": clientId})
	if err != nil {
		return nil, err
	}

	for cur.Next(context.Background()) {
		taskIn := Notes{}
		if err = cur.Decode(&taskIn); err != nil {
			return nil, err
		}
		notesList = append(notesList, *taskIn.toDomain())
	}

	err = cur.Close(context.Background())
	if err != nil {
		return nil, err
	}

	return &notesList, nil
}

func (t NotesRepository) DeleteById(id string) error {
	_, err := t.GetById(id)
	if err != nil {
		return err
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return repository.ErrNotesNotFound
	}

	_, err = t.Collection.DeleteOne(context.Background(), bson.M{"_id": objectID})
	if err != nil {
		return err
	}

	return nil
}

func (t NotesRepository) GetAllByDay(dateInput time.Time, clientId string) (*[]notes.Notes, error) {

	var notesList []notes.Notes

	cur, err := t.Collection.Find(context.Background(), bson.M{"id_user": clientId})
	if err != nil {
		return nil, err
	}

	for cur.Next(context.Background()) {
		taskIn := Notes{}
		if err = cur.Decode(&taskIn); err != nil {
			return nil, err
		}

		year, month, day := taskIn.DateCreated.Time().Date()
		yearInput, monthInput, dayInput := dateInput.Date()
		if (year == yearInput) && (month == monthInput) && (day == dayInput) {
			notesList = append(notesList, *taskIn.toDomain())
		}
	}

	err = cur.Close(context.Background())
	if err != nil {
		return nil, err
	}

	return &notesList, nil

}

func (t NotesRepository) GetByName(nameTask string, clientId string) (*[]notes.Notes, error) {
	var notesList []notes.Notes

	cur, err := t.Collection.Find(context.Background(), bson.M{"id_user": clientId})
	if err != nil {
		return nil, err
	}

	for cur.Next(context.Background()) {
		taskIn := Notes{}
		if err = cur.Decode(&taskIn); err != nil {
			return nil, err
		}

		if strings.Contains(strings.ToLower(taskIn.Title), strings.ToLower(nameTask)) {
			notesList = append(notesList, *taskIn.toDomain())
		}
	}

	err = cur.Close(context.Background())
	if err != nil {
		return nil, err
	}

	return &notesList, nil
}
