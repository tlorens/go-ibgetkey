package ibkey

/*
#include "cgetkey.h"
*/
import "C"

const HOME   = 149
const END    = 152
const PG_UP  = 153
const PG_DN  = 154
const KEY_UP = 165
const KEY_DN = 166
const KEY_RT = 167
const KEY_LF = 168

func ReadKey(raw bool) int {
	var ch int

	if raw  == true {
		ch = int(C.ibgetchraw())
	} else {
		ch = int(C.ibgetch())
	}

	return ch
}
