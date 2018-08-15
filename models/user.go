package models

type UserProduct struct {
	Model
	UserId int `json: "user_id" gorm: "index"`
	ProductId int `json: "product_id" gorm: "index"`
	Product Product `json: product`
}

func GetProductsByUserID(userId int) ([]map[string]interface{}, error){
	var results []map[string]interface{}

	rows, err := db.Table("user_product").Select("product.id, product.name").Joins(
		"join product on product.id = user_product.product_id").Where("user_product.user_id = ?", userId).Rows()
	if err != nil{
		return results, err
	}
	defer rows.Close()

	for rows.Next(){
		var id int
		var name string
		result := make(map[string]interface{})

		rows.Scan(&id, &name)
		result["id"] = id
		result["name"] = name

		results = append(results, result)
	}

	return results, nil
}