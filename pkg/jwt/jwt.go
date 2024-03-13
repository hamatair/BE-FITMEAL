package jwt

import (
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

	token := jwt.NewWithClaims(jwt.SigningMethodPS256, claim)

	tokenString, err := token.SignedString([]byte(j.SecretKey))
	if err != nil{
		return tokenString, err
	}

	return tokenString, nil
}

// GetLoginUser implements Interface.
func (*jsonWebToken) GetLoginUser(c *gin.Context) (entity.User, error) {
	panic("unimplemented")
}

// ValidateToken implements Interface.
func (*jsonWebToken) ValidateToken(token string) (uuid.UUID, error) {
	panic("unimplemented")
}