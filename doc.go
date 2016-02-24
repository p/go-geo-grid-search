/*

ggsearch is a package for performing fast K-closest lookups of places
on Earth.

The distinguishing feature of ggsearch is that it does not require a
bounding box for querying. ggsearch builds a sparse grid index of the places
and performs queries by iterating the grid tiles in a spiral fashion.

Performance

ggsearch was built for speed. On a t2.micro AWS instance a Web service
running ggsearch queried a data set of 40,000 places in 5 ms/request.

Caveats

ggsearch is built for querying populated areas and as such querying
near the poles is not going to produce the right results.

License

Released under the MIT license.

*/
package ggsearch
