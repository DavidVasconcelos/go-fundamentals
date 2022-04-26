package main

import "fmt"

func main() {
	leagueTitles := make(map[string]int)
	leagueTitles["Sunderland"] = 6
	leagueTitles["Newcastle"] = 4

	recenthead2HeadWins := map[string]int{
		"Sunderland": 6,
		"Newcastle":  0,
	}

	fmt.Printf("League titles: %v\nRecent Head to heads: %v\n", leagueTitles, recenthead2HeadWins)

	// League titles: map[Newcastle:4 Sunderland:6] -> key order
	// Recent Head to heads: map[Newcastle:0 Sunderland:6]
}
