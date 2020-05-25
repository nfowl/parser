package siso

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
)

func ReadXmlFile(path string) (SISOFile, error) {
	xmlFile, err := os.Open(path)
	var sisoParsed SISOFile
	if err != nil {
		log.Printf("Error Encountered reading xml %v", err)
		return sisoParsed, err
	}
	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)
	err = xml.Unmarshal(byteValue, &sisoParsed)
	if err != nil {
		log.Printf("Error Encountered reading xml %v", err)
	}
	return sisoParsed, err
}
