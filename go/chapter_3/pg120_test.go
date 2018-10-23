package pg120

import "testing"
import "github.com/collections/stack"

type (
	StackBasedQueue struct {
		mainStack *stack.Stack
		helperStack *stack.Stack
	}
)

func New() *StackBasedQueue {
	return &StackBasedQueue{stack.New(), stack.New()}
}

func (this *StackBasedQueue) Len() int {
	return this.mainStack.Len()
}

func (this *StackBasedQueue) Enqueue(value interface{}) {
	for ; this.mainStack.Len() > 0 ; {
		this.helperStack.Push(this.mainStack.Pop())
	}
	this.mainStack.Push(value)
}

func (this *StackBasedQueue) Dequeue() interface{} {
	for ; this.helperStack.Len() > 0 ; {
		this.mainStack.Push(this.helperStack.Pop())
	}
	if this.mainStack.Len() > 0 {
		return this.mainStack.Pop()
	}
	return nil
}


func TestQueue(t * testing.T) {
	queue := New()
	queue.Enqueue(1)
	queue.Enqueue(2)

	deq1 := queue.Dequeue()
	deq2 := queue.Dequeue()

	if deq1 != 1 || deq2 != 2 || queue.Len() != 0 || queue.Dequeue() != nil {
		t.Errorf("Err")
	}
}