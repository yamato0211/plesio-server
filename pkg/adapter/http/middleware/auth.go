package firebase

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
)

func InitializeAppWithServiceAccount(serviceAccountKeyFilePath string) *firebase.App {
	firebaseString := os.Getenv("FIREBASE_AUTH_KEY")
	opt := option.WithCredentialsJSON([]byte(firebaseString))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Error initializing Firebase app: %v", err)
	}
	return app
}

func NewFirebaseMiddleware(firebaseApp *firebase.App, db *sqlx.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			idToken := strings.TrimSpace(strings.Replace(authHeader, "Bearer", "", 1))

			// IDトークンの検証
			client, err := firebaseApp.Auth(context.Background())
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "Error initializing Firebase Auth client")
			}

			token, err := client.VerifyIDToken(context.Background(), idToken)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Firebase ID token")
			}

			log.Println(token)

			type UserInfo struct {
				UserID string `db:"id"`
				// 必要に応じて他のフィールドも追加
			}
			var user UserInfo
			githubID := token.UID
			err = db.Get(&user, "SELECT id FROM users WHERE git_id = ?", githubID)
			if err != nil {
				if err == sql.ErrNoRows {
					echo.NewHTTPError(http.StatusForbidden, "User not found")
				} else {
					echo.NewHTTPError(http.StatusInternalServerError, "Error fetching user")
				}
			} else {
				c.Set("user_id", user.UserID)
			}

			return next(c)
		}
	}
}
