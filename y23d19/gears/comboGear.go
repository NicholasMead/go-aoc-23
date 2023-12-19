package gears

type ComboGear struct {
	Properties map[Property][2]int
	Location   string
}

func (combo ComboGear) Combinations() int {
	total := 1

	for _, minMax := range combo.Properties {
		value := minMax[1] - minMax[0] + 1
		total *= value
	}

	return total
}

func (combo ComboGear) Copy() ComboGear {
	copy := ComboGear{Location: combo.Location, Properties: map[Property][2]int{}}
	for prop, value := range combo.Properties {
		copy.Properties[prop] = value
	}
	return copy
}

func (combo ComboGear) Split(property Property, value int) (ComboGear, ComboGear) {
	var (
		left  = ComboGear{Location: combo.Location, Properties: map[Property][2]int{}}
		right = ComboGear{Location: combo.Location, Properties: map[Property][2]int{}}
	)

	for prop, value := range combo.Properties {
		if prop != property {
			left.Properties[prop] = value
			right.Properties[prop] = value
		}
	}

	left.Properties[property] = [2]int{combo.Properties[property][0], value - 1}
	right.Properties[property] = [2]int{value, combo.Properties[property][1]}

	return left, right
}

func DefaultCombo(location string) ComboGear {
	return ComboGear{
		Properties: map[Property][2]int{
			ExtremelyCoolLooking: {1, 4000},
			Musical:              {1, 4000},
			Aerodynamics:         {1, 4000},
			Shiny:                {1, 4000},
		},
		Location: location,
	}
}
