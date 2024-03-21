package jwt

import (
	"errors"
	"intern-bcc/entity"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Interface interface {
	CreateJWTToken(userID uuid.UUID) (string, error)
	ValidateToken(token string) (uuid.UUID, error)
	GetLoginUser(c *gin.Context) (entity.User, error)
}

type jsonWebToken struct {
	SecretKey   string
	ExpiredTime time.Duration
}

type Claims struct {
	UserID uuid.UUID
	jwt.RegisteredClaims
}

func Init() Interface {
	secretkey := os.Getenv("JWT_SECRET_KEY")
	expiredTime, err := strconv.Atoi(os.Getenv("JWT_EXP_TIME"))
	if err != nil {
		log.Fatal("failed set expired time for jwt : ", err.Error())
	}

	return &jsonWebToken{
		SecretKey:   secretkey,
		ExpiredTime: time.Duration(expiredTime) * time.Hour,
	}
}

// CreateJWTToken implements Interface.
func (j *jsonWebToken) CreateJWTToken(userID uuid.UUID) (string, error) {
	claim := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.ExpiredTime)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenString, err := token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return tokenString, err
	}

	return tokenString, nil
}

// GetLoginUser implements Interface.
func (*jsonWebToken) GetLoginUser(c *gin.Context) (entity.User, error) {
	user, ok := c.Get("user")
	if !ok {
		return entity.User{}, errors.New("failed to get user")
	}

	return user.(entity.User), nil
}

// ValidateToken implements Interface.
func (j *jsonWebToken) ValidateToken(tokenString string) (uuid.UUID, error) {
	var (
		claims Claims
		userId uuid.UUID
	)

	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.SecretKey), nil
	})
	if err != nil {
		return userId, err
	}

	if !token.Valid {
		return userId, err
	}

	userId = claims.UserID

	return userId, nil
}
