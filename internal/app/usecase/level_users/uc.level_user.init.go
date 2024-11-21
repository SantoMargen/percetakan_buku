package levelusers

type UseCase struct {
	levelUserRepo levelUserRepo
}

func New(levelUserRepo levelUserRepo) *UseCase {
	return &UseCase{levelUserRepo: levelUserRepo}
}
