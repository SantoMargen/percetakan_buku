package paper

const (
	queryInsertPaper = `
	INSERT INTO papers (
		unique_id,
		user_id,
		title, 
		authors, 
		co_authors, 
		journal, 
		volume, 
		issue, 
		page_range, 
		doi, 
		abstract, 
		keywords, 
		research_type, 
		funding_info, 
		affiliations, 
		full_text_link, 
		language, 
		review_status, 
		license, 
		notes,
		created_at
	) 
	VALUES 
		($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, now() )`

	queryUpdatePaper = `
		UPDATE papers
		SET 
			title = $1, 
			authors = $2, 
			co_authors = $3, 
			journal = $4, 
			volume = $5, 
			issue = $6, 
			page_range = $7, 
			doi = $8, 
			abstract = $9, 
			keywords = $10, 
			research_type = $11, 
			funding_info = $12, 
			affiliations = $13, 
			full_text_link = $14, 
			language = $15, 
			review_status = $16, 
			license = $17, 
			notes = $18,
			updated_at = now()
		WHERE 
			id = $19`

	queryDeletePaper string = `
		DELETE FROM papers WHERE id = $1`

	querySelectPaper string = `
		SELECT 
			id,
			user_id,
			unique_id, 
			title, 
			authors, 
			co_authors, 
			publication_date, 
			journal, 
			volume, 
			issue, 
			page_range, 
			doi, 
			abstract, 
			keywords, 
			research_type, 
			funding_info, 
			affiliations, 
			full_text_link, 
			language, 
			review_status, 
			license, 
			notes, 
			created_at, 
			updated_at,
			approval_posisi,
			approval_list
		FROM 
			papers
		WHERE 
			id = $1`
)
