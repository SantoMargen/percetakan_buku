package papers

type UseCase struct {
	paperRepo paperRepo
}

func New(paperRepo paperRepo) *UseCase {
	return &UseCase{paperRepo: paperRepo}
}
