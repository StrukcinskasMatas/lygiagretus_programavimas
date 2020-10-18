package data_thread

import (
	"github.com/StrukcinskasMatas/lygiagretus_programavimas/L2/internal/config"
	"github.com/StrukcinskasMatas/lygiagretus_programavimas/L2/internal/movie"
)

type DataPayload struct {
	Movie    *movie.Movie
	Finished bool
}

func Start(dataInputChannel <-chan *movie.Movie, requestsChannel <-chan int, dataOutputChannel chan<- *movie.Movie) {
	var dataStorage []movie.Movie
	var inputFinished bool
	var outputFinished int

	for outputFinished != config.WorkerThreadCount {
		if len(dataStorage) == 0 && !inputFinished {
			movie := <-dataInputChannel
			if movie != nil {
				dataStorage = append(dataStorage, *movie)
			} else {
				inputFinished = true
			}

			continue
		}

		if len(dataStorage) == config.AllowedDataCount {
			<-requestsChannel
			if !inputFinished || len(dataStorage) != 0 {
				dataOutputChannel <- &dataStorage[0]
				dataStorage = dataStorage[1:]
			} else {
				dataOutputChannel <- nil
				outputFinished++
			}

			continue
		}

		select {
		case movie := <-dataInputChannel:
			if movie != nil {
				dataStorage = append(dataStorage, *movie)
			} else {
				inputFinished = true
			}

		case <-requestsChannel:
			if !inputFinished || len(dataStorage) != 0 {
				dataOutputChannel <- &dataStorage[0]
				dataStorage = dataStorage[1:]
			} else {
				dataOutputChannel <- nil
				outputFinished++
			}
		}
	}
}
