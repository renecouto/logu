package models

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/renecouto/logu/psql"
)

func convTaskFromPg(s psql.Task) Task {
	return Task{Id: s.ID, Description: s.Description, CreatedAt: s.CreatedAt, User: s.UserID, Done: s.Done}
}

func convTaskToPg(s Task) psql.Task {
	return psql.Task{ID: s.Id, Description: s.Description, CreatedAt: s.CreatedAt, UserID: s.User, Done: s.Done}
}

type PgItemsRepository struct {
	queries *psql.Queries
	db      *sql.DB
}

func NewPgItemsRepo(q *psql.Queries, db *sql.DB) *PgItemsRepository {
	return &PgItemsRepository{q, db}
}

func (i *PgItemsRepository) InitSchema() {
}

func (i *PgItemsRepository) GetAll(ctx context.Context) AllData {
	return AllData{}
}

func (i *PgItemsRepository) AddEvent(ctx context.Context, e Event) {
	_, err := i.queries.CreateEvent(ctx, psql.CreateEventParams{Description: e.Description, CreatedAt: e.CreatedAt, UserID: e.User, ScheduledFor: e.ScheduledFor})
	if err != nil {
		panic(err)
	}
}

func (i *PgItemsRepository) AddNote(ctx context.Context, e Note) {
	_, err := i.queries.CreateNote(ctx, psql.CreateNoteParams{Description: e.Description, CreatedAt: e.CreatedAt, UserID: e.User})
	if err != nil {
		panic(err)
	}
}

func (i *PgItemsRepository) AddTask(ctx context.Context, e Task) {

	_, err := i.queries.CreateTask(ctx, psql.CreateTaskParams{Description: e.Description, CreatedAt: e.CreatedAt, UserID: e.User})
	if err != nil {
		log.Println("got err from create task: ", err)
	}
}

func (i *PgItemsRepository) GetUserByUsername(ctx context.Context, username string) *User {
	res, err := i.queries.GetUser(ctx, username)
	if err != nil {
		panic(err)
	}
	return &User{Id: res.ID, FullName: res.Fullname, Username: res.Username}
}

func (i *PgItemsRepository) AddUser(ctx context.Context, user User) {
	i.queries.CreateUser(ctx, psql.CreateUserParams{ID: user.Id, Fullname: user.FullName, Username: user.Username})
}

func (i *PgItemsRepository) UpdateTask(ctx context.Context, userId int64, taskId int64, update Task) *Task {
	res, err := i.queries.UpdateTask(ctx, psql.UpdateTaskParams{UserID: userId, ID: taskId, Done: update.Done})
	if err != nil {
		panic(err)
	}
	r2 := convTaskFromPg(res)
	return &r2
}

func (i *PgItemsRepository) GetAllForDateAndUser(ctx context.Context, d time.Time, userId int64) AllData {

	tasksx, err := i.queries.ListTasks(ctx, psql.ListTasksParams{UserID: userId, Column2: d})
	if err != nil {
		panic(err)
	}
	var tasksf []Task
	for _, v := range tasksx {
		tasksf = append(tasksf, convTaskFromPg(v))
	}

	notesx, err := i.queries.ListNotes(ctx, psql.ListNotesParams{UserID: userId, Column2: d})
	if err != nil {
		panic(err)
	}
	var notesf []Note
	for _, v := range notesx {
		notesf = append(notesf, Note{Description: v.Description, User: v.UserID, Id: v.ID, CreatedAt: v.CreatedAt})
	}

	eventsx, err := i.queries.ListEvents(ctx, psql.ListEventsParams{UserID: userId, Column2: d})
	if err != nil {
		panic(err)
	}
	var eventsf []Event
	for _, v := range eventsx {
		eventsf = append(eventsf, Event{Description: v.Description, User: v.UserID, Id: v.ID, CreatedAt: v.CreatedAt, ScheduledFor: v.ScheduledFor})
	}

	return AllData{
		tasksf,
		eventsf,
		notesf,
	}
}
