package unit_tests

import (
	"os"
	"testing"
	helpers "theReader/services/helpers"
	constants "theReader/services/constants"
)

var csvFilePath = "/Users/phani/Desktop/fileReader/input_file_folder/roster4.csv"
var dummyFilePath = "/Users/phani/Desktop/fileReader/input_file_folder/roster4.dummy"

func TestGetinputFileFolderPath(t *testing.T){

	mydir, nerr := os.Getwd()
	if nerr != nil {
		panic(nerr)
	}

	if helpers.GetinputFileFolderPath(constants.InputFileFolder) != mydir + "/"+constants.InputFileFolder+"/" {
		t.Error("Error in TestGetinputFileFolderPath")
	}

	if helpers.GetinputFileFolderPath("") != mydir + "//" {
		t.Error("Error in TestGetinputFileFolderPath")
	}
}

func TestCheckFileFormatSupported(t *testing.T){

	fileExtension, processorName, fileFormatSupported  := helpers.CheckFileFormatSupported(csvFilePath)

	if !(fileExtension == ".csv" &&  processorName == "csvProcessor" && fileFormatSupported) {
		t.Error("Error in TestCheckFileFormatSupported")
	}

	fileExtension, processorName, fileFormatSupported  = helpers.CheckFileFormatSupported("")

	if !(len(fileExtension) == 0 &&  len(processorName) == 0 && !fileFormatSupported) {
		t.Error("Error in TestCheckFileFormatSupported")
	}

	fileExtension, processorName, fileFormatSupported  = helpers.CheckFileFormatSupported(dummyFilePath)

	if !(fileExtension == ".dummy" &&  len(processorName) == 0 && !fileFormatSupported) {
		t.Error("Error in TestCheckFileFormatSupported")
	}

}

func TestGetFileNameWithoutExt(t *testing.T){

	if helpers.GetFileNameWithoutExt(csvFilePath, ".csv") != "roster4" {
		t.Error("Error in TestGetinputFileFolderPath")
	}

	if helpers.GetFileNameWithoutExt("", ".csv") != "" {
		t.Error("Error in TestGetinputFileFolderPath")
	}

	if helpers.GetFileNameWithoutExt(csvFilePath, "") != "roster4.csv" {
		t.Error("Error in TestGetinputFileFolderPath")
	}

	if helpers.GetFileNameWithoutExt("", "") != "." {
		t.Error("Error in TestGetinputFileFolderPath")
	}
}