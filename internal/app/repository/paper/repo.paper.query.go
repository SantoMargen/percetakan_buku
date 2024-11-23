package paper

const (
	queryCekFileExist = `SELECT COUNT(*) FROM file_uploads WHERE unique_id = $1`

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
			license, 
			notes, 
			created_at, 
			updated_at,
			flag_assign,
			unique_id_paper,
			category
		FROM 
			papers
		WHERE 
			id = $1`

	queryAssignPaper = `
		INSERT INTO task_approval (
			paper_id,
			publisher_id,
			approval_posisi, 
			approval_list, 
			status,
			entry_user,
			entry_time
		
		) 
		VALUES 
			($1, $2, $3, $4, 2, $5, now())`

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
				task_approval 
			WHERE approval_posisi = $1 AND paper_id = $2 and status = 2
		`
	queryApprovalPaper = `
		UPDATE task_approval SET 
			approval_posisi = $1,
			approval_list = $2,
			catatan_tolakan = $3,
			status = $4
		WHERE paper_id = $5 and approval_posisi = $6 and status = 2`

	queryColumnDetailPaper = `
		papers.id, 
		papers.user_id, 
		papers.unique_id, 
		papers.title, 
		papers.authors, 
		papers.co_authors, 
		papers.publication_date, 
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
		papers.updated_at, 
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
		publishers.entry_time, 
		task_approval.approval_posisi, 
		task_approval.approval_list, 
		coalesce(task_approval.catatan_tolakan,'[]') as catatan_tolakan_task_approval ,
		task_approval.entry_user, 
		task_approval.entry_name, 
		task_approval.entry_time,
		task_publisher.entry_user, 
		task_publisher.entry_name, 
		task_publisher.entry_time,
		status_submission.id_status,
		status_submission.desc_status
	`

	queryDetailPaper = `
			SELECT 
				` + queryColumnDetailPaper + `
			from 
				papers 
			inner join task_approval on papers.id = task_approval.paper_id 
			inner join publishers on task_approval.publisher_id = publishers.publisher_id
			inner join task_publisher on task_publisher.publisher_id = publishers.publisher_id
			inner join status_submission on task_approval.status = status_submission.id_status

		WHERE papers.id = $1`

	queryDetailPaperByUserId = `
		SELECT 
			` + queryColumnDetailPaper + `
		from 
			papers 
		left join  task_approval on papers.id = task_approval.paper_id 
		inner join publishers on task_approval.publisher_id = publishers.publisher_id
		inner join task_publisher on task_publisher.publisher_id = publishers.publisher_id
		inner join status_submission on task_approval.status = status_submission.id_status

	WHERE  1 = 1`

	queryCountDetailPaperByUserId = `
		SELECT 
			COUNT(*)
		from 
			papers 
			inner join task_approval on papers.id = task_approval.paper_id 
			inner join publishers on task_approval.publisher_id = publishers.publisher_id
			inner JOIN task_publisher on task_publisher.publisher_id = publishers.publisher_id
		where 1 = 1
`
)
