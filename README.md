# elastic-type

Type generator that generates Go types from [Elasticsearch](https://www.elastic.co/guide/en/elastic-stack/index.html) mapping.

Also helper types.

## packages

### es_type

Helper types for Elasticsearch type.

Elasticsearch is somehow _elastic_ for its json value.

For example:

- The boolean type can accept true/false and "true"/"false"/""(which means false).
- The date type defaults to a format of `strict_date_optional_time||epoch_millis`, which means it can accept number that represent unix milli second, or string value that is formatted as `YYYY-MM-dd'T'HH:mm:ss.S[S...(up to 9 digits)]` or `YYYY-MM-dd`.
  - This is really helpful when you define a script or an ingest pipeline, as you can store simply a long value for date.
    - Namely, `new Date().getTime()`
- The geopoint type allows 6 different notations to store.
  - The doc says it is `historical reasons`.

#### MVPs

- [x] Marshalling/Unmarshalling helper (Field[T any])
- [x] binary
- [x] boolean
- [x] date for built-in es date formats
- [ ] histogram
- [x] geopoint
- [x] geoshape
- [ ] join
- [ ] ranges
- [ ] rank_feature/rank_features
- [ ] point
  - basically same as geopoint, but fewer supported data notations.
- [ ] shape
  - basically same as geoshape.
- [ ] version

### generate

Code generator. It generates go code from an es mapping.

#### MVPs

- [x] Generate high level / raw types
  - The high level type is a plain go struct type, which is similar to what you define every day. That probably can not be unmarshalled from / marshalled into a json to be stored in the Elasticsearch.
  - The raw type is strictly compliant to Elasticsearch json format. All fields can be `undefined`, `null`, T or an array of T.
    - Which is achieved in help of `estype.Field[T]`
- [x] Generate raw level types marshaller code.
- [x] Generate high level / raw conversion code.
- [ ] Test using a real Elasticsearch instance.

#### Optionally we would do

- [ ] Remove overlapping type definition.
  - If two type definitions are exactly same, then use same type.
  - Evaluating that 2 defs are semantically same is hard without ast parser. Maybe we should do it in post-process.

### mapping

Type definitions for Elasticsearch mappings.

Used to parse mapping.

#### MVPs

- [x] Cover all mappings.
