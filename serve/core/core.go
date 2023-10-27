package core

import "runtime"

func Bomb(err error) {
	if err != nil {
		panic(err)
	}
}

func Stack() string {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	return string(buf[:n])
}
