package model

import (
	config "saham_rakyat/config"
	src "saham_rakyat/models"
)

func Migrate() error {
	db, err := config.ConnectMysql()
	if err != nil {
		return err

	}
	err = db.AutoMigrate(&src.Users{})
	if err != nil {
		return err
	}
	return nil
}

func Store(data *src.Users) error {
	db, err := config.ConnectMysql()
	if err != nil {
		return err

	}

	if err = db.Create(data).Error; err != nil {
		return err
	}

	return nil
}

func Update(data *src.Users) error {
	db, err := config.ConnectMysql()
	if err != nil {
		return err

	}

	if err = db.Model(data).Updates(src.Users{FullName: data.FullName, FirstOrder: data.FirstOrder}).Error; err != nil {
		return err
	}

	return nil
}

func Delete(data *src.Users) error {
	db, err := config.ConnectMysql()
	if err != nil {
		return err

	}

	if err = db.Delete(data).Error; err != nil {
		return err
	}

	return nil
}

func GetDataById(id int) (src.Users, error) {
	var user src.Users
	user.Id = id
	db, err := config.ConnectMysql()
	if err != nil {
		return user, err

	}
	result := db.First(&user)
	if result.RowsAffected < 1 {
		return user, result.Error
	}

	return user, nil
}

func GetAllData(page int, limit int) ([]src.Users, error) {
	var users []src.Users
	db, err := config.ConnectMysql()
	if err != nil {
		return users, err

	}

	offset := (page - 1) * limit

	if err = db.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}
