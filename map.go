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

import "reflect"

// MapKeys returns all keys of the map
func MapKeys(m interface{}) []interface{} {
	v := reflect.ValueOf(m)
	if v.Kind() != reflect.Map {
		return nil
	}

	vkeys := v.MapKeys()
	keys := make([]interface{}, len(vkeys))
	for i, vkey := range vkeys {
		keys[i] = vkey.Interface()
	}
	return keys
}
