package matrix

type MatrixDimensionsError struct {}

func (err *MatrixDimensionsError) Error() string {
  return "Incorrect dimensions"
}

func (err *MatrixDimensionsError) String() string {
  return err.Error()
}

type MatrixIndexError struct {}

func (err *MatrixIndexError) Error() string {
  return "Index out of range"
}

func (err *MatrixIndexError) String() string {
  return err.Error()
}
