package main

import (
	"fmt"
	"os"
	"github.com/01-edu/z01" // Πακέτο για εκτύπωση χαρακτήρων στην κονσόλα
)

const size = 9 // Σταθερό μέγεθος πίνακα Sudoku

// Εκτυπώνει τον πίνακα Sudoku
func printBoard(board [size][size]int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			z01.PrintRune(rune(board[i][j] + '0')) // Μετατροπή αριθμού σε χαρακτήρα
			z01.PrintRune(' ') // Κενό μεταξύ των αριθμών
		}
		z01.PrintRune('\n') // Νέα γραμμή μετά από κάθε γραμμή του πίνακα
	}
}

// Ελέγχει αν ένας αριθμός μπορεί να τοποθετηθεί στη συγκεκριμένη θέση
func isValid(board [size][size]int, row, col, num int) bool {
	// Έλεγχος γραμμής και στήλης
	for i := 0; i < size; i++ {
		if board[row][i] == num || board[i][col] == num {
			return false
		}
	}

	// Υπολογισμός της αρχικής θέσης του 3x3 υποπίνακα
	startRow := row - row%3
	startCol := col - col%3

	// Έλεγχος του 3x3 τετραγώνου
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[startRow+i][startCol+j] == num {
				return false
			}
		}
	}

	return true // Αν δεν υπάρχει πρόβλημα, επιστρέφεται true
}

// Επίλυση του Sudoku μέσω αναδρομής και backtracking
func solveSudoku(board *[size][size]int) bool {
	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			if board[row][col] == 0 { // Αν η θέση είναι κενή
				for num := 1; num <= 9; num++ {
					if isValid(*board, row, col, num) {
						board[row][col] = num // Προσπάθεια τοποθέτησης αριθμού
						if solveSudoku(board) {
							return true // Αν λυθεί, επιστρέφει true
						}
						board[row][col] = 0 // Αν όχι, επαναφορά (backtrack)
					}
				}
				return false // Αν κανένας αριθμός δεν είναι έγκυρος
			}
		}
	}
	return true // Αν δεν υπάρχουν κενές θέσεις, το Sudoku λύθηκε
}

// Ανάλυση των ορισμάτων εισόδου και μετατροπή σε πίνακα Sudoku
func parseArgs(args []string) ([size][size]int, bool) {
	var board [size][size]int

	if len(args) != size {
		return board, false // Πρέπει να υπάρχουν 9 σειρές
	}

	for i := 0; i < size; i++ {
		if len(args[i]) != size {
			return board, false // Κάθε σειρά πρέπει να έχει 9 χαρακτήρες
		}
		for j := 0; j < size; j++ {
			ch := args[i][j]
			if ch == '.' {
				board[i][j] = 0 // Το '.' συμβολίζει κενό κελί
			} else if ch >= '1' && ch <= '9' {
				board[i][j] = int(ch - '0') // Μετατροπή χαρακτήρα σε ακέραιο
			} else {
				return board, false // Άκυρος χαρακτήρας
			}
		}
	}

	return board, true
}

// Κύρια συνάρτηση
func main() {
	args := os.Args[1:] // Παράλειψη του ονόματος του προγράμματος
	board, valid := parseArgs(args)

	// Αν τα δεδομένα δεν είναι έγκυρα ή δεν υπάρχει λύση
	if !valid || !solveSudoku(&board) {
		fmt.Println("Error$")
		return
	}

	// Εκτύπωση του λυμένου πίνακα
	printBoard(board)
}