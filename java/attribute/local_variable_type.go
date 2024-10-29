package attribute

func NewLocalVariableType(startPc int, length int, name string, signature string, index int) *LocalVariableType {
	return &LocalVariableType{startPc, length, name, signature, index}
}

type LocalVariableType struct {
	startPc   int
	length    int
	name      string
	signature string
	index     int
}

func (l LocalVariableType) StartPc() int {
	return l.startPc
}

func (l LocalVariableType) Length() int {
	return l.length
}

func (l LocalVariableType) Name() string {
	return l.name
}

func (l LocalVariableType) Signature() string {
	return l.signature
}

func (l LocalVariableType) Index() int {
	return l.index
}
