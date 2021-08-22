package spreadsheets

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	log "local.packages/game-information/lib/domain/log"
)

// InterfaceYoutube インターフェース
type InterfaceSpreadsheet interface {
	getService() *sheets.Service
	GetSheetData(sheetName string) *sheets.ValueRange
}

// SpreadSheet インスタンス
type Spreadsheet struct {
	DeveloperKey string
	SheetID      string
}

const credentialsFilePath = "assets/credentials/secret.json"

// GetInstance インスタンス生成
func GetInstance() *Spreadsheet {
	return &Spreadsheet{DeveloperKey: os.Getenv("SHEET_API_KEY"), SheetID: os.Getenv("SPREAD_SHEET_ID")}
}

func (domain *Spreadsheet) getService() *sheets.Service {
	ctx := context.Background()
	service, err := sheets.NewService(ctx, option.WithCredentialsFile(credentialsFilePath))
	if err != nil {
		log.Error(fmt.Sprintf("Error creating new Spreadsheet client: %v", err))
		return nil
	}
	return service
}

// GetSheetData spreadsheetからデータ取得
func (domain *Spreadsheet) GetSheetData(sheetRange string) *sheets.ValueRange {
	service := domain.getService()
	res, err := service.Spreadsheets.Values.Get(domain.SheetID, sheetRange).Do()
	if err != nil {
		log.Error(fmt.Sprintf("Unable to retrieve data from sheet: %v", err))
		return nil
	}

	if len(res.Values) == 0 {
		log.Info("No data found.")
		return nil
	}

	return res
}

// ConvertColNumberToAlphabet 数値をアルファベットに変換
// https://qiita.com/cyabane/items/ec2f17948e0f4070db87
func ConvertColNumberToAlphabet(col int) string {
	var colName string
	x := int((col - 1) / 26)
	y := col - (x * 26)
	if x > 0 {
		colName = string(x + 64)
	}
	if y > 0 {
		colName = colName + string(y+64)
	}
	return colName
}
