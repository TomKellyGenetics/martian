// Copyright (c) 2018 10X Genomics, Inc. All rights reserved.

package syntax

import (
	"testing"
)

func TestNextToken(t *testing.T) {
	check := func(s string, expect int) {
		t.Helper()
		tokid, tok := nextToken([]byte(s))
		if tokid != expect {
			t.Errorf("Wrong token type %d for %s", tokid, s)
		}
		if expect != INVALID && len(tok) != len(s) {
			t.Errorf("Only got %q back from %s", tok, s)
		}
	}
	check("# this is a comment\n", COMMENT)
	check(`"this/is/a/string"`, LITSTRING)
	check(`""`, LITSTRING)
	check(`"this\nis\na\nstring"`, LITSTRING)
	check(`"this \"is\" a string"`, LITSTRING)
	check(`"\"this \"is\" a string\""`, LITSTRING)
	check(`"this is a \U0001F9F5"`, LITSTRING)
	check(`"this is \u0146ot bad"`, LITSTRING)
	check(`"this is \167eird but fine"`, LITSTRING)
	check(`\"this is not a string"`, INVALID)
	check(`"\"this\" is not a string\"`, INVALID)
	check(`"Invalid unicode\u001!"`, INVALID)
	check(`"Invalid unicode\U0001!"`, INVALID)
	check(`"Invalid unicode\x1!"`, INVALID)
	check(`"Invalid unicode\xaz"`, INVALID)
	check(`@include`, INCLUDE_DIRECTIVE)
	check(`_INTERNAL_PIPELINE`, ID)
	check(`_type_name`, ID)
	check(`__type_name`, INVALID)
	check(`name_`, ID)
	check(`_name_`, ID)

	// int patterns
	check(`012345567`, NUM_INT)
	check(`-12345`, NUM_INT)
	check(`9223372036854775807`, NUM_INT)
	check(`00009223372036854775807`, NUM_INT)
	check(`-9223372036854775808`, NUM_INT)
	check(`-0`, NUM_INT)
	check(`0`, NUM_INT)

	// float patterns
	check(`0.0`, NUM_FLOAT)
	check(`0.1`, NUM_FLOAT)
	check(`0e0`, NUM_FLOAT)
	check(`0e10`, NUM_FLOAT)
	check(`0e01`, NUM_FLOAT)
	check(`1e0`, NUM_FLOAT)
	check(`10e0`, NUM_FLOAT)
	check(`01e0`, NUM_FLOAT)
	check(`0E0`, NUM_FLOAT)
	check(`0.0e0`, NUM_FLOAT)
	check(`0.0E0`, NUM_FLOAT)
	check(`0e+0`, NUM_FLOAT)
	check(`0E+0`, NUM_FLOAT)
	check(`0.0e+0`, NUM_FLOAT)
	check(`0.0E+0`, NUM_FLOAT)
	check(`0e-0`, NUM_FLOAT)
	check(`0E-0`, NUM_FLOAT)
	check(`0.0e-0`, NUM_FLOAT)
	check(`0.0E-0`, NUM_FLOAT)
	check(`-0.0`, NUM_FLOAT)
	check(`-0e0`, NUM_FLOAT)
	check(`-0e10`, NUM_FLOAT)
	check(`-0e01`, NUM_FLOAT)
	check(`-1e0`, NUM_FLOAT)
	check(`-10e0`, NUM_FLOAT)
	check(`-01e0`, NUM_FLOAT)
	check(`-0E0`, NUM_FLOAT)
	check(`-0.0e0`, NUM_FLOAT)
	check(`-0.0E0`, NUM_FLOAT)
	check(`-0e+0`, NUM_FLOAT)
	check(`-0E+0`, NUM_FLOAT)
	check(`-0.0e+0`, NUM_FLOAT)
	check(`-0.0E+0`, NUM_FLOAT)
	check(`-0e-0`, NUM_FLOAT)
	check(`-0E-0`, NUM_FLOAT)
	check(`-0.0e-0`, NUM_FLOAT)
	check(`-0.0E-0`, NUM_FLOAT)
}
