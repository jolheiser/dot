package dot

import (
	"strings"
	"testing"
)

func TestImport(t *testing.T) {
	defer func() {
		err := recover()
		errStr, ok := err.(string)
		if !ok {
			t.Log("This isn't the string we are looking for")
			t.FailNow()
		}
		if !strings.EqualFold(errStr, codeReviewComments) {
			t.Log("I'm honestly not sure where to go from here")
			t.FailNow()
		}
	}()
	Import()
}
