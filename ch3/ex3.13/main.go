package main

import (
	"fmt"
	"strconv"
)

/*
const (
	_   = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	GiB // 1073741824
	TiB // 1099511627776             (exceeds 1 << 32)
	PiB // 1125899906842624
	EiB // 1152921504606846976
	ZiB // 1180591620717411303424    (exceeds 1 << 64)
	YiB // 1208925819614629174706176
)
*/

// speed unit
const (
	KB = float64(1000)
	MB = KB * KB
	GB = MB * KB
	TB = GB * KB
	PB = TB * KB
	EB = PB * KB
	ZB = EB * KB
	YB = ZB * KB
)

func main() {
	fmt.Println(KB, MB, GB, TB, PB, EB, strconv.FormatFloat(ZB, 'g', 1, 64), strconv.FormatFloat(YB, 'g', 1, 64))
	//fmt.Println(KiB, MiB, GiB, TiB, PiB, EiB, string(ZiB), string(YiB))
}
