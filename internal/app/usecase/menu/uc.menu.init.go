package menu

type UseCase struct {
	menuRepo menuRepo
}

func New(menuRepo menuRepo) *UseCase {
	return &UseCase{menuRepo: menuRepo}
}
