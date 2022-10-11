# elastic-type

Type generator that generates Go types from [Elasticsearch](https://www.elastic.co/guide/en/elastic-stack/index.html) mapping.

Also helper types.

## packages

### mapping

Type definitions for Elasticsearch mappings.

Used to parse mapping.

### type

Helper types for Elasticsearch type.

Elasticsaerch is somehow _elastic_ for its json value.

For example:

- boolean value can accept true/false and "true"/"false".
- Date format defaults to `strict_date_optional_time||epoch_millis`, which means it can accept number that represent unix milli second, or string value that formatted as `YYYY-MM-dd'T'HH:mm:ss.S[S...(up to 9 digits)]` or `YYYY-MM-dd`.
