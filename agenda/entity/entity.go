package entity

import (
	"encoding/json"
	"fmt"
	"os"
)

//ShowHelp is a simple help function
func ShowHelp() {
	fmt.Println("Let me tell you how to use agenda!")
}

type userInfo struct {
	Un, Pw, Em, Ph string
}

//Reg is register function
func Reg(un string, pw string, em string, ph string) int {
	file, err := os.OpenFile("User", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(file)
	for decoder.More() {
		var uif userInfo
		decoder.Decode(&uif)
		if uif.Un == un {
			file.Close()
			return 1
		}
	}
	info := make(map[string]string)
	info["Un"], info["Pw"], info["Em"], info["Ph"] = un, pw, em, ph
	encoder := json.NewEncoder(file)
	encoder.Encode(info)
	file.Close()
	return 0

}

//Log in function , if you want to log out, just delete ./CurUser
func Log(un, pw string) int {
	_, err1 := os.Stat("CurUser")
	if !os.IsNotExist(err1) {
		return 2
	}
	file, err := os.OpenFile("User", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	for decoder.More() {
		var uif userInfo
		decoder.Decode(&uif)
		if uif.Un == un && uif.Pw == pw {
			cur, err2 := os.OpenFile("CurUser", os.O_RDWR|os.O_CREATE, 0666)
			if err2 != nil {
				panic(err2)
			}
			info := make(map[string]string)
			info["Un"], info["Pw"], info["Em"], info["Ph"] = uif.Un, uif.Pw, uif.Em, uif.Ph
			encoder := json.NewEncoder(cur)
			encoder.Encode(info)
			cur.Close()
			return 0
		}
	}
	return 1
}
