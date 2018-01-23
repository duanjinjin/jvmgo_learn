package rtda



// jvm stack
type Stack struct {
	maxSize uint
	size    uint
	_top    *Frame // stack is implemented as linked list
}


func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

// 后进先出，所以加入的全部挂在从头部插入
func (self *Stack) push(frame *Frame)  {
	//Thread限制Stack数量，最终通过Frame数量来限制
	if self.size >= self.maxSize{
		panic("java.lang.StackOverflowError")
	}

	if self._top != nil{
		frame.lower = self._top
	}

	self._top = frame
	self.size++
}

func (self *Stack) pop() *Frame  {
	if self._top == nil{
		panic("jvm stack is empty!")
	}

	top := self._top
	self._top = top.lower
	top.lower = nil// let GO gc
	self.size--
	return top
}

func (self *Stack) top() *Frame  {
	if self._top == nil{
		panic("jvm stack is empty!")
	}

	return self._top
}