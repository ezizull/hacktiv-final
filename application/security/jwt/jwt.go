// Package jwt implements the JWT authentication
package jwt

import (
	"errors"
	"fmt"
	errorDomain "hacktiv/final-project/domain/errors"
	secureDomain "hacktiv/final-project/domain/security"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// GenerateJWTToken generates a JWT token (refresh or access)
func GenerateJWTToken(userID int, tokenType string, roleName string, CSRF string) (appToken *secureDomain.AppToken, err error) {
	tokenTimeUnix, err := getTimeExpire(tokenType)
	if err != nil {
		return
	}

	nowTime := time.Now()
	expirationTokenTime := nowTime.Add(tokenTimeUnix)

	tokenClaims := &secureDomain.Claims{
		UserID: userID,
		Type:   tokenType,
		CSRF:   CSRF,
		Role:   roleName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTokenTime.Unix(),
			IssuedAt:  nowTime.UTC().Unix(),
		},
	}
	tokenWithClaims := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), tokenClaims)

	// Sign and get the complete encoded token as a string using the secret
	tokenStr, err := tokenWithClaims.SignedString(secureDomain.PrivateKey)
	if err != nil {
		return
	}

	secureDomain.CSRF = &CSRF
	appToken = &secureDomain.AppToken{
		Token:          tokenStr,
		TokenType:      tokenType,
		ExpirationTime: expirationTokenTime,
	}

	return
}

// GetClaimsAndVerifyToken verifies the token and returns the claims
func GetClaimsAndVerifyToken(tokenString string, tokenType string) (claims jwt.MapClaims, err error) {
	if secureDomain.CSRF == nil {
		return nil, errorDomain.NewAppError(errors.New("please login again"), errorDomain.NotAuthenticated)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			message := fmt.Sprintf("unexpected signing method: %v", token.Header["alg"])
			return nil, errorDomain.NewAppError(errors.New(message), errorDomain.NotAuthenticated)
		}

		return secureDomain.PublicKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["csrf"] != *secureDomain.CSRF {
			return nil, errorDomain.NewAppError(errors.New("invalid csrf"), errorDomain.NotAuthenticated)
		}

		if claims["type"] != tokenType {
			return nil, errorDomain.NewAppError(errors.New("invalid token type"), errorDomain.NotAuthenticated)
		}

		var timeExpire = claims["exp"].(float64)
		if time.Now().Unix() > int64(timeExpire) {
			return nil, errorDomain.NewAppError(errors.New("token expired"), errorDomain.NotAuthenticated)
		}

		return claims, nil
	}

	return nil, err
}

// ReGenerateCustomJWT regenerate jwt with custom modified data
func ReGenerateCustomJWT(tokenString string, oldClaims *secureDomain.Claims) (newClaims *secureDomain.Claims, err error) {
	_, err = GetClaimsAndVerifyToken(tokenString, oldClaims.Type)
	if err != nil {
		err = errorDomain.NewAppError(errors.New("error generate jwt"), errorDomain.TokenGeneratorError)
		return
	}

	if oldClaims.UserID == 0 {
		err = errorDomain.NewAppError(errors.New("error generate jwt"), errorDomain.TokenGeneratorError)
		return
	}

	if oldClaims.Type == "" {
		err = errorDomain.NewAppError(errors.New("error generate jwt"), errorDomain.TokenGeneratorError)
		return
	}

	if oldClaims.Role == "" {
		err = errorDomain.NewAppError(errors.New("error generate jwt"), errorDomain.TokenGeneratorError)
		return
	}

	newClaims.Type = oldClaims.Type
	newClaims.UserID = oldClaims.UserID
	newClaims.Role = oldClaims.Role

	if oldClaims.CSRF != "" {
		newClaims.CSRF = oldClaims.CSRF
	} else {
		var newCSRF string
		newCSRF, err = secureDomain.GenerateCSRF(32)

		if err != nil {
			err = errorDomain.NewAppError(errors.New(newCSRF), errorDomain.NotAuthenticated)
			return
		}

		newClaims.CSRF = newCSRF
	}

	// jwt.StandardClaims
	if oldClaims.Audience != "" {
		newClaims.Audience = oldClaims.Audience
	}

	if oldClaims.ExpiresAt != 0 {
		newClaims.ExpiresAt = oldClaims.ExpiresAt
	} else {
		var tokenTimeUnix time.Duration
		tokenTimeUnix, err = getTimeExpire(oldClaims.Type)

		if err != nil {
			err = errorDomain.NewAppError(errors.New("error generate jwt"), errorDomain.TokenGeneratorError)
			return
		}

		newClaims.ExpiresAt = time.Now().Add(tokenTimeUnix).Unix()
	}

	if oldClaims.IssuedAt != 0 {
		newClaims.IssuedAt = oldClaims.IssuedAt
	}

	if oldClaims.NotBefore != 0 {
		newClaims.NotBefore = oldClaims.NotBefore
	}

	if oldClaims.NotBefore != 0 {
		newClaims.NotBefore = oldClaims.NotBefore
	}

	if oldClaims.Issuer != "" {
		newClaims.Issuer = oldClaims.Issuer
	}

	if oldClaims.Subject != "" {
		newClaims.Subject = oldClaims.Subject
	}

	tokenWithClaims := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), newClaims)

	// Sign and get the complete encoded token as a string using the secret
	_, err = tokenWithClaims.SignedString(secureDomain.PrivateKey)
	secureDomain.CSRF = &newClaims.CSRF

	return
}
