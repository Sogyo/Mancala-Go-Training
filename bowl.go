package Mancala_Go

type Bowl struct {
	NumberOfBeads uint
}

func CreateBowl() *Bowl {
	return &Bowl{NumberOfBeads: 4}
}

func (bowl *Bowl) PassOn(numberOfBeads uint) {
	if numberOfBeads == 0 {
		return
	}
	bowl.NumberOfBeads++
}
