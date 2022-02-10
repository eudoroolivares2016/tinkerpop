/*
Licensed to the Apache Software Foundation (ASF) Under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

package gremlingo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTraversal(t *testing.T) {
	t.Run("Test clone traversal", func(t *testing.T) {
		g := NewGraphTraversalSource(&Graph{}, &TraversalStrategies{}, newBytecode(nil))
		original := g.V().Out("created")
		clone := original.Clone().Out("knows")
		cloneClone := clone.Clone().Out("created")

		assert.Equal(t, 2, len(original.t.bytecode.stepInstructions))
		assert.Equal(t, 3, len(clone.t.bytecode.stepInstructions))
		assert.Equal(t, 4, len(cloneClone.t.bytecode.stepInstructions))

		original.Has("person", "name", "marko")
		clone.V().Out()

		assert.Equal(t, 3, len(original.t.bytecode.stepInstructions))
		assert.Equal(t, 5, len(clone.t.bytecode.stepInstructions))
		assert.Equal(t, 4, len(cloneClone.t.bytecode.stepInstructions))
	})
}
