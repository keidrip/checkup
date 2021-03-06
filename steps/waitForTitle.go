package steps

import (
	"fmt"
	"time"
)

func init() {
	StepList["waitForTitle"] = waitForTitle
}

func waitForTitle() {

	fmt.Print("[waitForTitle]: " + Arg1)

	limit := SetStepTimeout("")
	latency := 0
	for {
		if limit < latency {
			StepFailure("It reached a time-out.")
			break
		}

		url, _ := WD.Title()
		if url == Arg1 {
			StepSuccess()
			break
		}

		time.Sleep(time.Millisecond * 500)
		latency += 500
	}

}
