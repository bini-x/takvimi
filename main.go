package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

var (
	next_prayer_name string
	time_left        time.Duration
)

func main() {
	time_now := time.Now()
	today := time_now.Format("2006-01-02")
	current_time_str := time_now.Format("15:04")
	parsed_current, _ := time.Parse("15:04", current_time_str)

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

	prayers := []struct {
		name  string
		index int
	}{
		{"Sabahu", 3},
		{"Dreka", 5},
		{"Ikindia", 6},
		{"Akshami", 7},
		{"Jacia", 8},
	}

	for _, record := range records {
		if record[0] == today {
			found_next := false

			for _, next_prayer := range prayers {
				prayer_time_str := record[next_prayer.index]
				parsed_prayer, err := time.Parse("15:04", prayer_time_str)
				if err != nil {
					fmt.Println("Error: ", err)
					return
				}

				if parsed_current.Before(parsed_prayer) {
					next_prayer_name = next_prayer.name
					time_left = parsed_prayer.Sub(parsed_current)
					found_next = true
					break
				}
			}

			if !found_next {
				next_prayer_name = "Sabahu (Neser)"

				parsed_sabahu, err := time.Parse("15:04", record[3])
				if err != nil {
					fmt.Println("Error:", err)
					return
				}

				time_until_midnight := 24*time.Hour -
					time.Duration(parsed_current.Hour())*time.Hour -
					time.Duration(parsed_current.Minute())*time.Minute -
					time.Duration(parsed_current.Second())*time.Second

				time_until_sabahu := time.Duration(parsed_sabahu.Hour())*time.Hour +
					time.Duration(parsed_sabahu.Minute())*time.Minute

				time_left = time_until_midnight + time_until_sabahu
			}

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

			fmt.Println("\n---------------------------------------------------------------------------------")
			fmt.Printf("Koha e ardhshme: %s\n", next_prayer_name)
			fmt.Printf("Koha e mbetur: %02d:%02d\n", int(time_left.Hours()), int(time_left.Minutes())%60)
			fmt.Println("---------------------------------------------------------------------------------")
		}
	}
}
