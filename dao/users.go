package dao

import (
	"database/sql"
	"github.com/pkg/errors"
)

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type userDaoInterface interface {
	//Create(u *User) error
	//Update(u *User) error
	//Delete(i int) error
	GetById(i int) (User, error)
	//GetAll() ([]User, error)
}

type userDao struct {
}

var UserDao = userDao{}

func (dao userDao) GetById(i int) (User, error) {
	user := User{}

	err := Dao.db.QueryRow("SELECT id, name FROM users WHERE id = ?", i).Scan(&user.ID, &user.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			// there were no rows, but otherwise no error occurred
			err = errors.Wrap(err, "no user found")
		} else {
			err = errors.Wrap(err, "when finding user")
		}
	}

	return user, err
}
