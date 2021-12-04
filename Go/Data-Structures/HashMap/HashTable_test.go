package HashTable

import (
	"fmt"
	"testing"
)

func TestHashTable(t *testing.T) {
	table := NewHashTable()
	fmt.Println("Before Add('test1', 1), Get('test1') ==> ", table.Get("test1"))
	table.Add("test1", 1)
	fmt.Println("After Add('test1', 1), Get('test1') ==> ", table.Get("test1"))
	table.Set("test1", 2)
	fmt.Println("After Set('test1', 2), Get('test1') ==> ", table.Get("test1"))
	table.Remove("test1")
	fmt.Println("After Remove('test1'), Get('test1') ==> ", table.Get("test1"))
}
