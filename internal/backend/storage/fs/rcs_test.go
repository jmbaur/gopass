package fs

import (
	"context"
	"testing"

	"github.com/blang/semver"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRCS(t *testing.T) {
	ctx := context.Background()
	path, cleanup := newTempDir(t)
	defer cleanup()

	g := New(path)
	assert.NoError(t, g.Add(ctx, "foo", "bar"))
	assert.NoError(t, g.Commit(ctx, "foobar"))
	assert.NoError(t, g.Push(ctx, "foo", "bar"))
	assert.NoError(t, g.Pull(ctx, "foo", "bar"))
	assert.NoError(t, g.Cmd(ctx, "foo", "bar"))
	assert.NoError(t, g.Init(ctx, "foo", "bar"))
	assert.NoError(t, g.InitConfig(ctx, "foo", "bar"))
	assert.Equal(t, g.Version(ctx), semver.Version{Minor: 1})
	assert.Equal(t, "fs", g.Name())
	assert.NoError(t, g.AddRemote(ctx, "foo", "bar"))
	revs, err := g.Revisions(ctx, "foo")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(revs))
	body, err := g.GetRevision(ctx, "foo", "bar")
	require.NoError(t, err)
	assert.Equal(t, "foo\nbar", string(body))
	assert.NoError(t, g.RemoveRemote(ctx, "foo"))
}
