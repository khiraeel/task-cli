package task

import (
	"encoding/json"
	"time"
)

type TaskStatus int

const (
	None TaskStatus = iota
	Todo
	InProgress
	Done
)

func (status *TaskStatus) FromString(value string) TaskStatus {
	switch value {
	case "todo":
		return Todo
	case "in_progress":
		return InProgress
	case "done":
		return Done
	default:
		return None
	}
}

func (status *TaskStatus) ToString() string {
	switch *status {
	case Todo:
		return "todo"
	case InProgress:
		return "in_progress"
	case Done:
		return "done"
	default:
		return "none"
	}
}

type Task struct {
	id          int
	title       string
	description string
	status      TaskStatus
	createdAt   time.Time
	updatedAt   time.Time
}

func NewTask(title, description string) *Task {
	return &Task{
		id:          0,
		title:       title,
		description: description,
		status:      Todo,
		createdAt:   time.Now(),
		updatedAt:   time.Now(),
	}
}

func (t *Task) SetID(id int) {
	t.id = id
}

func (t *Task) GetID() int {
	return t.id
}

func (t *Task) GetTitle() string {
	return t.title
}

func (t *Task) SetTitle(title string) {
	t.title = title
}

func (t *Task) GetDescription() string {
	return t.description
}

func (t *Task) SetDescription(description string) {
	t.description = description
}

func (t *Task) GetStatus() TaskStatus {
	return t.status
}

func (t *Task) SetStatus(status TaskStatus) {
	t.status = status
}

func (t *Task) GetCreatedAt() time.Time {
	return t.createdAt
}
func (t *Task) GetUpdatedAt() time.Time {
	return t.updatedAt
}

func (t *Task) SetUpdatedAt(updatedAt time.Time) {
	t.updatedAt = updatedAt
}

type taskDTO struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

func (t Task) FromEntity() ([]byte, error) {
	dto := taskDTO{
		Id:          t.id,
		Title:       t.title,
		Description: t.description,
		Status:      t.status,
		CreatedAt:   t.createdAt,
		UpdatedAt:   t.updatedAt,
	}

	return json.MarshalIndent(dto, "", "  ")
}

func (t *Task) ToEntity(data []byte) (*Task, error) {
	var dto taskDTO
	if err := json.Unmarshal(data, &dto); err != nil {
		return nil, err
	}

	t.id = dto.Id
	t.title = dto.Title
	t.description = dto.Description
	t.status = dto.Status
	t.createdAt = dto.CreatedAt
	t.updatedAt = dto.UpdatedAt

	return t, nil
}

func ToEntities(data []byte) (map[int]*Task, error) {
	var tasksDto []taskDTO
	if err := json.Unmarshal(data, &tasksDto); err != nil {
		return nil, err
	}

	// tasks := make([]*Task, len(dto))
	tasks := make(map[int]*Task)

	for _, dto := range tasksDto {
		task := &Task{
			id:          dto.Id,
			title:       dto.Title,
			description: dto.Description,
			status:      dto.Status,
			createdAt:   dto.CreatedAt,
			updatedAt:   dto.UpdatedAt,
		}

		tasks[task.id] = task
	}

	return tasks, nil
}

func FromEntities(tasks map[int]*Task) ([]byte, error) {
	var tasksDto []taskDTO

	for _, t := range tasks {
		dto := taskDTO{
			Id:          t.id,
			Title:       t.title,
			Description: t.description,
			Status:      t.status,
			CreatedAt:   t.createdAt,
			UpdatedAt:   t.updatedAt,
		}
		tasksDto = append(tasksDto, dto)
	}

	return json.MarshalIndent(tasksDto, "", "  ")
}
