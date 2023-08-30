package main

import "fmt"

/*

Состояние — это поведенческий паттерн проектирования, который позволяет объектам менять поведение в зависимости
от своего состояния. Извне создаётся впечатление, что изменился класс объекта.

Плюсы:
- Избавляет от множества больших условных операторов машины состояний.
- Концентрирует в одном месте код, связанный с определённым состоянием.
- Упрощает код контекста.

Минусы:
- Может неоправданно усложнить код, если состояний мало и они редко меняются.

*/

func main() {
	coffeeMachine := NewCoffeeMachine()

	coffeeMachine.PickDrink()
	coffeeMachine.PutCoin()
	coffeeMachine.PourDrink()
	coffeeMachine.PickDrink()
	coffeeMachine.PourDrink()
	coffeeMachine.PourDrink()
}

type CoffeeMachine struct {
	noCoinState     State
	hasCoinState    State
	pourPickedState State

	currentState State
}

func NewCoffeeMachine() *CoffeeMachine {
	machine := &CoffeeMachine{}

	machine.noCoinState = &NoCoinState{machine: machine}
	machine.hasCoinState = &HasCoinState{machine: machine}
	machine.pourPickedState = &PourPickedState{machine: machine}

	machine.SetState(machine.noCoinState)

	return machine
}

func (machine *CoffeeMachine) PutCoin() {
	machine.currentState.PutCoin()
}

func (machine *CoffeeMachine) TakeCoin() {
	machine.currentState.TakeCoin()
}

func (machine *CoffeeMachine) PickDrink() {
	machine.currentState.PickDrink()
}

func (machine *CoffeeMachine) PourDrink() {
	machine.currentState.PourDrink()
}

func (machine *CoffeeMachine) SetState(state State) {
	machine.currentState = state
}

type State interface {
	PutCoin()
	TakeCoin()
	PickDrink()
	PourDrink()
}

type NoCoinState struct {
	machine *CoffeeMachine
}

func (state *NoCoinState) PutCoin() {
	fmt.Println("[NoCoinState] coin putted")
	state.machine.SetState(state.machine.hasCoinState)
}

func (state *NoCoinState) TakeCoin() {
	fmt.Println("[NoCoinState] no coin")
}

func (state *NoCoinState) PickDrink() {
	fmt.Println("[NoCoinState] put coin")
}

func (state *NoCoinState) PourDrink() {
	fmt.Println("[NoCoinState] put coin")
}

type HasCoinState struct {
	machine *CoffeeMachine
}

func (state *HasCoinState) PutCoin() {
	fmt.Println("[HasCoinState] coin already putted")
}

func (state *HasCoinState) TakeCoin() {
	fmt.Println("[HasCoinState] coin taken")
	state.machine.SetState(state.machine.noCoinState)
}

func (state *HasCoinState) PickDrink() {
	fmt.Println("[HasCoinState] drink picked")
	state.machine.SetState(state.machine.pourPickedState)
}

func (state *HasCoinState) PourDrink() {
	fmt.Println("[HasCoinState] pick drink")
}

type PourPickedState struct {
	machine *CoffeeMachine
}

func (state *PourPickedState) PutCoin() {
	fmt.Println("[PourPickedState] coin already putted")
}

func (state *PourPickedState) TakeCoin() {
	fmt.Println("[PourPickedState] coin taken")
	state.machine.SetState(state.machine.noCoinState)
}

func (state *PourPickedState) PickDrink() {
	fmt.Println("[PourPickedState] drink already picked")
}

func (state *PourPickedState) PourDrink() {
	fmt.Println("[PourPickedState] drink poured")
	state.machine.SetState(state.machine.noCoinState)
}
