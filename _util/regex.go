// @author cold bin
// @date 2023/2/7

package _util

import "regexp"

func JudgePhone(phone string) (bool, error) {
	c, err := regexp.Compile("^((13[0-9])|(14[5|7])|(15([0-3]|[5-9]))|(18[0,5-9]))\\\\d{8}$")
	if err != nil {
		return false, err
	}
	return c.MatchString(phone), nil
}
