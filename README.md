https://roadmap.sh/projects/task-tracker

go build -o task-cli.exe main.go

.\task-cli add <title> <description>

.\task-cli update <title> <description>

.\task-cli update_status <id> <status>

.\task-cli delete <id>

.\task-cli list <status>
