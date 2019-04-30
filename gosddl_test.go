package gosddl

import (
	"testing"

	"io/ioutil"
	"os"
)

func TestProcessor(t *testing.T) {
	quit := make (chan bool)
	go func (){
		api := true
		url := ":8123"
		for {
			select {
			case <- quit:
				return
			default:
				err := Processor(api,url,"")
				if err != nil {
					t.Error("cant't run function Processor: ",err)
				}
			}

		}
		close(quit)
	}()
	quit <- true
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

func TestSidReplace(t *testing.T) {
	data := []byte("S-10-10,User\n")
	err := ioutil.WriteFile("test.txt", data, 0644)
	if err != nil {
		t.Error("can't write data test.txt", err)
	}
	str := checkSIDsFile("test.txt","S-10-10")
	if str == "User" {
		return
	}
	err = os.Remove("test.txt")
	if err != nil {
		t.Error("can't delete file", err)
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