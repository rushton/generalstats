package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/montanaflynn/stats"
)

var groupBy *int = flag.Int("group-by", 0, "Groups stats by the provided column index, starting at 1.")

func main() {
	in := bufio.NewReader(os.Stdin)
	dataMap := make(map[int]stats.Float64Data)
	for {
		line, err := in.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("Error reading line: %v", err)
		}
		fields := strings.Fields(strings.TrimSuffix(line, "\n"))

		for i := 0; i < len(fields); i++ {
			datum, err := strconv.ParseFloat(fields[i], 64)
			if err != nil {
				log.Fatalf("error parsing data: %v", err)
			}
			dataMap[i] = append(dataMap[i], datum)
		}
	}

	for _, data := range dataMap {
		count := len(data)

		min, err := data.Min()
		if err != nil {
			log.Fatalf("Error calculating min: %v", err)
		}
		max, err := data.Max()
		if err != nil {
			log.Fatalf("Error calculating max: %v", err)
		}
		mean, err := data.Mean()
		if err != nil {
			log.Fatalf("Error calculating mean: %v", err)
		}
		median, err := data.Median()
		if err != nil {
			log.Fatalf("Error calculating median: %v", err)
		}
		tp75, err := data.Percentile(75.0)
		if err != nil {
			log.Fatalf("Error calculating tp75: %v", err)
		}
		tp95, err := data.Percentile(95)
		if err != nil {
			log.Fatalf("Error calculating tp95: %v", err)
		}
		tp99, err := data.Percentile(99)
		if err != nil {
			log.Fatalf("Error calculating tp99: %v", err)
		}
		tp999, err := data.Percentile(99.9)
		if err != nil {
			log.Fatalf("Error calculating tp999: %v", err)
		}

		fmt.Printf(`count: %d
min: %.2f
max: %.2f
mean: %.2f
median: %.2f
tp75: %.2f
tp95: %.2f
tp99: %.2f
tp999: %.2f
`, count, min, max, mean, median, tp75, tp95, tp99, tp999)
	}
}
