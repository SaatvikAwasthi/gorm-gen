package repository

import (
	"log"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gen"
	"gorm.io/gorm"

	"gorm-gen/contract"
	"gorm-gen/dao"
)

type connection struct {
	gormDB *gorm.DB
}

type Connection interface {
	Init()        // Initialize the connection
	DB() *gorm.DB // Get the GORM DB connection
	Gen()         // Generate code using GORM Gen
}

// NewConnection creates a new connection instance for GORM Gen code generation.
func NewConnection() *connection {
	return &connection{
		gormDB: nil, // Initialize with nil, will be set in Init()
	}
}

func (c *connection) Init() {
	// We'll use an in-memory SQLite database for generation,
	// but you can connect to any database supported by GORM.
	gormDB, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database for generation: %v\n", err)
	}

	// Auto-migrate the tables in the in-memory database.
	err = gormDB.AutoMigrate(&dao.User{}, &dao.Product{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate tables for generation: %v\n", err)
	}

	c.gormDB = gormDB
}

func (c *connection) DB() *gorm.DB {
	// Return the GORM DB connection
	if c.gormDB == nil {
		log.Fatalf("GORM DB connection is not initialized. Call Init() first.")
	}
	return c.gormDB
}

func (c *connection) Gen() {
	// Configure the GORM Gen generator
	config := gen.Config{
		OutPath:           filepath.Join(".", "generated"),
		OutFile:           "generated_gen.go",
		ModelPkgPath:      filepath.Join(".", "generated", "model"),
		WithUnitTest:      true,
		FieldNullable:     true,
		FieldCoverable:    false,
		FieldSignable:     false,
		FieldWithIndexTag: false,
		FieldWithTypeTag:  false,
		Mode:              gen.WithDefaultQuery | gen.WithQueryInterface,
	}

	g := gen.NewGenerator(config)

	// Set the database connection for the generator
	g.UseDB(c.gormDB)

	// Apply basic code generation for our models
	// It will generate model structs in `generated/model` and query methods in `generated/query`.
	g.ApplyBasic(
		dao.User{},    // Generate code for the User model
		dao.Product{}, // Generate code for the Product model
	)

	g.ApplyInterface(func(contract.Querier) {}, dao.User{}) // Generate code for the Querier interface

	// Execute the code generation
	g.Execute()

	log.Println("GORM Gen code generated successfully in the 'generated' directory!")
	log.Println("You can now use the generated code in your application.")

	return
}
