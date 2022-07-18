/**
 * Spread Sheet.
 * nonporia, Jul 8.
 * **/
package main
import (
    "fmt"
    "strconv"
    "strings"
    "math"
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

func op_setths_cell (n_type CELL_TYPE, n_content string, thsCell *CELL) {
    /**
     * 'thsCell' will be as an auxiliar to get another cell value while
     * another operation is working, for example in the fstrings this function
     * will be called to get the value of some pointed cell.
     * **/
     thsCell.ctype = n_type
     thsCell.content = n_content
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
    if cpyCell.ctype == LIT_STRING {
        Op_fstring(cpyCell)
    }
    if cpyCell.ctype == OP_ARITH {
        Op_arith(cpyCell)
        thsCell.asnum = cpyCell.asnum
    }
    if cpyCell.ctype == OP_MAXMIN {
        Op_maxmin(cpyCell)
        thsCell.asnum = cpyCell.asnum
    }

    thsCell.content = cpyCell.content
    thsCell.ctype = cpyCell.ctype
}

func Op_fstring (thsCell *CELL) {
    var words []string = strings.Split(thsCell.content, " ")
    var newstr string

    for w := 0; w < len(words); w++ {
        if rgx_iscopyop.MatchString(words[w]) {
            op_setths_cell(OP_COPY, words[w], thsCell)
            Op_copy(thsCell)
            words[w] = thsCell.content
        }
        if rgx_isarith.MatchString(words[w]) {
            op_setths_cell(OP_ARITH, words[w], thsCell)
            Op_arith(thsCell)
            words[w] = thsCell.content
        }
        if rgx_ismaxminop.MatchString(words[w]) {
            op_setths_cell(OP_MAXMIN, words[w], thsCell)
            Op_maxmin(thsCell)
            words[w] = thsCell.content
        }

        newstr += words[w] + " "
    }
    thsCell.content = newstr
    thsCell.ctype = LIT_STRING
}

func Op_arith (thsCell *CELL) {
    var wholeOp []string = strings.Split(thsCell.content[1:], ";")
    wholeOp[len(wholeOp) - 1] = "END." /** To mark the end of the operation. **/
    var cu_opr, nxt_opr string

    for idx := 1; idx < len(wholeOp); idx += 2 {
        if (idx + 1) >= len(wholeOp) || wholeOp[idx] == "END." {
            break
        } else {
            /**
             * idx variable is always pointing to one operator, so:
             *     [NUMBER OPERATOR NUMBER OPERATOR NUMBER END]
             *               |                |
             *              idx            idx + 2
             * **/
            cu_opr = wholeOp[idx]
            nxt_opr = wholeOp[idx + 2]

            if arith_isoperator(cu_opr) == 0 || arith_isoperator(nxt_opr) == 0 {
                thsCell.content = "OPERA"
                thsCell.ctype = ERROR
                return
            }
        }

        /**
         * Getting the first number of the whole operation.
         * Setting the first value of the whole operation.
         * **/
        if idx == 1 {
            fvalue := arith_getnumber(thsCell, wholeOp[0])
            if thsCell.ctype == ERROR { return }
            thsCell.asnum = fvalue
        }
        /**
         * The next number of the operation is always at the rigth of current
         * operator.
         * **/
        thsnum := arith_getnumber(thsCell, wholeOp[idx + 1])
        if thsCell.ctype == ERROR { return }

        if cu_opr == "+" || cu_opr == "-" {
            var thenum float64 = thsnum
            /**
             * Respect to precedence.
             * if the next operations is not an addition and neither is a
             * substraction gotta make the next operation first.
             * **/
            if nxt_opr != "+" && nxt_opr != "-" && nxt_opr != "END." {
                if (idx + 3) >= len(wholeOp) {
                    thsCell.content = "INAOP"
                    thsCell.ctype = ERROR
                    return
                }
                thsnum = arith_getnumber(thsCell, wholeOp[idx + 3])
                if thsCell.ctype == ERROR { return }

                if nxt_opr == "*" { thenum *= thsnum } else { thenum /= thsnum; }
                idx += 2
            }
            /**
             * Sets the value of the next operation to the current
             * operation, and if there was not just make the current
             * operation.
             * **/
            if cu_opr == "+" { thsCell.asnum += thenum } else { thsCell.asnum -= thenum }
        } else {
            if cu_opr == "*" { thsCell.asnum *= thsnum } else { thsCell.asnum /= thsnum }
        }

        if thsCell.asnum == math.Inf(1) || thsCell.asnum == math.Inf(-1) {
            thsCell.content = "DIV_0"
            thsCell.ctype = ERROR
            return
        }
    }

    thsCell.content = fmt.Sprintf("%.2f", thsCell.asnum)
    thsCell.ctype = LIT_NUMBER
}

func Op_maxmin (thsCell *CELL) {
    var ismin bool = false
    if thsCell.content[:3] == "MIN" {
        ismin = true
    }

    var points []string = strings.Split(thsCell.content[4:], ":")
    var auxRow int = 0
    var fCell bool = true
    var cuCell CELL

    p1R, p1C := op_getcoords_cell_byref(points[0])
    p2R, p2C := op_getcoords_cell_byref(points[1])

    if (p1C > p2C) || (p1R > p2R) {
        thsCell.content = "RANGE"
        thsCell.ctype = ERROR
        return
    }

    for p1C <= p2C {
        auxRow = p1R
        for auxRow <= p2R {
            cuCell = Table[auxRow][p1C]
            if cuCell.ctype != LIT_NUMBER {
                thsCell.content = "REFER"
                thsCell.ctype = ERROR
                return
            }

            if fCell {
                thsCell.asnum = cuCell.asnum
                fCell = false
            }

            if !ismin && cuCell.asnum > thsCell.asnum { thsCell.asnum = cuCell.asnum  }
            if ismin && cuCell.asnum < thsCell.asnum { thsCell.asnum = cuCell.asnum  }
            auxRow++
        }
        p1C++
        auxRow = 0
    }

    thsCell.content = fmt.Sprintf("%.2f", thsCell.asnum)
    thsCell.ctype = LIT_NUMBER
}
