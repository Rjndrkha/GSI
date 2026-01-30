package usecase

import (
	"fmt"

	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

type ReportUsecase struct {
	DB *gorm.DB
}

func (u *ReportUsecase) GenerateExcelReport(pocketID string, reportType string, date string, fileID string) {
	f := excelize.NewFile()
	sheet := "Report"
	f.SetSheetName("Sheet1", sheet)
	f.SetCellValue(sheet, "A1", "Type")
	f.SetCellValue(sheet, "B1", "Date")
	f.SetCellValue(sheet, "A2", reportType)
	f.SetCellValue(sheet, "B2", date)

	filePath := fmt.Sprintf("./reports/%s.xlsx", fileID)
	if err := f.SaveAs(filePath); err != nil {
		fmt.Println("Error saving excel:", err)
	}
}
