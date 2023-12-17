package firebase

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
)

func InitializeAppWithServiceAccount(serviceAccountKeyFilePath string) *firebase.App {
	firebaseString := "{\n    \"type\": \"service_account\",\n    \"project_id\": \"plesio-cup\",\n    \"private_key_id\": \"db12c33c5a2618464344a231f7efa33dff01a4d5\",\n    \"private_key\": \"-----BEGIN PRIVATE KEY-----\\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCuddF9nRgU91YG\\nE8paTn3rDGonkakN6SqqVQ4npylo7VuNwTwDFrxAqF4a6i9E0spH1qH3FG5f8sB5\\nUrbOJCH4IHBOItvExsUJoaN3nQc0Iq9vcO/cFiIv3oHrVa3xihWbubB/U8gvpSD+\\nWbjItX5FMp02cdRusU1kJFRtyRYaeN6ZU4vPw8Oy3yJZkj/6WiI2psvXDYThH/ng\\nilwO1sRBCfRQiGzKjJX7tumVLfv6Kn9AnC74veFaVk9VZoE3KMD9XOmoNvZnBXe0\\nReeDIHWrS19P4ckJxZ7lbtMTW6HTGnXDiFgX/spmV3XdavC+aQ5fbVK9kkvYP1iC\\n5yloDhu1AgMBAAECggEAGJPjFzv4Lk34+cK0fXv029YEJ35rD1ljqBry802mtRBJ\\nV2NbccD0dPRzS33l+l9q3WuEplMjyJMhoAXx2IE5IuRpBez9wo9RJ2304k9GOkdP\\nqVbDB+G7X7yMV4vaxT4zlJx4m7uBJoswJn0b3fTJrPR4Nvka2RxT35GJrmzHlyNn\\nuaEB/qlDdAxtcfDH+lnEAnBYzElk4g7CcxJrkFDGRPRCgT0wyDA2NtRjz8VZ5544\\neoLraeh/GiOwEIPQVzGCipOiKslNza5vtG7s1jSlgCGGnKMMyC7G+RiXhb1HhnLC\\nSUzE7sAit/Mp15NZUFuyH4gNJPPpZHCPhDlQo+JQwQKBgQDpvA9p+aD1s1BrA3h5\\nx5v7SCHaCQdUAqrD1tiP+d29S+yw0Ij7t5K61DkDXgPNrefzfwdoz5fxSARv1Aau\\ndeTfiEPWHmp1MG6VjCKJ8ck4h//vhCPdzMPK5Z0bSkKkQMEnJq9qXmWx4CLPNATD\\n0EwvJFNBy1UlmxSTVFYLm+P9QQKBgQC/FEVT9LnYr3ooDwk6639YtCEiHbUVZj+u\\n5HCAm78GrIJkvPxXszLoaeqylu4+RQwfDvNHKssk6MTsFL+shMgyOZHAxG+mcdC7\\nmwCNn2KZ8M8u/hHL5mDTIuZEfim1sGXACzUvXnhX+6CiV+5HR33iOd67TZznO4oV\\nlyGpu2sddQKBgQC+XdXsCU2Ib/hCgHwiK2omWhSAIbRfPyDYCmcttGNeJrVrOR0o\\nbtZJ5kldLbhJT6fIESD2w4VMFczPS8/5TboQOUqO3Qz1z3FxycXjthKeQV45NNWF\\nc53P5f10WaI3Zyop+f8K0kFfg162hfaPiVpTVvGcd8AstccnS5wdLmWdQQKBgHxx\\nfT1Scv8WOZ6BPGetHADvP2zyoA7RysPibwqzZFO6Yimtk0KiqJMI+BtbhMn1OhG4\\nMOS2kBr3pdiDItn+mEtthulehcCG+4RTTKbvmM3dMXHPclOzNHgNkJ4m9I8p1Pqf\\n0gGJqyLR9CoGzl6Jxwhg3BhsUbWe5y1sgjYCo55JAoGBAKhXS8Ih06Y8f6aPiDa9\\nbSb8+gwP317HXMLSKp8IAQ1C0eNXqyGmYzPxdP/tlEJx1BMlnlqvk2652X2JaXGN\\nAp1Bp9Hl8NNcK/sA9vpxj+kHCYMwfFyKK2IU4JOpcl22ESMJT4b6JwpPz0ghI4bZ\\nNho+vbvDa/bL84DmRsjr/wC2\\n-----END PRIVATE KEY-----\\n\",\n    \"client_email\": \"firebase-adminsdk-fdn0f@plesio-cup.iam.gserviceaccount.com\",\n    \"client_id\": \"116071134463910163541\",\n    \"auth_uri\": \"https://accounts.google.com/o/oauth2/auth\",\n    \"token_uri\": \"https://oauth2.googleapis.com/token\",\n    \"auth_provider_x509_cert_url\": \"https://www.googleapis.com/oauth2/v1/certs\",\n    \"client_x509_cert_url\": \"https://www.googleapis.com/robot/v1/metadata/x509/firebase-adminsdk-fdn0f%40plesio-cup.iam.gserviceaccount.com\",\n    \"universe_domain\": \"googleapis.com\"\n}"
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

			log.Println(token.Subject)

			type UserInfo struct {
				UserID string `db:"id"`
				// 必要に応じて他のフィールドも追加
			}
			var user UserInfo
			if db.Ping() != nil {
				echo.NewHTTPError(http.StatusInternalServerError, "Error connecting to database")
			}
			err = db.Get(&user, "SELECT id FROM users WHERE git_id = ?", token.Subject)
			log.Println("===============-")
			log.Println(user)
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
