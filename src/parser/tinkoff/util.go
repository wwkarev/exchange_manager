package tinkoff

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func generateKey(datetime time.Time, sum float32) string {
	return datetime.Format(KeyDateLayout) + "_" + fmt.Sprintf("%f", sum)
}

func getIsinFromDescription(description string) (string, bool) {
	result := ""
	isFound := false
	regex := regexp.MustCompile(`(\w*)\/`)
	submatches := regex.FindAllStringSubmatch(description, -1)
	if len(submatches) > 0 && len(submatches[0]) > 0 {
		result = submatches[0][1]
		isFound = true
	}
	return result, isFound
}

func convertStringToFloat32(numberStr string) float32 {
	numberStr = strings.Replace(numberStr, ",", ".", -1)
	sum64, err := strconv.ParseFloat(numberStr, 32)
	if err != nil {
		fmt.Println(err)
		panic("Cannot convert refill to float 32. Value: " + numberStr + ". ")
	}
	return float32(sum64)
}
