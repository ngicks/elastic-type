package mapping

type EsType string

const (
	AggregateMetricDouble EsType = "aggregate_metric_double"
	Alias                 EsType = "alias"
	Binary                EsType = "binary" // Base64-encoded string
	Boolean               EsType = "boolean"
	Completion            EsType = "completion"
	Date                  EsType = "date"
	DateNanoseconds       EsType = "date_nanos"
	DenseVector           EsType = "dense_vector"
	Flattened             EsType = "flattened"
	Geopoint              EsType = "geo_point"
	Geoshape              EsType = "geo_shape"
	Histogram             EsType = "histogram"
	IP                    EsType = "ip"
	Join                  EsType = "join"
	Nested                EsType = "nested"
	Object                EsType = "object" // Default value. nil is object.
	Percolator            EsType = "percolator"
	Point                 EsType = "point"
	RankFeature           EsType = "rank_feature"
	RankFeatures          EsType = "rank_features"
	SearchAsYouType       EsType = "search_as_you_type"
	Shape                 EsType = "shape"
	TokenCount            EsType = "token_count"
	Version               EsType = "version"
)

// Text field types
const (
	Keyword         EsType = "keyword"
	ConstantKeyword EsType = "constant_keyword"
	Wildcard        EsType = "wildcard"
	Text            EsType = "text"
)

// Numerics field types
// see https://www.elastic.co/guide/en/elasticsearch/reference/8.4/number.html
const (
	Long         EsType = "long"
	Integer      EsType = "integer"
	Short        EsType = "short"
	Byte         EsType = "byte"
	Double       EsType = "double"
	Float        EsType = "float"
	HalfFloat    EsType = "half_float"
	ScaledFloat  EsType = "scaled_float"
	UnsignedLong EsType = "unsigned_long"
)

// Range field types
// see https://www.elastic.co/guide/en/elasticsearch/reference/8.4/range.html
const (
	IntegerRange EsType = "integer_range"
	FloatRange   EsType = "float_range"
	LongRange    EsType = "long_range"
	DoubleRange  EsType = "double_range"
	DateRange    EsType = "date_range"
	IpRange      EsType = "ip_range"
)
