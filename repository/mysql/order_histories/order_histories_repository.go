package model

import (
	DBConfig "saham_rakyat/config"
	src "saham_rakyat/models"
)

func Migrate() error {
	var db, err = DBConfig.ConnectMysql()
	if err != nil {
		return err

	}
	err = db.AutoMigrate(&src.OrderHistories{})
	if err != nil {
		return err
	}
	return nil
}

func Store(data *src.OrderHistories) error {
	db, err := DBConfig.ConnectMysql()
	if err != nil {
		return err

	}

	if err = db.Create(data).Error; err != nil {
		return err
	}

	return nil
}

func Update(data *src.OrderHistories) error {
	db, err := DBConfig.ConnectMysql()
	if err != nil {
		return err

	}

	if err = db.Model(data).Updates(src.OrderHistories{Description: data.Description}).Error; err != nil {
		return err
	}

	return nil
}

func Delete(data *src.OrderHistories) error {
	db, err := DBConfig.ConnectMysql()
	if err != nil {
		return err

	}

	if err = db.Delete(data).Error; err != nil {
		return err
	}

	return nil
}

func GetDataById(id int) (src.OrderHistories, error) {
	var order_histories src.OrderHistories
	order_histories.Id = id
	db, err := DBConfig.ConnectMysql()
	if err != nil {
		return order_histories, err

	}
	result := db.First(&order_histories)
	if result.RowsAffected < 1 {
		return order_histories, result.Error
	}
	return order_histories, nil
}

func GetAllData(page int, limit int) ([]src.OrderHistories, error) {
	var order_histories []src.OrderHistories
	db, err := DBConfig.ConnectMysql()
	if err != nil {
		return order_histories, err

	}

	offset := (page - 1) * limit

	if err = db.Offset(offset).Limit(limit).Find(&order_histories).Error; err != nil {
		return order_histories, err
	}
	return order_histories, nil
}
