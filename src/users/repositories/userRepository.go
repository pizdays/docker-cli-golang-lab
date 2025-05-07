package repositories

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/docker-cli-golang-lab/models"
	"github.com/docker-cli-golang-lab/src/users/domains"
	"github.com/docker-cli-golang-lab/src/users/handlers"
	"github.com/docker-cli-golang-lab/src/users/usecases"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserRepository is a mock implementation of domains.UserRepository
type Repository struct {
	users  map[uint]*models.User
	lastID uint
	mutex  sync.RWMutex
	// แค่เก็บ db ไว้แต่ไม่ใช้ เพราะตอนนี้เรายังใช้ mock data
	db *gorm.DB
}

// init Repository Handler
func NewRepositoryHandler(conn *gorm.DB) *handlers.UserHandler {
	useCase := usecases.NewUserUseCase(NewUserRepository(conn))
	handler := handlers.NewUserHandler(useCase)
	return handler
}

// NewUserRepository creates a new UserRepository with mock data
// รับ db เข้ามาเพื่อรองรับการใช้งานจริงในอนาคต แต่ตอนนี้ยังไม่ได้ใช้
func NewUserRepository(db *gorm.DB) domains.UserRepository {
	hashedPassword1, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	hashedPassword2, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)

	repo := &Repository{
		users:  make(map[uint]*models.User),
		lastID: 0,
		db:     db, // อาจเป็น nil ได้ตอนนี้ไม่ได้ใช้
	}

	// Add mock data
	mockUsers := []models.User{
		{
			Username:  "johndoe",
			Email:     "john.doe@example.com",
			Password:  string(hashedPassword1),
			FirstName: "John",
			LastName:  "Doe",
			Role:      "user",
			Active:    true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Username:  "janedoe",
			Email:     "jane.doe@example.com",
			Password:  string(hashedPassword1),
			FirstName: "Jane",
			LastName:  "Doe",
			Role:      "user",
			Active:    true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Username:  "admin",
			Email:     "admin@example.com",
			Password:  string(hashedPassword2),
			FirstName: "Admin",
			LastName:  "User",
			Role:      "admin",
			Active:    true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	for _, user := range mockUsers {
		repo.lastID++
		user.ID = repo.lastID
		repo.users[user.ID] = &user
	}

	return repo
}

// FindAll returns all users
func (r *Repository) FindAll(ctx context.Context) ([]models.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	users := make([]models.User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, *user)
	}
	return users, nil
}

// FindByID finds a user by ID
func (r *Repository) FindByID(ctx context.Context, id uint) (*models.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	user, ok := r.users[id]
	if !ok {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// FindByUsername finds a user by username
func (r *Repository) FindByUsername(ctx context.Context, username string) (*models.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for _, user := range r.users {
		if user.Username == username {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

// FindByEmail finds a user by email
func (r *Repository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

// Create adds a new user
func (r *Repository) Create(ctx context.Context, user *models.User) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// Check if username already exists
	for _, existingUser := range r.users {
		if existingUser.Username == user.Username {
			return errors.New("username already exists")
		}
		if existingUser.Email == user.Email {
			return errors.New("email already exists")
		}
	}

	r.lastID++
	user.ID = r.lastID
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	r.users[user.ID] = user
	return nil
}

// Update updates an existing user
func (r *Repository) Update(ctx context.Context, user *models.User) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, ok := r.users[user.ID]; !ok {
		return errors.New("user not found")
	}

	// Check if username already exists (but not by this user)
	for id, existingUser := range r.users {
		if id != user.ID {
			if existingUser.Username == user.Username {
				return errors.New("username already exists")
			}
			if existingUser.Email == user.Email {
				return errors.New("email already exists")
			}
		}
	}

	user.UpdatedAt = time.Now()
	r.users[user.ID] = user
	return nil
}

// Delete deletes a user
func (r *Repository) Delete(ctx context.Context, id uint) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, ok := r.users[id]; !ok {
		return errors.New("user not found")
	}

	delete(r.users, id)
	return nil
}
