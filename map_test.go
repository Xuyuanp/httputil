/*
 * Copyright 2015 Xuyuan Pang
 * Author: Xuyuan Pang
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package httputil

import "testing"

func TestMapKeys(t *testing.T) {
	m := map[interface{}]interface{}{
		1:   "A",
		"b": 2,
	}
	keys := MapKeys(m)
	if len(keys) != len(m) {
		t.Errorf("wrong length of keys. %d expected, %d got", len(m), len(keys))
	}

	for _, key := range keys {
		if _, ok := m[key]; !ok {
			t.Errorf("key: %v not in map", key)
		}
	}
}
