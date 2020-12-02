package Week02

import (
	"database/sql"
	"github.com/pkg/errors"
)

type User struct {
	Id uint64
	Name string
	Age uint8
}

func (u *User) GetUser(Id uint64) (*User,error) {
	// do stuff
	return nil, errors.Wrapf(sql.ErrNoRows, "user info is empty userId:%v", Id)
}

func IsErrorNoRows(err error) bool {
	if errors.Cause(err) == sql.ErrNoRows {
		return true
	}
	return false
}