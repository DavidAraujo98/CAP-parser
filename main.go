package cap

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

func readFileToCAP(filePath string) []byte{
    xmlFile, err := os.Open(filePath)
    if err != nil {
        fmt.Println(err)
    }
    defer xmlFile.Close()
    rawData, _ := ioutil.ReadAll(xmlFile)
    return rawData
}

func writeCAPToFile(rawData []byte, filePath string) { 
    ioutil.WriteFile(filePath, rawData, 0664)
}

func Parser(rawData []byte, alertType *Alert) {
    xml.Unmarshal(rawData, &alertType)
}

func Deparser(alertType Alert) []byte{
    rawData, _ := xml.Marshal(alertType) 
    return rawData
}
