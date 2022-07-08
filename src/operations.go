package main
import (
    _ "fmt"
    "strconv"
)

/**
 * When there is a call to copy operation could be 
 * a loop, this variable is to catch it if there
 * is one.
 * **/
var NL_lastvisited *CELL = nil

func op_getcoords_cell_byref (ref string) (int, int) {
    /**
     * This function is called when there is
     * a reference to another cell as A0, C5, Z54 etc.A
     * **/
     var col int = int(ref[1]) - 65
     row, _ := strconv.Atoi( ref[2:len(ref) - 1] )
     return row, col
}

func Op_copy (thsCell *CELL) {
    /**
     * Gotta set the current cells to avoid loops,
     * for example:
     *     | {B0} | {A0} |
     * (0, 0) trying to copy content of (0, 1), but
     * (0, 1) also is trying to copy the content of
     * (0, 0), loop detected.
     * **/
    rowCpy, colCpy := op_getcoords_cell_byref(thsCell.content)
    var cpyCell *CELL = &Table[rowCpy][colCpy]

    if NL_lastvisited != nil {
        if cpyCell.row == NL_lastvisited.row && cpyCell.col == NL_lastvisited.col {
            thsCell.content = "LOOP!"
            thsCell.ctype = ERROR
            return
        }
    }

    if cpyCell.ctype == OP_COPY {
        NL_lastvisited = thsCell
        Op_copy(cpyCell)
    }

    thsCell.content = cpyCell.content
    thsCell.ctype = cpyCell.ctype
}
