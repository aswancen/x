package x

import (
	"log"
)

// Result 结构体表示带有错误信息的结果
type Result[T any] struct {
	Value T
	Err   error
}

// Ok 返回一个成功的 Result
func Ok[T any](value T) Result[T] {
	return Result[T]{Value: value}
}

// Err 返回一个包含错误的 Result
func Err[T any](err error) Result[T] {
	return Result[T]{Err: err}
}

// Expect 返回包含在 Ok 中的值，如果是 Err 则输出错误信息并终止程序
func (this Result[T]) Expect(ex string) T {
	if this.Err != nil {
		log.Panicf("Expect failed: %s, Error: %v", ex, this.Err)
	}
	return this.Value
}

// Unwrap 返回包含在 Ok 中的值，如果是 Err 则输出默认错误信息并终止程序
func (this Result[T]) Unwrap() T {
	if this.Err != nil {
		log.Panicf("Unwrap failed: %v", this.Err)
	}
	return this.Value
}

// Process 处理 Result
func (this Result[T]) Process(f func(v T, e error) T) T {
	return f(this.Value, this.Err)
}
