// @author cold bin
// @date 2023/2/5

package _req

type GetCode struct {
	Phone string `json:"phone" binding:"required;len=11"`
}

type Register struct {
	Phone    string `json:"phone" binding:"required;len=11"`
	Password string `json:"password" binding:"required"`
	Code     string `json:"code" binding:"required;len=6"`
	NickName string `json:"nick_name" binding:"required"`
}

type LoginByPassword struct {
	Phone    string `json:"phone" binding:"required;len=11"`
	Password string `json:"password" binding:"required"`
}

type LoginByCode struct {
	Phone string `json:"phone" binding:"required;len=11"`
	Code  string `json:"code" binding:"required;len=6"`
}

type RegretPassword struct {
	Phone    string `json:"phone" binding:"required;len=11"`
	Password string `json:"password" binding:"required"`
	Code     string `json:"code" binding:"required;len=6"`
}
