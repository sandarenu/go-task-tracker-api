package main

type TaskStatus int

const (
	_ TaskStatus = iota
	Pending
	Completed
)

type Task struct {
	TaskId int32      `json:"task_id,omitempty"`
	Title  string     `json:"title,omitempty"`
	Status TaskStatus `json:"status,omitempty"`
}

type CmdData struct {
	Cmd  string
	Data string
}
