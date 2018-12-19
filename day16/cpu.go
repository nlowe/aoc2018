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
	candidates := [16]map[*instructionImplementation]bool{}
	for i := 0; i < 16; i++ {
		candidates[i] = map[*instructionImplementation]bool{}
	}

	for _, s := range ss {
		count := 0
		for _, i := range implementations {
			if i.candidate(&s.i, s.before, s.after) {
				count++
				candidates[s.i.opcode][i] = true
			}
		}

		if count >= 3 {
			threeOrMore++
		}
	}

	for opcodesRemaining := 16; opcodesRemaining > 0; opcodesRemaining-- {
		for i, m := range candidates {
			if len(m) == 1 {
				for k := range m {
					c.opcodes[i] = *k
					for j := 0; j < 16; j++ {
						delete(candidates[j], k)
					}
				}
			}
		}
	}

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
