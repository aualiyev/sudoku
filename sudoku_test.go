package piscine

import (
	"testing"
)

func TestSudoku1(t *testing.T) {
	expected := "3 9 6 2 4 5 7 8 1"
	expected += "\n1 7 8 3 6 9 5 2 4"
	expected += "\n5 2 4 8 1 7 3 9 6"
	expected += "\n2 8 7 9 5 1 6 4 3"
	expected += "\n9 3 1 4 8 6 2 7 5"
	expected += "\n4 6 5 7 2 3 9 1 8"
	expected += "\n7 1 2 6 3 8 4 5 9"
	expected += "\n6 5 9 1 7 4 8 3 2"
	expected += "\n8 4 3 5 9 2 1 6 7"
	expected += "\n"

	input := []string{
		".96.4...1",
		"1...6...4",
		"5.481.39.",
		"..795..43",
		".3..8....",
		"4.5.23.18",
		".1.63..59",
		".59.7.83.",
		"..359...7",
	}

	parsed, _ := ParseSudoku(input)
	solved, _ := SolveSudoku(parsed)
	res := GetPrintedBoard(solved)

	if res != expected {
		t.Errorf("Error: result and expected are not matched")
	}
}



func TestSudoku2(t *testing.T) {
	expected := "1 4 5 8 9 2 6 7 3"
	expected += "\n8 9 3 1 7 6 4 2 5"
	expected += "\n2 7 6 4 3 5 8 1 9"
	expected += "\n5 1 9 2 4 7 3 8 6"
	expected += "\n7 6 2 5 8 3 1 9 4"
	expected += "\n3 8 4 9 6 1 7 5 2"
	expected += "\n9 5 7 6 1 4 2 3 8"
	expected += "\n4 3 8 7 2 9 5 6 1"
	expected += "\n6 2 1 3 5 8 9 4 7"
	expected += "\n"

	input := []string{
		"1.58.2...",
		".9..764.5",
		"2..4..819",
		".19..73.6",
		"762.83.9.",
		"....61.5.",
		"..76...3.",
		"43..2.5.1",
		"6..3.89..",
	}

	parsed, _ := ParseSudoku(input)
	solved, _ := SolveSudoku(parsed)
	res := GetPrintedBoard(solved)

	if res != expected {
		t.Errorf("Error: result and expected are not matched")
	}
}

func TestSudoku3(t *testing.T) {
	expected := "3 4 7 9 1 5 6 2 8"
	expected += "\n2 9 6 7 8 3 5 4 1"
	expected += "\n5 1 8 6 2 4 9 7 3"
	expected += "\n8 6 4 1 5 7 2 3 9"
	expected += "\n1 3 2 4 6 9 7 8 5"
	expected += "\n9 7 5 8 3 2 1 6 4"
	expected += "\n4 5 9 2 7 8 3 1 6"
	expected += "\n7 8 1 3 9 6 4 5 2"
	expected += "\n6 2 3 5 4 1 8 9 7"
	expected += "\n"

	input := []string{"34.91..2.", ".96.8..41", "..8.2..7.", ".6..57.39", "1.2.6.7..", "97..3..64", "45.2.8..6", ".8..9..5.", "6.3..189."}
	parsed, _ := ParseSudoku(input)
	solved, _ := SolveSudoku(parsed)
	res := GetPrintedBoard(solved)

	if res != expected {
		t.Errorf("Error: result and expected are not matched")
	}
}

