package services

import (
	"crawler/bases"
	"crawler/models"
)

// CityService struct
type CityService struct{}

// FindAll func
func (c *CityService) FindAll(limit int32) ([]models.City, error) {
	orm := bases.GetEngine()
	city := make([]models.City, 0)
	err := orm.Find(&city)
	return city, err
}
