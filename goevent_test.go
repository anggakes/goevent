package goevent

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegisterListenerAndPublish(t *testing.T) {

	updateState := 1

	type MsgData struct {
		Increment int
	}

	listener1 := func(ctx context.Context, msg interface{}) error {

		m, ok := msg.(*MsgData)
		if !ok {
			return errors.New("invalid message")
		}

		fmt.Println("published event listener1")

		updateState += m.Increment

		return nil
	}

	RegisterListener("listener1", listener1)
	RegisterListener("listener1", listener1)

	// test publish
	ctx := context.Background()

	Publish(ctx, "listener1", &MsgData{Increment: 2})

	assert.Equal(t, updateState, 3)

	Publish(ctx, "listener1", &MsgData{Increment: 3})

	assert.Equal(t, updateState, 6)


}
