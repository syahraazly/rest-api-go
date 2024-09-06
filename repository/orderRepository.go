package repository

import (
	"rest-api/database"
	"rest-api/models"

	"gorm.io/gorm"
)

func CreateOrder(order *models.Order) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		// simpan order terlebih dahulu
		if err := tx.Create(order).Error; err != nil {
			return err
		}

		// simpan atau update items
		for i := range order.Items {
			order.Items[i].OrderID = order.ID
			if err := tx.Save(&order.Items[i]).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func GetAllOrders() ([]models.Order, error) {
	var orders []models.Order
	err := database.DB.Preload("Items").Find(&orders).Error
	return orders, err
}

func GetOrderById(id uint) (models.Order, error) {
	var order models.Order
	err := database.DB.Preload("Items").First(&order, id).Error
	return order, err
}

func UpdateOrder(order *models.Order) error {
	return database.DB.Save(order).Error
}

func DeleteOrder(order *models.Order) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		// delete semua item
		if err := tx.Where("order_id = ?", order.ID).Delete(&models.Item{}).Error; err != nil {
			return err
		}

		// delete order
		if err := tx.Delete(order).Error; err != nil {
			return err
		}
		return nil
	})
}
