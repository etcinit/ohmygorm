package ohmygorm

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

	db.Where("id = ?", id).Find(model)

	return nil
}
