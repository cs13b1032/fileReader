package helpers

import (
    "os"
    "path/filepath"
	constants "theReader/services/constants"
)

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

func CheckFileFormatSupported(fileName string)(string, string, bool){

	fileExtension := filepath.Ext(fileName)
	val, ok := constants.SupportedFileFormats[fileExtension]

	return fileExtension, val, ok
}

func GetFileNameWithoutExt(path string, extenstion string)(string){

	var filename = filepath.Base(path)
	if len(filename) > len(extenstion){
		return filename[0:len(filename)-len(extenstion)]
	}
	return ""	
}
