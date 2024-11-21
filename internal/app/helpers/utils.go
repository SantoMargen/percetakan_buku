package helpers

import (
	"fmt"
	"time"
)

func GetTableNameMonth(mstTable string, date *string) (string, error) {
	var tableName string
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return "", err
	}

	today := time.Now().In(location).Format("2006")
	tableName = fmt.Sprintf(`%s_%s`, mstTable, today)

	return tableName, nil
}
