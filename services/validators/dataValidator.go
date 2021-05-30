package validators

import (
	"strings"
	"regexp"
	constants "theReader/services/constants"
)

/**
* Takes emailId as the input and returns if it's a valid Email
* @Params {string} emailId
* @Returns {bool} isValidEmail
**/
func ValidateEmail(email string) bool {
	match, _ := regexp.MatchString(constants.EmailRegex, email)
	return match
}

/**
* Takes salary as the input and returns if it's a valid salary
* @Params {string} salarytext
* @Returns {bool} isValidSalary
**/
func ValidateSalary(text string)(bool){
	for i:=0;i<len(constants.CurrencyTypes);i++{
		if strings.Contains(text, constants.CurrencyTypes[i]){
			text = strings.ReplaceAll(text,constants. CurrencyTypes[i] , "")
			break
		}
	}
	match, _ := regexp.MatchString(constants.SalaryRegex, text)
	return match
}

/**
* Takes a record and it's name column indexes and other fields indexes and returns the processed values and if it is a vlaid record
* @Params {[]string} record
* @Params {[]int} nameIndexes
* @Params {map[string]int} indexesMap
* @Returns {[]string} processedValues
* @Returns {bool} isValidRecord
**/
func Validator(entry []string, nameIndexes []int,indexes map[string]int)([]string, bool){

	// var domain string
	var name string
	var email string
	var employeeId string
	var employeeSalary string
	validEntry := false
	
	// getting valid Name
	if len(nameIndexes) > 0{
		for i:=0; i< len(nameIndexes); i++{
			name = name + entry[nameIndexes[i]]
		}
	}

	// getting valid Email
	if indexes[constants.Email] != -1{
		value := entry[indexes[constants.Email]]
		if ValidateEmail(value){
			email = value
		}
	}

	// getting valid emaployee Id
	if indexes[constants.EmployeeId] != -1 {
		value := entry[indexes[constants.EmployeeId]]
		employeeId = value
	}

	// getting valid employee Salary
	if indexes[constants.EmployeeSalary] != -1 {
		value := entry[indexes[constants.EmployeeSalary]]
		if ValidateSalary(value){
			employeeSalary = value
		}
	}

	if len(name) > 0 && len(email) > 0 && len(employeeId) > 0 && len(employeeSalary) > 0{
		validEntry = true
	}

	// check if vlaid record is still present
	if !validEntry{
		for i:=0; i<len(entry); i++{
			if len(email) <= 0 {
				value := entry[i]
				if ValidateEmail(value){
					email = value
				}
			}
			if len(employeeSalary) <= 0 {
				value := entry[i]
				if ValidateSalary(value){
					employeeSalary = value
				}
			}
			if len(employeeId) <= 0 {
				// cannot determine as there will be no defined format
			}
			if len(name) <= 0 {
				// cannot determine as there will be no defined format and multiple fields can be considered as name
			}
		}
	}

	if len(name) > 0 && len(email) > 0 && len(employeeId) > 0 && len(employeeSalary) > 0{
		validEntry = true
	}

	return []string{name, email, employeeId, employeeSalary}, validEntry
}