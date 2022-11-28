package CAPParser

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

func Parser(inFilePath string, alertType *Alert) {
    xmlFile, err := os.Open(inFilePath)
    if err != nil {
        fmt.Println(err)
    }
    defer xmlFile.Close()
    byteValue, _ := ioutil.ReadAll(xmlFile)
    xml.Unmarshal(byteValue, &alertType)
}

func Deparser(alertType Alert, outFilePath string) {
    byteValue, _ := xml.Marshal(alertType)
    ioutil.WriteFile(outFilePath, byteValue, 0664)
}
