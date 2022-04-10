package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"time"
)

type User struct {
	name   string
	result int
	number int
	date   time.Time
}

func readCsvFile(fileName string) [][]string {
	_, currentFilePath, _, _ := runtime.Caller(0)
	filePath := filepath.Join(filepath.Dir(currentFilePath), fileName)
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+fileName, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+fileName, err)
	}

	return records
}

func mapResults(records [][]string) []User {
	results := make([]User, 0)
	for i := 1; i < len(records); i++ {
		result, _ := strconv.Atoi(records[i][1])
		number, _ := strconv.Atoi(records[i][2])
		date, _ := time.Parse("2006-01-02 15:04:05", records[i][3])
		user := User{records[i][0], result, number, date}
		results = append(results, user)
	}
	return results
}

func writeToCsvFile(fileName string, results []User) {
	_, currentFilePath, _, _ := runtime.Caller(0)
	filePath := filepath.Join(filepath.Dir(currentFilePath), fileName)
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal("Unable to open file "+fileName, err)
	}
	defer f.Close()
	f.Truncate(0)

	csvWriter := csv.NewWriter(f)
	defer csvWriter.Flush()

	headers := []string{"name", "result", "number", "date"}
	err = csvWriter.Write(headers)

	if err != nil {
		log.Fatal("Unable to write headers to file "+fileName, err)
	}

	for _, user := range results {
		record := []string{user.name, strconv.Itoa(user.result), strconv.Itoa(user.number), user.date.Format("2006-01-02 15:04:05")}
		err := csvWriter.Write(record)

		if err != nil {
			log.Fatal("Unable to write to file "+fileName, err)
		}
	}
}

func processResults(results []User, user User) []User {
	results = append(results, user)
	switch {
	case len(results) == 1:
		fmt.Println("Jesteś pierwszym graczem na liście!")
	case user.result < results[0].result:
		fmt.Printf("Gratulacje %s! Pobiłeś poprzedni rekord! Aktualny rekord to %d\n", user.name, user.result)
	case user.result == results[0].result:
		fmt.Printf("Gratulacje %s! Wyrównałeś się z poprzednim rekordem! Aktualny rekord to %d\n", user.name, user.result)
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].result < results[j].result
	})

	return results
}

func printResults(results []User) {
	fmt.Println("Hall of Fame:")
	for i, user := range results {
		if i > 2 {
			break
		}
		fmt.Printf("%d. %s: %d, %s\n", i+1, user.name, user.result, user.date.Format("2006-01-02 15:04:05"))
	}
}

func generateRandomNumber(start, end int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(end-start) + start
}

func replayAnswer() bool {
	fmt.Print("Czy gramy jeszcze raz? [T/N]: ")
	var answer string
	fmt.Scanf("%s", &answer)

	var replay bool
	switch answer {
	case "T", "t", "Y", "y":
		replay = true
	case "N", "n":
		replay = false
	default:
		fmt.Println("Kontynuujemy!")
		replay = true
	}
	return replay
}

func game() {
	records := readCsvFile("results.csv")
	results := mapResults(records)

	var guess string
	var replay bool = true

OuterLoop:
	for replay {
		random := generateRandomNumber(1, 100)
		fmt.Printf("Wylosowana liczba to: %d\n", random)
		fmt.Println("Aby zakończyć działanie programu wpisz 'koniec'.")
		fmt.Println("Teraz będziesz zgadywać liczbę, którą wylosowałem z zakresu 1-100: ")
		counter := 0

	InnerLoop:
		for {
			fmt.Print("Podaj liczbę: ")
			fmt.Scanf("%s", &guess)

			guessNumber, err := strconv.Atoi(guess)
			if err != nil {
				if guess == "koniec" {
					fmt.Println("Żegnaj!")
					break OuterLoop
				}
				fmt.Printf("%s nie jest liczbą!\n", guess)
				continue
			}

			counter++
			switch {
			case guessNumber == random:
				fmt.Printf("Brawo! Zgadłeś! Moja liczba to %d\n", guessNumber)

				var name string
				fmt.Print("Podaj swoje imię: ")
				fmt.Scanf("%s", &name)

				user := User{name, counter, random, time.Now()}
				results = processResults(results, user)
				break InnerLoop
			case guessNumber > random:
				fmt.Println("Za duża")
			case guessNumber < random:
				fmt.Println("Za mała")
			}
		}
		replay = replayAnswer()
	}

	printResults(results)
	writeToCsvFile("results.csv", results)
}

func main() {
	game()
}
