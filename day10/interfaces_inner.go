package main

type ReadWrite interface {
	Read(b Buffer) bool
	Write(b Buffer) bool
}

type Lock interface {
	Lock()
	Unlock()
}

type File interface {
	ReadWrite
	Lock
	Close()
}

// 一个接口可以包含其他接口，相当于内嵌接口
