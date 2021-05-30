package unit_tests

import (
	"log"
	"testing"
	"os"
	"bufio"
	"strings"
	processors "theReader/services/processors"
)

func readTxtFile(filePath string) [][]string {
    f, err := os.Open(filePath)

    if err != nil {
        log.Println(err)
     }

    defer f.Close()

    scanner := bufio.NewScanner(f)
    scanner.Split(bufio.ScanWords)

	var records [][]string
    for scanner.Scan() {
		records = append(records, strings.Split(scanner.Text(), ","))
    }
	return records
}

func CompareIfTwoTXTFilesAreEqual(pathOne string, pathTwo string)(bool){
	recordsOne := readTxtFile(pathOne)
	recordsTwo := readTxtFile(pathTwo)
	return StringArrayArrayEquals(recordsOne, recordsTwo)
}

func TestProcessTXTfile(t *testing.T){

	mydir, nerr := os.Getwd()
	if nerr != nil {
		panic(nerr)
	}

	inputFilePath := mydir + "/test_input_file_folder/roster1.txt"
	processors.ProcessTxtFile(inputFilePath)

	if !(CompareIfTwoTXTFilesAreEqual(mydir + "/roster1_inValidOutput.txt",mydir + "/test_output_file_compare_folder/roster1_inValidOutput.txt" ) && CompareIfTwoTXTFilesAreEqual(mydir + "/roster1_validOutput.txt",mydir + "/test_output_file_compare_folder/roster1_validOutput.txt" )){
		t.Error("Error in TestProcessCSVfile")
	}

}
