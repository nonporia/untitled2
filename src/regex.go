package main
import (
    "regexp"
)

/**
 * Any integer or any float.
 * **/
var rgx_isnumber,  _ = regexp.Compile("^(-|)(\\d+|\\d+.\\d+)$")
var rgx_isstring,  _ = regexp.Compile("^\".*\"$")
var rgx_isboolean, _ = regexp.Compile("^(true|false)$")
