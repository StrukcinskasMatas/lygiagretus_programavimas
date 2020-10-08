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
	Rating float32 `json:"rating"`
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

func (m Movie) CalculateMovieHash() string {
	return string(sha1.New().Sum([]byte(fmt.Sprintf("%s %d %f", m.Name, m.Year, m.Rating))))
}
