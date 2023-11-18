package main

import (
	"fmt"
	"strings"
)

type Question struct {
	Question      string
	Options       []string
	CorrectAnswer string
}

func main() {
	var playerName string
	fmt.Println("Selamat datang di Quiz!")
	fmt.Print("Masukkan nama Anda: ")
	fmt.Scanln(&playerName)

	questions := []Question{
		{
			Question:      "Siapakah ketua desa di serial animasi Upin&Ipin?",
			Options:       []string{"0. Saleh", "1. Upin&Ipin", "2. Tok Dalang", "3. Kak Ros"},
			CorrectAnswer: "2",
		},
		{
			Question:      "Siapakah member JKT 48?",
			Options:       []string{"0. Freya", "1. Jokowi", "2. Iqbal", "3. Igun"},
			CorrectAnswer: "0",
		},
		{
			Question:      "Apa ibukota Bali?",
			Options:       []string{"0. Buleleng", "1. Denpasar", "2. Gianyar", "3. Tabanan"},
			CorrectAnswer: "1",
		},
		{
			Question:      "Berapakah hasil dari 1 + 1?",
			Options:       []string{"0. 11", "1. 10", "2. 2", "3. 4"},
			CorrectAnswer: "2",
		},
		{
			Question:      "Berapakah hasil dari 1 + 1 + 1?",
			Options:       []string{"0. 11", "1. 10", "2. 3", "3. 4"},
			CorrectAnswer: "0",
		},
	}

	playerName = strings.TrimSpace(playerName)

	runQuiz(questions, playerName)
}

func runQuiz(questions []Question, playerName string) {
	var score int
	var salah int

	for _, q := range questions {
		fmt.Println(q.Question)
		for _, option := range q.Options {
			fmt.Println(option)
		}

		var userAnswer string
		fmt.Print("Masukkan pilihan jawaban (0, 1, 2, dan 3): ")
		fmt.Scanln(&userAnswer)

		userAnswer = strings.TrimSpace(strings.ToUpper(userAnswer))

		if userAnswer == q.CorrectAnswer {
			fmt.Println("Benar!")
			score++
		} else {
			fmt.Println("Salah!")
			salah++
		}
	}

	totalQuestions := len(questions)
	fmt.Printf("Statistic Kuis\n Nama : %s\n Score : %d\n Jawaban Benar : %d\n Jawaban Salah : %d\n Total Pertanyaan : %d\n", playerName, score, score, salah, totalQuestions)

}
