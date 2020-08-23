package goevent

import (
	"context"
	"github.com/pkg/errors"
)

type ListenerFunc func(ctx context.Context, msg interface{}) error

var l  map[string]ListenerFunc

var (
	ErrNoListener = errors.New("Error No listener")
)

func Publish(ctx context.Context, name string, msg interface{}) error {

	k, ok := l[name]
	if !ok {
		return ErrNoListener
	}

	return k (ctx, msg)
}


func RegisterListener(name string, listener ListenerFunc) {

	if l == nil {
		l = map[string]ListenerFunc{}
	}

	l[name] = listener
}

