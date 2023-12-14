package platform

import "strings"

type Platform interface {
	AddRock(Vector, RockType)
	ApplyTilt(direction Vector)
	SpinCycle()

	NorthWeight() int

	Hash() string
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

func LoadPlatform(input []string) Platform {
	height, width := len(input), len(input[0])
	p := NewPlatform(height, width)

	for y, line := range input {
		for x, char := range line {
			if !strings.ContainsRune(RockTypes, char) {
				continue
			}

			p.AddRock(Vector{x, y}, RockType(char))
		}
	}
	return p
}

func LoadPlatformFromHash(hash string) Platform {
	input := strings.Split(hash, "\n")
	return LoadPlatform(input)
}

func (p *platform) Width() int {
	return p.width
}
func (p *platform) Height() int {
	return p.height
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

func (p *platform) SpinCycle() {
	p.ApplyTilt(Vector{+0, -1})
	p.ApplyTilt(Vector{-1, +0})
	p.ApplyTilt(Vector{+0, +1})
	p.ApplyTilt(Vector{+1, +0})
}

func (p *platform) NorthWeight() int {
	total := 0
	for _, rock := range p.rocks {
		dist := p.Height() - rock.Position()[1]
		total += dist * rock.Weight()
	}
	return total
}

func (p *platform) String() string {
	platformString := ""

	for y := 0; y < p.height; y++ {
		for x := 0; x < p.width; x++ {
			rock, found := p.rocks[Vector{x, y}]

			if found {
				platformString += rock.String()
			} else {
				platformString += "."
			}
		}
		platformString += "\n"
	}

	return strings.TrimSpace(platformString)
}

func (p *platform) Hash() string {
	return p.String()
}
