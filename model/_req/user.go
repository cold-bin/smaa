// @author cold bin
// @date 2023/2/5

package _req

type GetCode struct {
	Phone string `json:"phone" binding:"len=11"`
}
