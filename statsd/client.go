package statsd

type Client interface {
	Incr(stat string, count int64) error
	Timing(stat string, delta int64) error
}

type Bucket interface {
	BucketName() string // Bucket name to increment.
}
