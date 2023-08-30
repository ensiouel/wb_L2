package main

import "fmt"

/*

Команда — это поведенческий паттерн, позволяющий заворачивать запросы или простые операции в отдельные объекты.

Плюсы:
- Убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют.
- Позволяет реализовать простую отмену и повтор операций.
- Позволяет реализовать отложенный запуск операций.
- Позволяет собирать сложные команды из простых.
- Реализует принцип открытости/закрытости.

Минусы:
- Усложняет код программы из-за введения множества дополнительных классов.

*/

func main() {
	tv := &TV{}

	onButton := Button{command: &OnCommand{device: tv}}
	offButton := Button{command: &OffCommand{device: tv}}

	onButton.Press()
	offButton.Press()
}

type Command interface {
	Execute()
}

type Device interface {
	On()
	Off()
}

type Button struct {
	command Command
}

func (button *Button) Press() {
	button.command.Execute()
}

type OnCommand struct {
	device Device
}

func (command *OnCommand) Execute() {
	command.device.On()
}

type OffCommand struct {
	device Device
}

func (command *OffCommand) Execute() {
	command.device.Off()
}

type TV struct {
	isOn bool
}

func (tv *TV) On() {
	tv.isOn = true

	fmt.Println("[TV] on")
}

func (tv *TV) Off() {
	tv.isOn = false

	fmt.Println("[TV] off")
}
