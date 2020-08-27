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

// Package gopool 提供一个协程池接口
package gopool

// Pool 协程池接口
type Pool interface {
	// 提交任务
	Go(f func())
	// 动态调整池大小
	Tune(size int)
	// 运行中的实例个数
	Running() int
	// 空闲空间大小
	Free() int
	// 池总大小
	Cap() int
}

// Go run on the goroutine pool,if pool is nil run on goroutine, never failed
func Go(p Pool, f func()) {
	if p != nil {
		p.Go(f)
	} else {
		go f()
	}
}
