package authentication

import (
	"context"
	"gqlGenTutorial/connection"
	"gqlGenTutorial/dataaccess/repositories/userrepositories"
	"gqlGenTutorial/models/usermodel"
	"gqlGenTutorial/settings"
	"net/http"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			// Allow unauthenticated users in
			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			//validate jwt token
			tokenStr := header
			username, err := ParseToken(tokenStr)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			// create user and check if user exists in db
			user := usermodel.UserModel{UserName: username}
			// Pedimos una conexion a la base de datos POSTGRES
			db := connection.OpenConnection(settings.PostgresDB)
			userrepositories.RetrieveUser(&user, db)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			// put it in context
			ctx := context.WithValue(r.Context(), userCtxKey, &user)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *usermodel.UserModel {
	raw, _ := ctx.Value(userCtxKey).(*usermodel.UserModel)
	return raw
}
