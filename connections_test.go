package ohmygorm

import (
	"os"
	"strconv"
	"testing"

	"github.com/jacobstr/confer"
	"github.com/stretchr/testify/assert"
)

func Test_MakeSqlite(t *testing.T) {
	config := confer.NewConfig()

	config.Set("database.driver", "sqlite")
	config.Set("database.file", ":memory:")

	connections := ConnectionsService{
		Config: config,
	}

	connectionOne, err := connections.Make()
	connectionTwo, errTwo := connections.Make()

	assert.Equal(t, &connectionOne, &connectionTwo)
	assert.Nil(t, err)
	assert.Nil(t, errTwo)
}

func Test_MakeMysql(t *testing.T) {
	// Only run this test if we provide the correct config
	if os.Getenv("TEST_MYSQL") == "" {
		return
	}

	config := confer.NewConfig()
	config.Set("database.driver", "mysql")
	config.Set("database.host", os.Getenv("WERCKER_MYSQL_HOST"))

	if port, err := strconv.Atoi(os.Getenv("WERCKER_MYSQL_PORT")); err == nil {
		config.Set("database.port", port)
	}

	config.Set("database.username", os.Getenv("WERCKER_MYSQL_USERNAME"))
	config.Set("database.password", os.Getenv("WERCKER_MYSQL_PASSWORD"))
	config.Set("database.name", os.Getenv("WERCKER_MYSQL_DATABASE"))

	connections := ConnectionsService{
		Config: config,
	}

	connectionOne, err := connections.Make()
	connectionTwo, errTwo := connections.Make()

	assert.Equal(t, &connectionOne, &connectionTwo)
	assert.Nil(t, err)
	assert.Nil(t, errTwo)
}

func Test_MakeInvalid(t *testing.T) {
	config := confer.NewConfig()

	config.Set("database.driver", "nope")
	config.Set("database.file", ":memory:")

	connections := ConnectionsService{
		Config: config,
	}

	_, err := connections.Make()

	assert.NotNil(t, err)
}
