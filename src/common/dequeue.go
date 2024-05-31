package common

import "errors"

type Dequeue struct {
    arr []any
    head int
    tail int
    size int
}

func NewDequeue(size int) *Dequeue {
    d := &Dequeue{
        arr: make([]any, size),
        head: 0,
        tail: 0,
        size: 0,
    }
    return d
}

func (dq *Dequeue) IsEmpty() bool {
    return dq.size == 0
}

func (dq *Dequeue) IsFull() bool {
    return dq.size == len(dq.arr)
}

func (dq *Dequeue) Size() int {
    return dq.size
}

func (dq *Dequeue) InsertHead(v any) error {
    if dq.IsFull() {
        return errors.New("Dequeue is full!")
    }
    if dq.IsEmpty() {
        dq.size ++
        dq.arr[dq.head] = v
        return nil
    }
    if dq.head == 0 {
        dq.head = len(dq.arr)-1
    } else {
        dq.head --
    }
    dq.size ++
    dq.arr[dq.head] = v
    return nil
}

func (dq *Dequeue) AppendTail(v any) error {
    if dq.IsFull() {
        return errors.New("Dequeue is full!")
    }
    if dq.IsEmpty() {
        dq.size ++
        dq.arr[dq.tail] = v
        return nil
    }
    if dq.tail == len(dq.arr)-1 {
        dq.tail = 0
    } else {
        dq.tail ++
    }
    dq.size ++
    dq.arr[dq.tail] = v
    return nil
}

func (dq *Dequeue) PopHead() *any {
    if dq.IsEmpty() {
        return nil
    }
    v := dq.arr[dq.head]
    dq.size --
    if dq.size == 0 {
        dq.head = 0
        dq.tail = 0
        return &v
    }
    if dq.head < len(dq.arr)-1 {
        dq.head ++
        return &v
    }
    dq.head = 0
    return &v
}

func (dq *Dequeue) PopTail() *any {
    if dq.IsEmpty() {
        return nil
    }
    v := dq.arr[dq.tail]
    dq.size --
    if dq.size == 0 {
        dq.head = 0
        dq.tail = 0
        return &v
    }
    if dq.tail > 0 {
        dq.tail --
        return &v
    }
    dq.tail = len(dq.arr)-1
    return &v
}

func (dq *Dequeue) GetHead() *any {
    if dq.IsEmpty() {
        return nil
    }
    v := dq.arr[dq.head]
    return &v
}

func (dq *Dequeue) GetTail() *any {
    if dq.IsEmpty() {
        return nil
    }
    v := dq.arr[dq.tail]
    return &v
}

