package http

import (
	"encoding/json"
	"errors"
	"time"
)

type TaskDTO struct {
	Title       string
	Discription string
}

type CompleteTaskDTO struct {
	Complete bool
}

func (t TaskDTO) validateForCreate() error {
	if t.Title == "" {
		return errors.New("title is empty")
	}
	if t.Discription == "" {
		return errors.New("description is empty")
	}

	return nil
}

type ErrorDTO struct {
	Message string
	Time    time.Time
}

func (e ErrorDTO) ToString() string {
	b, err := json.MarshalIndent(e, "", "    ")
	if err != nil {
		panic(err)
	}

	return string(b)
}
