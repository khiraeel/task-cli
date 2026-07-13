package task

import (
	"context"
	"errors"
	"fmt"
	"os"
)

type JSONTaskRepository struct {
	filePath string
}

func NewJSONRepository(filePath string) *JSONTaskRepository {
	return &JSONTaskRepository{
		filePath: filePath,
	}
}

func (r *JSONTaskRepository) Save(ctx context.Context, task *Task) error {
	fileData, err := r.getFileData()
	if err != nil {
		return errors.New("ошибка при чтении файла: " + err.Error())
	}

	tasks, err := ToEntities(fileData)
	lastTask, found := r.GetLastByMaxID(tasks)
	if (lastTask == nil && !found) || err != nil {
		task.SetID(1)
	} else {
		task.SetID(lastTask.GetID() + 1)
	}

	tasks[task.GetID()] = task
	json, _ := FromEntities(tasks)

	err = os.WriteFile(r.filePath, json, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (r *JSONTaskRepository) FindAll(ctx context.Context) (map[int]*Task, error) {
	fileData, err := r.getFileData()
	if err != nil {
		return nil, err
	}

	tasks, err := ToEntities(fileData)

	return tasks, err
}

func (r *JSONTaskRepository) FindById(ctx context.Context, id int) (*Task, error) {
	fileData, err := r.getFileData()
	if err != nil {
		return nil, err
	}

	tasks, err := ToEntities(fileData)

	if taskItem, ok := tasks[id]; ok {
		return taskItem, nil
	}

	return nil, nil
}

func (r *JSONTaskRepository) FindByStatus(ctx context.Context, status TaskStatus) (map[int]*Task, error) {
	fileData, err := r.getFileData()
	if err != nil {
		return nil, err
	}

	tasks, err := ToEntities(fileData)
	filteredTasks := make(map[int]*Task)

	for _, task := range tasks {
		if status == 0 || task.GetStatus() == status {
			filteredTasks[task.GetID()] = task
		}
	}

	return filteredTasks, nil
}

func (r *JSONTaskRepository) Update(ctx context.Context, task *Task) (*Task, error) {
	fileData, err := r.getFileData()
	if err != nil {
		return task, err
	}

	tasks, err := ToEntities(fileData)

	if taskItem, ok := tasks[task.GetID()]; ok {
		taskItem.SetTitle(task.GetTitle())
		taskItem.SetDescription(task.GetDescription())
		taskItem.SetStatus(task.GetStatus())
		taskItem.SetUpdatedAt(task.GetUpdatedAt())

		json, _ := FromEntities(tasks)

		err := os.WriteFile(r.filePath, json, 0644)
		if err != nil {
			return task, err
		}
		return task, nil
	}

	return task, errors.New("task not found")
}

func (r *JSONTaskRepository) Delete(ctx context.Context, id int) error {
	fileData, err := r.getFileData()
	if err != nil {
		return err
	}

	tasks, err := ToEntities(fileData)

	if ok := tasks[id]; ok != nil {
		delete(tasks, id)

		json, _ := FromEntities(tasks)

		err := os.WriteFile(r.filePath, json, 0644)
		if err != nil {
			return err
		}
		return nil
	}

	return errors.New("task not found")
}

func (r *JSONTaskRepository) getFileData() ([]byte, error) {
	fileData, err := os.ReadFile(r.filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			defaultData := []byte("[]")
			if err := os.WriteFile(r.filePath, defaultData, 0644); err != nil {
				return nil, fmt.Errorf("create default file: %w", err)
			}
			return defaultData, nil
		}
		return nil, err
	}

	return fileData, nil
}

func (r *JSONTaskRepository) GetLastByMaxID(tasks map[int]*Task) (*Task, bool) {
	if len(tasks) == 0 {
		return nil, false
	}

	var maxID int
	var last *Task
	first := true

	for id, user := range tasks {
		if first || id > maxID {
			maxID = id
			last = user
			first = false
		}
	}

	return last, true
}
