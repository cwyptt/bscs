# SysAdminQuiz

SysAdminQuiz is a terminal-based multiple-choice quiz application built in Go using the Charm libraries (`bubbletea` and `lipgloss`). It tests system administration knowledge with questions loaded from a `questions.txt` file. The quiz features a visually appealing TUI (Text User Interface) with score tracking (e.g., `15/20 (75%)`) and supports navigation via number keys (0–3), arrow keys, and Enter.

## Prerequisites

- The `sysadminquiz` executable binary.
- A `questions.txt` file in the same directory as the executable, formatted with 6 lines per question:
  - Line 1: Question text
  - Lines 2–5: Four answer options
  - Line 6: Correct answer index (1, 2, 3, or 4)
- A terminal that supports ANSI colors (most modern terminals do).

## How to run

- Precompiled binaries are provided for Windows and Linux operating systems.
- On Windows, navigate to the location of the repository and run ".\SysAdminQuizBSCS.exe" in CMD
- On Linux, cd into the repository and run "./sysadmin-quiz"

Ensure the binary and `questions.txt` are in the same directory.
