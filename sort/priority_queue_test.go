package sort

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestPriorityQ(t *testing.T) {
	// Some items and their priorities.
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	Init(&pq)

	// Insert a new item and then modify its priority.
	item := &Item{
		value:    "orange",
		priority: 3,
	}
	heap.Push(&pq, item)
	heap.Push(&pq, &Item{
		value:    "orange2",
		priority: 3,
	})

	heap.Push(&pq, &Item{
		value:    "orange3",
		priority: 3,
	})

	heap.Push(&pq, &Item{
		value:    "pear2",
		priority: 4,
	})

	//pq.update(item, item.value, 3)

	// Take the items out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%.2d:%s ", item.priority, item.value)
	}
	// Output:
	// 05:orange 04:pear 03:banana 02:apple

}
