package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
)

func main() {
	f, _ := os.Open("urls.csv")

	r := csv.NewReader(bufio.NewReader(f))

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		WriteDomain(record[2])
	}
}
