package middlewares

import (
	config "AuthInGo/config/db"
	env "AuthInGo/config/env"
	repo "AuthInGo/db/repositories"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

// contextKey is an unexported type for context keys in this package,
// preventing collisions with keys from other packages.
type contextKey string

const (
	ContextKeyUserID contextKey = "userId"
	ContextKeyEmail  contextKey = "email"
)

func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Authorization header must start with Bearer", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			http.Error(w, "Token is required", http.StatusUnauthorized)
			return
		}

		claims:=jwt.MapClaims{}

		_,err:=jwt.ParseWithClaims(token,&claims,func(t *jwt.Token) (any, error) {
			return []byte(env.GetString("JWT_SECRET","TOKEN")),nil
		})

		if err!=nil{
			http.Error(w,"Invalid token: "+err.Error(),http.StatusUnauthorized)
			return 
		}

		userId,okId:=claims["id"].(float64)

		email,okEmail:=claims["email"].(string)

		if !okId || !okEmail {
			http.Error(w,"Invalid token claims",http.StatusUnauthorized)
			return 
		}

		fmt.Println("Authenticated user ID:",userId,"Email:",email)

		ctx:=context.WithValue(r.Context(),ContextKeyUserID,strconv.FormatFloat(userId,'f',0,64))
		ctx=context.WithValue(ctx,ContextKeyEmail,email)
		next.ServeHTTP(w,r.WithContext(ctx))
	})
}

func RequireAllRoles(roles ...string) func(http.Handler) http.Handler{
	
	// function that can create a middleware for checking the above set of roles
	
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			userIdStr,ok:=r.Context().Value(ContextKeyUserID).(string)
			if !ok || userIdStr=="" {
				http.Error(w,"Unauthorized",http.StatusUnauthorized)
				return
			}

			userId,err:=strconv.ParseInt(userIdStr,10,64)
			if err!=nil{
				http.Error(w,"Invalid user ID",http.StatusBadRequest)
				return 
			}

			dbConn:=config.Db

			if dbConn==nil{
				http.Error(w,"Database connection error",http.StatusInternalServerError)
				return 
			}

			urr:=repo.NewUserRoleRepository(dbConn)

			allRoles,err:=urr.HasAllRoles(userId,roles)

			if err!=nil{
				http.Error(w,"Error checking roles: "+err.Error(),http.StatusInternalServerError)
				return 
			}
			if !allRoles {
				http.Error(w,"User does not have all required roles",http.StatusForbidden)
				return 
			}

			fmt.Println("User has all required roles:",roles)

			next.ServeHTTP(w,r)
		})
	}
}