package intro

import (
	"fmt"
	"sync"
)

/*
MemoryAccessSynchronization solves a data race: two concurrent processes which
are attempting to access the same area of memory, and the way they are accessing the memory
is not atomic.

While we have solved our data race, we haven’t actually solved our race condition! The order
of operations in this program is still nondeterministic; we’ve just narrowed the scope of the
nondeterminism a bit.
*/
func MemoryAccessSynchronization() {
	var memoryAccess sync.Mutex
	var data int
	go func() {
		memoryAccess.Lock()
		data++
		memoryAccess.Unlock()
	}()
	memoryAccess.Lock()
	if data == 0 {
		fmt.Println("the value is 0.")
	} else {
		fmt.Printf("the value is %v.\n", data)
	}
	memoryAccess.Unlock()
}
