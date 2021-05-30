package unit_tests

import (
	"testing"
	validators "theReader/services/validators"
	constants "theReader/services/constants"
)


func TestValidateEmail(t *testing.T){

	if !validators.ValidateEmail("phanikumar.tiragabattina@gmail.com") {
		t.Error("Error in TestvalidateEmail")
	}

	if validators.ValidateEmail(""){
		t.Error("Error in TestvalidateEmail")
	}

	if validators.ValidateEmail("phanikumar.tiragabattina@com") {
		t.Error("Error in TestvalidateEmail")
	}

	if validators.ValidateEmail("@gmail.com") {
		t.Error("Error in TestvalidateEmail")
	}

	if validators.ValidateEmail("gmail.com") {
		t.Error("Error in TestvalidateEmail")
	}
}

func TestValidateSalary(t *testing.T){

	if !validators.ValidateSalary("100") {
		t.Error("Error in TestvalidateEmail")
	}

	if !validators.ValidateSalary("100$"){
		t.Error("Error in TestvalidateEmail")
	}

	if !validators.ValidateSalary("100INR") {
		t.Error("Error in TestvalidateEmail")
	}

	if !validators.ValidateSalary("$100") {
		t.Error("Error in TestvalidateEmail")
	}

	if validators.ValidateSalary("$") {
		t.Error("Error in TestvalidateEmail")
	}

	if validators.ValidateSalary("NOMONEY") {
		t.Error("Error in TestvalidateEmail")
	}
}

func StringArrayEquals(a []string, b []string) bool {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}

func TestValidator(t *testing.T){

	var entryOne = []string{"Alfred Donald", "", "$11.5", "4"}
	var entryTwo = []string{"Jane Doe", "doe@test.com", "$8.45","5"}
	var nameIndexesOne = []int{0}
	var indexes = map[string]int{
		constants.Email : 1,
		constants.EmployeeId: 3,
		constants.EmployeeSalary : 2,
		constants.Mobile: -1,
	}

	arr, isValidEntry := validators.Validator(entryOne, nameIndexesOne, indexes)
	if !(StringArrayEquals(arr, []string{"Alfred Donald", "", "4", "$11.5"}) && !isValidEntry) {
		t.Error("Error in TestValidator")
	}

	arr, isValidEntry = validators.Validator(entryTwo, nameIndexesOne, indexes)

	if !(StringArrayEquals(arr, []string{"Jane Doe", "doe@test.com", "5", "$8.45"}) && isValidEntry) {
		t.Error("Error in TestValidator")
	}
	
}