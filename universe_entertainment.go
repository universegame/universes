package universes

type Universe_entertainment struct {

}

func (et Universe_entertainment) getGirl() {
	//自己卡住
	c := make(chan string)
	c <- "来个妹纸"
}