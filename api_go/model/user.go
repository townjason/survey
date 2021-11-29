package model

import (
	"api/database/mysql"
	"api/util/log"
	"time"
)

type Users struct {
	Id          int64     `table:"id"`
	NickName    string    `table:"nick_name"`
	CountryCode string    `table:"country_code"`
	Phone       string    `table:"phone"`
	Password    string    `table:"password"`
	BirthdayAt  string 	`table:"birthday_at"`
	CreatedAt   time.Time `table:"created_at"`
	UserCode    string    `table:"user_code"`
	Picture     string    `table:"picture"`
	Gender      int64     `table:"gender"`
	Email       string    `table:"email"`
	City        string    `table:"city"`
	Township    string    `table:"township"`
	Address     string    `table:"address"`
}

func (model *Users) SetId(id int64) *Users {
	model.Id = id
	return model
}

func (model *Users) SetNickName(nickName string) *Users {
	model.NickName = nickName
	return model
}

func (model *Users) SetPhone(phone string) *Users {
	model.Phone = mysql.Encode(phone)
	return model
}

func (model *Users) SetCountryCode(countryCode string) *Users {
	model.CountryCode = countryCode
	return model
}

func (model *Users) SetUserCode(userCode string) *Users {
	model.UserCode = userCode
	return model
}

func (model *Users) SetCreatedAt(createdAt time.Time) *Users {
	model.CreatedAt = createdAt
	return model
}



func (model *Users) SetPassword(password string) *Users {
	model.Password = password
	return model
}

func (model *Users) SetGender(gender int64) *Users {
	model.Gender = gender
	return model
}

func (model *Users) SetEmail(email string) *Users {
	model.Email = email
	return model
}

func (model *Users) SetCity(city string) *Users {
	model.City = city
	return model
}

func (model *Users) SetTownship(township string) *Users {
	model.Township = township
	return model
}

func (model *Users) SetAddress(address string) *Users {
	model.Address = address
	return model
}

func (model *Users) GetUsersIdByPhone() *Users {
	table := mysql.Model(model)

	if model.CountryCode != ""{
		table.Where("country_code", "=", model.CountryCode)
	}

	if model.Phone != ""{
		table.Where("phone", "=", model.Phone)
	}

	log.Error(table.
		Select([]string{"id", "password"}).
		Find().
		Scan(&model.Id, &model.Password))

	return model
}

func (model *Users) GetUsersInfoByUserCode() *Users {
	log.Error(mysql.Model(model).
		Where("user_code", "=", model.UserCode).
		Select([]string{"id", "password"}).
		Find().
		Scan(&model.Id, &model.Password))

	return model
}

func (model *Users) GetUsersIdById() *Users {
	log.Error(mysql.Model(model).
		Where("id", "=", model.Id).
		Select([]string{"id", "nick_name", "password"}).
		Find().
		Scan(&model.Id, &model.NickName, &model.Password))

	return model
}

func (model *Users) UpdateUserInfo(userId int64, birthday string) {
	model.BirthdayAt = mysql.Encode(birthday)
	log.Error(mysql.Model(model).
		Where("id", "=", userId).
		Update([]string{"birthday_at"}))
}
