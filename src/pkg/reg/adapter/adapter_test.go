// Copyright Project Harbor Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package adapter

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/goharbor/harbor/src/pkg/reg/model"
)

type fakedFactory struct {
}

func (fakedFactory) Create(*model.Registry) (Adapter, error) {
	return nil, nil
}

func (fakedFactory) AdapterPattern() *model.AdapterPattern {
	return nil
}

func TestRegisterFactory(t *testing.T) {
	// empty type
	assert.NotNil(t, RegisterFactory("", nil))
	// empty factory
	assert.NotNil(t, RegisterFactory("harbor", nil))
	// pass
	assert.Nil(t, RegisterFactory("harbor", new(fakedFactory)))
	// already exists
	assert.NotNil(t, RegisterFactory("harbor", new(fakedFactory)))
}

func TestGetFactory(t *testing.T) {
	registry = map[string]Factory{}
	require.Nil(t, RegisterFactory("harbor", new(fakedFactory)))
	// doesn't exist
	_, err := GetFactory("gcr")
	assert.NotNil(t, err)
	// pass
	_, err = GetFactory("harbor")
	assert.Nil(t, err)
}

func TestListRegisteredAdapterTypes(t *testing.T) {
	registry = map[string]Factory{}
	registryKeys = []string{}
	// not register, got nothing
	types := ListRegisteredAdapterTypes()
	assert.Equal(t, 0, len(types))

	// register one factory
	require.Nil(t, RegisterFactory("harbor", new(fakedFactory)))

	types = ListRegisteredAdapterTypes()
	require.Equal(t, 1, len(types))
	assert.Equal(t, "harbor", types[0])
}

func TestListRegisteredAdapterTypesOrder(t *testing.T) {
	registry = map[string]Factory{}
	registryKeys = []string{}
	require.Nil(t, RegisterFactory("a", new(fakedFactory)))
	require.Nil(t, RegisterFactory("c", new(fakedFactory)))
	require.Nil(t, RegisterFactory("b", new(fakedFactory)))

	types := ListRegisteredAdapterTypes()
	require.Equal(t, 3, len(types))
	require.Equal(t, "a", types[0])
	require.Equal(t, "b", types[1])
	require.Equal(t, "c", types[2])
}
