package myint

type MyInt int32

func (m MyInt) Divide(i int) MyInt {
	return m / MyInt(i)
}

func (m MyInt) Multiply(i int) MyInt {
	return m * MyInt(i)
}

func (m MyInt) Add(i int) MyInt {
	return m + MyInt(i)
}

func (m MyInt) Sub(i int) MyInt {
	return m - MyInt(i)
}
