package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	length              int
	FrontElem, BackElem *ListItem
}

func (l *list) Len() int {
	return l.length
}

func (l *list) Front() *ListItem {
	return l.FrontElem
}

func (l *list) Back() *ListItem {
	return l.BackElem
}

func (l *list) PushFront(v interface{}) *ListItem {
	ls := &ListItem{
		Value: v,
	}

	if l.length == 0 {
		ls.Prev = nil
		ls.Next = nil
		l.BackElem = ls
		l.FrontElem = ls
	} else {
		ls.Prev = nil
		l.FrontElem.Prev = ls
		ls.Next = l.Front()
		l.BackElem = l.Back()
		l.FrontElem = ls
	}
	l.length++
	return ls
}

func (l *list) PushBack(v interface{}) *ListItem {
	ls := &ListItem{
		Value: v,
	}

	if l.length == 0 {
		ls.Prev = nil
		ls.Next = nil
		l.BackElem = ls
		l.FrontElem = ls
	} else {
		ls.Prev = l.Back()
		ls.Prev.Next = ls
		l.FrontElem = l.Front()
		ls.Next = nil
		l.BackElem = ls
	}
	l.length++
	return ls
}

func (l *list) Remove(i *ListItem) {
	l.length--
	if l.length == 1 {
		l.FrontElem = nil
		return
	}

	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.FrontElem = i.Next
		l.FrontElem.Prev = nil
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.BackElem = i.Prev
		l.BackElem.Next = nil
	}
}

func (l *list) MoveToFront(i *ListItem) {
	l.Remove(i)
	if l.FrontElem != nil {
		l.FrontElem.Prev = i
		i.Next = l.FrontElem
	}
	l.FrontElem = i
}

func NewList() List {
	d := new(list)
	return d
}
