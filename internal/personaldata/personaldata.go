package personaldata

import "fmt"

type Personal struct {
	Name   string  // имя пользователя;
	Weight float64 // вес пользователя;
	Height float64 // рост пользователя.
}

/*
	Метод выводит даные в формате:
		Имя: <Имя_пользователя>
		Вес: <Вес_пользователя>
		Рост: <Рост_пользователя>
*/
func (p Personal) Print() {
	fmt.Printf("Имя: %s\nВес: %.2f\nРост: %.2f\n", p.Name, p.Weight, p.Height)
}
