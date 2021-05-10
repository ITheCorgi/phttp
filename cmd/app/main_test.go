package main

import (
	"os"
	"reflect"
	"testing"
)

func TestParsingCliArgs(t *testing.T) {
	testcases := []struct {
		name     string
		args     []string
		expected []string
	}{
		{name: "test-case-1", args: []string{"cmd-arg0", "http://google.com", "http://yandex.com", "http://facebook.com"}, expected: []string{"http://google.com", "http://yandex.com", "http://facebook.com"}},
		{name: "test-case-2", args: []string{"cmd-arg0"}, expected: []string{}},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			os.Args = testcase.args
			var got = GetFlags()
			if !reflect.DeepEqual(testcase.expected, got) {
				t.Errorf("Failed! Expected: %s, Got: %s", testcase.expected, got)
			}
		})
	}
}
