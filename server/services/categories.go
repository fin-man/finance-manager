package services

import "finance-manager/categories"

type CategoriesService struct {
}

func NewCategoriesService() *CategoriesService {
	return &CategoriesService{}
}

func (c *CategoriesService) GetCategories() []categories.Category {
	return []categories.Category{categories.Gas, categories.Groceries, categories.Personal, categories.Shopping, categories.Education, categories.Bills, categories.Health, categories.Automotive, categories.Travel, categories.Home, categories.Miscellaneous, categories.FoodAndDrink, categories.Fees, categories.Entertainment, categories.Other}
}
