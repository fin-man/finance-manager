package services

import "github.com/fin-man/finance-manager/server/models"

type ProcessorService struct {
	Processor *models.ProcessorModel
}

func NewProcessorService() *ProcessorService {
	processor := models.NewProcessorModel()
	return &ProcessorService{
		Processor: processor,
	}
}

func (c *ProcessorService) CreateProcessor(processor *models.ProcessorModel) error {
	return c.Processor.CreateProcessor(processor)
}

func (c *ProcessorService) GetAllProcessors() ([]models.ProcessorModel, error) {
	return c.Processor.GetAllProcessors()
}

func (c *ProcessorService) RemoveProcessor(processor *models.ProcessorModel) error {
	return c.Processor.DeleteProcessor(processor)
}
