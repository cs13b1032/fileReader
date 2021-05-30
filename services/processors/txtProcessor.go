package processors

import (
	"bufio"
	"log"
	"os"
	"strings"
	constants "theReader/services/constants"
	validators "theReader/services/validators"
	helpers "theReader/services/helpers"
	utils "theReader/services/utils"
)

func ProcessTxtFile(path string) {

  	var filenameWithoutExt = helpers.GetFileNameWithoutExt(path, constants.Txt)

	emailMap := make(map[string]int)
	employeeIdMap :=  make(map[string]int)
	txtReadFile, err := os.Open(path)

	if err != nil {
		log.Fatalf("failed to open")
	}

	defer txtReadFile.Close()

	scanner := bufio.NewScanner(txtReadFile)

	scanner.Split(bufio.ScanLines)

	txtWriteValidOutputfile, writevalidOutputErr := os.Create(filenameWithoutExt+"_validOutput.txt")
	txtWriteinValidOutputfile, writeinValidOutputErr := os.Create(filenameWithoutExt+"_inValidOutput.txt")

  	if writevalidOutputErr != nil {
		log.Fatalf("failed creating file: %s", writevalidOutputErr)
	}
	if writeinValidOutputErr != nil {
		log.Fatalf("failed creating file: %s", writeinValidOutputErr)
	}
	defer txtWriteValidOutputfile.Close()
	defer txtWriteinValidOutputfile.Close()

	txtValidOuptputWriter := bufio.NewWriter(txtWriteValidOutputfile)
	txtinValidOuptputWriter := bufio.NewWriter(txtWriteinValidOutputfile)

	headerString := strings.Join(constants.Header, "," )
	txtValidOuptputWriter.WriteString(headerString+ "\n")
	txtinValidOuptputWriter.WriteString(headerString+ "\n")

	if !scanner.Scan() {
		log.Fatalf("Error in Reading Header")
	}
	header := strings.Split(scanner.Text(), ",")
	headerNameIndexes, headerIndexes := utils.ProcessHeader(header)

	for scanner.Scan() {
		text := scanner.Text()
		record := strings.Split(text, ",")
		returnValues, validEntry := validators.Validator(record, headerNameIndexes, headerIndexes)

		if validEntry{
			_, okEmail := emailMap[returnValues[1]]
			_, okEmpId := employeeIdMap[returnValues[2]]

			if okEmail || okEmpId {
				txtinValidOuptputWriter.WriteString( strings.Join(returnValues, "," ) + "\n")
			}else{
				emailMap[returnValues[1]] = 1
				employeeIdMap[returnValues[2]] = 1
				txtValidOuptputWriter.WriteString( strings.Join(returnValues, "," ) + "\n")
			}
		}else{
			txtinValidOuptputWriter.WriteString( strings.Join(returnValues, "," ) + "\n")
		}
	}
	txtValidOuptputWriter.Flush()
	txtinValidOuptputWriter.Flush()
}
