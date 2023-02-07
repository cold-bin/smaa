// @author cold bin
// @date 2023/2/7

package num

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

// GetUniqueId 生成唯一的编号: 19位
func GetUniqueId() string {
	return fmt.Sprintf("%019d", node.Generate().Int64())
}

func init() {
	var err error
	node, err = snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}

}
