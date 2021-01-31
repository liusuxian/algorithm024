package Week_01

import "testing"

type myCircularDeque struct {
	cache    []int
	capacity int
	length   int
	front    int
	rear     int
}

func Constructor(k int) myCircularDeque {
	return myCircularDeque{
		cache:    make([]int, k),
		capacity: k,
		front:    1,
	}
}

func (d *myCircularDeque) insertFront(value int) bool {
	if d.length == d.capacity {
		return false
	}
	d.length++
	d.front--
	if d.front == -1 {
		d.front = d.capacity - 1
	}
	d.cache[d.front] = value
	return true
}

func (d *myCircularDeque) insertLast(value int) bool {
	if d.length == d.capacity {
		return false
	}
	d.length++
	d.rear++
	if d.rear == d.capacity {
		d.rear = 0
	}
	d.cache[d.rear] = value
	return true
}

func (d *myCircularDeque) deleteFront() bool {
	if d.length == 0 {
		return false
	}
	d.length--
	d.front++
	if d.front == d.capacity {
		d.front = 0
	}
	return true
}

func (d *myCircularDeque) deleteLast() bool {
	if d.length == 0 {
		return false
	}
	d.length--
	d.rear--
	if d.rear == -1 {
		d.rear = d.capacity - 1
	}
	return true
}

func (d *myCircularDeque) getFront() int {
	if d.length == 0 {
		return -1
	}
	return d.cache[d.front]
}

func (d *myCircularDeque) getRear() int {
	if d.length == 0 {
		return -1
	}
	return d.cache[d.rear]
}

func (d *myCircularDeque) isEmpty() bool {
	return d.length == 0
}

func (d *myCircularDeque) isFull() bool {
	return d.length == d.capacity
}

func TestMyCircularDeque(t *testing.T) {
	circularDeque := Constructor(3)

	t.Log(circularDeque.insertLast(1) == true)   // 返回 true
	t.Log(circularDeque.insertLast(2) == true)   // 返回 true
	t.Log(circularDeque.insertFront(3) == true)  // 返回 true
	t.Log(circularDeque.insertFront(4) == false) // 已经满了，返回 false

	t.Log(circularDeque.getRear() == 2)         // 返回 2
	t.Log(circularDeque.isFull() == true)       // 返回 true
	t.Log(circularDeque.deleteLast() == true)   // 返回 true
	t.Log(circularDeque.insertFront(4) == true) // 返回 true
	t.Log(circularDeque.getFront() == 4)        // 返回 4
}
