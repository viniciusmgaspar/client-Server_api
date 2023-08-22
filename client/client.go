package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Quotation struct {
	Dolar string
}

func Execute() {
	fmt.Println("Client Initiated!")
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, error := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if error != nil {
		panic(error)
	}

	fmt.Println("Request to server made!")
	res, error := http.DefaultClient.Do(req)
	if error != nil {
		fmt.Println("Erro to get quotation in client: ", error)
		return
	}

	fmt.Println("Response received!")
	body, error := io.ReadAll(res.Body)
	if error != nil {
		fmt.Println("Erro to read response body: ", error)
		return
	}

	fmt.Println("Response body readed!")

	var data Quotation
	error = json.Unmarshal(body, &data)
	if error != nil {
		fmt.Println("Error parsing JSON", error)
		return
	}

	defer res.Body.Close()
	f, error := os.Create("cotacao.txt")
	if error != nil {
		fmt.Println("Erro to create quotation archive: ", error)
		return
	}
	defer f.Close()
	f.WriteString(fmt.Sprintf("Dólar: %s", data.Dolar))
}
