package categories

type Category string

const (
	Gas               Category = "gas"
	Groceries         Category = "groceries"
	Personal          Category = "personal"
	Shopping          Category = "shopping"
	Education         Category = "education"
	Bills             Category = "bills"
	Health            Category = "health"
	Automotive        Category = "automotive"
	Travel            Category = "travel"
	Home              Category = "home"
	Miscellaneous     Category = "miscellaneous"
	FoodAndDrink      Category = "food_and_drink"
	Fees              Category = "fees"
	Entertainment     Category = "entertainment"
	GiftsAndDonations Category = "gifts_and_donations"
	Other             Category = "other"
)

var OverallTransactionTypes = map[string]Category{
	"gas":                 Gas,
	"groceries":           Groceries,
	"personal":            Personal,
	"shopping":            Shopping,
	"education":           Education,
	"bills":               Bills,
	"health":              Health,
	"automotive":          Automotive,
	"travel":              Travel,
	"home":                Home,
	"miscellaneous":       Miscellaneous,
	"food_and_drink":      FoodAndDrink,
	"gifts_and_donations": GiftsAndDonations,
	"fees":                Fees,
	"entertainment":       Entertainment,
	"other":               Other,
}
