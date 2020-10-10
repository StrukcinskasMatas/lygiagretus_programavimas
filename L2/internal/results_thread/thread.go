package results_thread

import (
	"github.com/StrukcinskasMatas/lygiagretus_programavimas/L2/internal/config"
	"github.com/StrukcinskasMatas/lygiagretus_programavimas/L2/internal/movie"
)

var resultsStorage []movie.Movie
var workersFinished int

func Start(workerChannel <-chan *movie.Movie, mainChannel chan<- []movie.Movie) {
	for workersFinished != config.WorkerThreadCount {
		result := <-workerChannel
		if result != nil {
			resultsStorage = append(resultsStorage, *result)
		} else {
			workersFinished++
		}
	}

	mainChannel <- resultsStorage
}
