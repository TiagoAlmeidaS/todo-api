package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"todo_project.com/internal/app/repository"
	"todo_project.com/internal/domain/task"
)

type TaskRepository struct {
	Collection *mongo.Collection
}

type Task struct {
	ID          primitive.ObjectID  `bson:"_id"`
	IDUser      string              `bson:"id_user"`
	Title       string              `bson:"title"`
	Description string              `bson:"description"`
	Status      string              `bson:"status"`
	DateInit    *primitive.DateTime `bson:"date_init"`
	DateEnd     *primitive.DateTime `bson:"date_end"`
	IDProject   string              `bson:"id_project"`
}

func domainToTask(taskDomain task.Task) *Task {
	objectId, err := primitive.ObjectIDFromHex(taskDomain.ID)
	if err != nil {
		objectId = primitive.NewObjectID()
	}

	taskModel := Task{
		ID:          objectId,
		Title:       taskDomain.Title,
		Description: taskDomain.Description,
		IDUser:      taskDomain.IDUser,
		Status:      string(taskDomain.Status),
	}

	if taskDomain.DateInit != nil {
		dateInit := primitive.NewDateTimeFromTime(*taskDomain.DateInit)
		taskModel.DateInit = &dateInit
	}

	if taskDomain.DateEnd != nil {
		dateEnd := primitive.NewDateTimeFromTime(*taskDomain.DateEnd)
		taskModel.DateEnd = &dateEnd
	}

	return &taskModel
}

func (t Task) toDomain() *task.Task {
	taskDomain := task.Task{
		ID:          t.ID.Hex(),
		Title:       t.Title,
		Description: t.Description,
		Status:      task.Status(t.Status),
		IDUser:      t.IDUser,
	}

	if t.DateInit != nil {
		dateInit := t.DateInit.Time()
		taskDomain.DateInit = &dateInit
	}

	if t.DateEnd != nil {
		dateEnd := t.DateEnd.Time()
		taskDomain.DateEnd = &dateEnd
	}

	return &taskDomain
}

func (t TaskRepository) GetById(id string) (*task.Task, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, repository.ErrTaskNotFound
	}

	taskModel := Task{}
	err = t.Collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&taskModel)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, repository.ErrTaskNotFound
		}
		return nil, err
	}

	return taskModel.toDomain(), nil
}

func (t TaskRepository) Insert(taskDomain task.Task) (*task.Task, error) {
	taskModel := domainToTask(taskDomain)
	_, err := t.Collection.InsertOne(context.Background(), taskModel)
	if err != nil {
		return nil, err
	}

	return taskModel.toDomain(), nil
}

func (t TaskRepository) Update(taskDomain task.Task) (*task.Task, error) {
	taskModel := domainToTask(taskDomain)
	_, err := t.Collection.UpdateOne(context.Background(), bson.M{"_id": taskModel.ID}, bson.M{"$set": taskModel})
	if err != nil {
		return nil, repository.ErrTaskUpdateNotFound
	}

	return taskModel.toDomain(), nil
}

func (t TaskRepository) GetAllByClientId(clientId string) (*[]task.Task, error) {
	var tasks []task.Task

	cur, err := t.Collection.Find(context.Background(), bson.M{"id_user": clientId})
	if err != nil {
		return nil, err
	}

	for cur.Next(context.Background()) {
		taskIn := Task{}
		if err = cur.Decode(&taskIn); err != nil {
			return nil, err
		}
		tasks = append(tasks, *taskIn.toDomain())
	}

	err = cur.Close(context.Background())
	if err != nil {
		return nil, err
	}

	return &tasks, nil
}

func (t TaskRepository) DeleteById(id string) error {
	_, err := t.GetById(id)
	if err != nil {
		return err
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return repository.ErrUserNotFound
	}

	_, err = t.Collection.DeleteOne(context.Background(), bson.M{"_id": objectID})
	if err != nil {
		return err
	}

	return nil
}

func (t TaskRepository) GetResumeStatus(clientId string) (*task.Resume, error) {
	var tasks []task.Task
	var open, process, completed int

	cur, err := t.Collection.Find(context.Background(), bson.M{"id_user": clientId})
	if err != nil {
		return nil, err
	}

	for cur.Next(context.Background()) {
		taskIn := Task{}
		if err = cur.Decode(&taskIn); err != nil {
			return nil, err
		}
		tasks = append(tasks, *taskIn.toDomain())
	}

	for _, taskFor := range tasks {
		if taskFor.Status == task.Open {
			open = open + 1
		}

		if taskFor.Status == task.Process {
			process = process + 1
		}

		if taskFor.Status == task.Completed {
			completed = completed + 1
		}
	}

	return &task.Resume{
		Completed: completed,
		Process:   process,
		Open:      open,
	}, nil

}

func (t TaskRepository) GetAllByDay(day time.Time, clientId string) (*[]task.Task, error) {

	var tasks []task.Task

	cur, err := t.Collection.Find(context.Background(), bson.M{"id_user": clientId})
	if err != nil {
		return nil, err
	}

	for cur.Next(context.Background()) {
		taskIn := Task{}
		if err = cur.Decode(&taskIn); err != nil {
			return nil, err
		}

		if taskIn.DateInit.Time().Before(day) && taskIn.DateEnd.Time().After(day) {
			tasks = append(tasks, *taskIn.toDomain())
		}
	}

	err = cur.Close(context.Background())
	if err != nil {
		return nil, err
	}

	return &tasks, nil

}
