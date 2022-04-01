package nba

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

const DATA_NBA_NET = "http://data.nba.net/10s/prod/"

func httpRequest(endpoint string, structure interface{}) error {
	client := &http.Client{
		Timeout: 20 * time.Second,
	}

	log.Println(DATA_NBA_NET + endpoint)
	response, err := client.Get(DATA_NBA_NET + endpoint)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	json.NewDecoder(response.Body).Decode(structure)
	return nil
}
