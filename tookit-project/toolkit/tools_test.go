package toolkit

import "testing"

func TestTools_randonString(t *testing.T) {

	var testTools Tools
	s := testTools.randomString(10)
	if len(s) != 10 {
		t.Error("wrong length random string returned")
	}

}
