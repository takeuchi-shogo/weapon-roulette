package product

type WeaponInteractor struct {
	DB     usecase.DBRepository
	Weapon usecase.WeaponRepository
}

func (interactor *WeaponInteractor) GetList(key string) (weapons []domain.Weapons, resultStatus *usecase.ResultStatus) {

	db := interactor.DB.Connect()

	foundWeapons, err := interactor.Weapon.FindByKey(db, key)

	if err != nil {
		return []domain.Weapons, usecase.NewResultStatus(400, err)
	}

	return foundWeapons, usecase.NewResultStatus(200, nil)
}
