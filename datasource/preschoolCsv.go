package datasource

import (
	"encoding/csv"
	"log"
	"os"
)

type PreschoolCsv struct{}

func (p *PreschoolCsv) LoadFromCsv() ([][]string, error) {
	file, err := os.Open("datasource/112038_preschool.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	rows, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return rows, nil
}
