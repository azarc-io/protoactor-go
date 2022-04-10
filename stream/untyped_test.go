package stream

import (
	"testing"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/stretchr/testify/assert"
)

var system = actor.NewActorSystem()

func TestReceiveFromStream(t *testing.T) {
	s := NewUntypedStream(system)
	go func() {
		rootContext := system.Root
		rootContext.Send(s.PID(), "hello")
		rootContext.Send(s.PID(), "you")
	}()
	res := <-s.C()
	res2 := <-s.C()
	assert.Equal(t, "hello", res.(string))
	assert.Equal(t, "you", res2.(string))
}
