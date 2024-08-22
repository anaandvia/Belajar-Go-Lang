package categoriesmodel

import (
	"web-golang/config"
	"web-golang/entities"
)

func GetAll() []entities.Categories {
	rows, err := config.DB.Query(`SELECT * FROM categories`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var categories []entities.Categories
	for rows.Next() {
		var category entities.Categories
		if err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpddatedAt); err != nil {
			panic(err)
		}

		categories = append(categories, category)
	}

	return categories
}
