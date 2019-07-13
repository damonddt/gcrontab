package gcrontab

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"testing"
	"time"
)

func TestCronParse(t *testing.T) {
	nextTime := cronexpr.MustParse("0 0 29 2 *").Next(time.Now())
	fmt.Printf("next time: %s\n", nextTime)

	expr, err := cronexpr.Parse("0 0 29 2 *")
	if err != nil {
		fmt.Println(err)
	}
	nextNextTime := expr.Next(nextTime)
	fmt.Printf("next next time: %s\n", nextNextTime)
}
