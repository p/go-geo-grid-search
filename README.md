# Go Geo Grid Search

[![GoDoc](https://godoc.org/github.com/p/go-geo-grid-search?status.svg)](https://godoc.org/github.com/p/go-geo-grid-search)

ggsearch is a package for performing fast K-closest lookups of places
on Earth.

The distinguishing feature of ggsearch is that it does not require a
bounding box for querying. ggsearch builds a sparse grid index of the places
and performs queries by iterating the grid tiles in a spiral fashion.

## Documentation

Documentation is available on [godoc.org](https://godoc.org/github.com/p/go-geo-grid-search).

Here is an [example](https://github.com/p/go-geo-grid-search/tree/master/examples/simple.go)
showing how to use ggsearch.

## Performance

ggsearch was built for speed. On a t2.micro AWS instance a Web service
running ggsearch queried a data set of 40,000 places in 5 ms/request.

## Caveats

ggsearch is built for querying populated areas and as such querying
near the poles is not going to produce the right results.

## License

Released under the MIT license.
