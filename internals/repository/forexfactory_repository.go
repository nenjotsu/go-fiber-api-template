package repository

import (
	"go-fiber-api-template/internals/entity"

	"gorm.io/gorm"
)

type ForexFactoryRepository interface {
	GetForexFactory(timeString string) ([]*entity.ForexFactory, error)
	Create(forexFactory *entity.ForexFactory) error
	FindByTime(time string) (*entity.ForexFactory, error)
	UpdateByTime(forexFactory *entity.ForexFactory) error
}

type forexFactoryRepository struct {
	db *gorm.DB
}

func NewForexFactoryRepository(db *gorm.DB) ForexFactoryRepository {
	return &forexFactoryRepository{db: db}
}

func (r *forexFactoryRepository) GetForexFactory(timeString string) ([]*entity.ForexFactory, error) {
	forexFactoryList := []*entity.ForexFactory{}
	err := r.db.Where("time LIKE ?", timeString+"%").Find(&forexFactoryList).Error
	if err != nil {
		return nil, err
	}
	return forexFactoryList, nil
}

func (r *forexFactoryRepository) Create(forexFactory *entity.ForexFactory) error {
	return r.db.Create(&forexFactory).Error
}

// Find by time
func (r *forexFactoryRepository) FindByTime(time string) (*entity.ForexFactory, error) {
	forexFactory := &entity.ForexFactory{}
	err := r.db.Where("time = ?", time).First(&forexFactory).Error

	if err != nil {
		return nil, err
	}
	return forexFactory, nil
}

// update by time
func (r *forexFactoryRepository) UpdateByTime(forexFactory *entity.ForexFactory) error {
	return r.db.Model(&entity.ForexFactory{}).Where("time = ?", forexFactory.Time).Updates(forexFactory).Error
}
