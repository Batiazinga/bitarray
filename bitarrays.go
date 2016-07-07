package bitarray

// New2D returns a two-dimensional array of single bit booleans with size nRows * nColumns.
// Initially, all values are false.
func New2D(nRows, nColumns int) Array2D {
	size := nRows * nColumns
	if size%8 == 0 {
		size = size / 8
	} else {
		size = size/8 + 1
	}
	return Array2D{
		nRows:    nRows,
		nColumns: nColumns,
		content:  make([]uint8, size),
	}
}

// Array2D is a two-dimensional array of single bit booleans.
type Array2D struct {
	nRows, nColumns int
	content         []uint8
}

// indexes finds the uint and the uint's bit corresponding to cell (i,j).
func (a Array2D) indexes(i, j int) (index, bit int) {
	absolute := i*a.nColumns + j
	return absolute / 8, absolute % 8
}

// NumRows returns the number of rows in the array.
func (a Array2D) NumRows() int {
	return a.nRows
}

// NumColumns returns the number of columns in the array.
func (a Array2D) NumColumns() int {
	return a.nColumns
}

// Get returns the (boolean) value stored in cell (i,j).
// It panics ("index out of range") if dimensions are invalid.
func (a Array2D) Get(i, j int) bool {
	// checks
	if i >= a.nRows || j >= a.nColumns {
		panic("bitarrays: index out of range")
	}

	index, bit := a.indexes(i, j)
	return a.content[index]&(1<<uint(bit)) != 0
}

// Set sets the value of cell (i,j).
// It panics ("index out of range") if dimensions are invalid.
func (a Array2D) Set(i, j int, val bool) {
	// checks
	if i >= a.nRows || j >= a.nColumns {
		panic("bitarrays: index out of range")
	}

	index, bit := a.indexes(i, j)
	if val {
		a.content[index] = a.content[index] | (1 << uint(bit))
	} else {
		a.content[index] = a.content[index] & ^(1 << uint(bit))
	}
}

func (a Array2D) String() string {
	s := ""
	for i := 0; i != a.NumRows(); i++ {
		for j := 0; j != a.NumColumns(); j++ {
			if a.Get(i, j) {
				s += "1"
			} else {
				s += "0"
			}
			if j != a.NumColumns()-1 {
				s += " "
			}
		}
		if i != a.NumRows()-1 {
			s += "\n"
		}
	}

	return s
}
