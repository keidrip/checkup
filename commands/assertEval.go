package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["assertEval"] = assertEval
}

func assertEval(_url interface{}) {

	url := _url.(string)

	fmt.Print("[assertEval]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}