package oop

type Human struct {
	high   float32
	weight float32
}

func NewHuman(h float32, w float32) *Human {
	return &Human{
		high:   h,
		weight: w,
	}
}

func (h *Human) GetH() float32 {
	return h.high
}

func (h *Human) SetH(hh float32) {
	h.high = hh
}
