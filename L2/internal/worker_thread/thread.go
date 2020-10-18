package worker_thread

import (
	"time"

	"github.com/StrukcinskasMatas/lygiagretus_programavimas/L2/internal/config"
	"github.com/StrukcinskasMatas/lygiagretus_programavimas/L2/internal/movie"
)

func Start(dataRequestChannel chan<- int, dateReceiveChannel <-chan *movie.Movie, resultsChannel chan<- *movie.Movie, number int) {
	finished := false
	for !finished {
		dataRequestChannel <- number
		response := <-dateReceiveChannel

		if response != nil {
			response.CalculateMovieHash()
			time.Sleep(10 * time.Millisecond) // Simulating more work

			if response.Rating >= config.MinMovieRating {
				resultsChannel <- response
			}
		} else {
			resultsChannel <- response
			finished = true
		}
	}
}
