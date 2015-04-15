package ohmygorm

// MigrationsService provides migration-related functions
type MigrationsService struct {
	Connections *ConnectionsService `inject:""`
}

// Run automatically runs migrations for the database. Note that some
// changes will not be applied automatically if they destroy data, such as
// removing a table or column.
func (m *MigrationsService) Run(models []interface{}) error {
	conn, err := m.Connections.Make()

	if err != nil {
		return err
	}

	conn.AutoMigrate(models...)
	return nil
}
