package main

import


func main() {
	db := infrastrcture.NewDB()

	r := infrastrcture.NewRouting()
	r.Run(r.Port)
}