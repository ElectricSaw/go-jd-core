package util

import "unicode/utf8"

func NewStringTokenizer(str string) *StringTokenizer {
	return NewStringTokenizer3(str, " \t\n\r\f", false)
}

func NewStringTokenizer2(str, delim string) *StringTokenizer {
	return NewStringTokenizer3(str, delim, false)
}

func NewStringTokenizer3(str string, delimiters string, retDelims bool) *StringTokenizer {
	tokenizer := &StringTokenizer{
		str:           str,
		delimiters:    delimiters,
		retDelims:     retDelims,
		currentPos:    0,
		maxPos:        len(str),
		delimitersSet: make(map[rune]bool),
	}

	for _, d := range delimiters {
		tokenizer.delimitersSet[d] = true
	}

	return tokenizer
}

// StringTokenizer 구조체 정의
type StringTokenizer struct {
	str           string
	delimiters    string
	retDelims     bool
	currentPos    int
	maxPos        int
	delimitersSet map[rune]bool
}

// skipDelimiters: 현재 위치에서 구분자를 건너뛰고 첫 번째 토큰 위치 반환
func (st *StringTokenizer) skipDelimiters(pos int) int {
	for pos < st.maxPos {
		r, size := utf8.DecodeRuneInString(st.str[pos:])
		if !st.delimitersSet[r] {
			break
		}
		pos += size
	}
	return pos
}

// scanToken: 다음 구분자까지의 위치 반환
func (st *StringTokenizer) scanToken(pos int) int {
	for pos < st.maxPos {
		r, size := utf8.DecodeRuneInString(st.str[pos:])
		if st.delimitersSet[r] {
			break
		}
		pos += size
	}
	return pos
}

// HasMoreTokens: 더 많은 토큰이 있는지 확인
func (st *StringTokenizer) HasMoreTokens() bool {
	pos := st.skipDelimiters(st.currentPos)
	return pos < st.maxPos
}

// NextToken: 다음 토큰 반환
func (st *StringTokenizer) NextToken() (string, error) {
	if st.currentPos >= st.maxPos {
		return "", &NoSuchElementException{"No more tokens available"}
	}

	if !st.retDelims {
		st.currentPos = st.skipDelimiters(st.currentPos)
	}

	start := st.currentPos
	if st.retDelims && st.delimitersSet[rune(st.str[st.currentPos])] {
		st.currentPos++
		return st.str[start:st.currentPos], nil
	}

	st.currentPos = st.scanToken(st.currentPos)
	return st.str[start:st.currentPos], nil
}

// CountTokens: 남은 토큰 수 계산
func (st *StringTokenizer) CountTokens() int {
	count := 0
	pos := st.currentPos
	for pos < st.maxPos {
		pos = st.skipDelimiters(pos)
		if pos >= st.maxPos {
			break
		}
		pos = st.scanToken(pos)
		count++
	}
	return count
}

// NoSuchElementException: 예외 정의
type NoSuchElementException struct {
	msg string
}

func (e *NoSuchElementException) Error() string {
	return e.msg
}

// 예제 사용 코드
func main() {
	tokenizer := NewStringTokenizer3("this is a test", " ", false)

	for tokenizer.HasMoreTokens() {
		token, err := tokenizer.NextToken()
		if err != nil {
			panic(err)
		}
		println(token)
	}

	println("Remaining tokens:", tokenizer.CountTokens())
}
