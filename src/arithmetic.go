/**
 * Spread Sheet.
 * nonporia, Jul 10.
 * **/
package main
import (
    "strconv"
)

func arith_isoperator (m_oprt string) int {
    if m_oprt == "END." { return 1 }
    if m_oprt == "+" || m_oprt == "-" || m_oprt == "/" || m_oprt == "*" { return 1 }
    return 0
}

func arith_getnumber (cuCell *CELL, refnum string) float64 {
    /**
     * To get a number there are only two posibilities:
     *    ~ Already the cell contains a literal number.
     *    ~ Is a copy cell operation.
     * If it is the second one, that cell must be a number.
     * **/
    if rgx_isnumber.MatchString(refnum) {
        thsNum, _ := strconv.ParseFloat(refnum, 64)
        return thsNum
    }
    if rgx_iscopyop.MatchString(refnum) {
        refRow, refCol := op_getcoords_cell_byref(refnum)
        if Table[refRow][refCol].ctype == LIT_NUMBER {
            return Table[refRow][refCol].asnum
        }
    }

    cuCell.content = "REFER"
    cuCell.ctype = ERROR
    return 0.0
}
