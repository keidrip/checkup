package commands

import (
//	"fmt"
//	"unidriver/Godeps/_workspace/src/github.com/tebeka/selenium"
//	"unidriver/errors"
)

func init() {
	CommandList["sendKeys"] = sendKeys
}

func sendKeys(args interface{}) {

	//       for t, v := range args.([]interface{}) {

	//target := t.(string)
	//value := v.(string)

	//		fmt.Print("[sendKeys]: " + target + " = " + value + " ")
	//		elem, err := WD.FindElement(selenium.ByCSSSelector, target)
	//		elem.Clear()
	//		elem.SendKeys(value)
	//
	//		errors.Fatal(err)
	//		fmt.Print("SUCCESS")
	//		fmt.Println("")
	//      }

}