package src

import (
	"errors"
	"github.com/labstack/gommon/log"
)

const RetryCount = 3

type Service interface {
	CreateRandomUsers(n int)
	CreateRandomUser() (User, error)
	GetRandomUnPlayedUsers(n int) []*User
	SavePlayersScore(users []*User) error
}

type service struct {
	db     *DB
	client Client
}

func (s *service) CreateRandomUsers(n int) {
	retryCount := RetryCount

	for i := 0; i < n; i++ {
		username := GenerateRandomString(10)
		password := GenerateRandomString(8)
		user, err := s.client.Signup(username, password)
		if err != nil && retryCount > 0 {
			log.Error(err, username)
			retryCount--
			i--
			return
		}
		s.db.Add(user.Result.ID, User{
			ID:       user.Result.ID,
			Username: user.Result.Username,
			IsPlayed: false,
			Score:    0,
		})
		retryCount = RetryCount
	}
}

func (s *service) CreateRandomUser() (User, error) {
	var user User

	for i := 0; i < RetryCount; i++ {
		username := GenerateRandomString(10)
		password := GenerateRandomString(8)

		res, err := s.client.Signup(username, password)

		if err != nil {
			log.Error(err, username)
			continue
		}

		return User{
			ID:       res.Result.ID,
			Username: res.Result.Username,
		}, nil
	}

	return user, errors.New("user couldn't created")
}

func (s *service) GetRandomUnPlayedUsers(n int) []*User {
	var result []*User
	var keys = make(map[int]int, 0)
	counter := 0

	for counter < n {
		user := s.db.GetRandom().(User)
		if user.IsPlayed == false {
			_, ok := keys[user.ID]
			if !ok {
				result = append(result, &user)
				keys[user.ID] = user.ID
				counter++
			}
		}
	}

	return result
}

func (s *service) SavePlayersScore(users []*User) error {
	var players Players
	for _, user := range users {
		players = append(players, Player{
			ID:    user.ID,
			Score: user.Score,
		})
	}

	err := s.client.EndGame(players)
	if err != nil {
		return err
	}

	return nil
}

func NewService(db *DB, client Client) Service {
	return &service{db: db, client: client}
}
