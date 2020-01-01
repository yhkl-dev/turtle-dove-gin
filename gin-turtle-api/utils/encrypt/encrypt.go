package encrypt

import (
	"crypto/md5"
	"fmt"
	"io"
)

// StringToMd5 encrypt string
func StringToMd5(str string) string {
	newMd5String := md5.New()
	io.WriteString(newMd5String, str)

	fmt.Println(newMd5String.Sum(nil))
	return fmt.Sprintf("%x", newMd5String.Sum(nil))
}

// CheckMd5Value return true or false
func CheckMd5Value(inputString, md5String string) bool {
	md5Object := md5.New()

	io.WriteString(md5Object, inputString)

	originMd5 := fmt.Sprintf("%x", md5Object.Sum(nil))
	if originMd5 == md5String {
		return true
	}

	return false
}
