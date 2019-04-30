package gosddl

import (
	"testing"

	"io/ioutil"
	"os"
)

func TestProcessor(t *testing.T) {
	var app ACLProcessor
	testStr := "{O:WA,G:SA}"
	err := app.Processor(testStr)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestFindGroupIndex(t *testing.T) {
	var app ACLProcessor
	testStr := "{O:WA,G:SA}"
	err := app.findGroupIndex(testStr)
	if err != nil {
		t.Error(err)
		return
	}
}


func TestFindGroupIndex2(t *testing.T) {
	var app ACLProcessor
	testStr := "{O:WA,G:SA,D:(SA;DA;;;;DA),S:AI(SA;DA;;;;ST)}"
	err := app.findGroupIndex(testStr)
	if err != nil {
		t.Error(err)
		return
	}
}


func TestSidReplace(t *testing.T) {
	data := []byte("S-10-10,User\n")
	err := ioutil.WriteFile("test.txt", data, 0644)
	if err != nil {
		t.Error("can't write data test.txt", err)
		return
	}
	str := checkSIDsFile("test.txt","S-10-10")
	err = os.Remove("test.txt")
	if err != nil {
		t.Error("can't delete file", err)
		return
	}
	if str == "User" {
		return
	}
	t.Errorf("replaced name doesn't match result: %s",str)
}

func TestReplacer(t *testing.T) {
	var app ACLProcessor
	testStr := "S-1-5-2"
	str := app.sidReplace(testStr)
	if str == "Network" {
		return
	}
	t.Errorf("replaced name doesn't match result: %s",str)
}

func TestSplitBodyACL(t *testing.T) {
	var app ACLProcessor
	testStr := "SA;DA;;;;ST"
	result := app.splitBodyACL(testStr)
	if result.AccountSid == "" {
		t.Error("function return nil data")
		return
	}
}