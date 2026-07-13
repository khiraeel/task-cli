package task

import (
	"fmt"
	"os"
	"strconv"
	task "task-cli/internal"
)

type CLIApp struct {
	taskService *task.Service
}

func NewCLIApp(svc *task.Service) *CLIApp {
	return &CLIApp{taskService: svc}
}

func (app *CLIApp) Run(args []string) {
	if len(args) < 2 {
		fmt.Println("Ошибка: выберите команду (add, update, delete, list)")
		os.Exit(1)
	}

	command := args[1]

	switch command {
	case "add":
		app.handleAdd(args)
	case "update":
		app.handleUpdate(args)
	case "update_status":
		app.handleUpdateStatus(args)
	case "delete":
		app.handleDelete(args)
	case "list":
		app.handleList(args)
	default:
		fmt.Printf("Неизвестная команда: %s\n", command)
		os.Exit(1)
	}
}

func (app *CLIApp) handleAdd(args []string) {
	if len(args) < 3 {
		fmt.Println("Ошибка: укажите описание задачи")
		os.Exit(1)
	}

	title := args[2]
	description := args[3]

	_, err := app.taskService.CreateTask(title, description)
	if err != nil {
		fmt.Printf("Ошибка при добавлении: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Задача успешно добавлена!")
}

func (app *CLIApp) handleUpdate(args []string) {
	if len(args) < 4 {
		fmt.Println("Ошибка: укажите ID задачи и новое описание")
		os.Exit(1)
	}

	id, _ := strconv.Atoi(args[2])
	title := args[3]

	var description string
	if len(args) > 4 {
		description = args[4]
	}

	result, err := app.taskService.UpdateTask(id, title, description)
	if err != nil {
		fmt.Printf("Ошибка при обновлении: %v\n", err)
		os.Exit(1)
	}

	if result == false {
		fmt.Println("Ошибка при обновлении задачи")
		os.Exit(1)
	}

	fmt.Println("Задача успешно обновлена!")
}

func (app *CLIApp) handleUpdateStatus(args []string) {
	if len(args) < 4 {
		fmt.Println("Ошибка: укажите ID задачи и новый статус")
		os.Exit(1)
	}

	id, _ := strconv.Atoi(args[2])
	var status task.TaskStatus

	status = status.FromString(args[3])

	result, err := app.taskService.UpdateTaskStatus(id, status)
	if err != nil {
		fmt.Printf("Ошибка при обновлении: %v\n", err)
		os.Exit(1)
	}

	if result == false {
		fmt.Println("Ошибка при обновлении статуса задачи")
		os.Exit(1)
	}

	fmt.Println("Статус задачи успешно обновлен!")
}

func (app *CLIApp) handleDelete(args []string) {
	if len(args) < 3 {
		fmt.Println("Ошибка: укажите ID задачи")
		os.Exit(1)
	}

	id, _ := strconv.Atoi(args[2])

	result, err := app.taskService.DeleteTask(id)
	if err != nil {
		fmt.Printf("Ошибка при удалении: %v\n", err)
		os.Exit(1)
	}

	if result == false {
		fmt.Println("Ошибка при удалении задачи")
		os.Exit(1)
	}

	fmt.Println("Задача успешно удалена!")
}

func (app *CLIApp) handleList(args []string) {
	var currentStatus task.TaskStatus

	if len(args) > 2 {
		currentStatus = currentStatus.FromString(args[2])
	}

	tasks, err := app.taskService.GetList(currentStatus)
	if err != nil {
		fmt.Printf("Ошибка при получении списка задач: %v\n", err)
		os.Exit(1)
	}

	for _, task := range tasks {
		status := task.GetStatus()

		fmt.Printf("ID: %d, Title: %s, Description: %s, Status: %s, CreatedAt: %s, UpdatedAt: %s\n",
			task.GetID(),
			task.GetTitle(),
			task.GetDescription(),
			status.ToString(),
			task.GetCreatedAt().Format("2006-01-02 15:04:05"),
			task.GetUpdatedAt().Format("2006-01-02 15:04:05"),
		)
	}
}
