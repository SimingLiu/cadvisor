// +build libpfm,cgo

// Copyright 2020 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Manager of perf events for containers.
package perf

import (
	"testing"

	"github.com/stretchr/testify/assert"

	info "github.com/google/cadvisor/info/v1"
	"github.com/google/cadvisor/stats"
)

func TestEmptyConfigPassed(t *testing.T) {
	manager, err := NewManager("testing/perf-no-events.json", []info.Node{})

	assert.NotNil(t, err)
	assert.Nil(t, manager)
}

func TestNoConfigFilePassed(t *testing.T) {
	manager, err := NewManager("", []info.Node{})

	assert.Nil(t, err)
	_, ok := manager.(*stats.NoopManager)
	assert.True(t, ok)
}

func TestNonExistentFile(t *testing.T) {
	manager, err := NewManager("this-file-is-so-non-existent", []info.Node{})

	assert.NotNil(t, err)
	assert.Nil(t, manager)
}

func TestMalformedJsonFile(t *testing.T) {
	manager, err := NewManager("testing/this-is-some-random.json", []info.Node{})

	assert.NotNil(t, err)
	assert.Nil(t, manager)
}

func TestNewManager(t *testing.T) {
	managerInstance, err := NewManager("testing/perf.json", []info.Node{})

	assert.Nil(t, err)
	_, ok := managerInstance.(*manager)
	assert.True(t, ok)
}
