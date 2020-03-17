package main

import (
	"fmt"
	owm "github.com/guaychou/openweatherapi"
	"os"
	"strconv"
	"strings"
)
var kotatmp string
var api_key = os.Getenv("OWM_TOKEN_API")
func cuaca(message string)string{
	split:=strings.Split(message," ")
	if len(split)<2{
		return "Some argument is missing. \nUse /cuaca <namaKota> to get the value."
	}else if len(split)>=2 {
		for i := 1;  i<len(split); i++ {
			fmt.Println(split[i])
			kotatmp+=split[i]+"%20"
		}
		kota:=kotatmp
		kotatmp=""
		result:=owm.GetWeather(kota,api_key)
		if result.Cod!=200{
			return "City not found."
		}else {
			city:=result.Name
			humidity:=strconv.Itoa(result.Humidity)
			description:=result.Weather[0].Description
			suhu:=fmt.Sprintf("%.2f",result.Temp)
			suhuMaks:=fmt.Sprintf("%.2f",result.Temp_max)
			suhuMin:=fmt.Sprintf("%.2f",result.Temp_min)
			kelembaban:=result.Kelembapan
			return "Kota: "+city+"\nCuaca: "+description+"\nSuhu: "+suhu+" °C\nSuhu Minimal: "+suhuMin+" °C\nSuhu Maksimal: "+suhuMaks+" °C\nAngka Kelembaban: "+humidity+"\nStatus Kelembaban: "+kelembaban
		}
	}
	return "Error: Something goes wrong"
}