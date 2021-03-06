/**
 * Spread Sheet.
 * nonporia, Jul 7.
 * **/
package main
import (
	"fmt"
	"strings"
	"strconv"
)

type CELL_TYPE int
const (
    LIT_STRING = iota
    LIT_NUMBER
    LIT_FSTRING
    LIT_BOOLEAN
    LIT_UNKNOWN

    OP_COPY
    OP_ARITH
    OP_MAXMIN

    ERROR
)

type CELL struct {
    /**
     * Information for all cells.
     * **/
    row int
    col int
    content string
    ctype CELL_TYPE

    /** Extra information if it is
     * a number.
     * **/
     asnum float64
}

var Table[500][26] CELL
var row_max int = 0
var col_max int = 0
var dig_max int = 5

func prs_cleanstr (clnstr string) string {
    /**
     * Is necessary clean the string of whitespaces
     * and be careful with string types (in literal strings
     * the spaces are saved).
     * **/
    var newstr string
    var instr bool = false
    for i := 0; i < len(clnstr); i++ {
        if clnstr[i] == '"' || clnstr[i] == '\'' {
            instr = !instr
        }
        if !instr && clnstr[i] != ' ' {
            newstr += string(clnstr[i])
        }
        if instr {
            newstr += string(clnstr[i])
        }
    }

    if len(newstr) > dig_max {
        dig_max = len(newstr)
    }
    return newstr
}

func prs_setcell (cont string, row, col int) {
    var cell CELL
    cell.row = row
    cell.col = col
    cell.content = prs_cleanstr(cont)

    if rgx_isnumber.MatchString(cell.content) {
        cell.ctype = LIT_NUMBER
        cell.asnum, _ = strconv.ParseFloat(cell.content, 64)
        goto SET_CELL
    }
    if rgx_isstring.MatchString(cell.content) {
        cell.ctype = LIT_STRING
        cell.content = cell.content[1:len(cell.content) - 1]
        goto SET_CELL
    }
    if rgx_isboolean.MatchString(cell.content) {
        cell.ctype = LIT_BOOLEAN
        goto SET_CELL
    }
    if rgx_iscopyop.MatchString(cell.content) {
        cell.ctype = OP_COPY
        goto SET_CELL
    }
    if rgx_isfstring.MatchString(cell.content) {
        cell.ctype = LIT_FSTRING
        cell.content = cell.content[1:len(cell.content) - 1]
        goto SET_CELL
    }
    if rgx_isarith.MatchString(cell.content) {
        cell.ctype = OP_ARITH
        goto SET_CELL
    }
    if rgx_ismaxminop.MatchString(cell.content) {
        cell.ctype = OP_MAXMIN
        goto SET_CELL
    }

    cell.ctype = LIT_UNKNOWN
    SET_CELL:
    Table[row][col] = cell
}

func prs_print (content string) {
    fmt.Printf("%s", content)
    for spc := 0; spc < (dig_max - len(content)); spc++ {
        fmt.Printf(" ")
    }
}

func Prs_setcolumns (content string, row_idx int) {
	if len(content) == 0 {
		fmt.Printf("%d: Empty string", row_idx)
		return
	}

	cells := strings.Split(content, "|")
    var col_idx int
    for col_idx = 0; col_idx < len(cells); col_idx++ {
        prs_setcell(cells[col_idx], row_idx - 1, col_idx)
    }

    if col_idx > col_max {
        col_max = col_idx
    }
    row_max++
}

func Prs_printable () {
    var cuCell *CELL
    for row := 0; row < row_max; row ++ {
        for col := 0; col < col_max; col ++ {
            cuCell = &Table[row][col]

            if cuCell.ctype == OP_COPY {
                Op_copy(cuCell)
                NL_lastvisited = nil
            }
            if cuCell.ctype == LIT_FSTRING { Op_fstring(cuCell) }
            if cuCell.ctype == OP_ARITH { Op_arith(cuCell) }
            if cuCell.ctype == OP_MAXMIN { Op_maxmin(cuCell) }

            prs_print(cuCell.content)
        }
        fmt.Println()
    }
}
