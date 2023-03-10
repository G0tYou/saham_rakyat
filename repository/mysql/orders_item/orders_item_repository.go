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
	err = db.AutoMigrate(&src.OrdersItem{})
	if err != nil {
		return err
	}
	return nil
}

func Store(data *src.OrdersItem) error {
	db, err := DBConfig.ConnectMysql()
	if err != nil {
		return err

	}

	if err = db.Create(data).Error; err != nil {
		return err
	}

	return nil
}

func Update(data *src.OrdersItem) error {
	db, err := DBConfig.ConnectMysql()
	if err != nil {
		return err

	}

	if err = db.Model(data).Updates(src.OrdersItem{Name: data.Name, Price: data.Price, ExpiredAt: data.ExpiredAt}).Error; err != nil {
		return err
	}

	return nil
}

func Delete(data *src.OrdersItem) error {
	db, err := DBConfig.ConnectMysql()
	if err != nil {
		return err

	}

	if err = db.Delete(data).Error; err != nil {
		return err
	}

	return nil
}

func GetDataById(id int) (src.OrdersItem, error) {
	var order_item src.OrdersItem
	order_item.Id = id
	db, err := DBConfig.ConnectMysql()
	if err != nil {
		return order_item, err

	}
	result := db.First(&order_item)
	if result.RowsAffected < 1 {
		return order_item, result.Error
	}
	return order_item, nil
}

func GetAllData(page int, limit int) ([]src.OrdersItem, error) {
	var order_items []src.OrdersItem
	db, err := DBConfig.ConnectMysql()
	if err != nil {
		return order_items, err

	}

	offset := (page - 1) * limit

	if err = db.Offset(offset).Limit(limit).Find(&order_items).Error; err != nil {
		return order_items, err
	}
	return order_items, nil
}
