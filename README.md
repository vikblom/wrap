# Wrap

Wrap prepends callsite info like `(file.go:123): ...` to Go errors.

It is intended to be a drop-in replacement for `fmt.Errorf`.
If used all the way through the call chain then errors will contain a chain of file+line information.
This "partial stacktrace" info is usually enough to satisfy colleagues complaining that Go does not capture stacktraces.

Note that only instrumented error wrapping adds file+line information.
There is no way to get the stack info at a "source" if that source is in some external module.
As such the chain of locations usually showswhere the error entered the module code, and the path it "bubbles up".

Note that wrapping errors like this adds a little extra work, use it judiciously and avoid any hot paths.
Benchmark your use-case if you need to be certain.
