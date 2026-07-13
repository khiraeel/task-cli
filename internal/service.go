package task

import (
	"context"
)

type taskRepository interface {
	Save(ctx context.Context, task *Task) error
	FindAll(ctx context.Context) (map[int]*Task, error)
	FindById(ctx context.Context, id int) (*Task, error)
	FindByStatus(ctx context.Context, status TaskStatus) (map[int]*Task, error)
	Update(ctx context.Context, task *Task) (*Task, error)
	Delete(ctx context.Context, id int) error
}

type Service struct {
	repo taskRepository
}

func NewService(repo taskRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateTask(title string, description string) (bool, error) {
	task := NewTask(title, description)

	err := s.repo.Save(context.Background(), task)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *Service) UpdateTask(id int, title string, description string) (bool, error) {
	task, err := s.repo.FindById(context.Background(), id)
	if task == nil || err != nil {
		return false, err
	}

	task.SetTitle(title)
	task.SetDescription(description)

	_, err = s.repo.Update(context.Background(), task)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *Service) UpdateTaskStatus(id int, status TaskStatus) (bool, error) {
	task, err := s.repo.FindById(context.Background(), id)
	if err != nil {
		return false, err
	}

	task.SetStatus(status)

	_, err = s.repo.Update(context.Background(), task)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *Service) DeleteTask(id int) (bool, error) {
	err := s.repo.Delete(context.Background(), id)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *Service) GetList(status ...TaskStatus) (map[int]*Task, error) {
	var tasks map[int]*Task
	var err error

	if len(status) == 0 {
		tasks, err = s.repo.FindAll(context.Background())
	} else {
		tasks, err = s.repo.FindByStatus(context.Background(), status[0])
	}

	if err != nil {
		return nil, err
	}
	return tasks, nil
}
