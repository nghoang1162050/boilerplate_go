package repository

import "gorm.io/gorm"

type BaseRepository[T any] interface {
    Search(condition string, pageNumber int, pageSize int, args ...interface{}) ([]T, int64, error)
	GetByID(id string) (*T, error)
	GetAll() ([]T, error)
	Create(entity *T) error
	Update(entity *T, id string) error
	Delete(id string) error
    First(condition string, args ...interface{}) (*T, error)
}

type baseRepository[T any] struct {
	db *gorm.DB
}

func NewRepository[T any](db *gorm.DB) BaseRepository[T] {
    return &baseRepository[T]{db: db}
}

// Search returns all entities that match the given condition.
func (r *baseRepository[T]) Search(condition string, pageNumber int, pageSize int, args ...interface{}) ([]T, int64, error) {
    var entities []T

    var count int64
    if err := r.db.
        Model(new(T)).
        Where(condition, args...).
        Count(&count).Error; err != nil {
        return nil, 0, err
    }

    query := r.db.Where(condition, args...)

    // Apply pagination if valid values are provided.
    if pageNumber > 0 && pageSize > 0 {
        offset := (pageNumber - 1) * pageSize
        query = query.Offset(offset).Limit(pageSize)
    }

    if err := query.Find(&entities).Error; err != nil {
        return nil, 0, err
    }
    
    return entities, count, nil
}

func (r *baseRepository[T]) GetByID(id string) (*T, error) {
    var entity T
    if err := r.db.First(&entity, "id = ?", id).Error; err != nil {
        return nil, err
    }
    return &entity, nil
}

func (r *baseRepository[T]) GetAll() ([]T, error) {
    var entities []T
    if err := r.db.Find(&entities).Error; err != nil {
        return nil, err
    }
    return entities, nil
}

func (r *baseRepository[T]) Create(entity *T) error {
    return r.db.Create(entity).Error
}

func (r *baseRepository[T]) Update(entity *T, id string) error {
    return r.db.Where("id = ?", id).Updates(entity).Error
}

func (r *baseRepository[T]) Delete(id string) error {
    var entity T
    return r.db.Delete(&entity, "id = ?", id).Error
}

func (r *baseRepository[T]) First(condition string, args ...interface{}) (*T, error) {
    var entity T
    if err := r.db.Where(condition, args...).First(&entity).Error; err != nil {
        return nil, err
    }
    return &entity, nil
}
