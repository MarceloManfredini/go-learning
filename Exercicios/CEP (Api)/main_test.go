// arquivo: main_test.go
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHandleCEP_Sucesso(t *testing.T) {
	// Inicializa o router Gin em modo teste
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/cep/:cep", handleCEP)

	// Cria uma requisição fake
	req, _ := http.NewRequest("GET", "/cep/01001000", nil)

	// Cria um gravador de resposta
	resp := httptest.NewRecorder()

	// Executa a requisição simulada
	r.ServeHTTP(resp, req)

	// Valida o resultado
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "logradouro")
}
