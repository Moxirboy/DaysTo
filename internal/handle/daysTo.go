package handle

import (
	_const "awesomeProject4/internal/const"
	"fmt"
	"net/http"
	"time"
)

func DaysTo(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()
	targetDate := time.Date(_const.TargetYear, _const.TargetMonth, _const.TargetDay, 0, 0, 0, 0, time.UTC)

	daysRemaining := int(targetDate.Sub(currentTime).Hours() / 24)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Days remaining until January 1, 2025: %d", daysRemaining)
}
