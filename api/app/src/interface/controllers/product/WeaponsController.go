package product

type WeaponsController struct {
	Interactor product.WeaponInteractor
}

func NewWeaponsController(db database.DB) *WeaponsController {
	return &WeaponsController{
		Interactor: product.WeaponInteractor{
			DB:     &database.DBRepository{DB: db},
			Weapon: &database.WeaponRepository{},
		},
	}
}

func (controller *WeaponsController) GetList(c controllers.Context) {

	key := c.Query("key")

	wepons, res := controller.Interactor.GetList(key)

	if res.Error != nil {
		c.JSON(res.StatusCode, controllers.NewH(res.Error.Error(), nil))
		return
	}

	c.JSON(res.StatusCode, controllers.NewH("success", weapons))
}
