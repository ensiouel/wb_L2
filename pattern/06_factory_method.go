package main

import "fmt"

/*

Фабричный метод — это порождающий паттерн проектирования, который определяет общий интерфейс для создания объектов в суперклассе,
позволяя подклассам изменять тип создаваемых объектов.

Плюсы:
- Избавляет класс от привязки к конкретным классам продуктов.
- Выделяет код производства продуктов в одно место, упрощая поддержку кода.
- Упрощает добавление новых продуктов в программу.
- Реализует принцип открытости/закрытости.

Минусы:
- Может привести к созданию больших параллельных иерархий классов, так как для каждого класса продукта надо создать свой подкласс создателя.

*/

func main() {
	var warriorFactory CharacterFactory = &WarriorFactory{}
	var mageFactory CharacterFactory = &MageFactory{}

	warrior := warriorFactory.Create()
	mage := mageFactory.Create()

	fmt.Println(warrior.Pick())
	fmt.Println(mage.Pick())
}

type Character interface {
	Pick() string
}

type Warrior struct{}

func (warrior *Warrior) Pick() string {
	return "now you pick warrior!"
}

type Mage struct{}

func (mage *Mage) Pick() string {
	return "now you pick mage!"
}

type CharacterFactory interface {
	Create() Character
}

type WarriorFactory struct{}

func (factory *WarriorFactory) Create() Character {
	return &Warrior{}
}

type MageFactory struct{}

func (factory *MageFactory) Create() Character {
	return &Mage{}
}
