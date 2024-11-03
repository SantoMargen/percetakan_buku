package category

const (
	columnSelectCategory string = `
			category_id,
			category_name, 
			description, 
			entry_user,
			entry_time
	`
	queryInsertCategory string = `
		INSERT INTO category(
			category_name, 
			description, 
			entry_user,
			entry_time)
		VALUES
		($1, $2, $3, now())`

	queryUpdateCategory string = `
		UPDATE category 
			set 
			category_name = $1, 
			description = $2,
			entry_time = now() 
		WHERE category_id = $3`

	queryDeleteCategory string = `
		DELETE FROM category WHERE category_id = $1`

	queryCategoryById string = `
		SELECT 
			category_id,
			category_name, 
			description, 
			entry_user,
			entry_time
		FROM category WHERE category_id = $1`

	queryCategoryByName string = `SELECT count(*) FROM publishers where name = $1`
)
