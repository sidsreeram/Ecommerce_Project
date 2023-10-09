package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           uint   `gorm:"primaryKey" json:"id"`
	Username     string `gorm:"not null;size:30" json:"username"`
	Password     string `gorm:"not null;size:255" json:"password"`
	Email        string `gorm:"not null;uniqueIndex;size:255" json:"email"`
	Phone        string `gorm:"not null;size:20" json:"phone"`
	Token        string `gorm:"size:255" json:"token"`
	RefreshToken string `gorm:"size:255" json:"refresh_token"`
	IsAdmin      bool   `json:"isadmin" validate:"required"`
	IsSuperAdmin bool   `json:"issuperadmin" validate:"required"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	UserCart     []ProductUser `gorm:"foreignKey:UserID" json:"user_cart"`
	AddressDetails []Address   `gorm:"foreignKey:UserID" json:"address_details"`
	OrderStatus  []Order      `gorm:"foreignKey:UserID" json:"order_status"`
}

type Product struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	ProductName string `json:"product_name"`
	Price       uint64 `json:"price"`
	Rating      uint8  `json:"rating"`
	Image       string `json:"image"`
}

type ProductUser struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	ProductName string `json:"product_name"`
	Price       int    `json:"price"`
	Rating      uint   `json:"rating"`
	Image       string `json:"image"`
}

type Address struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	House   string `json:"house_name"`
	Street  string `json:"street_name"`
	City    string `json:"city_name"`
	Pincode string `json:"pin_code"`
}

type Order struct {
	ID         uint         `gorm:"primaryKey" json:"id"`
	OrderCart  []ProductUser `gorm:"foreignKey:OrderID" json:"order_cart"`
	OrderedAt  time.Time    `json:"ordered_on"`
	TotalPrice int          `json:"total_price"`
	Discount   *int         `json:"discount"`
	PaymentMethod Payment   `json:"payment_method" gorm:"embedded"`
}

type Payment struct {
	Digital bool `json:"digital"`
	COD     bool `json:"cod"`
}
