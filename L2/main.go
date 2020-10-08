package main

import (
	"fmt"

	"github.com/StrukcinskasMatas/lygiagretus_programavimas/L2/internal/config"
	"github.com/StrukcinskasMatas/lygiagretus_programavimas/L2/internal/movie"
)

func main() {
	moviesData, err := movie.ReadMoviesJsonData(config.AllDataPassFileName)
	if err != nil {
		panic(err)
	}

	fmt.Println(moviesData[0].CalculateMovieHash())
}
