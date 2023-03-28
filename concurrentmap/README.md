# Concurrent Map

A generic concurrent map that uses a single mutex for locking of the underlying generic map.

## Install

`go get github.com/zarldev/zarlib/concurrentmap`

## Use

Example concurrent map which has string keys and integer values

```golang
cmap := NewConcurrentMap[string, int]()
cmap.Set("zero", 0)
val, found := cmap.Get("zero") // 0
```
