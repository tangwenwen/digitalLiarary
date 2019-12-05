package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/emicklei/go-restful"
)

func TokenFilter(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	tokenStr := req.Request.Header.Get("Authorization")
	if tokenStr == "" {
		resp.WriteHeader(401)
	} else {
		token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				resp.WriteHeader(401)
				return nil, fmt.Errorf("not authorization")
			}
			return []byte("hello"), nil
		})
		if !token.Valid {
			resp.WriteHeader(401)
		} else {
			chain.ProcessFilter(req, resp)
		}
	}
}
