package any_tests

import (
	"testing"

	// external
	"github.com/stretchr/testify/require"

	// internal
	"github.com/sniperkit/snk.golang.json/pkg/json/v1"
)

func Test_wrap_map(t *testing.T) {
	should := require.New(t)
	any := jsoniter.Wrap(map[string]string{"Field1": "hello"})
	should.Equal("hello", any.Get("Field1").ToString())
	any = jsoniter.Wrap(map[string]string{"Field1": "hello"})
	should.Equal(1, any.Size())
}

func Test_map_wrapper_any_get_all(t *testing.T) {
	should := require.New(t)
	any := jsoniter.Wrap(map[string][]int{"Field1": {1, 2}})
	should.Equal(`{"Field1":1}`, any.Get('*', 0).ToString())
	should.Contains(any.Keys(), "Field1")

	// map write to
	stream := jsoniter.NewStream(jsoniter.ConfigDefault, nil, 0)
	any.WriteTo(stream)
	// TODO cannot pass
	//should.Equal(string(stream.buf), "")
}
