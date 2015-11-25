package binmsg

import (
	"bytes"
	"fmt"
	"strings"
)

func errorCheck(i int, got, want error) (s string, ok bool) {
	var (
		gotStr  string
		wantStr string
	)

	s = ""
	ok = true

	if got != nil {
		gotStr = fmt.Sprintf("%s", got)
	} else {
		gotStr = "<nil>"
	}
	if want != nil {
		wantStr = fmt.Sprintf("%s", want)
	} else {
		wantStr = "<nil>"
	}
	if strings.Compare(wantStr, gotStr) != 0 {
		s = fmt.Sprintf("Test: %d::\n Got:\n\t '%s',\n Want:\n\t '%s'\n",
			i, gotStr, wantStr)
		ok = false
	}
	return
}

func byteSliceCheck(i int, got, want []byte) (s string, ok bool) {
	if bytes.Compare(got, want) != 0 {
		s = fmt.Sprintf("Test: %d::\n Got:\n\t '%#v'\n, Want:\n\t '%#v'\n", i, got, want)
		ok = false
	} else {
		s = ""
		ok = true
	}
	return
}

func float32Check(i int, got, want float32) (s string, ok bool) {
	if got != want {
		s = fmt.Sprintf("Test %d::\n\tWant: %f\n\tGot: %f", i, got, want)
		ok = false
	} else {
		s = ""
		ok = true
	}
	return
}

func mkIntErrString(name string, got int, want int) string {
	return fmt.Sprintf("%s: Got: %d, but want: %d", name, got, want)
}

func mkStrErrString(name, got, want string) string {
	return fmt.Sprintf("%s: Got: %s, but want: %s", name, got, want)
}
