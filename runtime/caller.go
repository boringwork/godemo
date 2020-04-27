package main

import (
	"fmt"
	"runtime"
)

func main() {
	// bar()
	fmt.Println(getFrame(0))
	// test()
}

func test() {
	fmt.Println(runtime.Caller(0))
	fmt.Println(runtime.Caller(1))
	fmt.Println(runtime.Caller(2))
	fmt.Println(runtime.Caller(3))
}

func getFrame(skipFrames int) string {
	// We need the frame at index skipFrames+2, since we never want runtime.Callers and getFrame
	targetFrameIndex := skipFrames + 2

	// Set size to targetFrameIndex+2 to ensure we have room for one more caller than we need
	programCounters := make([]uintptr, targetFrameIndex+2)
	n := runtime.Callers(0, programCounters)

	frame := runtime.Frame{Function: "unknown"}
	if n > 0 {
		frames := runtime.CallersFrames(programCounters[:n])
		for more, frameIndex := true, 0; more && frameIndex <= targetFrameIndex; frameIndex++ {
			var frameCandidate runtime.Frame
			frameCandidate, more = frames.Next()
			if frameIndex == targetFrameIndex {
				frame = frameCandidate
			}
		}
	}

	return frame.Function
}

// MyCaller returns the caller of the function that called it :)
// func MyCaller() string {
// 	// Skip GetCallerFunctionName and the function to get the caller of
// 	return getFrame(2).Function
// }

// foo calls MyCaller
// func foo() {
// 	fmt.Println(MyCaller())
// }

// // bar is what we want to see in the output - it is our "caller"
// func bar() {
// 	foo()
// }
