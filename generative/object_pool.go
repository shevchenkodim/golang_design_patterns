package generative

type Object struct {
	id        int
	someField map[string]interface{}
}

type Pool chan *Object

func NewObjects(total int) *Pool {
	p := make(Pool, total)

	for i := 0; i < total; i++ {
		p <- new(Object)
	}

	return &p
}

// Run Example
/*
func Run() {
	p := NewObjects(2)

	select {
	case obj := <-p:
		// obj.Do(...)
		p <- obj
	default:
		//
		return
	}
}
*/
