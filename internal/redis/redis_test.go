package redis

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	success := map[string]struct {
		key   string
		wants string
	}{
		"成功": {
			key:   "hoge",
			wants: "hoge",
		},
	}

	failed := map[string]struct {
		key string
		err error
	}{
		"値が存在しない": {
			key: "",
			err: nil,
		},
	}

	redisCli := New()
	ctx := context.Background()
	t.Run("成功", func(t *testing.T) {
		for tn, tt := range success {
			t.Run(tn, func(t *testing.T) {
				got, err := redisCli.Get(ctx, tt.key)
				assert.NoError(t, err)
				assert.Equal(t, got, tt.wants)
			})
		}
	})

	t.Run("失敗", func(t *testing.T) {
		for tn, tt := range failed {
			t.Run(tn, func(t *testing.T) {
				_, err := redisCli.Get(ctx, tt.key)
				assert.EqualError(t, err, "")
			})
		}
	})
}
