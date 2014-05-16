package stretch

import (
	"bytes"
	"fmt"
	"strings"
)

func (c *Cluster) GetHotThreads(nodes ...string) (string, error) {
	var nodestr string
	var buf bytes.Buffer

	if len(nodes) > 0 {
		nodestr = strings.Join(nodes, ",")
	} else {
		nodestr = "_all"
	}

	err := c.Client.Get(&buf, fmt.Sprintf("/_nodes/%s/hot_threads", nodestr))

	return buf.String(), err
}
