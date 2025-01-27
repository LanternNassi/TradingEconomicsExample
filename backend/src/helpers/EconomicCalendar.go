package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	_ "log"
	"net/http"
	"os"
	"osprey/src/models"
)

func FetchEventsByCountry(country string) ([]models.Event, error) {
	apiKey := os.Getenv("TE_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("missing API key")
	}
	url := fmt.Sprintf("https://api.tradingeconomics.com/calendar/country/%s?c=%s&f=json", country, "guest:guest")

	fmt.Println(url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("received non-2xx response: %d - %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var events []models.Event
	if err := json.Unmarshal(body, &events); err != nil {
		return nil, fmt.Errorf("failed to parse response JSON: %w", err)
	}

	return events, nil
}
