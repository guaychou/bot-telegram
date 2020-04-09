package main

import (
	ex "github.com/guaychou/exchangeAPI"
	"strconv"
	"strings"
)

func exchange(message string) string {
	message=goUppercase(message)
	split := strings.Split(message, " ")
	if len(split) == 3 {
		result,err := ex.GetExchange(split[1], split[2])
		if err!=nil {
			return err.Error()
		}
		strrate := strconv.Itoa(result.Rates[split[2]].(int))
		return "*INFORMATION*\n\nDate : "+result.Date+"\nFrom : " + result.Base + "\nTo : " + split[2] + "\nResult : "+ strrate
	}else {
		return "Please use /kurs <fromKurs> <toKurs>"
	}
}