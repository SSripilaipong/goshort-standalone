package either

type Of[Left, Right any] struct {
	left   Left
	right  Right
	isLeft bool
}

func (e Of[Left, Right]) IsRight() bool {
	return !e.IsLeft()
}

func (e Of[Left, Right]) IsLeft() bool {
	return e.isLeft
}

func (e Of[Left, Right]) Right() Right {
	return e.right
}

func (e Of[Left, Right]) Left() Left {
	return e.left
}

func Left[Left, Right any](left Left) Of[Left, Right] {
	return Of[Left, Right]{left: left, isLeft: true}
}

func Right[Left, Right any](right Right) Of[Left, Right] {
	return Of[Left, Right]{right: right, isLeft: false}
}
