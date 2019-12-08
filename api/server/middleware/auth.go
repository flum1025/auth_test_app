package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

var jwks = Jwks{}

func getPemCert(token *jwt.Token) (string, error) {
	cert := ""

	for k, _ := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		err := errors.New("Unable to find appropriate key.")
		return cert, err
	}

	return cert, nil
}

func getTokenPartFromHeader(tokenHeader string) (string, error) {
	splitted := strings.Split(tokenHeader, " ")
	if len(splitted) != 2 {
		return "", errors.New("invalid Authorization header")
	}

	return splitted[1], nil
}

func GetParsedToken(tokenHeader string) jwt.MapClaims {
	tokenPart, err := getTokenPartFromHeader(tokenHeader)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	token, err := jwt.Parse(tokenPart, func(token *jwt.Token) (interface{}, error) {
		cert, err := getPemCert(token)
		if err != nil {
			return token, err
		}

		return jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
	})

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return token.Claims.(jwt.MapClaims)
}

func checkJwt(claimsMap jwt.MapClaims, aud string, iss string) bool {
	checkAud := claimsMap.VerifyAudience(aud, false)
	if !checkAud {
		return false
	}

	checkIss := claimsMap.VerifyIssuer(iss, false)
	if !checkIss {
		return false
	}

	return true
}

func updateCache(jwksUrl string, interval time.Duration) {
	for {
		func() {
			resp, err := http.Get(jwksUrl)
			fmt.Println(resp)

			if err != nil {
				fmt.Println(err)
				return
			}

			defer resp.Body.Close()

			json.NewDecoder(resp.Body).Decode(&jwks)
			fmt.Println(jwks)
		}()

		fmt.Println("jwks updated")
		time.Sleep(interval)
	}
}

func parseRoles(claimsMap jwt.MapClaims) []string {
	authorization, ok := claimsMap["test_auth_app_authorization"]

	var parsedRoles = []string{}

	if !ok {
		return parsedRoles
	}

	roles, ok := authorization.(map[string]interface{})["roles"]

	if !ok {
		return parsedRoles
	}

	for _, role := range roles.([]interface{}) {
		parsedRoles = append(parsedRoles, role.(string))
	}

	return parsedRoles
}

func GenerateMiddleware() func(http.Handler) http.Handler {
	jwksUrl := os.Getenv("JWKS_URL")
	aud := os.Getenv("AUDIENCE")
	iss := os.Getenv("ISSUER")

	if jwksUrl == "" || aud == "" || iss == "" {
		panic("JWKS_URL, AUDIENCE, ISSUER environments is required")
	}

	go updateCache(jwksUrl, 1*time.Hour)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			claimsMap := GetParsedToken(token)
			isValid := checkJwt(claimsMap, aud, iss)

			if !isValid {
				w.WriteHeader(400)
				return
			}

			roles := parseRoles(claimsMap)

			ctx := context.WithValue(r.Context(), "roles", roles)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
