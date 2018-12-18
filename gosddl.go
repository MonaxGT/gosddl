package gosddl

import (
	"fmt"
	"log"
	"strings"
)

type entryACLInternal struct {
	AccountSid        string   `json:"accountsid"`
	AceType           string   `json:"aceType"`
	AceFlags          []string `json:"aceflags"`
	Rights            []string `json:"rights"`
	ObjectGuid        string   `json:"objectguid"`
	InheritObjectGuid string   `json:"InheritObjectGuid"`
}

type Permissons struct {
	Owner     string             `json:"owner"`
	Primary   string             `json:"primary"`
	Dacl      []entryACLInternal `json:"dacl"`
	DaclInher []string           `json:"daclInheritFlags"`
	Sacl      []entryACLInternal `json:"sacl"`
	SaclInger []string           `json:"saclInheritFlags"`
}

// replace identification account: sid/wellkhownsid/usersid
func sidReplace(str string) string {
	if len(str) > 2 {
		if x, ok := sddlWellKnownSidsRep[str]; ok {
			return x
		} else {
			return str
		}
		return replacer(sddlWellKnownSidsRep, str)[0]
	} else {
		return replacer(sddlSidsRep, str)[0]
	}
}

// Chunk string with 2 letters, add to array and then resolve
func replacer(maps map[string]string, str string) []string {
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

// Base format ACL: (ace_type;ace_flags;rights;object_guid;inherit_object_guid;account_sid)
// Convert values from string to struct with replace strings
func splitBodyACL(str string) entryACLInternal {
	temp := strings.Split(str, ";")
	return entryACLInternal{
		AceType:           replacer(sddlAceType, temp[0])[0],
		AceFlags:          replacer(sddlAceFlags, temp[1]),
		Rights:            replacer(sddlRights, temp[2]),
		ObjectGuid:        temp[3],
		InheritObjectGuid: temp[4],
		AccountSid:        sidReplace(temp[5]),
	}
}

func splitBody(body string) []entryACLInternal {
	var entryACLInternalArr []entryACLInternal
	for _, y := range strings.Split(body, "(") {
		if y != "" {
			ace := strings.TrimSuffix(y, ")")
			entryACLInternalArr = append(entryACLInternalArr, splitBodyACL(ace))
		}
	}
	return entryACLInternalArr
}

func (p *Permissons) parseBody(body string) ([]string, []entryACLInternal) {
	var inheritFlagArr []string
	var entryACLInternalArr []entryACLInternal
	if strings.Index(body, "(") != 0 {
		inheritFlag := body[0:strings.Index(body, "(")]
		ace := body[strings.Index(body, "("):]
		if len(inheritFlag) > 2 {
			for j := 0; j < len(inheritFlag)-1; j = j + 2 {
				inheritFlagArr = append(inheritFlagArr, replacer(sddlInheritanceFlags, fmt.Sprintf("%s%s", string(inheritFlag[j]), string(inheritFlag[j+1])))[0])
			}
		}
		entryACLInternalArr = splitBody(ace)
	} else {
		entryACLInternalArr = splitBody(body)
	}
	return inheritFlagArr, entryACLInternalArr
}

func (p *Permissons) parseSDDL(sddrArr []string) {
	for _, y := range sddrArr {
		sddlSplit := strings.Split(y, ":")
		letter := sddlSplit[0]
		body := sddlSplit[1]
		switch letter {
		case "O":
			p.Owner = sidReplace(body)
		case "G":
			p.Primary = sidReplace(body)
		case "D":
			p.DaclInher, p.Dacl = p.parseBody(body)
		case "S":
			p.SaclInger, p.Sacl = p.parseBody(body)
		default:
			log.Fatal("Unresolved group")
		}
	}

}

// create slice objects from str to array of strings
func (p *Permissons) sliceSDDL(indecs []int, str string) {
	var sddlArr []string
	for i := 0; i < len(indecs)-1; i++ {
		sl := str[indecs[i]:indecs[i+1]]
		sddlArr = append(sddlArr, sl)
	}
	p.parseSDDL(sddlArr)
}

func (p *Permissons) FindGroupIndex(str string) {
	groups := []string{"O:", "G:", "D:", "S:"}
	var result []int
	for _, i := range groups {
		if strings.Index(str, i) != -1 {
			result = append(result, strings.Index(str, i))
		}
	}
	result = append(result, len(str))
	p.sliceSDDL(result, str)
}
