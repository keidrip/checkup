package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["verifyElementAttribute"] = verifyElementAttribute
}

func verifyElementAttribute(_url interface{}) {

	url := _url.(string)

	fmt.Print("[verifyElementAttribute]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}