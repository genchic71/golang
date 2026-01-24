package main

import "fmt"

type Logger interface {
	Log(msg string)
}

// zad14
type StdoutLogger struct{}

func (l StdoutLogger) Log(msg string) {
	fmt.Println("REAL:", msg)
}

type MockLogger struct {
	Called      bool
	LastMessage string
}

func (m *MockLogger) Log(msg string) {
	m.Called = true
	m.LastMessage = msg
}

func Work(l Logger) {
	l.Log("test message")
}

func TestWork() {
	mock := &MockLogger{}
	Work(mock)

	if !mock.Called {
		fmt.Println("FAIL: Log не был вызван")
	} else if mock.LastMessage != "test message" {
		fmt.Printf("FAIL: неверное сообщение: %s\n", mock.LastMessage)
	} else {
		fmt.Println("PASS: тест пройден")
	}
}

func main() {

	real := StdoutLogger{}
	Work(real)

	TestWork()
}
