package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resposta, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, resposta.Body)
	resposta.Body.Close()

	//fmt.Println(resp)
	/* 	bs := make([]byte, 99999)
	   	resp.Body.Read(bs)
	   	fmt.Println(string(bs))
	   	resp.Body.Close() */
}
