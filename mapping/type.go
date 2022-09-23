package mapping

type esType string

// Keyword
const (
	AggregateMetricDouble esType = "aggregate_metric_double"
	Alias                 esType = "alias"
	Binary                esType = "binary" // Base64-encoded string
	Boolean               esType = "boolean"
	Completion            esType = "completion"
	Date                  esType = "date"
	DateNanoseconds       esType = "date_nanos"
	DenseVector           esType = "dense_vector"
	Flattened             esType = "flattened"
	Geopoint              esType = "geo_point"
	Geoshape              esType = "geo_shape"
	Histogram             esType = "histogram"
	IP                    esType = "ip"
	Join                  esType = "join"
	Nested                esType = "nested"
	Object                esType = "object" // Default value. nil is object.
	Percolator            esType = "percolator"
	Point                 esType = "point"
	Range                 esType = "Range"
	RankFeature           esType = "rank_feature"
	RankFeatures          esType = "rank_features"
	SearchAsYouType       esType = "search_as_you_type"
	Shape                 esType = "shape"
	TokenCount            esType = "token_count"
	Version               esType = "version"
)

// Text field types
const (
	Keyword         esType = "keyword"
	ConstantKeyword esType = "constant_keyword"
	Wildcard        esType = "wildcard"
	Text            esType = "text"
)

// Numerics field types
// see https://www.elastic.co/guide/en/elasticsearch/reference/8.4/number.html
const (
	Long         esType = "long"
	Integer      esType = "integer"
	Short        esType = "short"
	Byte         esType = "byte"
	Double       esType = "double"
	Float        esType = "float"
	HalfFloat    esType = "half_float"
	ScaledFloat  esType = "scaled_float"
	UnsignedLong esType = "unsigned_long"
)

// Range field types
// see https://www.elastic.co/guide/en/elasticsearch/reference/8.4/range.html
const (
	IntegerRange esType = "integer_range"
	FloatRange   esType = "float_range"
	LongRange    esType = "long_range"
	DoubleRange  esType = "double_range"
	DateRange    esType = "date_range"
	IpRange      esType = "ip_range"
)
