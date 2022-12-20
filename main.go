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
    byteValue, _ := ioutil.ReadAll(xmlFile)
    return byteValue
}

func writeCAPToFile(byteValue []byte, filePath string) { 
    ioutil.WriteFile(filePath, byteValue, 0664)
}

func Parser(byteValue []byte, alertType *Alert) {
    xml.Unmarshal(byteValue, &alertType)
}

func Deparser(alertType Alert) []byte{
    byteValue, _ := xml.Marshal(alertType) 
    return byteValue
}
