package createSeeds

import (
	"fmt"
	"strconv"

	"github.com/PBL1/model"
	"github.com/PBL1/service"
)

// CreateSeedDegreeOfContributions ...ここで記入したデータをDBにinsertする
func CreateSeedDegreeOfContributions() {
	degreeOfContributions := []map[string]string{
		map[string]string{
			"UserID":           "goya",
			"CumulativePoints": "1000",
		},
	}

	for _, info := range degreeOfContributions {
		userID, _ := info["UserID"]
		cumulativePoints, _ := strconv.Atoi(info["CumulativePoints"])
		createDegreeOfContribution(model.DegreeOfContribution{
			UserID:           string(userID),
			CumulativePoints: uint(cumulativePoints),
		})
	}
}

// createUser ...渡されたデータをserviceの関数へ渡し、DBにinsertさせる
func createDegreeOfContribution(degreeOfContribution model.DegreeOfContribution) {
	degreeOfContribution, err := service.CreateDegreeOfContribution(degreeOfContribution)
	if err != nil {
		panic(err)
	}
	fmt.Printf("degreeOfContribution created\n")
}
