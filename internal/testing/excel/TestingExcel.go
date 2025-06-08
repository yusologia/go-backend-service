package excel

import (
	"github.com/xuri/excelize/v2"
	"github.com/yusologia/go-core/v2"
	"github.com/yusologia/go-core/v2/pkg"
	"service/internal/pkg/constant"
)

type TestingExcel struct {
}

func (ex TestingExcel) Generate() error {
	sheets, properties := ex.setSheetsAndProperties()

	excel := ex.newFile(sheets, properties)
	excel = ex.modifySheet(excel)

	err := excel.Save(constant.PathExcelTesting(), "testing.xlsx")
	if err != nil {
		return err
	}

	return nil
}

func (ex TestingExcel) newFile(sheets []string, properties [][][]interface{}) logiacore.Excel {
	excel := logiacore.Excel{
		Sheets:     sheets,
		Properties: properties,
		IsPublic:   false,
	}

	excel.NewFile()

	return excel
}

func (ex TestingExcel) modifySheet(excel logiacore.Excel) logiacore.Excel {
	excel.SetStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "860A35", Style: 1},
			{Type: "right", Color: "860A35", Style: 1},
			{Type: "top", Color: "860A35", Style: 1},
			{Type: "bottom", Color: "860A35", Style: 1},
		},
	}, "A1:C1")

	excel.SetStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "A3B763", Style: 2},
			{Type: "right", Color: "A3B763", Style: 2},
			{Type: "top", Color: "A3B763", Style: 2},
			{Type: "bottom", Color: "A3B763", Style: 2},
		},
	}, "A3:C4")

	excel.MergeCells("A5:C5", "A6:A7")

	excel.SetWidthCols([]logiacore.ColWidth{
		{Cells: "A", Width: 5},
		{Cells: "B", Width: 15},
		{Cells: "C", Width: 5},
	})

	excel.SetHeightRows([]logiacore.RowHeight{
		{Row: 1, Height: 15},
	})

	return excel
}

func (TestingExcel) setSheetsAndProperties() ([]string, [][][]interface{}) {
	var sheets []string
	var properties [][][]interface{}

	sheets = append(sheets, "Testing 1", "Testing 2")
	for _ = range sheets {
		properties = append(properties, [][]interface{}{
			{"ID", "Name", "Age"},
		})
	}

	dataProperties := [][]int{
		{29, 23, 12},
		{19, 26, 14},
	}

	for sKey, property := range dataProperties {
		for pKey, val := range property {
			properties[sKey] = append(properties[sKey], []interface{}{pKey + 1, logiapkg.RandomString(10), val})
		}
	}

	return sheets, properties
}
