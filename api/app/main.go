package main

import "weapon-roulette/src/infrastructure"

func main() {
	db := infrastructure.NewDB()

	r := infrastructure.NewRouting(db)
	r.Run(r.Port)
}
