package services

import (
	"errors"
	"strings"

	"github.com/go-nine9/go-nine9/db"
	"github.com/go-nine9/go-nine9/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const SecretKey = "secret"

var ErrUnauthorized = errors.New("accès non autorisé")

func GenerateJWT(user models.User) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.ID,
		"exp":  time.Now().Add(time.Hour * 24).Unix(), // 1 jour
		"role": user.Roles,
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}

	return token, nil
}

func VerifyPassword(hashedPassword string, password string) error {

	if strings.HasPrefix(hashedPassword, "$2a$") {

		return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	}

	if hashedPassword != password {
		return errors.New("mot de passe incorrect")
	}

	return nil
}

func ParseJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}

func RegisterUser(user models.User) error {
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	user.Roles = "user"
	return db.DB.Create(&user).Error
}

// FindUserByEmail trouve un utilisateur par e-mail
func FindUserByEmail(email string) (models.User, error) {
	var user models.User
	if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// FindUserByID trouve un utilisateur par ID
func FindUserByID(userID string) (models.User, error) {
	var user models.User
	if err := db.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// AddUserByAdmin permet à un administrateur d'ajouter un autre utilisateur
func AddUserByAdmin(newUser models.User) error {
	if newUser.Roles == "admin" {

	}

	return nil
}

// hashPassword hache le mot de passe avec bcrypt
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
