package server

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Quotation struct {
	USDBRL struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

func saveDB(db *sql.DB, bid string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	_, err := db.ExecContext(ctx, fmt.Sprintf("INSERT INTO %s (currency, value, timestamp) VALUES (?, ?, ?)", "currencies"), "USD-BRL", bid, time.Now())
	if err != nil {
		return fmt.Errorf("error saving to database: %v", err)
	}

	fmt.Println("Quotation saved with successfully!")
	return nil
}

func Server(db *sql.DB) {
	http.HandleFunc("/cotacao", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request received!")
		ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
		defer cancel()

		req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
		if err != nil {
			fmt.Println("Request to API made!222222")
			panic(err)
		}
		fmt.Println("Request to API made!")

		resp, err := http.DefaultClient.Do(req)

		if err != nil {
			http.Error(w, "Error making request", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			http.Error(w, "Server returned non-200 status code", http.StatusInternalServerError)
			return
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Error reading response", http.StatusInternalServerError)
			return
		}
		
		var data Quotation
		err = json.Unmarshal(body, &data)
		if err != nil {
			http.Error(w, "Error unmarshal", http.StatusInternalServerError)
			return
		}
		
		bid := data.USDBRL.Bid
		fmt.Printf("bid:" + bid)
		if err := saveDB(db, bid); err != nil {
			http.Error(w, "Error saving to database", http.StatusInternalServerError)
			return
		}

		result := map[string]string{"dolar": data.USDBRL.Bid}
		response, err := json.Marshal(result)
		if err != nil {
			http.Error(w, "Error creating response JSON", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(response)

	})
}
