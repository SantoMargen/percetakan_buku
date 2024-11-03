package user

type UseCase struct {
	userRepo     userRepo
	redisRepo    redisRepo
	logLoginRepo logLoginRepo
}

func New(userRepo userRepo, redisRepo redisRepo, logLoginRepo logLoginRepo) *UseCase {
	return &UseCase{
		userRepo:     userRepo,
		redisRepo:    redisRepo,
		logLoginRepo: logLoginRepo,
	}
}
