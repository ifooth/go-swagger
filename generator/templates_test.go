// Copyright 2015 go-swagger maintainers
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

package generator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomTemplates(t *testing.T) {

	registry := NewTemplateRegistry()
	registry.LoadDefaults()

	err := registry.LoadDir("../fixtures/templates/")

	assert.NoError(t, err)

	for template, v := range registry.assets {
		if string(v) != fmt.Sprintf("./%s\n", template) {
			t.Errorf("Template %s wasn't loaded", template)
		}
	}

}
