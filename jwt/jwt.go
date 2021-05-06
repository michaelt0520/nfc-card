package jwt

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/michaelt0520/nfc-card/models"
)

const (
	// Expired : expired time for jwt
	Expired = time.Hour * 24 * 30
)

// PayLoad : struct payload jwt
type PayLoad struct {
	CompanyID uint
	UserID    uint
	UserRole  models.UserRole
	UserType  models.CardType
}

// CreateToken : ...
func CreateToken(user *models.User) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["company_id"] = user.CompanyID
	atClaims["user_id"] = user.ID
	atClaims["user_role"] = user.Role
	atClaims["user_type"] = user.Type
	atClaims["exp"] = time.Now().Add(Expired).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("jwt_key")))
	if err != nil {
		return "", err
	}
	return token, nil
}

// ExtractToken : extract valid token
func ExtractToken(tokenString string) (*PayLoad, error) {
	token, err := ValidateToken(tokenString)
	if err != nil {
		return nil, err
	}
	claim, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		companyID, _ := strconv.ParseUint(fmt.Sprintf("%.f", claim["company_id"]), 10, 32)

		userID, err := strconv.ParseUint(fmt.Sprintf("%.f", claim["user_id"]), 10, 32)
		if err != nil {
			return nil, err
		}

		userRole, err := strconv.ParseUint(fmt.Sprintf("%.f", claim["user_role"]), 10, 32)
		if err != nil {
			return nil, err
		}

		userType, err := strconv.ParseUint(fmt.Sprintf("%.f", claim["user_type"]), 10, 32)
		if err != nil {
			return nil, err
		}

		return &PayLoad{
			CompanyID: uint(companyID),
			UserID:    uint(userID),
			UserRole:  models.UserRole(userRole),
			UserType:  models.CardType(userType),
		}, nil
	}

	return nil, err
}

// ValidateToken : validate token
func ValidateToken(token string) (*jwt.Token, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("jwt_key")), nil
	})

	if err != nil {
		return nil, err
	}

	return parsedToken, nil
}
