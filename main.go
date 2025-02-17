package main

import (
	"context"
	"embed"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

type ToDoApp struct {
	Tasks []ToDo
}

type ToDo struct {
	Task string
	Done bool
}

// Инициализация нового приложения
func NewApp() *ToDoApp {
	return &ToDoApp{
		Tasks: []ToDo{},
	}
}
func (app *ToDoApp) GetTasks() []ToDo {
	return app.Tasks
}

// Метод для запуска приложения (теперь с контекстом)
func (app *ToDoApp) startup(ctx context.Context) {
	log.Println("ToDoApp started!")
}

// Добавить задачу
func (app *ToDoApp) AddTask(task string) {
	app.Tasks = append(app.Tasks, ToDo{Task: task, Done: false})
}

// Удалить задачу по индексу
func (app *ToDoApp) RemoveTask(index int) {
	if index < 0 || index >= len(app.Tasks) {
		return
	}
	app.Tasks = append(app.Tasks[:index], app.Tasks[index+1:]...)
}

// Изменить статус задачи (отметить выполненной)
func (app *ToDoApp) ToggleTask(index int) {
	if index < 0 || index >= len(app.Tasks) {
		return
	}
	app.Tasks[index].Done = !app.Tasks[index].Done
}

func main() {
	// Создаем новый экземпляр приложения
	app := NewApp()

	// Запуск Wails-приложения
	err := wails.Run(&options.App{
		Title:  "ToDo App", // Название приложения
		Width:  1024,       // Ширина окна
		Height: 768,        // Высота окна
		AssetServer: &assetserver.Options{
			Assets: assets, // Встраиваемый фронтенд
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1}, // Цвет фона
		OnStartup:        app.startup,                              // Функция старта приложения
		Bind: []interface{}{
			app, // Привязка методов Go (доступных для Vue)
		},
	})

	if err != nil {
		log.Fatalf("Error: %v", err) // Обработка ошибок
	}
}
