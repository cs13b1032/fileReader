package utils

import (
	"strings"
	constants "theReader/services/constants"
)

/**
* Takes a set of KeyWords and value, and returns if it is valid header
* @Params {[]string} keyWords
* @Params {string} headerKey
* @Returns {bool} isValidHeader
**/
func CheckIfRespectiveHeader(keyWords []string,value string)(bool){

	for i:=0;i<len(keyWords);i++{
		if strings.Contains(value, keyWords[i]){
			return true
		}
	}
	return false
}

/**
* Takes header record and process the header returns the respective indexes
* @Params {[]string} header
* @Returns {[]int} nameIndexes 
* @Returns {map[string]int} indexes
**/

func ProcessHeader(header []string)([]int, map[string]int){
	
	var nameIndexes []int
	var indexes = map[string]int{
		constants.Email: -1,
		constants.EmployeeId: -1,
		constants.EmployeeSalary: -1,
		constants.Mobile: -1,
	}

	for i := 0; i<len(header);i++{
		lowerCaseHeader := strings.ToLower(header[i])

		if CheckIfRespectiveHeader(constants.NameKeyWords, lowerCaseHeader){
			nameIndexes = append(nameIndexes, i)
		}else if indexes[constants.EmployeeId] == -1 && CheckIfRespectiveHeader(constants.EmployeeIdKeyWords, lowerCaseHeader){
			indexes[constants.EmployeeId] = i
		}else if indexes[constants.EmployeeSalary] == -1 && CheckIfRespectiveHeader(constants.SalaryKeyWords, lowerCaseHeader){
			indexes[constants.EmployeeSalary] = i
		}else if indexes[constants.Email] == -1 && CheckIfRespectiveHeader(constants.MailIDKeyWords, lowerCaseHeader){
			indexes[constants.Email] = i
		}else if indexes[constants.Mobile] == -1 && CheckIfRespectiveHeader(constants.MobileKeyWords, lowerCaseHeader){
			indexes[constants.Mobile] = i
		}
	}

	return nameIndexes, indexes
}
