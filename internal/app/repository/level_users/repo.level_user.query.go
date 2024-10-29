package leveluser

const (
	qryLevelUser string = `
		SELECT 
			id, 
			level_user, 
			keterangan, 
			created_at, 
			updated_at
		FROM level_users`

	qryLevelUserByID string = `
		SELECT 
			id, 
			level_user, 
			keterangan, 
			created_at, 
			updated_at
		FROM level_users
		WHERE id = $1`
)
