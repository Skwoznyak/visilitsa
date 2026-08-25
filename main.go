package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"unicode/utf8"
)

type Game struct {
	word       string
	hiddenWord string
	counter    int
	alphabet   map[string]bool
	scanner    *bufio.Scanner
}

func main() {
	game := NewGame()

	if !game.Agreement() {
		return
	}

	game.Run()
}

// Аналог конструктора
func NewGame() *Game {
	word := getWord()

	return &Game{
		word:       word,
		hiddenWord: strings.Repeat("-", utf8.RuneCountInString(word)),
		counter:    0,
		alphabet:   createAlphabet(),
		scanner:    bufio.NewScanner(os.Stdin),
	}
}

func (g *Game) Run() {
	fmt.Println("Длина слова:", utf8.RuneCountInString(g.word))
	fmt.Println(g.hiddenWord)

	for {
		if g.counter >= 6 {
			fmt.Println("Проигрыш")
			return
		}

		if g.word == g.hiddenWord {
			fmt.Println("Вы отгадали слово! Победа!", g.hiddenWord)
			return
		}

		g.MakeTurn()
	}
}

func (g *Game) MakeTurn() {
	g.PrintAlphabet()

	letter := g.AskLetter()

	if g.Guess(letter) {
		g.UpdateHiddenWord(letter)
	}

	g.alphabet[letter] = false

	fmt.Println("Слово:", g.hiddenWord)
	fmt.Println("Ошибок:", g.counter)
}

func (g *Game) AskLetter() string {
	for {
		fmt.Print("Введите одну букву: ")

		if !g.scanner.Scan() {
			fmt.Println("Ошибка ввода")
			continue
		}

		letter := strings.ToLower(g.scanner.Text())

		if utf8.RuneCountInString(letter) != 1 {
			fmt.Println("Ошибка: можно ввести только одну букву!")
			continue
		}

		return letter
	}
}

func (g *Game) Guess(letter string) bool {
	if strings.Contains(g.word, letter) {
		fmt.Println("Да, такая буква есть:", letter)
		return true
	}

	fmt.Println("Такой буквы нет:", letter)
	g.counter++

	return false
}

func (g *Game) UpdateHiddenWord(letter string) {
	wordRunes := []rune(g.word)
	hiddenRunes := []rune(g.hiddenWord)
	letterRune := []rune(letter)[0]

	for i, r := range wordRunes {
		if r == letterRune {
			hiddenRunes[i] = letterRune
		}
	}

	g.hiddenWord = string(hiddenRunes)
}

func (g *Game) PrintAlphabet() {
	fmt.Println("Доступные буквы:")

	for letter, available := range g.alphabet {
		if available {
			fmt.Print(letter, " ")
		}
	}

	fmt.Println()
}

func (g *Game) Agreement() bool {
	fmt.Println("Вы будете играть в виселицу? y/n")

	if !g.scanner.Scan() {
		fmt.Println("Ошибка ввода")
		return false
	}

	switch g.scanner.Text() {
	case "y":
		fmt.Println("Поехали")
		return true
	default:
		fmt.Println("Ну ок")
		return false
	}
}

func createAlphabet() map[string]bool {
	alphabet := make(map[string]bool)

	letters := "абвгдеёжзийклмнопрстуфхцчшщъыьэюя"

	for _, letter := range letters {
		alphabet[string(letter)] = true
	}

	return alphabet
}

func getWord() string {
	words := []string{
		"пицца",
		"снюс",
		"программист",
		"контейнер",
		"кубернетес",
	}

	return words[rand.Intn(len(words))]
}