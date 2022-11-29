package gohierarchyclient

import (
	"context"
	"fmt"
	"os/exec"
	"testing"

	rest "github.com/SKF/go-rest-utility/client"
	"github.com/SKF/go-rest-utility/client/auth"
	"github.com/SKF/go-utility/v2/uuid"
	"github.com/stretchr/testify/require"
)

func Test_client_getAncestors(t *testing.T) {
	t.SkipNow()

	stage := "test"
	c := NewClient(
		WithStage(stage),
		WithClientID("<client id>"),
		rest.WithTokenProvider(auth.RawToken(getToken(t, stage))),
	)

	nodeID := uuid.UUID("835f196b-2482-45e8-a303-390cf4e22775")
	ctx := context.Background()

	res, err := c.GetAncestorsIncludeSelf(ctx, nodeID, 0)
	require.NoError(t, err)

	for _, n := range res {
		fmt.Printf("%s: %s, %s\n", n.ID, *n.Label, *n.Type)
	}
}

func getToken(t *testing.T, stage string) string {
	tokenBytes, err := exec.Command("gob", "tokens", "-pis", stage).Output()
	require.NoError(t, err)

	return string(tokenBytes)
}
