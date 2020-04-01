package categories

type Category string

const (
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
)

var ChaseTransactionTypes = map[string]Category{
	"Gas":                Gas,
	"Groceries":          Groceries,
	"Personal":           Personal,
	"Shopping":           Shopping,
	"Education":          Education,
	"Bills & Utilities":  Bills,
	"Health & Wellness":  Health,
	"Automotive":         Automotive,
	"Travel":             Travel,
	"Home":               Home,
	"Miscellaneous":      Miscellaneous,
	"Food & Drink":       FoodAndDrink,
	"Fees & Adjustments": Fees,
	"Entertainment":      Entertainment,
}

var CapitalOneTransactionTypes = map[string]Category{
	"Payment/Credit":      Fees,
	"Health Care":         Health,
	"Other":               Other,
	"Phone/Cable":         Bills,
	"Gas/Automotive":      Gas,
	"Car Rental":          Travel,
	"Dining":              FoodAndDrink,
	"Merchandise":         Groceries,
	"Other Services":      Other,
	"Airfare":             Travel,
	"Lodging":             Travel,
	"Internet":            Bills,
	"Other Travel":        Travel,
	"Entertainment":       Entertainment,
	"Fee/Interest Charge": Fees,
}
