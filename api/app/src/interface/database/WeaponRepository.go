package database

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/takeuchi-shogo/weapon-roulette/api/app/src/domain"
)

type WeaponRepository struct{}

func (r *WeaponRepository) FindByKey(db *gorm.DB, key string) (weapons []domain.Weapons, err error) {

	weapons = []domain.Weapons{}

	db.Where("type = ?", key).Find(&weapons)
	if len(weapons) == 0 {
		return []domain.Weapons{}, errors.New("一致するものがありません")
	}

	return weapons, nil
}
