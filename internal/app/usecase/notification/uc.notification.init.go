package notification

type UseCase struct {
	notificationRepo notificationRepo
}

func New(notificationRepo notificationRepo) *UseCase {
	return &UseCase{notificationRepo: notificationRepo}
}
