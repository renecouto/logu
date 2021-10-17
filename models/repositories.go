package models

import (
	"context"
	"time"
)

type Task struct {
	Description string `form:"Description" json:"Description"`
	Id          int64  `form:"Id" json:"Id"`
	Done        bool   `form:"Done" json:"Done"`
	CreatedAt   time.Time
	User        int64
}

type Event struct {
	Description  string `form:"Description"`
	Id           int64  `form:"Id"`
	ScheduledFor time.Time
	CreatedAt    time.Time
	User         int64
}

type Note struct {
	Description string `form:"Description"`
	Id          int64  `form:"Id"`
	CreatedAt   time.Time
	User        int64
}

type ItemsRepository interface {
	InitSchema()
	GetAll(ctx context.Context) AllData
	AddEvent(ctx context.Context, e Event)
	AddNote(ctx context.Context, e Note)
	AddTask(ctx context.Context, e Task)
	GetUserByUsername(ctx context.Context, username string) *User
	AddUser(ctx context.Context, user User)
	UpdateTask(ctx context.Context, userId int64, taskId int64, update Task) *Task
	GetAllForDateAndUser(ctx context.Context, d time.Time, userId int64) AllData
}

type AllData struct {
	Tasks  []Task
	Events []Event
	Notes  []Note
}

type User struct {
	Id       int64
	Username string
	FullName string
}
