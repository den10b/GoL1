package main

import "fmt"

// Зарядка
type Charger interface {
	Charge() int
}

// Телефон
type Phone struct {
	model      string
	currCharge int
}

// Телефон должен заряжаться через любую зарядку
func (phone *Phone) Charge(charger Charger) {
	addedPower := charger.Charge()
	phone.currCharge += addedPower
	fmt.Println("Charged ", addedPower, "%")
}

// USB-зарядка с показателем мощности в int
type USB struct {
	power int
}

func (usb USB) Charge() int {
	fmt.Println("Charging via usb")
	return usb.power
}

// Адаптер
type usbToTypecAdapter struct{ TypeC }

func (adapter usbToTypecAdapter) Charge() int {
	speed := adapter.ConnectTYPEC()
	defer fmt.Println("Charging via TYPEC")
	if speed == "fast" {
		return 10
	} else {
		return 5
	}
}

// TypeC-зарядка с показателем мощности в string
type TypeC struct {
	power string
}

func (typeC TypeC) ConnectTYPEC() string {
	return typeC.power
}

func main() {
	myPhone := Phone{"galaxy s10", 25}
	oldUSB := USB{2}
	myPhone.Charge(oldUSB) //Заряжаем через старый USB
	fmt.Println("Charge: ", myPhone.currCharge)

	newTYPEC := TypeC{"fast"}                 //Купили новый провод
	newAdapter := usbToTypecAdapter{newTYPEC} //Воспользуемся адаптером
	myPhone.Charge(newAdapter)                //Заряжаем через Адаптер + новый провод
	fmt.Println("Charge: ", myPhone.currCharge)
}
