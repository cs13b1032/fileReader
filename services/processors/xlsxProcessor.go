
package processors
import "github.com/360EntSecGroup-Skylar/excelize"

// In Progress
func ProcessXLSXFile(path string) {
    f, err := excelize.OpenFile(path)
    if err != nil {
        println(err.Error())
        return
    }
    cell:= f.GetCellValue("Sheet1", "B2")
    println(cell)
    rows:= f.GetRows("Sheet1")
    for _, row := range rows {
        for _, colCell := range row {
            print(colCell, "\t")
        }
        println()
    }
}