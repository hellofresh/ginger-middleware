package statsd

const (
	DefaultResponseTimeEnabled = true
	DefaultThroughputEnabled   = true
	DefaultStatusCodeEnabled   = true
	DefaultSuccessEnabled      = true
	DefaultErrorEnabled        = true

	DefaultResponseTimeBucket = "request.response_time"
	DefaultThroughputBucket   = "request.throughput"
	DefaultStatusCodeBucket   = "request.status_code"
	DefaultSuccessBucket      = "request.success"
	DefaultErrorBucket        = "request.error.default"
)

// Options represents the configuration of statsd
type Options struct {
	Client Client
	ResponseTimeEnabled,
	ThroughputEnabled,
	StatusCodeEnabled,
	SuccessEnabled,
	ErrorEnabled bool

	// Bucket names for each metric.
	ResponseTimeBucket,
	ThroughputBucket,
	StatusCodeBucket,
	SuccessBucket,
	ErrorBucket string
}

func newConfiguredClient(client Client) *Options {
	return &Options{
		Client: client,

		ResponseTimeEnabled: DefaultResponseTimeEnabled,
		ThroughputEnabled:   DefaultThroughputEnabled,
		StatusCodeEnabled:   DefaultStatusCodeEnabled,
		SuccessEnabled:      DefaultSuccessEnabled,
		ErrorEnabled:        DefaultErrorEnabled,

		ResponseTimeBucket: DefaultResponseTimeBucket,
		ThroughputBucket:   DefaultThroughputBucket,
		StatusCodeBucket:   DefaultStatusCodeBucket,
		SuccessBucket:      DefaultSuccessBucket,
		ErrorBucket:        DefaultErrorBucket,
	}
}



func (c *Options) IncrError(errors []*gin.Error, handler string) {
	if c.ErrorEnabled {
		for _, err := range errors {
			// If the gin.Error.Meta implements the Bucket interface,
			// increment its specific bucket by its given increment amount.
			// Otherwise, increment the default error bucket with the default amount.
			b, ok := err.Meta.(Bucket)
			if !ok {
				c.Client.Incr(join(handler, c.ErrorBucket), 1)
				continue
			}
			c.Client.Incr(join(handler, b.BucketName()), 1)
		}
	}
}

func (c *Options) Timing(start time.Time, handler string) {
	if c.ResponseTimeEnabled {
		c.Client.Timing(join(handler, c.ResponseTimeBucket),
			// Convert to milliseconds.
			time.Now().Sub(start).Nanoseconds()/time.Millisecond.Nanoseconds())
	}
}

func (c *Options) IncrThroughput(handler string) {
	if c.ThroughputEnabled {
		c.Client.Incr(join(handler, c.ThroughputBucket), 1)
	}
}

// IncrStatusCode increments the context's response status code bucket.
func (c *Options) IncrStatusCode(status int, handler string) {
	if c.StatusCodeEnabled {
		c.Client.Incr(join(handler, c.StatusCodeBucket, strconv.Itoa(status)), 1)
	}
}

// IncrSuccess increments the success bucket
// if no errors were attached to the context.
func (c *Options) IncrSuccess(errors []*gin.Error, handler string) {
	if c.SuccessEnabled && len(errors) == 0 {
		c.Client.Incr(join(handler, c.SuccessBucket), 1)
	}
}
