package authorization

import (
	"encoding/base64"
	"errors"
	"strings"
	"time"

	"authentication/src/core/domain/account"
	"authentication/src/core/messages"

	"github.com/golang-jwt/jwt"
)

var SecretKey string

func GenerateTokenFromAccount(account account.Account) (string, error) {
	claims := buildClaimsFromAccount(account)
	token, err := getSignedTokenFromClaims(*claims)
	if err != nil {
		return "", errors.New(messages.UnauthorizedErr)
	}

	return token, nil
}

func IsSessionValid(authToken string) bool {
	token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		return false
	}

	return token.Valid && token.Claims.Valid() == nil
}

func buildClaimsFromAccount(account account.Account) *jwt.MapClaims {
	roleNames := make([]string, 0)
	for _, role := range account.Roles() {
		roleNames = append(roleNames, role.Name())
	}
	concatenatedRoles := strings.Join(roleNames, ",")
	return &jwt.MapClaims{
		messages.AccountIDClaim:    account.ID().String(),
		messages.ExpirationClaim:   time.Now().Add(time.Minute * 60).Unix(),
		messages.IssuedAtClaim:     time.Now().Unix(),
		messages.TokenTypeClaim:    messages.BearerTokenType,
		messages.AccountRolesClaim: base64.StdEncoding.EncodeToString([]byte(concatenatedRoles)),
	}
}

func getSignedTokenFromClaims(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
