# Reference types - Pass by reference (không cần pointer)
map[K]V
[]T          // slice
chan T       // channel
interface{}  // interface
func()       // function

#  Value types - Pass by value (cần pointer để modify)
struct
array [N]T
int, string, bool, float64, etc.

# Buffered Channel KHÔNG block luồng chính 
channel := make(chan int, 3)
# Unbuffered Channel BLOCK luồng chính
 channel := make(chan int)