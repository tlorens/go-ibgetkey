package ibkey

/*
#include "cgetkey.h"
*/
import "C"

func ReadKey() int {
	return int(C.ibgetch())
}
