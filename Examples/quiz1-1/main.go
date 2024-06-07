package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const defaultProblemsFile = "problems2.csv"

var (
	correctAnswers int
	totalQuestions int
)

func main() {
	var (
		// If you're using the flags themselves, they are all pointers; if you bind to variables, they're values.
		flagProblemsFilename = flag.String("p", defaultProblemsFile, "the filename of the problems csv")
		flagTimer            = flag.Duration("t", 30*time.Second, "time to wait for a quiz answer")
		flagShuffle          = flag.Bool("s", false, "shuffle the quiz questions")
	)
	flag.Parse() //to parse the command line into the defined flags.
	if flagProblemsFilename == nil ||
		flagTimer == nil ||
		flagShuffle == nil {
		fmt.Println("Missing required flags")
		return
	}

	fmt.Printf("Hit enter to start quiz from %q in %v?\n",
		*flagProblemsFilename, *flagTimer)
	fmt.Scanln() // make user to hit enter before starting
	f, err := os.Open(*flagProblemsFilename)
	if err != nil {
		fmt.Printf("Error opening %q: %v", *flagProblemsFilename, err)
		return
	}
	defer f.Close()
	r := csv.NewReader(f)
	questions, err := r.ReadAll()
	totalQuestions = len(questions)
	if *flagShuffle {
		fmt.Println("Shuffle the quiz")
		rand.Shuffle(totalQuestions, func(i, j int) {
			questions[i], questions[j] = questions[j], questions[i]
		})
	}
	if err != nil {
		fmt.Printf("Error reading CSV: %v", err)
		return
	}
	quizDone := startQuiz(questions)
	quizTimer := time.NewTimer(*flagTimer).C // chanel time
	//wait for quiz or timer are done
	select { //like switch but used for chanel
	case <-quizDone:
	case <-quizTimer:
	}

	fmt.Printf("You scored %d out of %d!\n", correctAnswers, totalQuestions)

}

func startQuiz(questions [][]string) chan bool {
	done := make(chan bool)
	go func() {
		for i, record := range questions {
			question, correctAnswer := record[0], record[1]
			fmt.Printf("&d.&s?\n", i+1, question)
			var answer string
			if _, err := fmt.Scanf("%s", &answer); err != nil {
				fmt.Printf("Error reading answer: %v", err)
				return
			}
			answer = strings.TrimSpace(answer)
			answer = strings.ToLower(answer)
			if answer == correctAnswer {
				correctAnswers++
			}
		}
		done <- true // notify main thread that we are running quiz
	}()
	return done
}
