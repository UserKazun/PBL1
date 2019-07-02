package service

import (
	"log"

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

// UpdateCumulativePoints ...購入ボタンを押したユーザーの飢餓対策貢献Pointを更新
func UpdateCumulativePoints(userID string, recipeCount uint, recipePoint uint) error {
	degreeOfContribution := model.DegreeOfContribution{}
	var err error

	err = db.Where("user_id = ?", userID).First(&degreeOfContribution).Error
	if err != nil {
		return err
	}

	err = db.Model(&degreeOfContribution).Update("cumulative_points", degreeOfContribution.CumulativePoints+(recipeCount*recipePoint)).Error
	if err != nil {
		return err
	}

	return nil
}

func ReduceCumulativePointsByUserID(userID string, tradeItemID uint) (error, bool) {
	degreeOfContribution := model.DegreeOfContribution{}
	tradeItem := model.TradeItem{}
	var err error

	err = db.Where("user_id = ?", userID).First(&degreeOfContribution).Error
	if err != nil {
		return err, false
	}

	err = db.Where("id = ?", tradeItemID).First(&tradeItem).Error
	if err != nil {
		return err, false
	}

	newDegreeOfContribution := degreeOfContribution
	if newDegreeOfContribution.CumulativePoints < tradeItem.Point {
		return err, true
	}
	newDegreeOfContribution.CumulativePoints = degreeOfContribution.CumulativePoints - tradeItem.Point

	log.Println(newDegreeOfContribution.CumulativePoints)

	err = db.Model(&degreeOfContribution).Update("cumulative_points", newDegreeOfContribution.CumulativePoints).Error
	if err != nil {
		return err, false
	}

	return nil, false
}
