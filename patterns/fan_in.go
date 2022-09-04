package patterns

/*
Fan In Pattern

Multiplex the channel into single channel
*/

func FanIn[Type interface{}](inputChannels ...<-chan Type) <-chan Type {
	fanInChannel := make(chan Type)
	return fanIn(fanInChannel, inputChannels...)
}

func BufferedFanIn[Type interface{}](bufferSize int, inputChannels ...<-chan Type) <-chan Type {
	fanInChannel := make(chan Type, bufferSize)
	return fanIn(fanInChannel, inputChannels...)
}

func fanIn[Type interface{}](fanInChannel chan Type, inputChannels ...<-chan Type) <-chan Type {
	for _, ch := range inputChannels {
		// Generator pattern with input source as another input channel
		go func(in <-chan Type) {
			for {
				fanInChannel <- <-in
			}
		}(ch)
	}
	return fanInChannel
}
