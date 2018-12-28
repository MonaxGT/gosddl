package gosddl

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"encoding/json"
	"github.com/pkg/errors"
)

// ACLProcessor main struct with methods
type ACLProcessor struct {
	Rights permissons
	File   string
}

type entryACL struct {
	AccountSid        string   `json:"accountSID,omitempty"`
	AceType           string   `json:"aceType,omitempty"`
	AceFlags          []string `json:"aceFlags,omitempty"`
	Rights            []string `json:"rights,omitempty"`
	ObjectGUID        string   `json:"objectGUID,omitempty"`
	InheritObjectGUID string   `json:"inheritObjectGUID,omitempty"`
}

type permissons struct {
	Owner     string     `json:"owner,omitempty"`
	Primary   string     `json:"primary,omitempty"`
	Dacl      []entryACL `json:"dacl,omitempty"`
	DaclInher []string   `json:"daclInheritFlags,omitempty"`
	Sacl      []entryACL `json:"sacl,omitempty"`
	SaclInger []string   `json:"saclInheritFlags,omitempty"`
}

// checkSIDsFile check file of SIDs where data saved in SID,User
func checkSIDsFile(filePath string, sid string) string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Split(scanner.Text(), ",")[0] == sid {
			return strings.Split(scanner.Text(), ",")[1]
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return sid
}

// sidReplace replace identification account: sid/wellkhownsid/usersid
func (app *ACLProcessor) sidReplace(str string) string {
	if len(str) > 2 {

		if x, ok := sddlWellKnownSidsRep[str]; ok {
			return x
		} else if app.File != "" {
			return checkSIDsFile(app.File, str)
		}
		return str
	}
	return app.replacer(sddlSidsRep, str)[0]
}

// replacer chunk string with 2 letters, add to array and then resolve
func (app *ACLProcessor) replacer(maps map[string]string, str string) []string {
	var temp, result []string
	if len(str) > 2 {
		for j := 0; j < len(str)-1; j = j + 2 {
			temp = append(temp, fmt.Sprintf("%s%s", string(str[j]), string(str[j+1])))
		}
	} else {
		temp = append(temp, str)
	}
	for _, v := range temp {
		if x, ok := maps[v]; ok {
			result = append(result, x)
		} else {
			result = append(result, v)
		}
	}
	return result
}

/* splitBodyACL Convert values from string to struct with replace strings
Base format Rights: (ace_type;ace_flags;rights;object_guid;inherit_object_guid;account_sid)
*/
func (app *ACLProcessor) splitBodyACL(str string) entryACL {
	splitACL := strings.Split(str, ";")
	return entryACL{
		AceType:           app.replacer(sddlAceType, splitACL[0])[0],
		AceFlags:          app.replacer(sddlAceFlags, splitACL[1]),
		Rights:            app.replacer(sddlRights, splitACL[2]),
		ObjectGUID:        splitACL[3],
		InheritObjectGUID: splitACL[4],
		AccountSid:        app.sidReplace(splitACL[5]),
	}
}

func (app *ACLProcessor) splitBody(body string) []entryACL {
	var entryACLInternalArr []entryACL
	for _, y := range strings.Split(body, "(") {
		if y != "" {
			ace := strings.TrimSuffix(y, ")")
			entryACLInternalArr = append(entryACLInternalArr, app.splitBodyACL(ace))
		}
	}
	return entryACLInternalArr
}

func (app *ACLProcessor) parseBody(body string) ([]string, []entryACL) {
	var inheritFlagArr []string
	var entryACLInternalArr []entryACL
	if strings.Index(body, "(") != 0 {
		inheritFlag := body[0:strings.Index(body, "(")]
		ace := body[strings.Index(body, "("):]
		if len(inheritFlag) > 2 {
			for j := 0; j < len(inheritFlag)-1; j = j + 2 {
				inheritFlagArr = append(inheritFlagArr, app.replacer(sddlInheritanceFlags, fmt.Sprintf("%s%s", string(inheritFlag[j]), string(inheritFlag[j+1])))[0])
			}
		}
		entryACLInternalArr = app.splitBody(ace)
	} else {
		entryACLInternalArr = app.splitBody(body)
	}
	return inheritFlagArr, entryACLInternalArr
}

func (app *ACLProcessor) parseSDDL(sddrArr []string) {
	for _, y := range sddrArr {
		sddlSplit := strings.Split(y, ":")
		letter := sddlSplit[0]
		body := sddlSplit[1]
		switch letter {
		case "O":
			app.Rights.Owner = app.sidReplace(body)
		case "G":
			app.Rights.Primary = app.sidReplace(body)
		case "D":
			app.Rights.DaclInher, app.Rights.Dacl = app.parseBody(body)
		case "S":
			app.Rights.SaclInger, app.Rights.Sacl = app.parseBody(body)
		default:
			log.Fatal("Unresolved group")
		}
	}

}

// slice SDDL create slice objects from str to array of strings
func (app *ACLProcessor) sliceSDDL(indecs []int, str string) {
	var sddlArr []string
	for i := 0; i < len(indecs)-1; i++ {
		sl := str[indecs[i]:indecs[i+1]]
		sddlArr = append(sddlArr, sl)
	}
	app.parseSDDL(sddlArr)
}

// FindGroupIndex used for find index of group Owner, Primary, DACL, SACL
func (app *ACLProcessor) findGroupIndex(str string) error {
	groups := []string{"O:", "G:", "D:", "S:"}
	var result []int
	for _, i := range groups {
		if strings.Index(str, i) != -1 {
			result = append(result, strings.Index(str, i))
		}
	}
	if result == nil {
		return errors.New("Can't find any group")
	}
	result = append(result, len(str))
	app.sliceSDDL(result, str)
	return nil
}

// Processor main function in gosddl package
func Processor(api bool, port string, file string) error {
	var app ACLProcessor
	app.File = file
	if api {
		fmt.Println("API Interface started on port", port)
		app.httpHandler(port)
	} else if flag.Args() != nil {
		err := app.findGroupIndex(flag.Args()[0])
		if err != nil {
			return err
		}
		body, err := json.Marshal(app.Rights)
		if err != nil {
			log.Fatal(err)
			return err
		}
		fmt.Println(string(body))
		return nil
	}
	log.Fatal("You should give me SDDL string or use API mode")
	return nil
}