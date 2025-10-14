package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ExchangeResponse struct {
	Rates map[string]float64 `json:"rates"`
}

func main() {
	var from, to string
	var amount float64

	fmt.Println("Conversor de moeda")
	fmt.Print("Moeda de origem (ex: USD): ")
	fmt.Scan(&from)
	fmt.Print("Moeda de destino (ex: BRL): ")
	fmt.Scan(&to)
	fmt.Print("Digite o valor a ser convertido: ")
	fmt.Scan(&amount)

	//Monta a URL da API
	url := fmt.Sprintf("https://api.exchangerate-api.com/v4/latest/%s", from)
	fmt.Printf("URL: %v\n", url)

	//Faz a requisição HTTP
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Erro ao fazer a requisição: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Erro na resposta da API: %v", resp.Status)
	}
	/* 	fmt.Println("Body:")
	   	body, err := io.ReadAll(resp.Body)
	   	if err != nil {
	   		log.Fatalf("Erro ao ler o corpo da resposta: %v", err)
	   	}
	   	fmt.Println(string(body)) */

	var exchangeResponse ExchangeResponse
	if err := json.NewDecoder(resp.Body).Decode(&exchangeResponse); err != nil {
		log.Fatalf("Erro ao decodificar a resposta: %v", err)
	}

	rate := exchangeResponse.Rates[to]
	converted := amount * rate
	fmt.Printf("%.2f %s equivalem a %.2f %s\n", amount, from, converted, to)
}
