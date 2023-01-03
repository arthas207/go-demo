package enums

import "fmt"

type WeekDay int

const (
	MONDAY WeekDay = iota + 1
	TUESDAY
	WEDNESDAY
	THUSDAY
	FRIDAY
	SATURDAY
	SUNDAY
)

func Example() {
	weekDays := make([]WeekDay, int(SUNDAY))
	for i := 0; i < int(SUNDAY); i++ {
		weekDays[i] = WeekDay(i + 1)
	}
	fmt.Println("weekdays:", weekDays)
	// stringer -type=Pill代码生成
}
