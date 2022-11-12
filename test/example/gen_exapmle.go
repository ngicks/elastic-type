package example

//go:generate go run ../../cmd/generate/main.go -prefix-with-index-name -i ./object.json -out-high ./object_high.go -out-raw ./object_raw.go -map-option ./object_map_option.json
//go:generate go run ../../cmd/generate/main.go -prefix-with-index-name -i ./object_w_overlap.json -out-high ./object_w_overlap_high.go -out-raw ./object_w_overlap_raw.go
