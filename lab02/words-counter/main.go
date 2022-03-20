package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(wordCount("Lubię bardzo placki"))
	fmt.Println(wordsCountMap("Lubię lubię placki placki a!"))
}

/*
   Napisz funkcję która liczy ilość wyrazów w zdaniu
   (nie musisz usuwać żadnych znaków interpunkcyjnych)
*/
func wordCount(s string) int {
	s = strings.ToLower(s)
	r, _ := regexp.Compile(`[\p{L}-]+`)
	words := r.FindAllString(s, -1)

	return len(words)
}

/* trudniejszy wariant*/

/*
Teraz stwórz funckję, która liczy wystąpienia każdego wyrazu w zdaniu (stringu)
i zwraca je w formie mapy (słownika)
np.:
fmt.Println(wordsCountMap(""Lubię lubię placki placki a!""))
map[a:1 lubię:2 placki:2]
*/
func wordsCountMap(s string) map[string]int {
	/*
	   Teraz stwórz funckję, która liczy wystąpienia każdego wyrazu w zdaniu (stringu)
	   i zwraca je w formie mapy (słownika)
	   np.:
	   fmt.Println(wordsCountMap(""Lubię lubię placki placki a!""))
	   map[a:1 lubię:2 placki:2]
	*/
	s = strings.ToLower(s)
	r, _ := regexp.Compile(`[\p{L}-]+`)
	words := r.FindAllString(s, -1)

	wordsMap := make(map[string]int)

	for _, word := range words {
		wordsMap[word]++
	}

	return wordsMap

}
