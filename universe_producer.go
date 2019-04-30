package universes

import "time"

type Universe_producer struct {

}

func (prd Universe_producer) Product() {
	time.Sleep(100000)
	Wormhole <- 1
}