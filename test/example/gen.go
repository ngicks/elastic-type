package example

//go:generate go run ../../cmd/generate-es-type/main.go -prefix-with-index-name -i ./object.json -out-high ./object_high.go -out-raw ./object_raw.go -out-test ./object_test.go -map-option ./object_map_option.json
//go:generate go run ../../cmd/generate-es-type/main.go -prefix-with-index-name -i ./object_w_overlap.json -out-high ./object_w_overlap_high.go -out-raw ./object_w_overlap_raw.go -out-test ./object_w_overlap_test.go
//go:generate go run ../../cmd/generate-es-type/main.go -prefix-with-index-name -i ./all.json -out-high ./all_high.go -out-raw ./all_raw.go -out-test ./all_test.go -global-option ./all_global_option.json
//go:generate go run ../../cmd/generate-es-type/main.go -prefix-with-index-name -i ./object_dynamic_inheritance.json -out-high ./object_dynamic_inheritance_high.go -out-raw ./object_dynamic_inheritance_raw.go -out-test ./object_dynamic_inheritance_test.go
//go:generate go run ../../cmd/generate-es-type/main.go -prefix-with-index-name -i ./example.json -out-high ./example_high.go -out-raw ./example_raw.go -out-test ./example_test.go -global-option ./example_global_option.json -map-option ./example_map_option.json
