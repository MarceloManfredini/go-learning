package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"task-api/database"
	"task-api/handlers"
	"task-api/models"
	"task-api/repository"
	"task-api/services"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// setupTestDB cria uma conexão SQLite em memória
func setupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to test database")
	}
	// Cria a tabela Task apenas em memória
	db.AutoMigrate(&models.Task{})
	return db
}

// setupRouter cria uma instância do Gin e injeta dependências para os testes
func setupRouter(db *gorm.DB) *gin.Engine {
	gin.SetMode(gin.TestMode)

	// Substitui o DB global pela instância de teste
	database.DB = db

	repo := repository.NewTaskRepository()
	svc := services.NewTaskService(repo)
	handler := handlers.NewTaskHandler(svc)

	r := gin.Default()
	r.POST("/api/tasks/", handler.CreateTask)
	r.GET("/api/tasks/", handler.GetTasks)
	return r
}

func TestCreateAndListTask(t *testing.T) {
	db := setupTestDB()
	r := setupRouter(db)

	// Cria uma nova tarefa
	body := []byte(`{"title":"Test task","detail":"abc"}`)
	req, _ := http.NewRequest("POST", "/api/tasks/", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	// Lista as tarefas
	req2, _ := http.NewRequest("GET", "/api/tasks/", nil)
	w2 := httptest.NewRecorder()
	r.ServeHTTP(w2, req2)
	assert.Equal(t, http.StatusOK, w2.Code)
	assert.Contains(t, w2.Body.String(), "Test task")
}
