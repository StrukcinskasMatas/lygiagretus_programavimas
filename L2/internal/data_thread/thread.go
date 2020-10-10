package data_thread

import (
	"github.com/StrukcinskasMatas/lygiagretus_programavimas/L2/internal/config"
	"github.com/StrukcinskasMatas/lygiagretus_programavimas/L2/internal/movie"
)

type DataPayload struct {
	Movie    *movie.Movie
	Finished bool
}

var dataStorage []movie.Movie
var inputFinished bool
var outputFinished int

func Start(dataInputChannel <-chan *movie.Movie, requestsChannel <-chan bool, dataOutputChannel chan<- *movie.Movie) {
	for !inputFinished || outputFinished != config.WorkerThreadCount {
		if len(dataStorage) < config.AllowedDataCount {
			select {
			case movie := <-dataInputChannel:
				if movie != nil {
					dataStorage = append(dataStorage, *movie)
				} else {
					inputFinished = true
				}

			default:
				// do nothing
			}
		}

		if len(dataStorage) > 0 || inputFinished {
			select {
			case <-requestsChannel:
				if len(dataStorage) > 0 {
					dataOutputChannel <- &dataStorage[0]
					dataStorage = dataStorage[1:]
				} else {
					dataOutputChannel <- nil
					outputFinished++
				}

			default:
				// do nothing
			}
		}
	}
}
