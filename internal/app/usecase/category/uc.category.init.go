package category

type UseCase struct {
	categoryRepo categoryRepo
}

func New(categoryRepo categoryRepo) *UseCase {
	return &UseCase{categoryRepo: categoryRepo}
}
