package main

//Импорт библиотек
import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Основная функция в которой используется для запуска роутер и инициализируется локальный порт
func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getOneTodo)
	router.POST("/todos", addTodo)
	router.Run("localhost:8082")

}

// Простая структура списка дел и их формат отображения
type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

// Переменная содержащая массив, с данными списка дел
var todos = []todo{
	{"1", "Clean room", false},
	{"2", "Read book", false},
	{"3", "Record video", false},
}

// Метод получения задач и отображения их в json-формате
func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

// Функция поиска задачи по ее ID
func getTodoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found ")
}

// Функция получения задачи по ее ID
func getOneTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, todo)

}

// Функция для добавления задачи
func addTodo(context *gin.Context) {
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}
	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)

}
