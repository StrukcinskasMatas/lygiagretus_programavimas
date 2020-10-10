package movie

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Movie struct {
	Name   string  `json:"name"`
	Year   int     `json:"year"`
	Rating float64 `json:"rating"`
}

func (m Movie) CalculateMovieHash() string {
	return string(sha1.New().Sum([]byte(fmt.Sprintf("%s %d %f", m.Name, m.Year, m.Rating))))
}

func ReadMoviesJsonData(fileName string) ([]Movie, error) {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	movies := make([]Movie, 0)
	err = json.Unmarshal(jsonData, &movies)
	if err != nil {
		return nil, err
	}

	return movies, nil
}

func WriteResultsToFile(results []Movie, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer file.Close()

	file.WriteString(fmt.Sprintf("%50s %8s %8s\n", "Name", "Year", "Rating"))
	for _, r := range results {
		file.WriteString(fmt.Sprintf("%50s %8d %8.2f\n", r.Name, r.Year, r.Rating))
	}

	return nil
}
