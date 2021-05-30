package main

import (
	"os"
	"path/filepath"
	"log"
	processors "theReader/services/processors"
	constants "theReader/services/constants"
	helpers "theReader/services/helpers"
)

// All the files that are to be processed currently reside in the folder input_file_folder in the current directory
func main() {

	log.Println("Getting folder path for input files")
	inputFileFolderPath := helpers.GetinputFileFolderPath(constants.InputFileFolder)
	log.Println("Input folder path : ", inputFileFolderPath)
	
	log.Println("Walking through the files sequentially")
    err := filepath.Walk(inputFileFolderPath, func(path string, info os.FileInfo, err error) error {

		log.Println("Cheking if file ", path, "can be processed")
		fileExtension, processorName, fileFormatSupported := helpers.CheckFileFormatSupported(path)
		// fileExtension, _, fileFormatSupported := helpers.CheckFileFormatSupported(path)

		if !fileFormatSupported {
			log.Println("File format not Supported yet ", fileExtension)
		} else {
			log.Println("Processing file ", path, "with ", processorName)

			switch fileExtension {
			case constants.Csv:
				processors.ProcessCSVfile(path)
			case constants.Txt:
				processors.ProcessTxtFile(path)
			// case constants.Excel:
			// 	constants.ProcessXLSXFile(path)
			default:
				log.Println("No Processor!! Strange")
			}
		
		}

        return nil
    })
    if err != nil {
        panic(err)
    }
}