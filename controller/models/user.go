package models

type User struct {
	Id       int64
	UserName string `xorm:"varchar(30) unique"`
	NickName string `xorm:"varchar(30) unique"`
	Password string `xorm:"varchar(128)"`
}

func GetUserById(id int64) (*User, error) {
	var user User
	has, err := orm.Id(id).Get(&user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, ErrNotExist
	}
	return &user, nil
}

func GetUserByName(username string) (*User, error) {
	var user = User{
		UserName: username,
	}
	has, err := orm.Get(&user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, ErrNotExist
	}
	return &user, nil
}
