package upload

type UseCase struct {
	uploadRepo uploadRepo
}

func New(uploadRepo uploadRepo) *UseCase {
	return &UseCase{uploadRepo: uploadRepo}
}
