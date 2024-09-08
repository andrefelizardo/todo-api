package domain

// User represents a user in the system
type User struct {
	ID       UUID
	Name string
	Email string
	Password string
	CreatedAt string
}