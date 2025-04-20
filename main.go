package main

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

func main() {
	ctx, timeoutCtx := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer timeoutCtx()

	req, err := http.NewRequestWithContext(timeoutCtx, http.MethodGet, "http://localhost:8080/cotacao", nil)
	if err != nil {
		panic("Failed to access cotacao")
	}

	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			panic("Timeout")
		}
		panic("Failed to access cotacao")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic("Failed to read response body")
	}
	defer res.Body.Close()

	var exchange exchangeTo.exchangeTo
	json.Unmarshal(body, &exchange)
	saveExchange(exchange.USDBRL.Bid)
}
