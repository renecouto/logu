package models

import (
	"fmt"
	"time"
)

type Task struct {
	Description string `form:"Description" json:"Description"`
	Id          int    `form:"Id" json:"Id"`
	Done        bool   `form:"Done" json:"Done"`
	CreatedAt   time.Time
	User        int
}

type Event struct {
	Description  string `form:"Description"`
	Id           int    `form:"Id"`
	ScheduledFor time.Time
	CreatedAt    time.Time
	User         int
}

type Note struct {
	Description string `form:"Description"`
	Id          int    `form:"Id"`
	CreatedAt   time.Time
	User        int
}

type ItemsRepository struct {
	tasks  []Task
	events []Event
	notes  []Note
	users  map[string]User
}

type AllData struct {
	Tasks  []Task
	Events []Event
	Notes  []Note
}

type User struct {
	Id       int
	Username string
	FullName string
}

type Date struct {
	year  int
	month int
	day   int
}

func NewItemsRepository() ItemsRepository {
	return ItemsRepository{users: make(map[string]User)}
}

func (i *ItemsRepository) GetAll() AllData {
	return AllData{i.tasks, i.events, i.notes}
}

func (i *ItemsRepository) AddEvent(e Event) {
	e.Id = len(i.events)
	i.events = append(i.events, e)
}

func (i *ItemsRepository) AddNote(e Note) {
	e.Id = len(i.notes)
	i.notes = append(i.notes, e)
}

func (i *ItemsRepository) AddTask(e Task) {
	e.Id = len(i.tasks)
	i.tasks = append(i.tasks, e)
}

func dateEquals(time1, time2 time.Time) bool {
	y1, m1, d1 := time1.Date()
	y2, m2, d2 := time2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

func (i *ItemsRepository) GetUserByUsername(username string) *User {
	if val, ok := i.users[username]; ok {
		return &val
	} else {
		return nil
	}
}

func (i *ItemsRepository) AddUser(user User) {
	i.users[user.Username] = user
}

func (i *ItemsRepository) GetTaskById(userId int, taskId int) *Task {
	for ff, t := range i.tasks {
		if t.Id == taskId && t.User == userId {
			return &i.tasks[ff]
		}
	}
	text := fmt.Sprintln("task of id", taskId, userId, "was not found")
	panic(text)
}

func (i *ItemsRepository) GetAllForDateAndUser(d time.Time, userId int) AllData {
	var tasksf []Task
	for _, t := range i.tasks {
		if dateEquals(t.CreatedAt, d) && t.User == userId {
			tasksf = append(tasksf, t)
		}
	}

	var eventsf []Event
	for _, t := range i.events {
		if dateEquals(t.CreatedAt, d) && t.User == userId {
			eventsf = append(eventsf, t)
		}
	}

	var notesf []Note
	for _, t := range i.notes {
		if dateEquals(t.CreatedAt, d) && t.User == userId {
			notesf = append(notesf, t)
		}
	}

	return AllData{
		tasksf,
		eventsf,
		notesf,
	}
}
