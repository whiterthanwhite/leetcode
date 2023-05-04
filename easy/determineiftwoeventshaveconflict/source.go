package determineiftwoeventshaveconflict

import (
	"time"
)

func haveConflict(event1 []string, event2 []string) bool {
	timeLayout := `15:04`

	t1, _ := time.Parse(timeLayout, event1[0])
	t2, _ := time.Parse(timeLayout, event1[1])

	t3, _ := time.Parse(timeLayout, event2[0])
	t4, _ := time.Parse(timeLayout, event2[1])

	if (t3.After(t1) && t3.Before(t2)) || (t4.After(t1) && t4.Before(t2)) {
		return true
	} else if t3.Equal(t1) || t3.Equal(t2) || t4.Equal(t1) || t4.Equal(t2) {
		return true
	} else if t1.After(t3) && t2.Before(t4) {
		return true
	}

	return false
}
