// Напишите программу, подсчитывающую сколько раз буква встречается в предложении, а также частоту встречаемости в процентах.

package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	fmt.Println("Введите предложение")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	myMap := make(map[rune]int, 100)
	lenText := 0

	for _,v := range text {
		if unicode.IsLetter(v) {
			lenText ++
			myMap[unicode.ToLower(v)]++
		}	
	}

	for k,v := range myMap {
		percLetter := float64(v) / float64(lenText)
		fmt.Printf("%c - %d %0.1f \n", k, v, percLetter)
	}
}