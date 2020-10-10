package results_thread

import (
	"sort"

	"github.com/StrukcinskasMatas/lygiagretus_programavimas/L2/internal/config"
	"github.com/StrukcinskasMatas/lygiagretus_programavimas/L2/internal/movie"
)

var resultsStorage []movie.Movie
var workersFinished int

func Start(workerChannel <-chan *movie.Movie, mainChannel chan<- []movie.Movie) {
	for workersFinished != config.WorkerThreadCount {
		result := <-workerChannel
		if result != nil {
			resultsStorage = insertSorted(resultsStorage, *result)
		} else {
			workersFinished++
		}
	}

	mainChannel <- resultsStorage
}

func insertSorted(slice []movie.Movie, element movie.Movie) []movie.Movie {
	i := sort.Search(len(slice), func(i int) bool { return slice[i].Rating < element.Rating })
	slice = append(slice, movie.Movie{Rating: 0})
	copy(slice[i+1:], slice[i:])
	slice[i] = element
	return slice
}
