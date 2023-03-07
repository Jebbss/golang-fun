package main

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

type listItemDO struct {
	shortDescription string
	price            float32
}

type receiptDO struct {
	points         int
	retailer       string
	purchasedYear  int
	purchasedMonth int
	purchasedDay   int
	purchasedHour  int
	purchasedMin   int
	items          []listItemDO
	total          float32
}

func (r *receiptDTO) toReceiptDO() (receiptDO, error) {
	var receipt = receiptDO{
		points:         0,
		retailer:       r.Retailer,
		purchasedYear:  0,
		purchasedMonth: 0,
		purchasedDay:   0,
		purchasedHour:  0,
		purchasedMin:   0,
		items:          r.toListItemDO(),
		total:          r.Total,
	}

	year, month, day, err := getPurchasedDateNumbers(r.PurchasedDate)
	if err != nil {
		return receipt, err
	}

	hour, min, err := getPurchasedTime(r.PurchasedTime)
	if err != nil {
		return receipt, err
	}
	receipt.purchasedYear = year
	receipt.purchasedMonth = month
	receipt.purchasedDay = day
	receipt.purchasedHour = hour
	receipt.purchasedMin = min
	receipt.calculatePointsForReceipt()
	return receipt, nil
}

func getPurchasedDateNumbers(purchasedDate string) (int, int, int, error) {
	splitDate := strings.Split(purchasedDate, "-")
	if len(splitDate) != 3 {
		return 0, 0, 0, errors.New("invalid date format")
	}
	year, err := strconv.Atoi(splitDate[0])
	if err != nil {
		return 0, 0, 0, err
	}
	month, err := strconv.Atoi(splitDate[1])
	if err != nil {
		return 0, 0, 0, err
	}
	day, err := strconv.Atoi(splitDate[2])
	if err != nil {
		return 0, 0, 0, err
	}
	return year, month, day, nil
}

func getPurchasedTime(purchasedTime string) (int, int, error) {
	splitTime := strings.Split(purchasedTime, ":")
	if len(splitTime) != 2 {
		return 0, 0, errors.New("invalid time format")
	}
	hour, err := strconv.Atoi(splitTime[0])
	if err != nil {
		return 0, 0, err
	}
	if hour < 0 || hour > 23 {
		return 0, 0, errors.New("invalid time format")
	}
	min, err := strconv.Atoi(splitTime[1])
	if err != nil {
		return 0, 0, err
	}
	if min < 0 || min > 59 {
		return 0, 0, errors.New("invalid time format")
	}
	return hour, min, nil
}

func (r *receiptDO) calculatePointsForReceipt() {
	r.points += pointsForRetailerName(r.retailer)
	r.points += pointsForTotalAmount(r.total)
	r.points += pointsForListItems(r.items)
	r.points += pointsForPurchasedTime(r.purchasedHour)
	r.points += pointsForPurchasedDate(r.purchasedDay)
}

func pointsForPurchasedTime(purchaseHour int) int {
	if purchaseHour >= 14 && purchaseHour < 16 {
		return 10
	}
	return 0
}

func pointsForPurchasedDate(purchasedDay int) int {
	if purchasedDay%2 == 1 {
		return 6
	}
	return 0
}

func pointsForTotalAmount(total float32) int {
	points := pointsForRoundTotal(total)
	points += pointsForDivisibleByOneQuarter(total)
	return points
}

func pointsForListItems(listItems []listItemDO) int {
	points := pointsForListItemPairs(listItems)
	points += pointsForListItemDescriptions(listItems)
	return points
}

func pointsForListItemDescriptions(listItems []listItemDO) int {
	var points = 0
	for _, item := range listItems {
		if len(strings.TrimSpace(item.shortDescription))%3 == 0 {
			points += int(math.Ceil(float64(item.price * 0.2)))
		}
	}
	return points
}

func pointsForListItemPairs(listItems []listItemDO) int {
	return 5 * (len(listItems) / 2)
}

func pointsForDivisibleByOneQuarter(total float32) int {
	var points = 0
	if total*4 == float32(int(total*4)) {
		points += 25
	}
	return points
}

func pointsForRoundTotal(total float32) int {
	var points = 0
	if total == float32(int(total)) {
		points += 50
	}
	return points
}

func pointsForRetailerName(retailerName string) int {
	var points = 0
	for _, char := range retailerName {
		if isAlphanumeric(char) {
			points++
		}
	}
	return points
}

func isAlphanumeric(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')
}

func (r *receiptDTO) toListItemDO() []listItemDO {
	var items []listItemDO
	for _, item := range r.Items {
		items = append(items, listItemDO{
			shortDescription: item.ShortDescription,
			price:            item.Price,
		})
	}
	return items
}
