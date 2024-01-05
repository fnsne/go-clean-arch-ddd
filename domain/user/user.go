package user

// User is a entity in Domain Driven Design
type User struct {
	ID           string
	Name         string
	Email        string
	Status       string
	HashPassword string
}
