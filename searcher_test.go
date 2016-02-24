package ggsearch

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchTrivial(t *testing.T) {
	searcher := NewSearcher(10, 20)

	locatable := SampleLocatable{"test1", 0, 0}
	searcher.AddLocatable(locatable)

	results := searcher.Search(nil, 0, 0, 10)
	assert.Equal(t, len(results), 1)
	assert.Equal(t, results[0].Locatable().(SampleLocatable), locatable)
}

func TestSearchTwoEntries(t *testing.T) {
	searcher := NewSearcher(10, 20)

	locatable1 := SampleLocatable{"test1", 0, 0}
	searcher.AddLocatable(locatable1)
	locatable2 := SampleLocatable{"test2", 0, 1}
	searcher.AddLocatable(locatable2)

	results := searcher.Search(nil, 0, 0, 10)
	assert.Equal(t, len(results), 2)
	assert.Equal(t, results[0].Locatable().(SampleLocatable), locatable1)
	assert.Equal(t, results[1].Locatable().(SampleLocatable), locatable2)
}
