package universes

import "time"

type universe_producer struct {

}

func (prd universe_producer) product() {
	time.Sleep(100000)
	wormhole <- 1
}