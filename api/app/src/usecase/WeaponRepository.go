package usecase

import (
	"weapon-roulette/src/domain"

	"github.com/jinzhu/gorm"
)

type WeaponRepository interface {
	FindByKey(db *gorm.DB, key string) (weapons []domain.Weapons, err error)
}
