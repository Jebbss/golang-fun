package main

import "testing"

var listItemshortDescriptionDivisibleBy3 = listItemDO{shortDescription: "abc", price: 11}
var listItemshortDescriptionNotDivisibleBy3 = listItemDO{shortDescription: "abcd", price: 10}

func TestIsAlphaNumericLowercaseATrue(t *testing.T) {
	if !isAlphanumeric('a') {
		t.Error("Expected true, got false")
	}
}

func TestIsAlphaNumericUppercaseATrue(t *testing.T) {
	if !isAlphanumeric('A') {
		t.Error("Expected true, got false")
	}
}

func TestIsAlphaNumericOneTrue(t *testing.T) {
	if !isAlphanumeric('1') {
		t.Error("Expected true, got false")
	}
}

func TestIsAlphaNumericSpaceFalse(t *testing.T) {
	if isAlphanumeric(' ') {
		t.Error("Expected false, got true")
	}
}

func TestPointsForCompanyName7(t *testing.T) {
	var points = pointsForRetailerName("Walmart")
	if points != 7 {
		t.Error("Expected 7, got ", points)
	}
}

func TestPointsForCompanyNameShouldBe7WithSpaceInName(t *testing.T) {
	var points = pointsForRetailerName(" Walmart")
	if points != 7 {
		t.Error("Expected 7, got ", points)
	}
}

func TestPointsForRoundTotalRoundNumber(t *testing.T) {
	var points = pointsForRoundTotal(100)
	if points != 50 {
		t.Error("Expected 50, got ", points)
	}
}

func TestPointsForRoundTotalDecimalNumber(t *testing.T) {
	var points = pointsForRoundTotal(100.1)
	if points != 0 {
		t.Error("Expected 0, got ", points)
	}
}

func TestPointsForDivisibleByOneQuarterTrue(t *testing.T) {
	var points = pointsForDivisibleByOneQuarter(100.25)
	if points != 25 {
		t.Error("Expected 25, got ", points)
	}
}

func TestPointsForDivisibleByOneQuarterFalse(t *testing.T) {
	var points = pointsForDivisibleByOneQuarter(100.1)
	if points != 0 {
		t.Error("Expected 0, got ", points)
	}
}

func TestPointsForTotalAmountRoundNumberDivisibleByAQuarter(t *testing.T) {
	var points = pointsForTotalAmount(100)
	if points != 75 {
		t.Error("Expected 75, got ", points)
	}
}

func TestPointsForTotalAmountDecimalNumberDivisibleByAQuarter(t *testing.T) {
	var points = pointsForTotalAmount(100.25)
	if points != 25 {
		t.Error("Expected 25, got ", points)
	}
}

func TestPointsForTotalAmountDecimalNumberNotDivisibleByAQuarter(t *testing.T) {
	var points = pointsForTotalAmount(100.26)
	if points != 0 {
		t.Error("Expected 0, got ", points)
	}
}

func TestPointsForListItemPairsOddNumberOfItems(t *testing.T) {
	var points = pointsForListItemPairs([]listItemDO{{}, {}, {}, {}, {}})
	if points != 10 {
		t.Error("Expected 10, got ", points)
	}
}

func TestPointsForListItemPairsEvenNumberOfItems(t *testing.T) {
	var points = pointsForListItemPairs([]listItemDO{{}, {}, {}, {}, {}, {}})
	if points != 15 {
		t.Error("Expected 15, got ", points)
	}
}

func TestPointsForListItemDescriptionsshortDescriptionDivisibleBy3(t *testing.T) {
	var points = pointsForListItemDescriptions([]listItemDO{listItemshortDescriptionDivisibleBy3})
	if points != 3 {
		t.Error("Expected 3, got ", points)
	}
}

func TestPointsForListItemDescriptionsshortDescriptionNotDivisibleBy3(t *testing.T) {
	var points = pointsForListItemDescriptions([]listItemDO{listItemshortDescriptionNotDivisibleBy3})
	if points != 0 {
		t.Error("Expected 0, got ", points)
	}
}

func TestPointsForListItemEvenNumberOfItemsAndDescriptionDivisibleBy3(t *testing.T) {
	var points = pointsForListItems([]listItemDO{listItemshortDescriptionDivisibleBy3, listItemshortDescriptionDivisibleBy3})
	if points != 11 {
		t.Error("Expected 11, got ", points)
	}
}

func TestPointsForListItemEvenNumberOfItemsAndDescriptionNotDivisibleBy3(t *testing.T) {
	var points = pointsForListItems([]listItemDO{listItemshortDescriptionNotDivisibleBy3, listItemshortDescriptionNotDivisibleBy3})
	if points != 5 {
		t.Error("Expected 5, got ", points)
	}
}

func TestPointsForPurchasedTime12PM(t *testing.T) {
	var points = pointsForPurchasedTime(12)
	if points != 0 {
		t.Error("Expected 0, got ", points)
	}
}

func TestPointsForPurchasedTime2PM(t *testing.T) {
	var points = pointsForPurchasedTime(14)
	if points != 10 {
		t.Error("Expected 10, got ", points)
	}
}

func TestPointsForPurchasedTime24(t *testing.T) {
	var points = pointsForPurchasedTime(16)
	if points != 0 {
		t.Error("Expected 0, got ", points)
	}
}

func TestPointsForPurchasedTime8AM(t *testing.T) {
	var points = pointsForPurchasedTime(8)
	if points != 0 {
		t.Error("Expected 10, got ", points)
	}
}

func TestPointsForPurchasedDateEven(t *testing.T) {
	var points = pointsForPurchasedDate(2)
	if points != 0 {
		t.Error("Expected 0, got ", points)
	}
}

func TestPointsForPurchasedDateOdd(t *testing.T) {
	var points = pointsForPurchasedDate(21)
	if points != 6 {
		t.Error("Expected 6, got ", points)
	}
}

func TestCalculatePointsForReceipt(t *testing.T) {
	var receipt = receiptDO{
		retailer:       "Target",
		purchasedDay:   1,
		purchasedMonth: 1,
		purchasedYear:  2022,
		purchasedHour:  13,
		purchasedMin:   1,
		total:          35.35,
		items: []listItemDO{
			{
				shortDescription: "Mountain Dew 12PK",
				price:            6.49,
			}, {
				shortDescription: "Emils Cheese Pizza",
				price:            12.25,
			}, {
				shortDescription: "Knorr Creamy Chicken",
				price:            1.26,
			}, {
				shortDescription: "Doritos Nacho Cheese",
				price:            3.35,
			}, {
				shortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
				price:            12.00,
			},
		},
	}
	receipt.calculatePointsForReceipt()
	if receipt.points != 28 {
		t.Error("Expected 28, got ", receipt)
	}
}
