package numeric

import (
	"testing"
)

func TestToFixedDecimal(t *testing.T) {
	num := 1.5346
	t.Log(ToFixedDecimal(num, 3)) // print 1.525
	t.Log(ToFixedDecimal(num, 0)) //print 2
}

func TestFloorToFixedDecimal(t *testing.T) {
	num := 1.5346
	t.Log(FloorToFixedDecimal(num, 3)) //print 1.534
	t.Log(FloorToFixedDecimal(num, 0)) //print 1
}

func TestCeilToFixedDecimal(t *testing.T) {
	num := 1.0341
	t.Log(CeilToFixedDecimal(num, 3)) //print 1.035
	t.Log(CeilToFixedDecimal(num, 0)) //print 2
}

func TestToInt64(t *testing.T) {
	t.Log(ToInt64(12.1))
	t.Log(ToInt64("0xaf"))
}

func TestToUint64(t *testing.T) {
	t.Log(ToUint64(123))
	t.Log(ToUint64(123.23))
	t.Log(ToUint64("123"))
	t.Log(ToUint64(0x123))
}