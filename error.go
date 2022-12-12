package spingo

import "log"

func check(err error, origin string) {
	if err != nil {
		log.Printf("Error at %s\nDetails:\n", origin)
		panic(err.Error())
	}
}
