package api

import (
	"fmt"
	"time"
)

// 地址 -------------------

func formatAddr(region string, detail string) string {
	return region + detail
}

// 时间 ---------------------------
const timeLayout = "2006-01-02 15:04:05"
const dateLayout = "2006-01-02"

func formatTime(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	return t.Format(timeLayout)
}

func formatDate(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	return t.Format(dateLayout)
}

// 金额与货币 ---------------------

// 不含货币单位
func formatMoney(amt int64) string {
	res := float64(amt) / float64(100)
	return fmt.Sprintf("%.2f", res)
}
