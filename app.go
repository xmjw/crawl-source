package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"sync"
)

func main() {
	f, _ := os.Open("urls.csv")

	r := csv.NewReader(bufio.NewReader(f))

	var wg sync.WaitGroup

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		wg.Add(1)
		go WriteDomain(record[2], &wg)
	}
	wg.Wait()
}
