package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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

func main() {
	r := gin.Default()
	r.GET("/cep/:cep", handleCEP)
	log.Println("Servidor iniciado na porta 8080")
	r.Run(":8080")
}

// Handler da rota /cep/:cep
func handleCEP(c *gin.Context) {
	cep := c.Param("cep")
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
}
