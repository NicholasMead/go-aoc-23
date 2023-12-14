package platform

type Platform interface {
	AddRock(Vector, RockType)
	ApplyTilt(direction Vector)

	GetRocksOfType(RockType) []Vector
}

type platform struct {
	height, width int
	rocks         map[Vector]Rock
}

func NewPlatform(height, width int) Platform {
	return &platform{
		height,
		width,
		map[Vector]Rock{},
	}
}

func (p *platform) AddRock(position Vector, rockType RockType) {
	if _, found := p.rocks[position]; found {
		panic("Space taken")
	}

	p.rocks[position] = makeRock(p, position, rockType)
}

func (p *platform) ApplyTilt(direction Vector) {
	didMove := true
	count, max := 0, p.width*p.height*len(p.rocks)

	for didMove && count < max {
		count++
		didMove = false

		for _, rock := range p.rocks {
			err := rock.Move(direction)

			if err == nil {
				didMove = true
			}
		}
	}
}

func (p *platform) GetRocksOfType(targetType RockType) (out []Vector) {
	for pos, rock := range p.rocks {
		if rock.Type() == targetType {
			out = append(out, pos)
		}
	}
	return
}

func (p *platform) String() string {
	platformString := ""
	platformString += "********\n"

	for y := 0; y < p.height; y++ {
		for x := 0; x < p.width; x++ {
			rock, found := p.rocks[Vector{x, y}]

			if found {
				platformString += string(rock.Type())
			} else {
				platformString += "."
			}
		}
		platformString += "\n"
	}
	platformString += "********\n"

	return platformString
}
