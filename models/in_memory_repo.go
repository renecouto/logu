package models

import (
	"context"
	"fmt"
	"time"

	"github.com/renecouto/logu/utils"
)

type InMemoryItemsRepository struct {
	tasks  []Task
	events []Event
	notes  []Note
	users  map[string]User
}

func (i *InMemoryItemsRepository) InitSchema() {
	i.users = make(map[string]User)
}

func (i *InMemoryItemsRepository) GetAll(ctx context.Context) AllData {
	return AllData{i.tasks, i.events, i.notes}
}

func (i *InMemoryItemsRepository) AddEvent(ctx context.Context, e Event) {
	e.Id = int64(len(i.events))
	i.events = append(i.events, e)
}

func (i *InMemoryItemsRepository) AddNote(ctx context.Context, e Note) {
	e.Id = int64(len(i.notes))
	i.notes = append(i.notes, e)
}

func (i *InMemoryItemsRepository) AddTask(ctx context.Context, e Task) {
	e.Id = int64(len(i.tasks))
	i.tasks = append(i.tasks, e)
}

func (i *InMemoryItemsRepository) GetUserByUsername(ctx context.Context, username string) *User {
	if val, ok := i.users[username]; ok {
		return &val
	} else {
		return nil
	}
}

func (i *InMemoryItemsRepository) AddUser(ctx context.Context, user User) {
	i.users[user.Username] = user
}

func (i *InMemoryItemsRepository) GetTaskById(ctx context.Context, userId int64, taskId int64) *Task {
	for ff, t := range i.tasks {
		if t.Id == taskId && t.User == userId {
			return &i.tasks[ff]
		}
	}
	text := fmt.Sprintln("task of id", taskId, userId, "was not found")
	panic(text)
}

func (i *InMemoryItemsRepository) UpdateTask(ctx context.Context, userId int64, taskId int64, update Task) *Task {
	t := i.GetTaskById(ctx, userId, taskId)
	t.Done = update.Done
	return t
}

func (i *InMemoryItemsRepository) GetAllForDateAndUser(ctx context.Context, d time.Time, userId int64) AllData {
	var tasksf []Task
	for _, t := range i.tasks {
		if utils.DateEquals(t.CreatedAt, d) && t.User == userId {
			tasksf = append(tasksf, t)
		}
	}

	var eventsf []Event
	for _, t := range i.events {
		if utils.DateEquals(t.CreatedAt, d) && t.User == userId {
			eventsf = append(eventsf, t)
		}
	}

	var notesf []Note
	for _, t := range i.notes {
		if utils.DateEquals(t.CreatedAt, d) && t.User == userId {
			notesf = append(notesf, t)
		}
	}

	return AllData{
		tasksf,
		eventsf,
		notesf,
	}
}
