package main

import "bytes"

func re2Email3RuleBV(input []byte) int {
	/*!max:re2c*/
	var count int
	var cur, mar int
	input = append(input, bytes.Repeat([]byte{0}, YYMAXFILL)...)
	lim := len(input)

	for { /*!re2c
		re2c:bit-vectors        = 1;
		re2c:posix-captures     = 0;
		re2c:case-insensitive   = 0;
		re2c:define:YYCTYPE     = byte;
		re2c:define:YYPEEK      = "input[cur]";
		re2c:define:YYSKIP      = "cur += 1";
		re2c:define:YYBACKUP    = "mar = cur";
		re2c:define:YYRESTORE   = "cur = mar";
		re2c:define:YYLESSTHAN  = "lim - cur < @@";
		re2c:define:YYSTAGP     = "@@{tag} = cur";
		re2c:define:YYSTAGN     = "@@{tag} = -1";
		re2c:define:YYSHIFTSTAG = "@@{tag} += @@{shift}";
		re2c:define:YYFILL      = "panic(`fill unexpected`)";

		email   = [+\-.0-9A-Z_a-z]+@[\-.0-9A-Z_a-z]+[.][\-.0-9A-Z_a-z]+;
		base64  = [;]base64[,][+0-9A-Za-z/]+[=]*;
		tooLong = [+\-.0-9A-Z_a-z]{100,};

		{email}   { count += 1; continue }
		{base64}  { continue }
		{tooLong} { continue }
		*         { continue }
		[\x00]    { return count }
		*/
	}
}
