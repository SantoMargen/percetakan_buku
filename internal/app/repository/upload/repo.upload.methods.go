package upload

import (
	"context"
	"fmt"
	"siap_app/internal/app/entity/upload"
	"strconv"
	"strings"
	"time"
)

func (r *repository) UploadFile(ctx context.Context, input []upload.RequestUpload) error {

	var (
		values []string
		args   []interface{}
	)

	tx, err := r.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}
	defer tx.Rollback()

	for i, getValue := range input {

		values = append(values, fmt.Sprintf("($%d,$%d, $%d, $%d, $%d)",
			(i*5)+1, (i*5)+2, (i*5)+3, (i*5)+4, (i*5)+5))

		getCounter := strconv.Itoa(i)

		args = append(args,
			getValue.IDFile+getCounter,
			getValue.Filename,
			getValue.Filetype,
			getValue.Path,
			time.Now(),
		)
	}

	query := queryInsertFile + " " + strings.Join(values, ", ")

	_, errInsert := tx.ExecContext(ctx, query, args...)

	if errInsert != nil {
		tx.Rollback()
		return errInsert
	}
	return tx.Commit()
}
