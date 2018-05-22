package openshift

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPrintNode(t *testing.T) {
	assert := assert.New(t)
	node := Node{
		Region: "test-region",
		InternalIp: "127.0.0.1",
		Zone: "test-region-1a",
		InternalHostname: "localhost.internal",
	}

	ret := printNode(node)
	assert.NotEmpty(ret)

	expected := "127.0.0.1 openshift_node_labels=\"{'region':'test-region','zone':'test-region-1a'}\" openshift_ip=127.0.0.1 openshift_hostname=localhost.internal"
	assert.Equal(expected, ret)
}

