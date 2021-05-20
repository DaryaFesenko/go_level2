package task2

import (
	"fmt"
	"time"
)

type MyError struct {
	date    time.Time
	message string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("error: %s\ndatetime: %s", e.message, e.date)
}

func New(message string) error {
	return &MyError{
		message: message,
		date:    time.Now(),
	}
}
