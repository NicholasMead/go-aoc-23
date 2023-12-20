package machine

type Broadcast struct{}

// Connect implements Module.
func (b Broadcast) Connect(string) {}

// GetConnection implements Module.
func (b Broadcast) GetConnection(string) Wire {
	return broadcastInput{}
}

func (b Broadcast) Type() string {
	return "Broadcast"
}

type broadcastInput struct{}

func (broadcastInput) Send(p Pulse) *Pulse {
	return &p
}
