package main

import (
	"fmt"
	"time"

	"github.com/StrukcinskasMatas/lygiagretus_programavimas/L2/internal/config"
	"github.com/StrukcinskasMatas/lygiagretus_programavimas/L2/internal/data_thread"
	"github.com/StrukcinskasMatas/lygiagretus_programavimas/L2/internal/movie"
	"github.com/StrukcinskasMatas/lygiagretus_programavimas/L2/internal/results_thread"
	"github.com/StrukcinskasMatas/lygiagretus_programavimas/L2/internal/worker_thread"
)

func main() {
	moviesData, err := movie.ReadMoviesJsonData(config.HalfDataPassFileName)
	if err != nil {
		panic(err)
	}

	mainToDataChannel := make(chan *movie.Movie)
	dataToWorkersChannel := make(chan *movie.Movie)
	workersToDataChannel := make(chan bool)
	workersToResultsChannel := make(chan *movie.Movie)
	resultsToMainChannel := make(chan []movie.Movie)

	go data_thread.Start(mainToDataChannel, workersToDataChannel, dataToWorkersChannel)
	go results_thread.Start(workersToResultsChannel, resultsToMainChannel)

	for i := 1; i <= config.WorkerThreadCount; i++ {
		go worker_thread.Start(workersToDataChannel, dataToWorkersChannel, workersToResultsChannel, i)
	}

	startTime := time.Now()
	for _, movie := range moviesData {
		mainToDataChannel <- &movie
	}

	// signaling that we've finished
	mainToDataChannel <- nil
	fmt.Printf("Results received in %f seconds.\n", time.Since(startTime).Seconds())

	result := <-resultsToMainChannel
	movie.WriteResultsToFile(result, config.ResultsFileName)

	fmt.Println("The application has finished.")
}
