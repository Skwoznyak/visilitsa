package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {

	counter := 0
	agreement()
	word := *getWord()
	fmt.Println(word)
	num := getNumLetters(word)

	fmt.Println("Длина слова:", num)

	myHiddenWord := createHiddenWord(num)

	fmt.Println(myHiddenWord)

	alph := createAlph()

	fmt.Println(alph)

	gameLogick(word, &counter, myHiddenWord)

	// ans := gessing(letter, word, &counter)

	// if ans {
	// 	posIdx := getLetterIndices(word, letter)
	// 	myHiddenWord = *updateHiddenWord(posIdx, &myHiddenWord, letter)
	// 	fmt.Println(myHiddenWord)
	// }

	// getCounter(&counter)

}


// func askLetter() string {
// 	fmt.Println("Загадайте букву")

// 	scanner := bufio.NewScanner((os.Stdin))

// 	ok := scanner.Scan()
// 	if !ok {
// 		fmt.Println("ошибка ввода")
// 		return "-1"
// 	}

// 	letter := scanner.Text()
// 	ok = validLetter(letter); if !ok {return "-1"}

// 	return letter
// }

func askLetter() string {

	for {
		scanner := bufio.NewScanner((os.Stdin))

		fmt.Print("Введите одну букву: ")

		if !scanner.Scan() {
			fmt.Println("Ошибка ввода, попробуйте еще раз.")
			continue 
		}

		letter := scanner.Text()

		if validLetter(letter) {
			return letter
		}
		
	}
}

func validLetter(letter string) bool {
	if utf8.RuneCountInString(letter) == 1 {
		return true
	} 
	
	fmt.Println("Ошибка: Можно ввести только одну букву!")
	return false
}


func gameLogick(word string, counter *int, myHiddenWord string) {

	letter := askLetter()
	if letter == "-1"{
		askLetter()
	}

	ans := gessing(letter, word, counter)

	if ans {
		posIdx := getLetterIndices(word, letter)
		myHiddenWord = *updateHiddenWord(posIdx, &myHiddenWord, letter)
		fmt.Println(myHiddenWord)
	}

	getCounter(counter)

}

func agreement() {
	fmt.Println("Вы будете играть в виселицу? y/n")

	scanner := bufio.NewScanner((os.Stdin))

	ok := scanner.Scan()
	if !ok {
		fmt.Println("ошибка ввода")
		return
	}

	ans := scanner.Text()
	if ans == "y" {
		fmt.Println("поехали")
		return
	} else {
		fmt.Println("ну ок")
	}

}

func getWordList() *[]string {

	words := []string{"пицца", "снюс"}

	return &words
}

func getWord() *string {

	words := *getWordList()

	size := len(words)

	randomWord := rand.Intn(size)

	return &words[randomWord]
}

func getNumLetters(w string) int {
	realLength := utf8.RuneCountInString(w)

	return realLength
}

func createHiddenWord(n int) string {
	myHiddenWord := strings.Repeat("-", n)

	return myHiddenWord
}

func updateHiddenWord(posIdx []int, myHiddenWord *string, letter string) *string {
	fmt.Println(posIdx)

	runes := []rune(*myHiddenWord)

	// Берем именно руну (символ Unicode), а не байт
	r := []rune(letter)[0]

	for i := range posIdx {
		runes[posIdx[i]] = r
	}

	str := string(runes)

	return &str
}

func getAlph(alph *map[string]bool) {

}

func createAlph() *map[string]bool {

	alphabetMap := make(map[string]bool)

	letters := "абвгдеёжзийклмнопрстуфхцчшщъыьэюя"
	for _, char := range letters {
		alphabetMap[string(char)] = true
	}

	return &alphabetMap
}

func gessing(letter string, word string, c *int) bool {

	success := false

	if strings.Contains(word, letter) {
		fmt.Println("Да, такая буква есть: ", letter)
		success = true
	} else {
		fmt.Println("Такой буквы нет: ", letter)
		*c += 1

	}

	return success
}

func getCounter(c *int) {

	fmt.Println("Ошибок:", *c)
}

func getLetterIndices(word string, letter string) []int {
	runes := []rune(word)

	letterRune := []rune(letter)[0]

	var indices []int

	for i, r := range runes {
		if r == letterRune {
			indices = append(indices, i)
		}
	}

	return indices
}
