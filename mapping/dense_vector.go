package mapping

// https://www.elastic.co/guide/en/elasticsearch/reference/8.4/dense-vector.html#dense-vector-params
type DenseVectorParams struct {
	// Type is type of this property. Automatically filled if zero.
	Type esType `json:"type,omitempty"`
	Dims int    `json:"dims"`
	// If true, you can search this field using the kNN search API.
	// Defaults to false.
	Index *bool `json:"index,omitempty"`
	// Required if Index is true.
	Similarity   *denseVectorSimilarity  `json:"similarity,omitempty"`
	IndexOptions *DenseVectorIndexOption `json:"index_options,omitempty"`
}

func (p *DenseVectorParams) FillType() {
	p.Type = DenseVector
}

type denseVectorSimilarity string

const (
	L2Norm     denseVectorSimilarity = "l2_norm"
	DotProduct denseVectorSimilarity = "dot_product"
	Cosine     denseVectorSimilarity = "cosine"
)

type DenseVectorIndexOption struct {
	// Currently only hnsw is supported.
	Type string `json:"type"`
	// Defaults to 16.
	M int `json:"m"`
	// Defaults to 100.
	EfConstruction int `json:"ef_construction"`
}
