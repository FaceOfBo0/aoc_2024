package utils

type SqrMatrix struct {
	Dims int
	Data []int
}

type Vector struct {
	Size int
	Data []int
}

func NewSqrMatrix(dim int, data []int) *SqrMatrix {
	return &SqrMatrix{
		Dims: dim,
		Data: data,
	}
}

func NewVector(size int, data []int) *Vector {
	return &Vector{
		Size: size,
		Data: data,
	}
}

func (m *SqrMatrix) Det() int {
	if m.Dims == 2 {
		return m.Data[0]*m.Data[3] - m.Data[1]*m.Data[2]
	}
	return 0
}
