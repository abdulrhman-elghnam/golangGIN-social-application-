package repository

import (
	"gorm.io/gorm"
)

type BaseRepository[T any] struct {
	db *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) *BaseRepository[T] {
	return &BaseRepository[T]{
		db: db,
	}
}

func (r *BaseRepository[T]) Create(data T) error {
	return r.db.Create(&data).Error
}

func (r *BaseRepository[T]) FindByID(id uint) (*T, error) {
	var data T

	err := r.db.First(&data, id).Error

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *BaseRepository[T]) Find(by string, value string) ([]T, error) {
	var data []T

	err := r.db.
		Where(by+" = ?", value).
		Find(&data).
		Error

	return data, err
}

func (r *BaseRepository[T]) FindAll() ([]T, error) {
	var data []T

	err := r.db.Find(&data).Error
	return data, err
}

func (r *BaseRepository[T]) Update(data T) error {
	return r.db.Save(&data).Error
}

func (r *BaseRepository[T]) Delete(id uint) error {
	var data T

	return r.db.Delete(&data, id).Error
}
