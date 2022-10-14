// Copyright 2022 TriggerMesh Inc.
// SPDX-License-Identifier: Apache-2.0

package lib

import (
	"context"
	"testing"
)

type TestsFlags struct {
	RedisAddress      string
	RedisPassword     string
	RedisStreamPrefix string
}

type TestManager struct {
	flags   TestsFlags
	runners map[string]*BrokerTestRunner
}

func NewTestManager(flags TestsFlags) *TestManager {
	return &TestManager{
		flags:   flags,
		runners: make(map[string]*BrokerTestRunner),
	}
}

func (tm *TestManager) NewBrokerTestRunner(ctx context.Context, name string, t *testing.T) *BrokerTestRunner {
	btr := NewBrokerTestRunner(ctx, t)
	tm.runners[name] = btr

	return btr
}
