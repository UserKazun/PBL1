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

// GetCumulativePointsByUserID ...対象ユーザーの飢餓対策合計ポイントを返す
func GetCumulativePointsByUserID(userID string) (uint, error) {
	degreeOfContribution := model.DegreeOfContribution{}

	err := db.Where("user_id = ?", userID).First(&degreeOfContribution).Error
	if err != nil {
		return 0, err
	}

	return degreeOfContribution.CumulativePoints, nil
}

// UpdateCumCumlativePoints ...購入ボランを押したユーザーの貢献Pointを更新
func UpdateCumlativePoints(cumulativePoints uint) (model.DegreeOfContribution, error) {
	degreeOfContribution := model.DegreeOfContribution{}

	err := db.Model(&degreeOfContribution).Update("cumulative_points", cumulativePoints).Error
	if err != nil {
		return model.DegreeOfContribution{}, err
	}

	return degreeOfContribution, err
}