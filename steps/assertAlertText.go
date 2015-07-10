package steps

import (
	"fmt"
	"unidriver/Godeps/_workspace/src/github.com/mattn/go-scan"
)

func init() {
	StepList["assertAlertText"] = assertAlertText
}

func assertAlertText(a interface{}) {

	var target string
	scan.ScanTree(a, "/target", &target)

	fmt.Print("[assertAlertText]: " + target)
	text, err := WD.AlertText()
	StepFailure(err)

	if text == target {
		StepSuccess()
	} else {
		AssertionFailure()
	}

}