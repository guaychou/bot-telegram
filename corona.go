package main

import (
	"fmt"
	cga "github.com/guaychou/corona-api"
	"strconv"
	"strings"
)

func corona(message string)string{
	split:=strings.Split(message," ")
	if len(split)<2{
		return "Some argument is missing. \nUse /corona <namaNegara> to get the value."
	}else if len(split)>=2 {
		country:=split[1]
		result:=cga.GetCorona(country)
		fmt.Println(result.Recovered.Value)
		if result.Recovered.Value==-1 || result.Deaths.Value==-1 || result.Confirmed.Value==-1{
			return "Country not found."
		}else {
			confirmed:=strconv.Itoa(result.Confirmed.Value)
			recovered:=strconv.Itoa(result.Recovered.Value)
			deaths:=strconv.Itoa(result.Deaths.Value)
			return "*INFORMASI*\n\nNegara: "+country+"\nJumlah terkonfirmasi: "+confirmed+"\nJumlah berhasil sembuh: "+recovered+"\nJumlah kematian: "+deaths+"\n\nStay safe :)"
		}
	}
	return "Error: Something goes wrong"
}