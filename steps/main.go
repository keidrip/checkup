package steps

import (
	"checkup/Godeps/_workspace/src/github.com/k0kubun/pp"
	"checkup/Godeps/_workspace/src/github.com/mattn/go-scan"
	"checkup/Godeps/_workspace/src/github.com/tebeka/selenium"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

var (
	WD               selenium.WebDriver
	StepList         = map[string]func(){}
	Browser          = "firefox"
	DefaultTimeout   = "60000"
	Arg1, Arg2, Arg3 string
	BaseUrl          string
	SeleniumSelector string
	Datas            map[interface{}]interface{}
)

func Init(remote string, data interface{}) {

	scan.ScanTree(data, "/testcase[0]/browser/", &Browser)
	scan.ScanTree(data, "/testcase[0]/baseurl/", &BaseUrl)
	scan.ScanTree(data, "/testcase[0]/selector/", &SeleniumSelector)

	caps := selenium.Capabilities{"browserName": Browser}
	wd, err := selenium.NewRemote(caps, remote)
	WD = wd
	WDFatal(err)
	defer WD.Quit()

	SetAroundTimeout(data)
	SetSelector()

	Do(data)
}

func Do(data interface{}) {

	steps := data.(map[interface{}]interface{})["testcase"].([]interface{})[0].(map[interface{}]interface{})["steps"].([]interface{})

	for _, ss := range steps {

		stepSet := NormalizeStepSet(ss)

		for s, a := range stepSet.(map[interface{}]interface{}) {

			step := s.(string)
			Arg1, Arg2, Arg3 = "", "", ""
			var arg1, arg2, arg3 interface{}

			scan.ScanTree(a, "/target", &arg1)
			scan.ScanTree(a, "/value", &arg2)
			scan.ScanTree(a, "/option", &arg3)
			Arg1 = SimplifyTypeAttributeValue(arg1).(string)
			Arg2 = SimplifyTypeAttributeValue(arg2).(string)
			Arg3 = SimplifyTypeAttributeValue(arg3).(string)

			if _, x := a.(string); x {
				Arg1 = a.(string)
			}

			StepList[step]()
		}
	}
}

const SCRIPT_getElementsByXPath = `document.getElementsByXPath = function(expression, parentElement) {
  var r = [];
  var x = document.evaluate(expression, parentElement || document, null, XPathResult.ORDERED_NODE_SNAPSHOT_TYPE, null);
  for (var i = 0, l = x.snapshotLength; i < l; i++) {
    r.push(x.snapshotItem(i));
  }
  return r;
}
return `

func WDFatal(err error) {
	if err != nil {
		fmt.Println("")
		fmt.Println(err)
		os.Exit(WDQuit())
	}
}

func WDQuit() int {
	defer WD.Quit()
	return 2
}

func StepFailure(err interface{}) {

	if err != nil {
		fmt.Print(" Failure")
		fmt.Println("")
		fmt.Println(err)
		os.Exit(WDQuit())
	}

}

func StepSuccess() {

	fmt.Print(" Success")
	fmt.Println("")

}

func AssertionFailure() {

	fmt.Print(" Assertion Failure")
	fmt.Println("")
	os.Exit(WDQuit())

}

func VerificationFailure() {

	fmt.Print(" Verification Failure")
	fmt.Println("")

}

func SetStepTimeout(value string) int {

	if value == "" {
		value = DefaultTimeout
	}
	var err1 error
	limit, err1 := strconv.Atoi(value)
	StepFailure(err1)

	return limit
}

func SetSelector() {
	if m, _ := regexp.MatchString("(?i)^xpath$", SeleniumSelector); m {
		SeleniumSelector = "xpath"
	} else if m, _ := regexp.MatchString("(?i)^css ?selector$", SeleniumSelector); m {
		SeleniumSelector = "css selector"
	} else {
		SeleniumSelector = "xpath"
	}
}

func SetAroundTimeout(data interface{}) {

	var timeout time.Duration
	scan.ScanTree(data, "/testcase[0]/timeout/", &timeout)
	if timeout == 0 {
		i, err1 := strconv.Atoi(DefaultTimeout)
		WDFatal(err1)
		timeout = time.Duration(i)
	} else {
		DefaultTimeout = strconv.Itoa(int(timeout))
	}
	err2 := WD.SetImplicitWaitTimeout(timeout)
	WDFatal(err2)
}

func SimplifyTypeAttributeTarget(o interface{}) string {
	var v string
	switch o.(type) {
	case float64:
		v = strconv.Itoa(int(o.(float64)))
	case float32:
		v = strconv.Itoa(int(o.(float32)))
	case int:
		v = strconv.Itoa(o.(int))
	case bool:
		v = strconv.FormatBool(o.(bool))
	case map[interface{}]interface{}:
		// none
	case nil:
		v = ""
	default:
		v = o.(string)
	}
	return v
}
func SimplifyTypeAttributeValue(o interface{}) interface{} {
	var v interface{}
	switch o.(type) {
	case float64:
		v = strconv.Itoa(int(o.(float64)))
	case float32:
		v = strconv.Itoa(int(o.(float32)))
	case int:
		v = strconv.Itoa(o.(int))
	case bool:
		v = strconv.FormatBool(o.(bool))
	case nil:
		v = ""
	default:
		v = o
	}
	return v
}
func NestingTarget(o interface{}) interface{} {

	switch o.(type) {
	case int, string, bool:
		o = map[interface{}]interface{}{"target": o}
	}
	return o
}
func NormalizeStepSet(o interface{}) interface{} {
	var result map[interface{}]interface{}
	switch o.(type) {
	case map[interface{}]interface{}:
		result = o.(map[interface{}]interface{})
	default:
		result = map[interface{}]interface{}{o: nil}
	}
	return result
}

func PrintStep(name string, a ...interface{}) {

	var target, value interface{}

	ok1 := scan.ScanTree(a, "/[0]/target", &target)
	ok2 := scan.ScanTree(a, "/[0]/value", &value)
	pp.Println(a)
	pp.Println(target)
	pp.Println(ok1)
	pp.Println(value)
	pp.Println(ok2)
	os.Exit(9)

	for i, arg := range a {
		pp.Println("----")
		pp.Println(i)
		pp.Println(arg)
	}
	os.Exit(9)
}
