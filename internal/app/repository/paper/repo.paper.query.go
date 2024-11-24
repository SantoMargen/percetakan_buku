package paper

const (
	queryCekFileExist    = `SELECT COUNT(*) FROM file_uploads WHERE unique_id = $1`
	queryColumnTblPapers = `id,
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
			license, 
			notes, 
			created_at, 
			updated_at,
			flag_assign,
			unique_id_paper,
			category_name`

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
		license, 
		notes,
		unique_id_paper,
		created_at
	) 
	VALUES 
		($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19,$20, now() )`

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
			license = $16, 
			notes = $17,
			unique_id_paper = $18,
			updated_at = now()
		WHERE 
			id = $18`

	queryDeletePaper string = `
		DELETE FROM papers WHERE id = $1`

	querySelectPaper string = `
		SELECT 
			` + queryColumnDetailPaper + `
		FROM 
				"public"."papers"
			INNER JOIN 
				"public"."publishers" ON papers.publisher_id = publishers.publisher_id
			INNER JOIN 
				"public"."status_submission" ON papers.status = status_submission.id_status

		WHERE papers.id = $1`

	queryAssignPaper = `
		UPDATE papers SET
			publisher_id = $1,
			approval_posisi = $2,
			approval_list = $3,
			status = 3,
			entry_assign = $4,
			flag_assign = 1,
			entry_assign_name = $5,
			entry_assign_time = now(),
			catatan_assignment = $6

		WHERE id = $7`

	queryAssignPaperPublisher = `
		INSERT INTO task_publisher (
			paper_id,
			publisher_id,
			user_id,
			entry_time
		) 
		VALUES 
			($1, $2, $3, now())`

	queryUpdateFlagSubmission = `
		UPDATE papers SET flag_assign = 1
		WHERE id = $1`

	queryCheckTask = `
		SELECT count(*)
			FROM
				papers 
			WHERE approval_posisi = $1 AND id = $2 and status = 3
		`
	queryApprovalPaper = `
		UPDATE papers SET 
			approval_posisi = $1,
			approval_list = $2,
			catatan_tolakan = $3,
			status = $4
		WHERE id = $5 and approval_posisi = $6 and status = $7`

	queryColumnDetailPaper = ` 
			papers.id, 
			papers.user_id, 
			papers.unique_id, 
			papers.title, 
			papers.authors, 
			papers.co_authors, 
			COALESCE(papers.publication_date, '1970-01-01 00:00:00'::timestamp) AS publication_date, 
			papers.journal, 
			papers.volume, 
			papers.issue, 
			papers.page_range, 
			papers.doi, 
			papers.abstract, 
			papers.keywords, 
			papers.research_type, 
			papers.funding_info, 
			papers.affiliations, 
			papers.full_text_link, 
			papers.language, 
			papers.license, 
			papers.notes, 
			papers.created_at, 
			COALESCE(papers.updated_at, '1970-01-01 00:00:00'::timestamp) as updated_at, 
			papers.flag_assign, 
			papers.id_category,
			papers.category_name,
			publishers.name, 
			publishers.address, 
			publishers.phone, 
			publishers.email, 
			publishers.website, 
			publishers.founded_year, 
			publishers.country, 
			publishers.contact_person_1, 
			publishers.contact_person_2, 
			publishers.fax, 
			publishers.fb_link, 
			publishers.twitter_link, 
			publishers.web_link, 
			publishers.join_date, 
			publishers.entry_user, 
			publishers.entry_name,
			COALESCE(publishers.entry_time, '1970-01-01 00:00:00'::timestamp) AS entry_assign_time,
			papers.approval_posisi, 
			COALESCE(papers.approval_list, '[]'::jsonb) AS approval_list,  -- Ensure proper JSON casting
			COALESCE(papers.catatan_tolakan, '[]'::jsonb) AS catatan_tolakan_task_approval,  -- Ensure proper JSON casting
			papers.entry_assign, 
			papers.entry_assign_name, 
			COALESCE(papers.entry_assign_time, '1970-01-01 00:00:00'::timestamp) AS entry_assign_time,
			papers.catatan_assignment,
			status_submission.id_status,
			status_submission.desc_status
		`

	queryListPaper = `
			SELECT 
				` + queryColumnDetailPaper + `
			FROM 
				"public"."papers"
			INNER JOIN 
				"public"."publishers" ON papers.publisher_id = publishers.publisher_id
			INNER JOIN 
				"public"."status_submission" ON papers.status = status_submission.id_status

		WHERE 1 = 1`

	queryDetailPaper = `
		SELECT 
			` + queryColumnDetailPaper + `
		FROM 
			"public"."papers"
		INNER JOIN 
			"public"."publishers" ON papers.publisher_id = publishers.publisher_id
		INNER JOIN 
			"public"."status_submission" ON papers.status = status_submission.id_status

	WHERE papers.id = $1`

	queryCountDetailPaper = `
		SELECT count(*)
		from 
			papers 
		INNER JOIN 
			"public"."publishers" ON papers.publisher_id = publishers.publisher_id
		INNER JOIN 
			"public"."status_submission" ON papers.status = status_submission.id_status
		WHERE 1 = 1`
)
