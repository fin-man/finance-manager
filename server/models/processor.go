package models

import (
	"github.com/jinzhu/gorm"
)

type ProcessorModel struct {
	gorm.Model
	URL           string `json:"url" gorm:"unique"`
	ProcessorName string `json:"processor_name" gorm:"unique"`
}

func (p ProcessorModel) TableName() string {
	return "processors"
}

func NewProcessorModel() *ProcessorModel {
	return &ProcessorModel{}
}

func (p *ProcessorModel) CreateProcessor(processor *ProcessorModel) error {
	return DB.Create(processor).Error
}

func (p *ProcessorModel) GetAllProcessors() ([]ProcessorModel, error) {
	var processors []ProcessorModel

	if err := DB.Find(&processors).Error; err != nil {
		return processors, err
	}

	return processors, err
}

func (p *ProcessorModel) DeleteProcessor(processor *ProcessorModel) error {
	return DB.Unscoped().Delete(&processor).Error
}
