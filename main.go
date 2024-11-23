package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type problem struct {
	q string
	a string
}

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timePerQuestion := flag.Int("time", 30, "the time limit per question in seconds")
	flag.Parse()
	if *timePerQuestion <= 0 {
		*timePerQuestion = 10
		fmt.Println("⚠️ Invalid time limit (must be positive). Using default of 10 seconds.")
	}

	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open csv file: %s \n", *csvFileName))
	}
	defer file.Close()

	r := csv.NewReader(file)
	// ReadAll is used as csv file will be small
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file")
	}

	// Add validation and get problems
	problems, err := parseAndValidateCSV(lines)
	if err != nil {
		exit(fmt.Sprintf("Invalid CSV file: %v", err))
	}

	// problems := parseLines(lines)

	fmt.Println("Quiz will begin in 3 seconds...")
	time.Sleep(3 * time.Second)

	correct := runQuiz(problems, *timePerQuestion)

	// Print final results
	fmt.Printf("\n--- Quiz Finished ---\n")
	fmt.Printf("Final Score: %d out of %d (%.1f%%)\n",
		correct,
		len(problems),
		float64(correct)/float64(len(problems))*100,
	)
}

func parseAndValidateCSV(lines [][]string) ([]problem, error) {
	if len(lines) == 0 {
		return nil, fmt.Errorf("CSV file is empty")
	}

	problems := make([]problem, len(lines))

	for i, line := range lines {
		if len(line) != 2 {
			return nil, fmt.Errorf("invalid format at line %d: expected 2 fields, got %d", i+1, len(line))
		}
		if strings.TrimSpace(line[0]) == "" || strings.TrimSpace(line[1]) == "" {
			return nil, fmt.Errorf("invalid format at line %d: empty question or answer", i+1)
		}
		problems[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return problems, nil
}

func runQuiz(problems []problem, timePerQuestion int) int {
	correct := 0
	total := len(problems)

	for i, p := range problems {
		fmt.Printf("\nQuestion %d/%d\n", i+1, total)
		fmt.Printf("Problem: %s = ", p.q)

		// Create timer for this question
		timer := time.NewTimer(time.Duration(timePerQuestion) * time.Second)
		answerCh := make(chan string)

		// Create warning timer
		warningTimer := time.NewTimer(time.Duration(timePerQuestion/2) * time.Second)
		go func() {
			<-warningTimer.C
			fmt.Printf("\n⚠️  %d seconds remaining!\n", timePerQuestion/2)
		}()

		// Handle user input in goroutine
		go func() {
			var answer string
			reader := bufio.NewReader(os.Stdin)
			answer, _ = reader.ReadString('\n')
			answer = strings.TrimSpace(answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\n⌛ Time's up for this question!\n")
			timer.Stop()
			warningTimer.Stop()
			// Drain answer channel if needed
			go func() {
				<-answerCh
			}()
			continue

		case answer := <-answerCh:
			timer.Stop()
			warningTimer.Stop()
			if strings.EqualFold(answer, p.a) { // Case-insensitive comparison
				correct++
				fmt.Println("✅ Correct!")
			} else {
				fmt.Printf("❌ Wrong! The correct answer was: %s\n", p.a)
			}
		}
	}

	return correct
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
