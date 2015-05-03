package ohmygorm

import "github.com/jinzhu/gorm"

// RepositoryService provides function for building structures following the
// repository pattern.
type RepositoryService struct {
	Connections *ConnectionsService `inject:""`
}

// Exists checks if a model instance exists in the database by its ID
func (r *RepositoryService) Exists(model interface{}, id int) bool {
	db, err := r.Connections.Make()

	if err != nil {
		return false
	}

	itemCount := 0
	db.Model(model).Where("id = ?", id).Count(&itemCount)

	return itemCount > 0
}

// Find attempts to find a specific model instance by its ID.
func (r *RepositoryService) Find(model interface{}, id int) error {
	db, err := r.Connections.Make()

	if err != nil {
		return err
	}

	if err := db.Where("id = ?", id).Find(model).Error; err != nil {
		return err
	}

	return nil
}

// FirstOrFail is a shortcut for using First and checking for the RecordNotFound error
func (r *RepositoryService) FirstOrFail(model interface{}, query *gorm.DB) error {
	if err := query.First(model).Error; err != nil {
		return err
	}

	return nil
}
