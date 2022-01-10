package encoding

import (
	"testing"

	"github.com/parca-dev/parca/pkg/columnstore/types"
	"github.com/stretchr/testify/require"
)

func TestPlain(t *testing.T) {
	p := NewPlain(types.String, 10)

	count, err := p.Insert(0, types.Value{Data: "test"})
	require.NoError(t, err)
	require.Equal(t, 1, count)
	require.Equal(t, []types.Value{
		{Data: "test"},
		{},
		{},
		{},
		{},
		{},
		{},
		{},
		{},
		{},
	}, p.values)
}

func TestPlainInsertTwo(t *testing.T) {
	p := NewPlain(types.String, 10)

	count, err := p.Insert(0, types.Value{Data: "test1"})
	require.NoError(t, err)
	require.Equal(t, 1, count)

	count, err = p.Insert(1, types.Value{Data: "test3"})
	require.NoError(t, err)
	require.Equal(t, 2, count)

	require.Equal(t, []types.Value{
		{Data: "test1"},
		{Data: "test3"},
		{},
		{},
		{},
		{},
		{},
		{},
		{},
		{},
	}, p.values)
}

func TestPlainInsertMany(t *testing.T) {
	p := NewPlain(types.String, 10)

	count, err := p.Insert(0, types.Value{Data: "test1"})
	require.NoError(t, err)
	require.Equal(t, 1, count)

	count, err = p.Insert(1, types.Value{Data: "test3"})
	require.NoError(t, err)
	require.Equal(t, 2, count)

	count, err = p.Insert(1, types.Value{Data: "test2"})
	require.NoError(t, err)
	require.Equal(t, 3, count)

	require.Equal(t, []types.Value{
		{Data: "test1"},
		{Data: "test2"},
		{Data: "test3"},
		{},
		{},
		{},
		{},
		{},
		{},
		{},
	}, p.values)
}

func TestPlainFind(t *testing.T) {
	p := NewPlain(types.String, 10)

	count, err := p.Insert(0, types.Value{Data: "test1"})
	require.NoError(t, err)
	require.Equal(t, 1, count)

	count, err = p.Insert(1, types.Value{Data: "test3"})
	require.NoError(t, err)
	require.Equal(t, 2, count)

	indexRange, err := p.Find(types.Value{Data: "test2"})
	require.NoError(t, err)
	require.Equal(t, IndexRange{Start: 1, End: 1}, indexRange)
}