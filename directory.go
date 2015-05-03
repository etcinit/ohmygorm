package ohmygorm

import (
	"fmt"

	"github.com/codegangsta/cli"
)

// A ModelDirectory lists an application models, which can be used for
// migrations and administration of models.
type ModelDirectory interface {
	GetModels() []interface{}
}

// DirectoryMigratorService provides the migration command
type DirectoryMigratorService struct {
	Migrations  *MigrationsService  `inject:""`
	Connections *ConnectionsService `inject:""`
	Directory   ModelDirectory      `inject:""`
}

// Run performs all the migrations for the models listed in the directory.
func (m *DirectoryMigratorService) Run() error {
	return m.Migrations.Run(m.Directory.GetModels())
}

// RunCommand performs all the migrations for the models listed in the
// directory and provides some feedback through standard output.
func (m *DirectoryMigratorService) RunCommand(c *cli.Context) {
	fmt.Println("Running migrations...")

	db, _ := m.Connections.Make()
	db.LogMode(true)

	m.Migrations.Run(m.Directory.GetModels())

	fmt.Println("✔︎ Done!")
}
