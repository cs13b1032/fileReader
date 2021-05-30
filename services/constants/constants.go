package constants

const Csv = ".csv"
const Txt = ".txt"
const Excel = ".xlsx"
const Name = "Name"
const Email = "Email"
const EmployeeId = "EmployeeId"
const EmployeeSalary = "EmployeeSalary"
const Mobile = "Mobile"
const InputFileFolder = "input_file_folder"

var Header = []string{"Name", "Email", "EmployeeId", "EmployeeSalary"}
var CurrencyTypes = []string{"$", "INR", "£", "€"}
var SalaryRegex = "\\d*\\.?\\d+"
const EmailRegex = "^[a-zA-Z0-9._%+\\-]+@[a-zA-Z0-9.\\-]+\\.[a-zA-Z]{2,4}$"

var SupportedFileFormats  = map[string]string{
	".csv": "csvProcessor",
	".txt": "txtProcessor",
	".xlsx": "excelProcessor",
}

type Entry struct {
	name string
	email string
	employeeId string 
	employeeSalary string
}

var SalaryKeyWords = []string{"wage", "salary", "rate"}
var EmployeeIdKeyWords = []string{"id", "number"}
var MailIDKeyWords = []string{"mail"}
var NameKeyWords = []string{"name", "first", "last"}
var MobileKeyWords = []string{"phone", "mobile"}