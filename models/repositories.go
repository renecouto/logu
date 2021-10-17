package models

import (
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
	GetAll() AllData
	AddEvent(e Event)
	AddNote(e Note)
	AddTask(e Task)
	GetUserByUsername(username string) *User
	AddUser(user User)
	UpdateTask(userId int64, taskId int64, update Task) *Task
	GetAllForDateAndUser(d time.Time, userId int64) AllData
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
