package models

import "github.com/jinzhu/gorm"

type Product struct{
	Model

	Name string `json: "name" gorm: "size:100;not null"`
	ModelType int `json: "model_type" gorm: "SMALLINT;default:0"`
}

func (Product) TableName() string{
	return "product"
}

func GetProduct(id int) (*Product, error) {
	var product Product
	err := db.Where("id = ?", id).First(&product).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &product, nil
}

func GetProducts()  ([] *Product, error){
	var products []*Product
	err := db.Find(&products).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return nil, err
	}

	return products, nil
}

func ExistProductByID(id int) (bool, error)  {
	var count int
	err := db.Model(&Product{}).Where("id = ?", id).Count(&count).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return false, err
	}

	if count > 0{
		return true, nil
	}

	return false, nil
}