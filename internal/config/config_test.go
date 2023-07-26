package config

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadCfg(t *testing.T) {
	conf, err := LoadCfg("")
	assert.NoError(t, err)
	b, _ := json.Marshal(conf)
	t.Log(string(b))
}
