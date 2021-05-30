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

/**
* Takes the path of the file as input, reads the file and then creates files for valid output and invlaid output
* @Params {string} path
*/

func ProcessTxtFile(path string) {

  	var filenameWithoutExt = helpers.GetFileNameWithoutExt(path, constants.Txt)
	
	  // for unique emailIds and employeeIds 
	emailMap := make(map[string]int)
	employeeIdMap :=  make(map[string]int)

	//opening a file to read with given path
	txtReadFile, err := os.Open(path)

	if err != nil {
		log.Fatalf("failed to open")
	}

	defer txtReadFile.Close()

	scanner := bufio.NewScanner(txtReadFile)

	scanner.Split(bufio.ScanLines)

	//creating validOutput and invalidOutput files to write the data
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

	//processing the header
	headerNameIndexes, headerIndexes := utils.ProcessHeader(header)

	for scanner.Scan() {
		text := scanner.Text()
		record := strings.Split(text, ",")

		// validating the record
		returnValues, validEntry := validators.Validator(record, headerNameIndexes, headerIndexes)

		// checking if the record is valid and no repeated emailId and EmployeeId
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
