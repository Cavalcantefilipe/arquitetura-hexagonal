package handler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandler_JsonError(t *testing.T) {
	msg := "test"
	r := jsonError(msg)
	require.Equal(t, []byte(`{"message":"test"}`), r)
}
