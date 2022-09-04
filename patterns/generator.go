package patterns

type Iterator[Type interface{}] interface {
	HasNext() bool
	Next() Type
}

func Generator[Type interface{}](itr Iterator[Type]) <-chan Type {
	generatorChannel := make(chan Type)
	go func() {
		for itr.HasNext() {
			generatorChannel <- itr.Next()
		}
		close(generatorChannel)
	}()
	return generatorChannel
}
