package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	d "todo/internal/db"
)

func ValidatePassword(password string) error {
	switch {
	case len(password) < 6:
		return errors.New("password must be at least 6 characters")
	case len(password) > 72:
		return errors.New("password must be less than 72 characters")
	default:
		return nil
	}
}

func ValidateUsername(username string) error {
	switch {
	case len(username) < 3:
		return errors.New("username must be at least 3 characters")
	case len(username) > 32:
		return errors.New("username must be less than 32 characters")
	default:
		return nil
	}
}

func Register(username, password string) error {
	db, err := d.GetDB()
	if err != nil {
		return err
	}

	if err := ValidateUsername(username); err != nil {
		return err
	}
	if err := ValidatePassword(password); err != nil {
		return err
	}


	var existingUser User
	if err := db.Where("username = ?", username).First(&existingUser).Error; err == nil {
		return errors.New("username already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user := User{
		Username: username,
		Password: string(hash),
	}

	return db.Create(&user).Error
}

func Login(username, password string) (*User, error) {
	db, err := d.GetDB()
	if err != nil {
		return nil, err
	}

	var user User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	return &user, nil
}
