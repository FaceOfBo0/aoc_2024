package utils

import "fmt"

type SqrMatrix interface {
	Inv() (SqrMatrix, error)
	Det() float64
	Dim() int
	Data() []float64
}

type SqrMat2 struct {
	Dims    int
	matData []float64
}

type Vector struct {
	Size   int
	vecDat []float64
}

func NewSqrMatrix(dim int, data []float64) SqrMatrix {
	return SqrMat2{
		Dims:    2,
		matData: data,
	}
}

func NewVector(size int, data []float64) Vector {
	return Vector{
		Size:   size,
		vecDat: data,
	}
}

func SolveLinEq(matA SqrMatrix, vecB Vector) (Vector, error) {
	invA, err := matA.Inv()
	if err != nil {
		return Vector{}, err
	}
	return MatVecMult(invA, vecB)
}

func MatVecMult(matA SqrMatrix, vecB Vector) (Vector, error) {
	if matA.Dim() != vecB.Size {
		return Vector{}, fmt.Errorf("Dimension mismatch between matA and vecB!")
	}

	newData := make([]float64, vecB.Size)
	for i := range matA.Dim() {
		for j := range matA.Dim() {
			newData[i] += matA.Data()[i*matA.Dim()+j] * vecB.vecDat[j]
		}
	}

	return NewVector(vecB.Size, newData), nil
}

func (sm SqrMat2) Inv() (SqrMatrix, error) {
	if sm.Det() == 0 {
		return SqrMat2{}, fmt.Errorf("There is no inverse for that matrix! (Det(M) == 0)")
	}
	newData := make([]float64, len(sm.matData))
	newData[0], newData[1], newData[2], newData[3] =
		sm.matData[3], -1*sm.matData[1], -1*sm.matData[2], sm.matData[0]
	for i := range newData {
		newData[i] = newData[i] / sm.Det()
	}
	return NewSqrMatrix(sm.Dims, newData), nil

}

func (sm SqrMat2) Dim() int {
	return sm.Dims
}

func (sm SqrMat2) Data() []float64 {
	return sm.matData
}

func (sm SqrMat2) Det() float64 {
	if sm.Dims == 2 {
		return sm.matData[0]*sm.matData[3] - sm.matData[1]*sm.matData[2]
	}
	return 0
}
