package helpers

import (
    "os"
    "path/filepath"
	constants "theReader/services/constants"
)

/**
* Takes folder where the input files to be processed as input returns the total path to be processed
* @Params {string} inputFileFolder
* @Returns {string} inputFileFolderPath
**/
func GetinputFileFolderPath(inputFileFolder string)(string){

	// Get current directory Path
	mydir, nerr := os.Getwd()
	if nerr != nil {
		panic(nerr)
	}

	// get the input file folder in the current directory
	inputFileFolderPath := mydir + "/"+inputFileFolder+"/"

	return inputFileFolderPath
}

/**
* Takes fileName as the input and returns it's extension, processorName and if it can be processed
* @Params {string} fileName
* @Returns {string} fileExtension
* @Returns {string} processorName
* @Returns {bool} isfileProcessable
**/
func CheckFileFormatSupported(fileName string)(string, string, bool){

	fileExtension := filepath.Ext(fileName)
	val, ok := constants.SupportedFileFormats[fileExtension]

	return fileExtension, val, ok
}

/**
* Takes path and extension of the file as input and returns the fileName without extension
* @Params {string} path
* @Params {string} extenstion
* @Returns {string} fileNameWihoutExtension
**/

func GetFileNameWithoutExt(path string, extenstion string)(string){

	var filename = filepath.Base(path)
	if len(filename) > len(extenstion){
		return filename[0:len(filename)-len(extenstion)]
	}
	return ""	
}
