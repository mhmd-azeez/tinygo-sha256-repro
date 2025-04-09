# TinyGo Reactor Module Sha256 bug repro

```
➜  sha256-repro ./build.sh 
➜  sha256-repro extism call --wasi command.wasm _start
hash: b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9
➜  sha256-repro extism call --wasi reactor.wasm hash  
panic: runtime error: nil pointer dereference
Error: wasm error: unreachable
wasm stack trace:
	main.runtime.runtimePanicAt(i32,i32)
		0xda05: /usr/local/lib/tinygo/src/runtime/runtime_tinygowasm.go:78:6
	main.runtime.nilPanic()
		0x2ee3: /usr/local/lib/tinygo/src/runtime/panic.go:180:16
	main.crypto/internal/fips140.RecordApproved()
		0x1d0a0: /usr/local/lib/tinygo/src/runtime/scheduler.go:37:24 (inlined)
	main.(*crypto/internal/fips140/sha256.Digest).Sum(i32,i32,i32,i32)
		0x1da16: /usr/local/go/src/crypto/internal/fips140/sha256/sha256.go:185:24 (inlined)
	main.hash()
		0x354da: /usr/local/lib/tinygo/src/sync/pool.go:29:16 (inlined)
		         /usr/local/go/src/encoding/hex/hex.go:164:24 (inlined)
		         /home/mo/work/x/wasm/sha256-repro/main.go:892:13 (inlined)
		         /usr/local/go/src/encoding/hex/hex.go:164:24 (inlined)
		         /home/mo/work/x/wasm/sha256-repro/main.go:875:14 (inlined)
returned non-zero exit code: 1
```