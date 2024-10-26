package user

type UseCase struct {
	userRepo userRepo
}

func New(userRepo userRepo) *UseCase {
	return &UseCase{userRepo: userRepo}
}
