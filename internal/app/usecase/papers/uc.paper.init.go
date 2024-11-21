package papers

type UseCase struct {
	paperRepo        paperRepo
	publisherRepo    publisherRepo
	notificationRepo notificationRepo
}

func New(
	paperRepo paperRepo,
	publisherRepo publisherRepo,
	notificationRepo notificationRepo,

) *UseCase {
	return &UseCase{
		paperRepo:        paperRepo,
		publisherRepo:    publisherRepo,
		notificationRepo: notificationRepo,
	}
}
