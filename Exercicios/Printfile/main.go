package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	param := os.Args

	for i := 0; i < len(param); i++ {
		fmt.Printf("%d - %s\n", i, param[i])
	}

	file, err := os.Open(param[1]) // For read access.
	if err != nil {
		fmt.Println("Error:", err)
	}
	io.Copy(os.Stdout, file)

	// data := make([]byte, 100)
	// count, err := file.Read(data)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }
	// fmt.Printf("read %d bytes: %q\n", count, data[:count])
}
