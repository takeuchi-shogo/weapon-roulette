package infrastrcture

import "github.com/gin-gonic/gin"

type Routing struct {
	DB   *DB
	Gin  *gin.Engine
	Port string
}

func NewRouting(db *DB) *Routing {
	c := NewConfig()
	r := &Routing{
		DB: db,
	}

	r.setRouting()
	return r
}

func (r *Routing) setRouting() {

	weaponsController := product.NewWeaponsController(r.DB)

	v1 := r.Gin.Group("/v1/product")
	{
		/*
		* roulette
		 */
		v1.GET("/roulette", func(c *gin.Context) { eaponsController.GetList(c) })
	}
}
