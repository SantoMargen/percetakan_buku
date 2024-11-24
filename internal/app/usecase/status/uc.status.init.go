package status

type UseCase struct {
	statusRepo statusRepo
}

func New(statusRepo statusRepo) *UseCase {
	return &UseCase{statusRepo: statusRepo}
}
