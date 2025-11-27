package types

// Pointer
type PointerType *string // var pointerName *string

// Array
type ArrayType [3]string // var arrayName = [3]string{"a", "b", "e"}

// Slice
type SliceType []string // var sliceName = []string{"a", "b", "c"}

// Map
type MapType map[string]string // var mapName = map[string]string{"a": "1", "b": "2"}

// Struct
type StructType struct {
	Name string
	Age  int
}

// Function
type FuncType func(string) (int, string) //var funcName func(string) (int, string)

// Interface
type InterfaceType interface {
	Method1(int) string
	Method2() int
} // var interfaceName InterfaceType

// Channel
type ChannelType chan string        // var channelName chan string
type ReadChannelType <-chan string  // var readChannelName <-chan string
type WriteChannelType chan<- string // var writeChannelName chan<- string
