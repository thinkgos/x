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
