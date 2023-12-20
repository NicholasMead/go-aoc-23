package machine

import "testing"

func TestBroadcast(t *testing.T) {
	var b Module = &Broadcast{}

	out := b.GetConnection("a").Send(LOW)

	if out == nil || *out != LOW {
		t.Errorf("got %v expected %v", out, LOW)
	}
}
