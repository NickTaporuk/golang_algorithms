package queue

import "fmt"

func QueueExample() {
	q := New()

	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)

	for q.Len() != 0  {
		val := q.Dequeue()
		fmt.Println(val)
	}

}
