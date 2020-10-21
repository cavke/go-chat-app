package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"

	"github.com/cavke/go-chat-app"
)

// UserService represents a MySQL implementation of chatapp.UserService.
type UserService struct {
	DB *sql.DB
}

// User returns a user for a given id.
func (s *UserService) User(id int) (*chatapp.User, error) {
	var u chatapp.User
	err := s.DB.QueryRow("SELECT id, name FROM users WHERE id = ?", id).Scan(&u.ID, &u.Name)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &u, nil
}

// Users returns list of users.
func (s *UserService) Users() ([]*chatapp.User, error) {
	rows, err := s.DB.Query("SELECT id, name FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []*chatapp.User
	for rows.Next() {
		u := chatapp.User{}
		err := rows.Scan(&u.ID, &u.Name)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, &u)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return users, nil
}

// implement remaining chatapp.UserService interface...
