package keyboard

/*
#include "cgetkey.h"
*/
import "C"

const HOME   = 149 // DOS Keycode 71
const END    = 152 // DOS Keycode 79
const PG_UP  = 153 // DOS Keycode 73
const PG_DN  = 154 // DOS Keycode 81
const KEY_UP = 165 // DOS Keycode 72
const KEY_DN = 166 // DOS Keycode 80
const KEY_RT = 167 // DOS Keycode 77
const KEY_LF = 168 // DOS Keycode 75

func ReadKey() int {
	var ch int

	ch = int(C.readkey())
	if ch == 27 { // ESC
		if int(C.readkey()) == 91 { // [ bracket
			ch = int(C.readkey()) + 100
		}
	}

	return ch
}

func RawKey() int {
	return int(C.readkey())
}
