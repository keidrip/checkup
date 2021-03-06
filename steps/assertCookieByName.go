package steps

import (
	"fmt"
)

func init() {
	StepList["assertCookieByName"] = assertCookieByName
}

func assertCookieByName() {

	fmt.Print("[assertCookieByName]: " + Arg1 + ", " + Arg2)

	c, err := WD.GetCookies()
	StepFailure(err)

	for _, cookie := range c {
		if cookie.Name == Arg1 && cookie.Value == Arg2 {
			StepSuccess()
			return
		}

	}

	AssertionFailure()
}
