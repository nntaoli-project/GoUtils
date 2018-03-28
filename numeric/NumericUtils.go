package numeric

import (
	"fmt"
	"math"
	"strconv"
)

/**
 * 四舍五入
 * @param n 要保留的小数点位数
 */
func ToFixedDecimal(num float64, n int) string {
	return fmt.Sprintf("%."+fmt.Sprint(n)+"f", num)
}

func FloorToFixedDecimal(num float64, n int) string {

	var factor float64 = 1
	for i := n; i > 0; i-- {
		factor = factor * 10
	}

	return ToFixedDecimal(math.Floor(num*factor)/factor, n)
}

func CeilToFixedDecimal(num float64, n int) string {
	var factor float64 = 1
	for i := n; i > 0; i-- {
		factor = factor * 10
	}
	return ToFixedDecimal(math.Ceil(num*factor)/factor, n)
}

func ToFloat64(v interface{}) float64 {
	if v == nil {
		return 0.0
	}

	switch v.(type) {
	case float64:
		return v.(float64)
	case string:
		vStr := v.(string)
		vF, _ := strconv.ParseFloat(vStr, 64)
		return vF
	case int:
		return float64(v.(int))
	default:
		panic("to float64 error.")
	}
}

func ToInt(v interface{}) int {
	if v == nil {
		return 0
	}

	switch v.(type) {
	case string:
		vStr := v.(string)
		vInt, _ := strconv.Atoi(vStr)
		return vInt
	case int:
		return v.(int)
	case float64:
		vF := v.(float64)
		return int(vF)
	default:
		panic("to int error.")
	}
}

func ToUint64(v interface{}) uint64 {
	if v == nil {
		return 0
	}

	switch v.(type) {
	case int:
		return uint64(v.(int))
	case float64:
		return uint64(v.(float64))
	case string:
		uV, _ := strconv.ParseUint(v.(string), 10, 64)
		return uV
	default:
		panic("to uint64 error.")
	}
}
