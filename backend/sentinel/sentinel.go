package sentinel

import "fmt"

func Assert(mustBeTrue bool, otherwisePanicWith string) {
	if !mustBeTrue {
		panic(otherwisePanicWith)
	}
}

func AssertError(mustBeNil error, otherwisePanicWith string) {
	if mustBeNil != nil {
		msg := fmt.Sprintf("%s: %s", otherwisePanicWith, mustBeNil.Error())
		panic(msg)
	}
}
