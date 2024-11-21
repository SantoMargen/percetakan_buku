package publisher

type UseCase struct {
	publishRepo publisherRepo
}

func New(publishRepo publisherRepo) *UseCase {
	return &UseCase{publishRepo: publishRepo}
}
