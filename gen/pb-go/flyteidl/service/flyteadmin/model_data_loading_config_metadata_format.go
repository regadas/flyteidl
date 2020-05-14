/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin
// DataLoadingConfigMetadataFormat : - JSON: JSON / YAML for the metadata (which contains inlined primitive values). The representation is inline with the standard json specification as specified - https://www.json.org/json-en.html  - PROTO: Proto is a serialized binary of `core.LiteralMap` defined in flyteidl/core
type DataLoadingConfigMetadataFormat string

// List of DataLoadingConfigMetadataFormat
const (
	DataLoadingConfigMetadataFormatJSON DataLoadingConfigMetadataFormat = "JSON"
	DataLoadingConfigMetadataFormatYAML DataLoadingConfigMetadataFormat = "YAML"
	DataLoadingConfigMetadataFormatPROTO DataLoadingConfigMetadataFormat = "PROTO"
)