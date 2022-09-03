package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const SEARCH_URL = "https://api.boardgameatlas.com/api/search"

type BGAtlas struct {
	// private
	clientId string
	//public
	// Name string
}

// GAME
type Game struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Price         string `json:"price"`
	YearPublished uint   `json:"year_published"`
	Description   string `json:"description"`
	ImageUrl      string `json:"image_url"`
	RulesUrl      string `json:"rules_url"`
	Url           string `json:"official_url"`
}

type SerachResult struct {
	Games []Game `json:"games"`
	Count uint   `json:"count"`
}

// func NewBGA(cid string) BGAtlas {
// 	return BGAtlas{clientId: cid}
// }

// Functions as a consctructor to initialize
func NewBGA(clientId string) BGAtlas {
	return BGAtlas{clientId}
}

// Method
func (b BGAtlas) Search(ctx context.Context, query string, limit uint, skip uint) (*SerachResult, error) {
	// create a http client
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, SEARCH_URL, nil)

	//check if there is an error
	if err != nil {
		// return an error obj
		return nil, fmt.Errorf("ERROR: Cannot create http client: %v", err)
	}

	// get the query string obj
	params := req.URL.Query()

	// Populate the url with query params
	params.Add("name", query)
	params.Add("limit", fmt.Sprintf("%d", limit))
	// params.Add("skip", fmt.Sprintf("%d", skip))
	params.Add("skip", strconv.Itoa(int(skip)))
	params.Add("client_id", b.clientId)

	// Encode the query params and add it back to the request
	req.URL.RawQuery = params.Encode()

	fmt.Printf("Url = %s\n", req.URL.String())

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, fmt.Errorf("ERROR: http GET invocation error : %v", err)
	}

	// HTTP status code >=400 is error
	if res.StatusCode >= 400 {
		return nil, fmt.Errorf("ERROR: HTTP status : %s", res.Status)
	}

	var result SerachResult
	// deserialize the JSON payload to struct
	if err := json.NewDecoder(res.Body).Decode(&result); nil != err { // note & inside Decode - address of result
		return nil, fmt.Errorf("ERROR: cannot deserialize JSON payload : %v", err)
	}

	return &result, nil
}
