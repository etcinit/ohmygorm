package main

import (
	"github.com/etcinit/ohmygorm"
	"github.com/facebookgo/inject"
	"github.com/jacobstr/confer"
	"github.com/kr/pretty"
)

// User is a simple example user model
type User struct {
	ID    int
	Name  string
	Email string
}

// Post is a simple blog post
type Post struct {
	ID       int
	Author   User
	AuthorID int
	Title    string
	Content  string
}

func main() {
	// Create the configuration
	// In this case, we will be using the environment and some safe defaults
	config := confer.NewConfig()
	config.SetDefault("database.driver", "sqlite")
	config.SetDefault("database.file", ":memory:")
	config.AutomaticEnv()

	// Next, we setup the dependency graph
	// In this example, the graph won't have many nodes, but on more complex
	// applications it becomes more useful.
	var g inject.Graph
	var connections ohmygorm.ConnectionsService
	var migrator ohmygorm.MigrationsService
	g.Provide(
		&inject.Object{Value: config},
		&inject.Object{Value: &connections},
		&inject.Object{Value: &migrator},
	)
	g.Populate()

	// At this point, the DI library has automatically set the dependencies of
	// both structs (connections and migrator), so we can start using them!

	// Run migrations
	migrator.Run([]interface{}{&User{}, &Post{}})

	// Get a connection
	db, err := connections.Make()

	if err != nil {
		panic(err)
	}

	// Create a user
	user := User{Name: "Bobby Tables", Email: "bobby@tables.inc"}
	db.Create(&user)

	// Create a post
	post := Post{
		Author:  user,
		Title:   "SQL injection is bad for you",
		Content: "Its bad. Really bad.",
	}
	db.Create(&post)

	// Query all posts
	var posts []Post
	db.Preload("Author").Find(&posts)

	pretty.Println(posts)
}
