# Quiz CLI Application

A command-line quiz application written in Go that reads questions and answers from a CSV file and presents them as an interactive timed quiz.

Inspired from [gophercises/quiz](https://github.com/gophercises/quiz) by [Jon Calhoun](https://github.com/joncalhoun) Exercise

# Features

- Read quiz questions from CSV files
- Configurable time limit per question
- Mid-question time warnings
- Case-insensitive answer checking
- Progress tracking
- Detailed score reporting
- Input validation for both CSV format and time settings

# Installation

To install the application, you need to have Go installed on your system. Then:

```bash
# Clone the repository
git clone https://github.com/ZnarKhalil/quiz-cli.git
cd quiz-app

# Build the application
go build -o quiz
```

# Usage

```bash
# Run with default settings (30 seconds per question, using problems.csv)
./quiz

# Run with custom time limit (e.g., 20 seconds per question)
./quiz -time 20

# Run with custom CSV file
./quiz -csv your-questions.csv

# Run with both custom time and CSV file
./quiz -time 15 -csv custom-quiz.csv
```

# CSV File

**Format**

The CSV file should follow this format:

```
question,answer
5+5,10
what is the capital of France?,Paris
```

**Requirements**

- Each line must contain exactly two fields: question and answer
- Neither question nor answer can be empty
- No header row is required
- Answers are case-insensitive

# Command Line Flags

| Flag  | Default        | Description                                           |
| ----- | -------------- | ----------------------------------------------------- |
| -csv  | "problems.csv" | Path to the CSV file containing questions and answers |
| -time | 30             | Time limit per question in seconds (minimum 1 second) |

# Output

**During the Quiz**

```bash
Question 1/10
Problem: 5+5 =
⚠️  15 seconds remaining!
10
✅ Correct!

Question 2/10
Problem: 7+3 =
⌛ Time's up for this question!
```

**Final Result**

```bash
--- Quiz Finished ---
Final Score: 8 out of 10 (80.0%)
```

# Notes

- The timer warning appears halfway through each question's time limit
- If a time limit of 0 or negative is provided, it defaults to 10 seconds
- The application trims whitespace from answers
  Answer comparison is case-insensitive
