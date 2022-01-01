package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("purge logic", func(t *testing.T) {
		c := NewCache(5)
		c.Set("aaa", 100)
		c.Clear()
		val, wasInCache := c.Get("aaa")
		require.False(t, wasInCache)
		require.Nil(t, val)
	})

	t.Run("purge length logic", func(t *testing.T) {
		c := NewCache(3)

		wasInCache := c.Set("fsab", 400)
		require.False(t, wasInCache)

		wasInCache = c.Set("ccasdas", 300)
		require.False(t, wasInCache)

		wasInCache = c.Set("dgfhdf", 400)
		require.False(t, wasInCache)

		val, wasInCache := c.Get("fsab")
		require.True(t, wasInCache)
		require.Equal(t, 400, val)
	})

	t.Run("purge cache logic", func(t *testing.T) {
		c := NewCache(3)

		wasInCache := c.Set("bbb", 200) // {bbb, aaa}
		require.False(t, wasInCache)

		wasInCache = c.Set("ccc", 300) // {ccc, bbb, aaa}
		require.False(t, wasInCache)

		val, wasInCache := c.Get("bbb") // {bbb, ccc, aaa}
		require.True(t, wasInCache)
		require.Equal(t, 200, val)

		val, wasInCache = c.Get("aaa")
		require.False(t, wasInCache)
		require.Nil(t, val)

		val, wasInCache = c.Get("ccc")
		require.True(t, wasInCache)
		require.Equal(t, 300, val)
	})
}

func TestCacheMultithreading(t *testing.T) {
	t.Skip() // Remove me if task with asterisk completed.

	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
