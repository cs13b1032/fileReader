package unit_tests

import (
	"log"
	"testing"
	"os"
	"encoding/csv"
	processors "theReader/services/processors"
)

func readCsvFile(filePath string) [][]string {
    f, err := os.Open(filePath)
    if err != nil {
        log.Fatal("Unable to read input file " + filePath, err)
    }
    defer f.Close()

    csvReader := csv.NewReader(f)
    records, err := csvReader.ReadAll()
    if err != nil {
        log.Fatal("Unable to parse file as CSV for " + filePath, err)
    }

    return records
}

func StringArrayArrayEquals(a [][]string, b [][]string) bool {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
		if !StringArrayEquals(v, b[i]){
            return false
        }
    }
    return true
}

func CompareIfTwoCSVFilesAreEqual(pathOne string, pathTwo string)(bool){
	recordsOne := readCsvFile(pathOne)
	recordsTwo := readCsvFile(pathTwo)

	return StringArrayArrayEquals(recordsOne, recordsTwo)
}

func TestProcessCSVfile(t *testing.T){

	mydir, nerr := os.Getwd()
	if nerr != nil {
		panic(nerr)
	}

	inputFilePath := mydir + "/test_input_file_folder/roster1.csv"
	processors.ProcessCSVfile(inputFilePath)

	if !(CompareIfTwoCSVFilesAreEqual(mydir + "/roster1_inValidOutput.csv",mydir + "/test_output_file_compare_folder/roster1_inValidOutput.csv" ) && CompareIfTwoCSVFilesAreEqual(mydir + "/roster1_validOutput.csv",mydir + "/test_output_file_compare_folder/roster1_validOutput.csv" )){
		t.Error("Error in TestProcessCSVfile")
	}

}
