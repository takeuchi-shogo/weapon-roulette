package main

import


func main() {
	db := infrastructure.NewDB()

	r := infrastructure.NewRouting()
	r.Run(r.Port)
}