package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Estrutura que representa o retorno da API ViaCEP
type CepResponse struct {
	CEP         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	Erro        string `json:"erro,omitempty"` // ViaCEP retorna "true" como string
}

var cache = make(map[string]CepResponse)

func main() {
	r := gin.Default()
	r.Use(RequestLogger())
	// Rota para consultar CEP
	r.GET("/cep/:cep", handleCEP)
	log.Println("Servidor iniciado na porta 8080")
	r.Run(":8080")
}

// Valida se o CEP possui apenas números e 8 dígitos
func IsValidCEP(cep string) bool {
	if len(cep) != 8 {
		return false
	}
	for _, c := range cep {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

// Handler da rota /cep/:cep
func handleCEP(c *gin.Context) {
	if !IsValidCEP(c.Param("cep")) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "CEP inválido. Deve conter apenas números e ter 8 dígitos."})
		return
	}

	cep := c.Param("cep")

	// Verifica se o CEP está no cache e retorna em val o cepresponse
	if val, ok := cache[cep]; ok {
		fmt.Println("Cache hit:", cep)
		c.JSON(http.StatusOK, val)
		return
	}

	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	log.Printf("Consultando API ViaCEP: %s", url)

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Erro na requisição HTTP: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao consultar API ViaCEP"})
		return
	}
	defer resp.Body.Close()

	log.Printf("Status da resposta HTTP: %d", resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Falha ao obter resposta da API ViaCEP"})
		return
	}

	var cepResponse CepResponse
	if err := json.NewDecoder(resp.Body).Decode(&cepResponse); err != nil {
		log.Printf("Erro ao decodificar resposta: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao decodificar resposta: %v", err)})
		return
	}

	log.Printf("Resposta da API ViaCEP: %+v", cepResponse)

	// Verifica se o CEP não existe
	if cepResponse.Erro == "true" {
		c.JSON(http.StatusNotFound, gin.H{"error": "CEP não encontrado"})
		return
	}

	c.JSON(http.StatusOK, cepResponse)
	cache[cep] = cepResponse // Armazena no cache
}

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Antes de processar a requisição
		log.Printf("Início da requisição: %s %s", c.Request.Method, c.Request.URL.Path)
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		log.Printf("Duração da requisição: %v", duration)

		// Após processar a requisição
		log.Printf("Fim da requisição: %s %s - Status: %d", c.Request.Method, c.Request.URL.Path, c.Writer.Status())
	}
}
