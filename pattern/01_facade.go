package main

import (
	"crypto/rand"
	"fmt"
)

/*

Фасад — это структурный паттерн проектирования, который предоставляет простой интерфейс к сложной системе классов,
библиотеке или фреймворку.

Плюсы:
- Изолирует клиентов от компонентов сложной подсистемы.

Минусы:
- Фасад рискует стать слишком большим.

*/

func main() {
	computer := NewComputer()
	computer.Start()
}

type Computer struct {
	hardDrive HardDrive
	memory    Memory
	cpu       CPU
}

func NewComputer() Computer {
	hardDrive := NewHardDrive()
	memory := NewMemory()
	cpu := NewCPU()

	return Computer{
		hardDrive: hardDrive,
		memory:    memory,
		cpu:       cpu,
	}
}

func (computer *Computer) Start() {
	var bootAddress int64 = 0x0005
	var bootSector int64 = 0x001
	var sectorSize int64 = 32

	computer.cpu.Freeze()
	computer.memory.Load(bootAddress, computer.hardDrive.Read(bootSector, sectorSize))
	computer.cpu.Jump(bootAddress)
	computer.cpu.Execute()

	fmt.Printf("[Computer] start\n")
}

type HardDrive struct{}

func NewHardDrive() HardDrive {
	return HardDrive{}
}

func (HardDrive *HardDrive) Read(lba int64, size int64) []byte {
	data := make([]byte, size)
	_, _ = rand.Read(data)

	fmt.Printf("[HardDrive] read lba = %d, size = %d\n", lba, size)

	return data
}

type Memory struct{}

func NewMemory() Memory {
	return Memory{}
}

func (Memory *Memory) Load(position int64, data []byte) {
	fmt.Printf("[Memory] load position = %d, data = %x\n", position, data)
}

type CPU struct{}

func NewCPU() CPU {
	return CPU{}
}

func (CPU *CPU) Freeze() {
	fmt.Printf("[CPU] freeze\n")
}

func (CPU *CPU) Jump(position int64) {
	fmt.Printf("[CPU] jump position = %d\n", position)
}

func (CPU *CPU) Execute() {
	fmt.Printf("[CPU] execute\n")
}
