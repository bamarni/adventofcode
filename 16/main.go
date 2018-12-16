package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Instruction func(a, b, c int)

type CPU struct {
	registers [4]int
}

func (cpu *CPU) addr(a, b, c int) { cpu.registers[c] = cpu.registers[a] + cpu.registers[b] }
func (cpu *CPU) addi(a, b, c int) { cpu.registers[c] = cpu.registers[a] + b }
func (cpu *CPU) mulr(a, b, c int) { cpu.registers[c] = cpu.registers[a] * cpu.registers[b] }
func (cpu *CPU) muli(a, b, c int) { cpu.registers[c] = cpu.registers[a] * b }
func (cpu *CPU) banr(a, b, c int) { cpu.registers[c] = cpu.registers[a] & cpu.registers[b] }
func (cpu *CPU) bani(a, b, c int) { cpu.registers[c] = cpu.registers[a] & b }
func (cpu *CPU) borr(a, b, c int) { cpu.registers[c] = cpu.registers[a] | cpu.registers[b] }
func (cpu *CPU) bori(a, b, c int) { cpu.registers[c] = cpu.registers[a] | b }
func (cpu *CPU) setr(a, b, c int) { cpu.registers[c] = cpu.registers[a] }
func (cpu *CPU) seti(a, b, c int) { cpu.registers[c] = a }
func (cpu *CPU) gtir(a, b, c int) {
	if a > cpu.registers[b] {
		cpu.registers[c] = 1
	} else {
		cpu.registers[c] = 0
	}
}
func (cpu *CPU) gtri(a, b, c int) {
	if cpu.registers[a] > b {
		cpu.registers[c] = 1
	} else {
		cpu.registers[c] = 0
	}
}
func (cpu *CPU) gtrr(a, b, c int) {
	if cpu.registers[a] > cpu.registers[b] {
		cpu.registers[c] = 1
	} else {
		cpu.registers[c] = 0
	}
}
func (cpu *CPU) eqir(a, b, c int) {
	if a == cpu.registers[b] {
		cpu.registers[c] = 1
	} else {
		cpu.registers[c] = 0
	}
}
func (cpu *CPU) eqri(a, b, c int) {
	if cpu.registers[a] == b {
		cpu.registers[c] = 1
	} else {
		cpu.registers[c] = 0
	}
}
func (cpu *CPU) eqrr(a, b, c int) {
	if cpu.registers[a] == cpu.registers[b] {
		cpu.registers[c] = 1
	} else {
		cpu.registers[c] = 0
	}
}

func main() {
	cpu := &CPU{}
	instructions := [16]Instruction{cpu.addr, cpu.addi, cpu.mulr, cpu.muli, cpu.banr, cpu.bani, cpu.borr, cpu.bori, cpu.setr, cpu.seti, cpu.gtir, cpu.gtri, cpu.gtrr, cpu.eqir, cpu.eqri, cpu.eqrr}
	instructionsMap := make(map[int][]int)

	var threeOrMore int
	samples := strings.Split(rawSamples, "\n")
	for i := 0; i < len(samples); i += 4 {
		before := [4]int{int(samples[i][9]) - 48, int(samples[i][12]) - 48, int(samples[i][15]) - 48, int(samples[i][18] - 48)}
		after := [4]int{int(samples[i+2][9]) - 48, int(samples[i+2][12]) - 48, int(samples[i+2][15]) - 48, int(samples[i+2][18] - 48)}
		sample := strings.Fields(samples[i+1])
		sampleOpcode, _ := strconv.Atoi(sample[0])
		sampleA, _ := strconv.Atoi(sample[1])
		sampleB, _ := strconv.Atoi(sample[2])
		sampleC, _ := strconv.Atoi(sample[3])
		var sum int
		for i, instruction := range instructions {
			cpu.registers = before
			instruction(sampleA, sampleB, sampleC)
			if cpu.registers == after {
				sum++
				instructionsMap[sampleOpcode] = add(instructionsMap[sampleOpcode], i)
			}
		}
		if sum >= 3 {
			threeOrMore++
		}
	}
	fmt.Println(threeOrMore)
	instructionSet := make(map[int]Instruction)
	toDelete := make(map[int]bool)
	for len(instructionSet) != 16 {
		for opcode, ids := range instructionsMap {
			if len(ids) == 1 {
				toDelete[ids[0]] = true
				instructionSet[opcode] = instructions[ids[0]]
				continue
			}
			for i := 0; i < len(ids); i++ {
				if toDelete[ids[i]] {
					instructionsMap[opcode] = append(instructionsMap[opcode][:i], ids[i+1:]...)
				}
			}
		}
	}

	cpu.registers = [4]int{}
	for _, line := range strings.Split(program, "\n") {
		fields := strings.Fields(line)
		opcode, _ := strconv.Atoi(fields[0])
		a, _ := strconv.Atoi(fields[1])
		b, _ := strconv.Atoi(fields[2])
		c, _ := strconv.Atoi(fields[3])
		instructionSet[opcode](a, b, c)
	}
	fmt.Println(cpu.registers[0])
}

func add(set []int, i int) []int {
	for _, val := range set {
		if val == i {
			return set
		}
	}
	return append(set, i)
}
