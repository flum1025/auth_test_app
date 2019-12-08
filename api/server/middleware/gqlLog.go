package middleware

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func GqlLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			panic(err)
		}

		r.Body = ioutil.NopCloser(bytes.NewReader(body))

		token := r.Header.Get("Authorization")
		userEmail := GetParsedToken(token)["email"].(string)

		log.Println("gql query from (" + userEmail + "): " + strings.Replace(string(body), `\n`, "\n", -1))
		next.ServeHTTP(w, r)
	})
}
