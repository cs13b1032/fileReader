package unit_tests

import (
	"testing"
	constants "theReader/services/constants"
	utils "theReader/services/utils"
)


func TestCheckIfRespectiveHeader(t *testing.T){

	if !utils.CheckIfRespectiveHeader(constants.SalaryKeyWords, "wage") {
		t.Error("Error in TestCheckIfRespectiveHeader")
	}

	if utils.CheckIfRespectiveHeader(constants.SalaryKeyWords, "") {
		t.Error("Error in TestCheckIfRespectiveHeader")
	}
}


func TestProcessHeader(t *testing.T){

	arrRes, mapRes := utils.ProcessHeader([]string{"Name", "Email", "Wage" ,"Number"})
	if !(arrRes[0] == 0 && mapRes["Email"] == 1 && mapRes["EmployeeId"] == 3 && mapRes["EmployeeSalary"] == 2 && mapRes["Mobile"] == -1) {
		t.Error("Error in TestProcessHeader")
	}

	if (arrRes[0] == 0 && mapRes["Email"] == 1 && mapRes["EmployeeId"] == 2 && mapRes["EmployeeSalary"] == 2 && mapRes["Mobile"] == -1) {
		t.Error("Error in TestProcessHeader")
	}

}
