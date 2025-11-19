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