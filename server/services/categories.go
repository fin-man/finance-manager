package services

import "finance-manager/categories"

type CategoriesService struct {
}

func NewCategoriesService() *CategoriesService {
	return &CategoriesService{}
}

func (c *CategoriesService) GetCategories() []categories.Category {
	/*
		Gas           Category = "gas"
		Groceries     Category = "groceries"
		Personal      Category = "personal"
		Shopping      Category = "shopping"
		Education     Category = "education"
		Bills         Category = "bills"
		Health        Category = "health"
		Automotive    Category = "automotive"
		Travel        Category = "travel"
		Home          Category = "home"
		Miscellaneous Category = "miscellaneous"
		FoodAndDrink  Category = "food_and_drink"
		Fees          Category = "fees"
		Entertainment Category = "entertainment"
		Other         Category = "other"
	*/

	return []categories.Category{categories.Gas, categories.Groceries, categories.Personal, categories.Shopping, categories.Education, categories.Bills, categories.Health, categories.Automotive, categories.Travel, categories.Home, categories.Miscellaneous, categories.FoodAndDrink, categories.Fees, categories.Entertainment, categories.Other}
}
