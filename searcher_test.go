package ggsearch

import (
	"github.com/stretchr/testify/assert" 
	"testing"
)

func TestSearch(t *testing.T) {
	searcher := NewSearcher(10, 20)
	
	locatable := SampleLocatable{"test1", 0, 0}
	searcher.AddLocatable(locatable)
	
	results := searcher.Search(nil, 0, 0, 10)
	assert.Equal(t, len(results), 1)
	assert.Equal(t, results[0].GetLocatable().(SampleLocatable), locatable)
}
