package point

type Point struct {
	x, y float64 //Инкапсулируем поля, так что их название с маленькой буквы
}

func (p *Point) Set(x, y float64) {
	p.y = y
	p.x = x
}
func NewPoint(x, y float64) *Point { //Конструктор должен быть доступен извне, так что его название с большой буквы
	p := &Point{}
	p.Set(x, y)
	return p
}

func (p *Point) GetX() float64 {
	return p.x
}
func (p *Point) GetY() float64 {
	return p.y
}
