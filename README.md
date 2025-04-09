# TinyGo Reactor Module Sha256 bug repro

```
➜  sha256-repro git:(main) ✗ ./build.sh                            
➜  sha256-repro git:(main) ✗ extism call --wasi command.wasm _start
sha1: 2aae6c35c94fcfb415dbe95f408b9ce91ee846ed
sha256: b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9

➜  sha256-repro git:(main) ✗ extism call --wasi reactor.wasm sha1  
sha1: 2aae6c35c94fcfb415dbe95f408b9ce91ee846ed

➜  sha256-repro git:(main) ✗ extism call --wasi reactor.wasm sha256
panic: runtime error: nil pointer dereference
Error: wasm error: unreachable
wasm stack trace:
	main.runtime.runtimePanicAt(i32,i32)
		0xda08: /usr/local/lib/tinygo/src/runtime/runtime_tinygowasm.go:78:6
	main.runtime.nilPanic()
		0x2ee6: /usr/local/lib/tinygo/src/runtime/panic.go:180:16
	main.crypto/internal/fips140.RecordApproved()
		0x1d0bc: /usr/local/lib/tinygo/src/runtime/scheduler.go:37:24 (inlined)
	main.(*crypto/internal/fips140/sha256.Digest).Sum(i32,i32,i32,i32)
		0x1dcc9: /usr/local/go/src/crypto/internal/fips140/sha256/sha256.go:185:24 (inlined)
	main.interface:{BlockSize:func:{}{basic:int},Reset:func:{}{},Size:func:{}{basic:int},Sum:func:{slice:basic:uint8}{slice:basic:uint8},Write:func:{slice:basic:uint8}{basic:int,named:error}}.Sum$invoke(i32,i32,i32,i32,i32)
		0x210db: /usr/local/go/src/crypto/sha1/sha1.go:155:15 (inlined)
		         /usr/local/go/src/io/fs/fs.go:110:34 (inlined)
		         <Go interface method>:101:33 (inlined)
		         <Go interface method>:103:29 (inlined)
		         /usr/local/lib/tinygo/src/runtime/slice.go:180:33 (inlined)
		         /usr/local/lib/tinygo/src/runtime/slice.go:298:36
	main.sha256()
		0x38228: /home/mo/work/x/wasm/sha256-repro/main.go:28:34 (inlined)
		         /home/mo/work/x/wasm/sha256-repro/main.go:12:18 (inlined)
returned non-zero exit code: 1
➜  sha256-repro git:(main) ✗ tinygo version
tinygo version 0.37.0 linux/amd64 (using go version go1.24.2 and LLVM version 19.1.2)
```