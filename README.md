# Task Tracker CLI

Task Tracker CLI is a simple and lightweight command-line tool written in Go that helps you manage your daily tasks directly from the terminal. It supports adding, updating, deleting, and listing tasks, with data stored in a local `tasks.json` file.

## 🚀 Features

- ✅ Add new tasks with descriptions
- ✏️ Update task descriptions or status
- 🗑️ Delete tasks by ID
- 📋 List all tasks or filter by status (`todo`, `in-progress`, `done`)
- 🕹️ Simple CLI interface that works anywhere once installed

## 📦 Prerequisites

- [Go](https://golang.org/dl/) (version 1.18 or later)

## 🔧 Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/task-tracker.git
   cd task-tracker
   ```

2. Build the binary:

   ```bash
   go build -o task-cli
   ```

3. (Optional) Install globally:

   Move the binary to a directory in your `PATH` to use `task-cli` as a global command:

   ```bash
   sudo mv task-cli /usr/local/bin/
   ```

   Now you can use `task-cli` from anywhere on your system!

## 🛠️ Usage

Once installed, use the CLI like this:

### ➕ Adding a task

```bash
task-cli add "Buy groceries"
```

### ✏️ Updating a task

```bash
task-cli update 1 "Buy groceries and cook dinner"
```

### ✅ Marking a task as done

```bash
task-cli mark-done 1
```

### 🔄 Marking a task as in progress

```bash
task-cli mark-in-progress 1
```

### 🗑️ Deleting a task

```bash
task-cli delete 1
```

### 📋 Listing tasks

List all tasks:

```bash
task-cli list
```

List by status:

```bash
task-cli list todo
task-cli list in-progress
task-cli list done
```

## 💡 Example Workflow

```bash
task-cli add "Finish the Go CLI project"
task-cli list
task-cli mark-in-progress 1
task-cli update 1 "Finish and polish the Go CLI project"
task-cli mark-done 1
task-cli delete 1
```
