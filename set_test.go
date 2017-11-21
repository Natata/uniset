package uniset

import "testing"

func TestBasicThreeFunc(t *testing.T) {
	us := make(Uniset)

	l := us.List()
	if len(l) != 0 {
		t.Error("expect 0, got ", len(l))
	}

	ok := us.Add("test")
	if !ok {
		t.Error("expect true, got", ok)
	}

	ok = us.Add("test")
	if ok {
		t.Error("expect false, got", ok)
	}

	l = us.List()
	if len(l) != 1 {
		t.Error("expect 1, got", len(l))
	}

	us.Remove("test")
	l = us.List()
	if len(l) != 0 {
		t.Error("expect 0, got ", len(l))
	}
}
