// Create a program that will read in a quiz provided via a CSV file (more
// details below) and will then give the quiz to a user keeping track of how
// many questions they get right and how many they get incorrect. Regardless of
// whether the answer is correct or wrong the next question should be asked
// immediately afterwards.
//
// The CSV file should default to problems.csv (example shown below), but the
// user should be able to customize the filename via a flag.
//
// The CSV file will be in a format like below, where the first column is a
// question and the second column in the same row is the answer to that
// question.
//
//     5+5,10
//     7+3,10
//     1+1,2
//     8+3,11
//     1+2,3
//     8+6,14
//     3+1,4
//     1+4,5
//     5+1,6
//     2+3,5
//     3+3,6
//     2+4,6
//     5+2,7
//
// You can assume that quizzes will be relatively short (< 100 questions) and
// will have single word/number answers.
//
// At the end of the quiz the program should output the total number of
// questions correct and how many questions there were in total. Questions given
// invalid answers are considered incorrect.
//
// NOTE: CSV files may have questions with commas in them. Eg: "what 2+2,
// sir?",4 is a valid row in a CSV. I suggest you look into the CSV package in
// Go and don’t try to write your own CSV parser.

package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	// "errors"
)

type QuizItem struct {
	Solution string
	Question string
}

type Quiz struct {
	Items []QuizItem
}

func main() {

	fileStr, err := ReadFile("./problems.csv", 4096)
	HandleErr(err, "Read File Fail")
	items, err := ParseCsv(fileStr)
	HandleErr(err, "Parse Csv Fail")

	var quiz Quiz
	quiz.Items = items
	p := quiz.Start()

	fmt.Println("Your Score: ", p)
}

// Readfile Reades the specified file name and returns a string
// of the file contents
func ReadFile(name string, size int) (string, error){
	f, err := os.Open(name)
	store := make([]byte, size)
	_, err = f.Read(store)
	return string(store), err
}

// ParseCsv reads the csv string and returns a slice of QuizItem 
// Assumes the format of `question, answer`.
func ParseCsv(data string) ([]QuizItem, error){
	r := csv.NewReader(strings.NewReader(data))
	var q []QuizItem
	var qi QuizItem
	for {
		d, err := r.Read()

		if err  == io.EOF {
			break
		}
		HandleErr(err, "CSV Read Error")

		qi.Question = d[0]
		qi.Solution = d[1]
		q = append(q, qi)
	}
	return q, nil
}

// Quiz.Start() starts asking the questions, returns the total points
func (q *Quiz) Start() int {
	var points int
	for _, v := range q.Items {

		correct, err := Question(v.Question, v.Solution)
		HandleErr(err, "Question fail")

		if correct {
			points += 1
		}
	}
	return points
}

// Question asks the question and waits for the user to answer. Returns
// true or false if the answer is right or wrong.
func Question(q, s string) (bool, error) {
	fmt.Println(q, "?")
	reader := bufio.NewReader(os.Stdin)
	answer, err := reader.ReadString('\n')
	answer = strings.TrimSuffix(answer, "\n")
	
	if answer != s {
		fmt.Println("Wrong. Correct: ", s)
		return false, err
	}
	return true, err
}

func HandleErr(e error, msg string) {
	if e != nil {
		log.Fatal(e, msg)
	}
}
