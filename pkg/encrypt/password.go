package encrypt

import (
	"strconv"
	"time"
)

//创建加密后的密码及加密因子，用于注册用户
func CreatePassword(password string) (hash string, encrypt string) {
	encrypt = createEncrypt()
	hash = GetPassword(password, encrypt)
	return
}

//验证密码是否正确
func VerifyPassword(password, encrypt, correctHash string) bool {
	return correctHash == GetPassword(password, encrypt)
}

//根据随机因子生成密码
func GetPassword(password, encrypt string) string {
	return Md5(Md5(password) + encrypt)
}

//创建随机因子
func createEncrypt() string {
	unixNano := time.Now().UnixNano()
	str := strconv.FormatInt(unixNano, 10)
	return Md5(str)[0:10]
}
