package patterns

func FanOut[Type interface{}](inputChannel <-chan Type, outputChannel ...chan<- Type) {
	for _, ch := range outputChannel {
		go func(in <-chan Type, out chan<- Type) {
			for {
				out <- <-in
			}
		}(inputChannel, ch)
	}
}
