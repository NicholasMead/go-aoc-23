package machine

import "fmt"

type Pulse bool

var (
	LOW  Pulse = false
	HIGH Pulse = true
)

func (p Pulse) String() string {
	switch p {
	case HIGH:
		return "HIGH"
	case LOW:
		return "LOW"
	default:
		err := fmt.Errorf("unknown pulse %v", bool(p))
		panic(err)
	}
}
