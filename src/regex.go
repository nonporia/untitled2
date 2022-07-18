/**
 * Spread Sheet.
 * nonporia, Jul 7.
 * **/
package main

import (
	"regexp"
)

/**
 * Any integer or any float.
 * **/
var rgx_isnumber,  _ = regexp.Compile("^(-|)(\\d+|\\d+.\\d+)$")
var rgx_isstring,  _ = regexp.Compile("^\".+\"$")
var rgx_isboolean, _ = regexp.Compile("^(TRUE|FALSE)$")
var rgx_isfstring, _ = regexp.Compile("^'.+'$")
var rgx_iscopyop,  _ = regexp.Compile("^{[A-Z][0-9](|[0-9])(|[0-9])}$")
/**
 * To make some arithemetic operation or bitwise operation (AND, XOR and OR one of those)
 * this "formula" must be used, is something like:
 *     [4; +; 6; *; 1;] (The spaces are not required).
 * As you can see every element in the element is separated by one comma, that is just to
 * interpret it easier.
 **/
var rgx_isarith,    _ = regexp.Compile("^\\[(.*;){3,}\\]$")
var rgx_ismaxminop, _ = regexp.Compile("^(MAX|MIN)\\(({[A-Z][0-9](|[0-9])(|[0-9])}:){2}\\)$")
