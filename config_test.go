package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewConfigEnvLoad(t *testing.T) {
	conf, err := NewConfig(defaultConfFileName)
	require.NoError(t, err)
	require.NotEmpty(t, conf.Jira.Username)
	require.NotEmpty(t, conf.Jira.URL)
	require.NotEmpty(t, conf.Jira.APIToken)

}
