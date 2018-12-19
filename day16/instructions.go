package day16

type instructionImplementation struct {
	candidate func(i *instruction, before, after [4]int) bool
	exec      func(i *instruction, c *cpu)
}

var (
	addr = &instructionImplementation{
		candidate: func(i *instruction, before, after [4]int) bool {
			return after[i.c] == before[i.a]+before[i.b]
		},
		exec: func(i *instruction, c *cpu) {
			c.registers[i.c] = c.registers[i.a] + c.registers[i.b]
		},
	}

	addi = &instructionImplementation{
		candidate: func(i *instruction, before, after [4]int) bool {
			return after[i.c] == before[i.a]+i.b
		},
		exec: func(i *instruction, c *cpu) {
			c.registers[i.c] = c.registers[i.a] + i.b
		},
	}

	mulr = &instructionImplementation{
		candidate: func(i *instruction, before, after [4]int) bool {
			return after[i.c] == before[i.a]*before[i.b]
		},
		exec: func(i *instruction, c *cpu) {
			c.registers[i.c] = c.registers[i.a] * c.registers[i.b]
		},
	}

	muli = &instructionImplementation{
		candidate: func(i *instruction, before, after [4]int) bool {
			return after[i.c] == before[i.a]*i.b
		},
		exec: func(i *instruction, c *cpu) {
			c.registers[i.c] = c.registers[i.a] * i.b
		},
	}

	banr = &instructionImplementation{
		candidate: func(i *instruction, before, after [4]int) bool {
			return after[i.c] == before[i.a]&before[i.b]
		},
		exec: func(i *instruction, c *cpu) {
			c.registers[i.c] = c.registers[i.a] & c.registers[i.b]
		},
	}

	bani = &instructionImplementation{
		candidate: func(i *instruction, before, after [4]int) bool {
			return after[i.c] == before[i.a]&i.b
		},
		exec: func(i *instruction, c *cpu) {
			c.registers[i.c] = c.registers[i.a] & i.b
		},
	}

	borr = &instructionImplementation{
		candidate: func(i *instruction, before, after [4]int) bool {
			return after[i.c] == before[i.a]|before[i.b]
		},
		exec: func(i *instruction, c *cpu) {
			c.registers[i.c] = c.registers[i.a] | c.registers[i.b]
		},
	}

	bori = &instructionImplementation{
		candidate: func(i *instruction, before, after [4]int) bool {
			return after[i.c] == before[i.a]|i.b
		},
		exec: func(i *instruction, c *cpu) {
			c.registers[i.c] = c.registers[i.a] | i.b
		},
	}

	setr = &instructionImplementation{
		candidate: func(i *instruction, before, after [4]int) bool {
			return after[i.c] == before[i.a]
		},
		exec: func(i *instruction, c *cpu) {
			c.registers[i.c] = c.registers[i.a]
		},
	}

	seti = &instructionImplementation{
		candidate: func(i *instruction, before, after [4]int) bool {
			return after[i.c] == i.a
		},
		exec: func(i *instruction, c *cpu) {
			c.registers[i.c] = i.a
		},
	}

	gtir = &instructionImplementation{
		candidate: func(i *instruction, before, after [4]int) bool {
			if i.a > before[i.b] {
				return after[i.c] == 1
			} else {
				return after[i.c] == 0
			}
		},
		exec: func(i *instruction, c *cpu) {
			if i.a > c.registers[i.b] {
				c.registers[i.c] = 1
			} else {
				c.registers[i.c] = 0
			}
		},
	}

	gtri = &instructionImplementation{
		candidate: func(i *instruction, before, after [4]int) bool {
			if before[i.a] > i.b {
				return after[i.c] == 1
			} else {
				return after[i.c] == 0
			}
		},
		exec: func(i *instruction, c *cpu) {
			if c.registers[i.a] > i.b {
				c.registers[i.c] = 1
			} else {
				c.registers[i.c] = 0
			}
		},
	}

	gtrr = &instructionImplementation{
		candidate: func(i *instruction, before, after [4]int) bool {
			if before[i.a] > before[i.b] {
				return after[i.c] == 1
			} else {
				return after[i.c] == 0
			}
		},
		exec: func(i *instruction, c *cpu) {
			if c.registers[i.a] > c.registers[i.b] {
				c.registers[i.c] = 1
			} else {
				c.registers[i.c] = 0
			}
		},
	}

	eqir = &instructionImplementation{
		candidate: func(i *instruction, before, after [4]int) bool {
			if i.a == before[i.b] {
				return after[i.c] == 1
			} else {
				return after[i.c] == 0
			}
		},
		exec: func(i *instruction, c *cpu) {
			if i.a == c.registers[i.b] {
				c.registers[i.c] = 1
			} else {
				c.registers[i.c] = 0
			}
		},
	}

	eqri = &instructionImplementation{
		candidate: func(i *instruction, before, after [4]int) bool {
			if before[i.a] == i.b {
				return after[i.c] == 1
			} else {
				return after[i.c] == 0
			}
		},
		exec: func(i *instruction, c *cpu) {
			if c.registers[i.a] == i.b {
				c.registers[i.c] = 1
			} else {
				c.registers[i.c] = 0
			}
		},
	}

	eqrr = &instructionImplementation{
		candidate: func(i *instruction, before, after [4]int) bool {
			if before[i.a] == before[i.b] {
				return after[i.c] == 1
			} else {
				return after[i.c] == 0
			}
		},
		exec: func(i *instruction, c *cpu) {
			if c.registers[i.a] == c.registers[i.b] {
				c.registers[i.c] = 1
			} else {
				c.registers[i.c] = 0
			}
		},
	}

	implementations = []*instructionImplementation{addr, addi, mulr, muli, banr, bani, borr, bori, setr, seti, gtir, gtri, gtrr, eqir, eqri, eqrr}
)
