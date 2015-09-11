package main

import (
	"fmt"
)

func Season (month int) string {
	var season string
	switch {
	case month < 3:
		season = "Winter"
	case month >= 3 && month < 6:
		season = "Spring"
	case month >= 6 && month < 9:
		season = "Summer"
	case month >= 9 && month < 12:
		season = "Fall"
	}
	return season
}

func main(){
	
	fmt.Printf("it is this season: %s\n", Season(8))
}

