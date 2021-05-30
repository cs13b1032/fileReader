package processors

import (
	"encoding/csv"
    "os"
	"log"
	"io"
	constants "theReader/services/constants"
	helpers "theReader/services/helpers"
	utils "theReader/services/utils"
	validators "theReader/services/validators"
)

func ProcessCSVfile(path string){

	var filenameWithoutExt = helpers.GetFileNameWithoutExt(path, constants.Csv)

	emailMap := make(map[string]int)
	employeeIdMap :=  make(map[string]int)

	csvReadfile, readErr := os.Open(path)
	if readErr != nil {
		log.Fatalln("Couldn't open the csv file", readErr)
	}

	defer csvReadfile.Close()
	csvReader := csv.NewReader(csvReadfile)

	csvWriteValidOutputfile, writevalidOutputErr := os.Create(filenameWithoutExt+"_validOutput.csv")
	csvWriteinValidOutputfile, writeinValidOutputErr := os.Create(filenameWithoutExt+"_inValidOutput.csv")

	if writevalidOutputErr != nil {
		log.Fatalf("failed creating file: %s", writevalidOutputErr)
	}
	if writeinValidOutputErr != nil {
		log.Fatalf("failed creating file: %s", writeinValidOutputErr)
	}

	defer csvWriteValidOutputfile.Close()
	defer csvWriteinValidOutputfile.Close()

	csvValidOutputWriter := csv.NewWriter(csvWriteValidOutputfile)
	csvinValidOutputWriter := csv.NewWriter(csvWriteinValidOutputfile)
	
	header, err := csvReader.Read()
	if (err == io.EOF) || (err != nil) {
		log.Fatalf("Error in Reading Header")
	}

	headerNameIndexes, headerIndexes := utils.ProcessHeader(header)
	csvValidOutputWriter.Write(constants.Header)
	csvinValidOutputWriter.Write(constants.Header)

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		returnValues, validEntry := validators.Validator(record, headerNameIndexes, headerIndexes)
		
		if validEntry{
			_, okEmail := emailMap[returnValues[1]]
			_, okEmpId := employeeIdMap[returnValues[2]]

			if okEmail || okEmpId {
				csvinValidOutputWriter.Write(returnValues)
			}else{
				emailMap[returnValues[1]] = 1
				employeeIdMap[returnValues[2]] = 1
				csvValidOutputWriter.Write(returnValues)
			}
		}else{
			csvinValidOutputWriter.Write(returnValues)
		}
	}

	csvValidOutputWriter.Flush()
	csvinValidOutputWriter.Flush()
	
}