package services

import (
	"crawler/bases"
	"crawler/models"
	"fmt"
)

// MachineInfoService ...
type MachineInfoService struct{}

// FindOne ...
func (s *MachineInfoService) FindOne(id int64) (models.MachineInfo, error) {
	orm := bases.GetEngine()
	var machaineInfo models.MachineInfo
	orm.Id(id).Get(&machaineInfo)
	return machaineInfo, nil
}

// FindAll ...
func (s *MachineInfoService) FindAll(limit int) ([]models.MachineInfo, error) {
	orm := bases.GetEngine()
	machineInfos := make([]models.MachineInfo, 0)
	err := orm.Find(&machineInfos)
	fmt.Println(len(machineInfos))
	for key, value := range machineInfos {
		fmt.Println(key, value)
	}
	return machineInfos, err
}
