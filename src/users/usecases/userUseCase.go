package usecases

import (
	"context"
	"errors"

	"github.com/docker-cli-golang-lab/models"
	"github.com/docker-cli-golang-lab/src/users/domains"
	"golang.org/x/crypto/bcrypt"
)

// UserUseCase implements the domains.UserUseCase interface
type UserUseCase struct {
	userRepo domains.UserRepository
}

// NewUserUseCase creates a new UserUseCase
func NewUserUseCase(userRepo domains.UserRepository) domains.UserUseCase {
	return &UserUseCase{
		userRepo: userRepo,
	}
}

// GetUsers returns all users
func (uc *UserUseCase) GetUsers(ctx context.Context) ([]models.User, error) {
	return uc.userRepo.FindAll(ctx)
}

// GetUserByID returns a user by ID
func (uc *UserUseCase) GetUserByID(ctx context.Context, id uint) (*models.User, error) {
	return uc.userRepo.FindByID(ctx, id)
}

// CreateUser creates a new user
func (uc *UserUseCase) CreateUser(ctx context.Context, user *models.User) error {
	// Validate user data
	if user.Username == "" || user.Email == "" || user.Password == "" {
		return errors.New("username, email, and password are required")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}
	user.Password = string(hashedPassword)

	// Set default values if not provided
	if user.Role == "" {
		user.Role = "user"
	}
	user.Active = true

	return uc.userRepo.Create(ctx, user)
}

// UpdateUser updates an existing user
func (uc *UserUseCase) UpdateUser(ctx context.Context, user *models.User) error {
	// Get existing user
	existingUser, err := uc.userRepo.FindByID(ctx, user.ID)
	if err != nil {
		return err
	}

	// Update password if provided
	if user.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return errors.New("failed to hash password")
		}
		user.Password = string(hashedPassword)
	} else {
		// Keep existing password
		user.Password = existingUser.Password
	}

	return uc.userRepo.Update(ctx, user)
}

// DeleteUser deletes a user
func (uc *UserUseCase) DeleteUser(ctx context.Context, id uint) error {
	return uc.userRepo.Delete(ctx, id)
}

// AuthenticateUser authenticates a user by username and password
func (uc *UserUseCase) AuthenticateUser(ctx context.Context, username, password string) (*models.User, error) {
	// Find user by username
	user, err := uc.userRepo.FindByUsername(ctx, username)
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	// Check if user is active
	if !user.Active {
		return nil, errors.New("user account is inactive")
	}

	// Verify password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	return user, nil
}
