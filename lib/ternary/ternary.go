// Copyright [2020] [thinkgos] thinkgo@aliyun.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package ternary implement like condition ? trueVal : falseVal
package ternary

// If like condition ? trueVal : falseVal
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

// IfInt like condition ? trueVal : falseVal
func IfInt(condition bool, trueVal, falseVal int) int {
	if condition {
		return trueVal
	}
	return falseVal
}

// IfInt64 like condition ? trueVal : falseVal
func IfInt64(condition bool, trueVal, falseVal int64) int64 {
	if condition {
		return trueVal
	}
	return falseVal
}

// IfFloat like condition ? trueVal : falseVal
func IfFloat(condition bool, trueVal, falseVal float64) float64 {
	if condition {
		return trueVal
	}
	return falseVal
}

// IfString like condition ? trueVal : falseVal
func IfString(condition bool, trueVal, falseVal string) string {
	if condition {
		return trueVal
	}
	return falseVal
}
