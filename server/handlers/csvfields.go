package handlers 




type CSVFieldHandler struct {
	CSVFieldService *services.CSVFieldService
}


func NewCSVFieldHandler() *CSVfieldHandler {

	csvFieldService := services.NewCSVFieldService() 

	return CSVFieldHandler{
		CSVFieldService: csvFieldService,
	}
}

func ( c *CSVFieldHandler) GetAllFields