package stretch

import (
	"bytes"
)

func (c *Cluster) GetHotThreads() (responseBody string) {
	var buf bytes.Buffer
	c.Client.Get(&buf, "/_nodes/_all/hot_threads")

	return buf.String()
}
