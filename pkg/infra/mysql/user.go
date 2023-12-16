package mysql

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/yamato0211/plesio-server/pkg/domain/entity"
	"github.com/yamato0211/plesio-server/pkg/domain/repository"
	"github.com/yamato0211/plesio-server/pkg/utils/uuid"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) repository.UserRepository {
	return &userRepository{
		db: db,
	}
}

type GraphQLResponse struct {
	Data struct {
		User struct {
			ContributionsCollection struct {
				ContributionCalendar struct {
					TotalContributions int `json:"totalContributions"`
					Weeks              []struct {
						ContributionDays []struct {
							ContributionCount int    `json:"contributionCount"`
							Date              string `json:"date"`
						} `json:"contributionDays"`
					} `json:"weeks"`
				} `json:"contributionCalendar"`
			} `json:"contributionsCollection"`
		} `json:"user"`
	} `json:"data"`
}

// idによるuserの単一取得
func (ur *userRepository) Select(ctx echo.Context, id string) (*entity.User, error) {
	sql := `SELECT * FROM users WHERE id = ?`
	user := entity.User{}
	err := ur.db.Get(&user, sql, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepository) Insert(ctx echo.Context, name string, email string, git_id string) error {
	sql := `INSERT INTO users (id, name, email, git_id) VALUES (:id, :name, :email, :git_id);`
	in := entity.User{
		ID:    uuid.NewUUID(),
		Name:  name,
		Email: email,
		GitID: git_id,
	}
	_, err := ur.db.NamedExec(sql, in)
	if err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) LoginBonus(ctx echo.Context, id string, git_id string) (*entity.User, error) {
	githubPAT := "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" // GitHub Personal Access Token
	if githubPAT == "" {
		fmt.Println("GITHUB_PAT environment variable is not set.")
		log.Fatal("githubPAT is not set.")
		return nil, nil
	}

	sql := `SELECT * FROM users WHERE id = ?`
	user := entity.User{}
	err := ur.db.Get(&user, sql, id)

	if err != nil {
		log.Fatal(id)
		log.Fatal("select error")
		return nil, err
	}
	if user.IsLogined == false {
		sql := `UPDATE users SET is_logined = true WHERE id = ?`
		_, err := ur.db.Exec(sql, id)
		if err != nil {
			log.Fatal("update error")
			return nil, err
		} else {
			requestBody, err := json.Marshal(map[string]interface{}{
				"query": `
				query($userName:String!) {
					user(login: $userName) {
						contributionsCollection {
							contributionCalendar {
								totalContributions
								weeks {
									contributionDays{
										contributionCount
										date
									}
								}
							}
						}
					}
				}`,
				"variables": map[string]string{
					"userName": "IsseTeruhi-uni",
				},
			})
			if err != nil {
				log.Fatal("request error")
				fmt.Println("Error marshalling request body:", err)

				return nil, err
			}
			//fmt.Println(string(requestBody))
			// HTTP リクエストの作成
			req, err := http.NewRequest("POST", "https://api.github.com/graphql", bytes.NewBuffer(requestBody))
			if err != nil {
				log.Fatal("request error2")
				fmt.Println("Error creating request:", err)
				return nil, err
			}
			req.Header.Set("Authorization", "Bearer "+githubPAT)
			req.Header.Set("Content-Type", "application/json")

			// HTTP クライアントでリクエストを送信
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				log.Fatal("request error3")
				fmt.Println("Error sending request:", err)
				return nil, err
			}
			defer resp.Body.Close()

			// レスポンスを読み取る
			respBody, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatal("read error4")
				fmt.Println("Error reading response:", err)
				return nil, err
			}

			var response GraphQLResponse
			if err = json.Unmarshal(respBody, &response); err != nil {
				log.Fatal(err)
				log.Fatal("json error5")
			}

			sql := `UPDATE users SET coin = coin + ? WHERE id = ?`
			weeks := response.Data.User.ContributionsCollection.ContributionCalendar.Weeks
			lastWeek := weeks[len(weeks)-1]
			contributionDays := lastWeek.ContributionDays
			if len(contributionDays) == 1 {
				lastWeek = weeks[len(weeks)-2]
				contributionDays = lastWeek.ContributionDays
				lastDay := contributionDays[len(contributionDays)-1]
				cnt, _ := strconv.Atoi(strconv.Itoa(lastDay.ContributionCount))
				_, err = ur.db.Exec(sql, cnt, id)
				if err != nil {
					log.Fatal("update error")
					return nil, err
				}
				log.Fatal("update ok")
				return &user, nil
			}
			lastDay := contributionDays[len(contributionDays)-1]
			cnt := lastDay.ContributionCount
			_, err = ur.db.Exec(sql, cnt, id)
			// if len(contributionDays) == 1 {
			// 	lastWeek := weeks[len(weeks)-2]
			// 	contributionDays := lastWeek.ContributionDays
			// 	lastDay := contributionDays[len(contributionDays)-1]
			// 	cnt := lastDay.ContributionCount
			// 	_, err = ur.db.Exec(sql, cnt, id)
			// } else {

			// }

			if err != nil {
				log.Fatal("update error")
				return nil, err
			}
			//fmt.Println(response.Data.User.ContributionsCollection.ContributionCalendar.Weeks.contributionDays.contributionCount)
			log.Fatal("update ok")
			return &user, nil
		}
	} else {
		log.Fatal("?")
		return nil, err
	}
}
