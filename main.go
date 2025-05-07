package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	questionStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFD700")). // Gold
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#4CAF50"))

	optionStyle = lipgloss.NewStyle().
			PaddingLeft(2).
			Foreground(lipgloss.Color("#00BFFF")) // Blue

	selectedOptionStyle = optionStyle.Copy().
				Background(lipgloss.Color("#4CAF50")). // Green background
				Foreground(lipgloss.Color("#FFFFFF"))  // White text

	feedbackStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("#FF4500")) // Orange

	correctStyle = feedbackStyle.Copy().
			Foreground(lipgloss.Color("#32CD32")) // Lime Green

	incorrectStyle = feedbackStyle.Copy().
			Foreground(lipgloss.Color("#FF0000")) // Red

	scoreStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFA500")). // Orange
			Padding(1, 2).
			Border(lipgloss.DoubleBorder()).
			BorderForeground(lipgloss.Color("#FFD700")) // Gold
)

type Question struct {
	Text          string
	Options       [4]string
	CorrectAnswer int
}

type Model struct {
	questions  []Question
	current    int
	selected   int // Currently selected option
	feedback   string
	showResult bool
	quizEnded  bool
	correct    int // Number of correct answers
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			if m.quizEnded {
				return m, tea.Quit
			}
			if m.showResult {
				// Move to next question
				m.current++
				m.selected = 0
				m.feedback = ""
				m.showResult = false
				if m.current >= len(m.questions) {
					m.quizEnded = true
				}
			} else {
				// Submit answer
				if m.selected+1 == m.questions[m.current].CorrectAnswer {
					m.feedback = "Correct!"
					m.correct++
				} else {
					correct := m.questions[m.current].Options[m.questions[m.current].CorrectAnswer-1]
					m.feedback = fmt.Sprintf("Incorrect. The answer is: %s", correct)
				}
				m.showResult = true
			}
		case "up":
			if !m.showResult && !m.quizEnded {
				m.selected = (m.selected - 1 + 4) % 4 // Wrap around
			}
		case "down":
			if !m.showResult && !m.quizEnded {
				m.selected = (m.selected + 1) % 4 // Wrap around
			}
		case "1", "2", "3", "4":
			if !m.showResult && !m.quizEnded {
				sel, _ := strconv.Atoi(msg.String())
				m.selected = sel - 1
			}
		}
	}
	return m, nil
}

func (m Model) View() string {
	if m.quizEnded {
		// Calculate percentage
		percentage := float64(m.correct) / float64(len(m.questions)) * 100
		scoreText := fmt.Sprintf("Quiz Complete!\nScore: %d/%d (%.0f%%)\nPress Enter to exit.", m.correct, len(m.questions), percentage)
		return scoreStyle.Render(scoreText)
	}

	var s strings.Builder
	// Display question
	s.WriteString(questionStyle.Render(fmt.Sprintf("Question %d: %s", m.current+1, m.questions[m.current].Text)))
	s.WriteString("\n")

	// Display options or result
	if m.showResult {
		style := incorrectStyle
		if m.feedback == "Correct!" {
			style = correctStyle
		}
		// Include current score in feedback
		percentage := float64(m.correct) / float64(m.current+1) * 100
		s.WriteString(style.Render(fmt.Sprintf("%s\nCurrent Score: %d/%d (%.0f%%)", m.feedback, m.correct, len(m.questions), percentage)))
		s.WriteString("\nPress Enter to Continue...\n")
	} else {
		// Display options
		for i, option := range m.questions[m.current].Options {
			prefix := fmt.Sprintf("%d. ", i+1)
			if i == m.selected {
				s.WriteString(selectedOptionStyle.Render(prefix + option))
			} else {
				s.WriteString(optionStyle.Render(prefix + option))
			}
			s.WriteString("\n")
		}
		s.WriteString("\nUse 1-4 or arrow keys to select, Enter to submit, q to quit.\n")
	}

	return s.String()
}

func readQuestions(filename string) ([]Question, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var questions []Question
	scanner := bufio.NewScanner(file)
	var currentQuestion Question
	lineNum := 0
	optionIndex := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		switch lineNum % 6 {
		case 0: // Question text
			currentQuestion = Question{}
			currentQuestion.Text = line
		case 1, 2, 3, 4: // Options
			currentQuestion.Options[optionIndex] = line
			optionIndex++
		case 5: // Correct answer index
			correct, err := strconv.Atoi(line)
			if err != nil || correct < 1 || correct > 4 {
				return nil, fmt.Errorf("Invalid correct answer index at line %d: %s", lineNum+1, line)
			}
			currentQuestion.CorrectAnswer = correct
			questions = append(questions, currentQuestion)
			optionIndex = 0
		}
		lineNum++
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if lineNum%6 != 0 {
		return nil, fmt.Errorf("File format error: Incomplete question block")
	}

	return questions, nil
}

func main() {
	// Read questions from file
	questions, err := readQuestions("questions.txt")
	if err != nil {
		fmt.Printf("Error reading questions: %v\n", err)
		os.Exit(1)
	}
	if len(questions) == 0 {
		fmt.Println("No questions found in file.")
		os.Exit(1)
	}

	// Initial model
	m := Model{
		questions: questions,
	}

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v\n", err)
		os.Exit(1)
	}
}
