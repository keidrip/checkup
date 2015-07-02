package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["storeBodyText"] = storeBodyText
}

func storeBodyText(_url interface{}) {

	url := _url.(string)

	fmt.Print("[storeBodyText]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}