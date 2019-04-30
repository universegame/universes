package universes

type universe_entertainment struct {

}

func (et universe_entertainment) getGirl() {
	//自己卡住
	c := make(chan string)
	c <- "来个妹纸"
}