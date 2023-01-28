package main

import (
	"awesomeProject/wangdejiang/src/router"
)

func main() {
	r := router.Router()
	err := r.Run(":9000")
	if err != nil {
		return
	}
}
