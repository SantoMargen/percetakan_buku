package upload

const (
	queryInsertFile string = `
		INSERT INTO file_uploads(
			unique_id,
			filename, 
			filetype, 
			path,
			created_time
				)
		VALUES `
)
