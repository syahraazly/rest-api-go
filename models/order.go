package models

import (
	"strconv"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID           uint       `gorm:"primaryKey;autoIncrement"`
	CustomerName string     `gorm:"type:varchar(255);not null"`
	OrderedAt    *time.Time `gorm:"not null"`
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
	Items        []Item `gorm:"foreignKey:OrderID"`
}

func (order *Order) BeforeSave(tx *gorm.DB) (err error) {

	// jika order ID == 0, maka ini adalah operasi create dan tidak dijalankan
	if order.ID == 0 {
		return nil
	}

	// ambil item order dari database
	var existingItems []Item
	if err := tx.Where("order_id = ?", order.ID).Find(&existingItems).Error; err != nil {
		return err
	}

	// fungsi untuk menghasilkan key unik berdasarkan name, quantity, dan description
	generateItemKey := func(item Item) string {
		return item.Name + "|" + item.Description + "|" + strconv.Itoa(item.Quantity)
	}

	// map data item dalam payload berdasarkan key unik
	newItemsMap := make(map[string]bool)
	for _, item := range order.Items {
		key := generateItemKey(item)
		newItemsMap[key] = true
	}

	// hapus item yang ada di database tapi tidak ada dalam payload terbaru (berdasarkan key unik)
	for _, existingItem := range existingItems {
		key := generateItemKey(existingItem)
		if _, found := newItemsMap[key]; !found {
			if err := tx.Delete(&existingItem).Error; err != nil {
				return err
			}
		}
	}

	// update atau tambahkan item baru
	for i := range order.Items {
		var existingItem Item
		_ = generateItemKey(order.Items[i])
		if err := tx.Where("order_id = ? AND name = ? AND description = ? AND quantity = ?", order.ID, order.Items[i].Name, order.Items[i].Description, order.Items[i].Quantity).First(&existingItem).Error; err != nil {
			// jika item tidak ditemukan, tambahkan item baru
			order.Items[i].OrderID = order.ID
			if err := tx.Create(&order.Items[i]).Error; err != nil {
				return err
			}
		} else {
			// jika item ditemukan, lakukan update
			existingItem.Description = order.Items[i].Description
			existingItem.Quantity = order.Items[i].Quantity
			if err := tx.Save(&existingItem).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
