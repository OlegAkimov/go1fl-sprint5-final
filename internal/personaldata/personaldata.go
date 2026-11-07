package personaldata

import "fmt"

type Personal struct {
	// TODO: добавить поля
	Name   string  // имя пользователя;
	Weight float64 // вес пользователя;
	Height float64 // рост пользователя.
}

func (p Personal) Print() {
	// TODO: реализовать функцию
	fmt.Printf("Имя: %s\n", p.Name)
	fmt.Printf("Вес: %.2f\n", p.Weight)
	fmt.Printf("Height: %.2f\n", p.Height)
}
