package component

import (
	"context"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
)

var ErrInterrupted = errors.New("component: interrupted")

type Component interface {
	Run(ctx context.Context) error
}

type options struct {
	retryDelay time.Duration
}

type OptionsModifier func(opts *options)

func WithRetryDelay(duration time.Duration) OptionsModifier {
	return func(opts *options) {
		opts.retryDelay = duration
	}
}

func Start(ctx context.Context, component Component, optsModifiers ...OptionsModifier) {
	opts := options{
		retryDelay: 10 * time.Second,
	}
	for _, modifier := range optsModifiers {
		modifier(&opts)
	}

	go func() {
		logrus.WithContext(ctx).Debugf("starting the %T component", component)

		for {
			if err := component.Run(ctx); err != nil {
				if err == ErrInterrupted {
					logrus.WithContext(ctx).Debugf("the %T component has been interrupted", component)
					return
				}
				logrus.WithContext(ctx).WithError(err).Errorf("error while executing the %T component", component)
			}

			logrus.WithContext(ctx).Debugf("waiting %d seconds before the next execution", opts.retryDelay/time.Second)
			select {
			case <-time.After(opts.retryDelay):
				continue
			case <-ctx.Done():
				logrus.WithContext(ctx).Debugf("stopping the %T component", component)
				return
			}
		}
	}()
}
