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
var rgx_isfstring, _ = regexp.Compile("^'.+'$");
var rgx_iscopyop,  _ = regexp.Compile("^{[A-Z][0-4](|[0-9])(|[0-9])}$");
