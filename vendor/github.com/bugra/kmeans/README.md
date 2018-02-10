![Testing Image](http://img.shields.io/travis/bugra/kmeans.svg?style=flat)
![Issues](http://img.shields.io/github/issues/bugra/kmeans.svg?style=flat)  

K-Means
===
kmeans is BSD licensed fast k-means implementation(with better initialization, 
known as k-means++) written in Go. It supports various distance functions out
of the box for convenience and experimentation. It has large coverage for tests.
The algorithm is tested on Iris dataset and distance functions have full test coverage.

## Documentation
[Godoc](https://godoc.org/github.com/bugra/kmeans)


## License
[BSD License](https://github.com/bugra/kmeans/blob/master/LICENSE)

### Distance Functions
It supports various distance functions:

- LP norms (manhattan, euclidean distances including)
- SquaredEuclideanDistance
- Minkowski Distance
- Weighted Minkowski Distance
- Chebyshev Distance
- Hamming Distance
- Bray Curtis Distance
- CanberraDistance
