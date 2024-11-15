// By default channels are _unbuffered_, meaning that they
// will only accept sends (`chan <-`) if there is a
// corresponding receive (`<- chan`) ready to receive the
// sent value. _Buffered channels_ accept a limited
// number of  values without a corresponding receiver for
// those values.

package referenceprogramsexternal

import "fmt"

func channelSample() {

	// Here we `make` a channel of strings buffering up to
	// 2 values.
	messages := make(chan string, 2)

	// Because this channel is buffered, we can send these
	// values into the channel without a corresponding
	// concurrent receive.
	messages <- "buffered"
	messages <- "channel"

	// Later we can receive these two values as usual.
	fmt.Println(<-messages)
	fmt.Println(<-messages)

	// If we add a 3rd message to the channel, this goroutine, which is main
	// would block untill another goroutine receives the value.

	// messages <- "blocking" // this line would raise deadlock

	// to avoid this we can add the writing of the 3rd message to another goroutine
	// so that main is free

	go func() {
		messages <- "blocking"
	}()
	fmt.Println(<-messages)
}
