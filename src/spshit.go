package main
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Not enough arguments.")
		os.Exit(1)
	}

	tblf, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("(%s) does not exist.\n", os.Args[1])
		os.Exit(1)
	}

	scnner := bufio.NewScanner(tblf)
	var row_idx int = 1
	for scnner.Scan() {
		Prs_setcolumns(scnner.Text(), row_idx)
		row_idx++
	}

    Prs_printable();
}
