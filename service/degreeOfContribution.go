package service

import (
	"github.com/PBL1/model"
)

// CreateDegreeOfContribution ...DBに与えられたデータをinsertする
func CreateDegreeOfContribution(degreeOfContribution model.DegreeOfContribution) (model.DegreeOfContribution, error) {
	err := db.Create(&degreeOfContribution).Error
	if err != nil {
		return model.DegreeOfContribution{}, err
	}
	return degreeOfContribution, nil
}
