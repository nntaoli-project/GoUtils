package GoUtils

import (
	"testing"
	"time"
)

func TestDateFormat(t *testing.T) {
	t.Log(DateFormat(1511765558806649344, time.Nanosecond, "yyyy/MM/dd HH:mm:ss")) // output 2017/11/27 14:52:38
}

func TestDateFormat2(t *testing.T) {
	t.Log(DateFormat2(time.Now(), "yyyy-MM-dd HH:mm"))
}
