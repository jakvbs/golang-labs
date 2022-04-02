package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	adjectivesStr, _ := ioutil.ReadFile("adjectives.txt")
	adjectivesArr := strings.Split(string(adjectivesStr), "\n")
	nounsStr, _ := ioutil.ReadFile("nouns.txt")
	nounsArr := strings.Split(string(nounsStr), "\n")
	termsStr, _ := ioutil.ReadFile("terms.txt")
	termsArr := strings.Split(string(termsStr), "\n")
	nouns := make(map[string]string)
	terms := make(map[string]string)
	for i := 0; i < 26; i++ {
		nouns[alphabet[i]] = nounsArr[i]
		terms[alphabet[i]] = termsArr[i]
	}

	var day int
	var name, sureName string
	flag.IntVar(&day, "day", 0, "dzień urodzenia")
	flag.StringVar(&name, "name", "", "imie")
	flag.StringVar(&sureName, "surename", "", "naziwsko")
	flag.Parse()

	if day == 0 {
		fmt.Print("Podaj dzień urodzenia: ")
		fmt.Scanf("%d", &day)
	}
	if name == "" {
		fmt.Print("Podaj imie: ")
		fmt.Scanf("%s", &name)
	}
	if sureName == "" {
		fmt.Print("Podaj nazwisko: ")
		fmt.Scanf("%s", &sureName)
	}

	name = strings.ToLower(name)
	sureName = strings.ToLower(sureName)

	if day < 1 || day > 31 {
		fmt.Println("Nieprawidłowy dzień urodzenia")
		return
	}
	adjective := adjectivesArr[day-1]
	noun, found := nouns[name[0:1]]
	if !found {
		fmt.Println("Pierwsza litera imienia nie jest w alfabecie")
		return
	}
	term, found := terms[sureName[0:1]]
	if !found {
		fmt.Println("Pierwsza litera nazwiska nie jest w alfabecie")
		return
	}

	println("Twój kryptonim to:", adjective, noun, term)
}
