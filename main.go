package main

import (
	"os"
	cmd "task-cli/cmd/task-cli"
	task "task-cli/internal"
)

func main() {
	repo := task.NewJSONRepository("tasks.json")

	svc := task.NewService(repo)

	app := cmd.NewCLIApp(svc)

	app.Run(os.Args)
}
