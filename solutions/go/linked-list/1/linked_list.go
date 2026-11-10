package linkedlist
import "errors"

// Define List and Node types here.
// Note: The tests expect Node type to include an exported field with name Value to pass.
type Node struct{
    Value any
    next *Node
    prev *Node
}

type List struct{
    Head *Node
    Tail *Node
}


func NewList(elements ...any) *List {
    list := &List{}
	for _,value := range elements{
        node := &Node{Value: value}
    	if list.Head == nil {
            list.Head = node
            list.Tail = node
        } else {
            list.Tail.next = node
            node.prev = list.Tail 
            list.Tail = node 
        }
    }
    return list 
}

func (n *Node) Next() *Node {
    return n.next 	
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (l *List) Unshift(v any) {
    node := &Node{Value: v}
    if l.Head != nil {
    	node.next = l.Head
        l.Head.prev = node
        l.Head = node
    } else {
        l.Head = node
        l.Tail = node
    }
}

func (l *List) Push(v any) {
    node := &Node{Value: v}
    if l.Head != nil {
        l.Tail.next = node
        node.prev = l.Tail
        l.Tail = node
    } else {
        l.Head = node
        l.Tail = node
    }
}

var errEmptyList = errors.New("empty list")

func (l *List) Shift() (any, error) {
	if l.Head != nil{
        rmv := l.Head
        if rmv == l.Tail{
            l.Head = nil
            l.Tail = nil
            return rmv.Value, nil
        } else {
            l.Head = l.Head.next
            l.Head.prev = nil
            return rmv.Value, nil
        }
    }
    return nil, errEmptyList
}

func (l *List) Pop() (any, error) {
	if l.Tail != nil{
        rmv := l.Tail
        if rmv == l.Head{
            l.Head = nil
            l.Tail = nil
            return rmv.Value, nil
        } else {
            l.Tail = l.Tail.prev
            l.Tail.next = nil
            return rmv.Value, nil
        }
    }
    return nil, errEmptyList
}

func (l *List) Reverse() {
    current := l.Head
    if l.Head != nil {
        for current != nil{
            next := current.next
            current.next = current.prev
            current.prev = next 
            current = next 
        }
        temp := l.Head
        l.Head = l.Tail
        l.Tail = temp 
    }
}

func (l *List) First() *Node {
	return l.Head
}

func (l *List) Last() *Node {
	return l.Tail
}

func (l *List) Count() int {
	count := 0
    current := l.Head 
    for current != nil {
        count ++
        current = current.next 
    }
    return count 
}

// Delete removes the first node in a list with a given value.
// Returns true if a node was removed.
func (ll *List) Delete(v any) bool {
	current := ll.Head
    for current != nil{
        if current.Value == v {
             if current == ll.Head && current == ll.Tail{
                ll.Head = nil
                ll.Tail = nil
                return true 
            }else if current == ll.Head{
                ll.Head = current.next
                ll.Head.prev = nil
                return true
            }else if current == ll.Tail{
                ll.Tail = current.prev
                ll.Tail.next = nil
                return true 
            }else {
                current.prev.next = current.next
                current.next.prev = current.prev 
                return true
             }
        }
        current = current.next 
    }
    return false
}
