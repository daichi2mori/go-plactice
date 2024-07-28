package foo

//変数名の最初の文字が、大文字だとpublicで他のファイルから参照可能
//小文字だとprivateになり参照不可
const (
	Max = 100
	min = 1
)

func ReturnMin() int {
	return min
}

//関数名も同じく、最初の文字が小文字だと他のファイルから参照不可
func appName() string {
	return "Golang"
}
