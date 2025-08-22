package todo

import "time"

type Task struct {
	Title       string
	Discription string
	IsDone      bool
	CreatedAt   time.Time
	DoneAt      *time.Time
}

func NewTask(title string, discription string) Task {
	return Task{
		Title:       title,
		Discription: discription,
		IsDone:      false,
		CreatedAt:   time.Now(),
		DoneAt:      nil,
	}
}

func (t *Task) Done() {
	t.IsDone = true
	DoneTime := time.Now()
	t.DoneAt = &DoneTime
}

func (t *Task) UnDone() {
	t.IsDone = false
	t.DoneAt = nil
}
