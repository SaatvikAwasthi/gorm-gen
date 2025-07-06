package service

import (
	"context"
	"log"

	"gorm-gen/dao"
	query "gorm-gen/generated"
	"gorm-gen/repository"
)

const (
	name        = "Alice"
	email       = "alice@example.com"
	updatedName = "Alicia"
	name2       = "Bob"
	email2      = "bob@example.com"
)

func UserService(connect repository.Connection) {
	db := connect.DB()
	ctx := context.Background()

	// Initialize the generated query objects
	queryUser := query.Use(db).User
	queryUserWithContext := queryUser.WithContext(ctx)

	// --- C: Create ---
	user := &dao.User{Name: name, Email: email}
	err := queryUserWithContext.Create(user) // Type-safe create
	if err != nil {
		log.Fatalf("Error creating user: %v", err)
	}
	log.Printf("Created user: %v\n", user)

	// --- R: Read (Type-Safe Queries) ---
	foundUser, err := queryUserWithContext.Where(queryUser.Email.Eq(email)).First()
	if err != nil {
		log.Fatalf("Error finding user: %v", err)
	}
	log.Printf("Found user: %s\n", foundUser.Name)

	// --- U: Update ---
	_, err = queryUserWithContext.Where(queryUser.ID.Eq(user.ID)).Update(queryUser.Name, updatedName)
	if err != nil {
		log.Fatalf("Error updating user: %v", err)
	}
	log.Printf("Updated user: %s to %s\n", user.Name, updatedName)

	// --- Transactions ---
	if err = query.Use(db).Transaction(func(tx *query.Query) error {
		// All operations within this block are part of the transaction
		newUserInTx := &dao.User{Name: name2, Email: email2}
		if err := tx.User.WithContext(ctx).Create(newUserInTx); err != nil {
			log.Fatalf("Error creating user in transaction: %v", err)
			return err // Rollback the transaction if there's an error
		}
		return nil
	}); err != nil {
		log.Fatalf("Transaction failed: %v", err)
	}
	log.Println("Transaction succeeded")

	// --- Preloading (Relations) ---
	usersWithProducts, err := queryUserWithContext.Preload(queryUser.Products).Find()
	if err != nil {
		log.Fatalf("Error preloading users with products: %v", err)
	}
	for _, u := range usersWithProducts {
		log.Printf("User %s has %d products.\n", u.Name, len(u.Products))
	}

	// --- D: Delete ---
	_, err = queryUserWithContext.Where(queryUser.ID.Eq(user.ID)).Delete()
	if err != nil {
		log.Fatalf("Error deleting user: %v", err)
	}
	log.Printf("Deleted user: %s\n", user.Name)
}
