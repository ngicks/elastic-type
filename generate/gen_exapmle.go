package generate

//go:generate go run ../cmd/generate/main.go -prefix-with-index-name -i ./example/object.json -out-high ./example/object_high.go -out-raw ./example/object_raw.go -map-option ./example/object_map_option.json
//go:generate go run ../cmd/generate/main.go -prefix-with-index-name -i ./example/object_w_overlap.json -out-high ./example/object_w_overlap_high.go -out-raw ./example/object_w_overlap_raw.go
