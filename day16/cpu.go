package day16

type instruction struct {
	opcode int

	a int
	b int
	c int
}

type snapshot struct {
	before [4]int
	i      instruction
	after  [4]int
}

type cpu struct {
	registers [4]int
	opcodes   [16]instructionImplementation

	ip      int
	program []instruction
}

func (c *cpu) train(ss []snapshot) (threeOrMore int) {
	for _, s := range ss {
		count := 0
		for _, i := range implementations {
			if i.candidate(&s.i, s.before, s.after) {
				count++
			}
		}

		if count >= 3 {
			threeOrMore++
		}
	}

	// TODO: Determine implementation
	return
}

func (c *cpu) step() {
	i := c.program[c.ip]
	c.opcodes[i.opcode].exec(&i, c)
	c.ip++
}

func (c *cpu) run() {
	for c.ip < len(c.program) {
		c.step()
	}
}
