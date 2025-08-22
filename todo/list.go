package todo

import "sync"

type List struct {
	tasks map[string]Task
	mtx   sync.RWMutex
}

func NewList() *List {
	list := List{
		tasks: make(map[string]Task),
	}

	return &list
}

func (l *List) AddTask(task Task) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	if _, ok := l.tasks[task.Title]; ok {
		return ErrTaskAlreadyExists
	}
	l.tasks[task.Title] = task

	return nil
}

func (l *List) GetTask(title string) (Task, error) {
	l.mtx.RLock()
	defer l.mtx.RUnlock()
	task, ok := l.tasks[title]
	if !ok {
		return Task{}, ErrTaskNotFound
	}

	return task, nil
}

func (l *List) ListTasks() map[string]Task {
	l.mtx.RLock()
	defer l.mtx.RUnlock()
	tmp := make(map[string]Task, len(l.tasks))
	for k, v := range l.tasks {
		tmp[k] = v
	}
	return tmp
}

func (l *List) ListNotComletedTasks() map[string]Task {
	l.mtx.RLock()
	defer l.mtx.RUnlock()
	notComletedTasks := make(map[string]Task)

	for title, task := range l.tasks {
		if !task.IsDone {
			notComletedTasks[title] = task
		}
	}
	return notComletedTasks
}

func (l *List) DoneTask(title string) (Task, error) {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	task, ok := l.tasks[title]
	if !ok {
		return Task{}, ErrTaskNotFound
	}

	task.Done()
	l.tasks[title] = task

	return l.tasks[title], nil
}

func (l *List) UnDoneTask(title string) (Task, error) {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	task, ok := l.tasks[title]
	if !ok {
		return Task{}, ErrTaskNotFound
	}

	task.UnDone()
	l.tasks[title] = task

	return l.tasks[title], nil
}

func (l *List) DeleteTask(title string) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	_, ok := l.tasks[title]
	if !ok {
		return ErrTaskNotFound
	}

	delete(l.tasks, title)

	return nil
}
