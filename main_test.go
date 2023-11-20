package main

import "testing"

func TestHello(t *testing.T) {
	///dckr_pat_iEnUI5hWoG3hiiJR8Ep2G016Kig
	want := "Hello Golang!"
	got := hello()

	if want != got {
		t.Fatalf("want %s, got %s\n", want, got)
	}
}
