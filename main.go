package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

func main() {
	time_now := time.Now()
	today := time_now.Format("2006-01-02")

	file, err := os.Open("data/kosovo-prayer-times.csv")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	for _, record := range records {
		if record[0] == today {
			fmt.Printf("| %-8s | %-8s | %-17s | %-8s | %-8s | %-8s | %-8s | %-16s |\n",
				"Imsaku",
				"Sabahu",
				"Lindja e Diellit",
				"Dreka",
				"Ikindia",
				"Akshami",
				"Jacia",
				"Gjatesia e Dites",
			)

			fmt.Printf("| %-8v | %-8v | %-17v | %-8v | %-8v | %-8v | %-8v | %-16v |\n",
				record[2],
				record[3],
				record[4],
				record[5],
				record[6],
				record[7],
				record[8],
				record[9],
			)
		}
	}
}
