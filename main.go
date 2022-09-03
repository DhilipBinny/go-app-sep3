package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/DhilipBinny/go-app-sep3/api"
	"github.com/fatih/color"
)

// board_game_atlas
// --query "ticket to time" --limit 5 --skip 1 --cid aqkbqPB6Gj
// https://api.boardgameatlas.com/api/search?name=Catan&client_id=<client_id_from>&limit=5&skip=5

func main() {

	fmt.Println("BGA")

	query := flag.String("query", "", "Name of a board game to search for.")
	cid := flag.String("cid", "", "ClientID")
	skip := flag.Uint("skip", 0, "Skips the number of results provided, It's generally used for paging results")
	limit := flag.Uint("limit", 0, "Limits the number of results returned")
	timeout := flag.Uint("timeout", 10, "Timeout")

	flag.Parse()

	fmt.Printf("name=%s & cid=%s & skip=%d & limit %d \n", *query, *cid, *skip, *limit)

	if isNull(*query) {
		log.Fatalln("Please use -query to set the boardgame name to seach")
	}
	if isNull(*cid) {
		log.Fatalln("Please use -cid to set the boardgame name to seach")
	}

	// Instantiate the BGAtlas struct
	bga := api.NewBGA((*cid))

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(*timeout*uint(time.Second)))

	defer cancel()

	// time.Sleep(3 * time.Second)

	// Make the invocation
	result, err := bga.Search(ctx, *query, *limit, *skip)

	if err != nil {
		log.Fatalf("ERROR: cannot search boardgame : %v", err)
	}

	// Color
	boldgreen := color.New(color.Bold).Add(color.FgHiGreen).SprintFunc()

	// for _, game := range result.Games {
	// 	fmt.Printf("Name: %s\n", game.Name)
	// 	fmt.Printf("%s: %s\n", boldgreen("Name"), game.Name)
	// 	fmt.Printf("Url: %s\n", game.Url)
	// 	fmt.Printf("Price: %s\n", game.Price)
	// 	fmt.Printf("Year: %d\n", game.YearPublished)
	// 	fmt.Printf("Description: %s\n\n", game.Description)
	// }

	for _, game := range result.Games {
		fmt.Printf("%s: %s\n", boldgreen("Name"), game.Name)
		fmt.Printf("%s: %s\n", boldgreen("Url"), game.Url)
		fmt.Printf("%s: %s\n", boldgreen("Price"), game.Price)
		fmt.Printf("%s: %d\n", boldgreen("YaearPublished"), game.YearPublished)
		fmt.Printf("%s: %s\n\n", boldgreen("Description"), game.Description)
	}

}

func isNull(s string) bool {
	return len(strings.TrimSpace(s)) <= 0
}

// // build
// go build -o bga .

// // cross-sompile
// GOOS=darwin GOARCH=arm64 go build -o bga-darwin-arm64 .

// // test
// bga-darwin-arm64 --query "Catan" --limit 5 --skip 1 --cid <client_id_from> --timeout=5
