/*
 * Copyright (c) 2020. Temple3x (temple3x@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package config

import (
	"time"

	"github.com/templexxx/crank/typeutil"
)

// Adjust val in config to default value if need.
// val must be a pointer, and the type of defValue must be
// as same as *val.
func Adjust(val interface{}, defValue interface{}) {
	switch v := val.(type) {
	case *string:
		if *v == "" {
			*v = defValue.(string)
		}
	case *int:
		if *v == 0 {
			*v = defValue.(int)
		}
	case *int64:
		if *v == 0 {
			*v = defValue.(int64)
		}
	case *uint32:
		if *v == 0 {
			*v = defValue.(uint32)
		}
	case *float64:
		if *v == 0 {
			*v = defValue.(float64)
		}
	case *uint64:
		if *v == 0 {
			*v = defValue.(uint64)
		}
	case *time.Duration:
		if *v <= 0 {
			*v = defValue.(time.Duration)
		}
	case *typeutil.Duration:
		if v.Duration <= 0 {
			v.Duration = defValue.(time.Duration)
		}
	case *typeutil.ByteSize:
		if *v <= 0 {
			*v = defValue.(typeutil.ByteSize)
		}
	case *[]string:
		if len(*v) == 0 {
			*v = defValue.([]string)
		}
	}
}
