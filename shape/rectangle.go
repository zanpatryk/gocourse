package shape

type Rectangle struct {
	width  float32
	height float32
}

func (r Rectangle) Area() float32 {
	return r.width * r.height
}
