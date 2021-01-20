# Dialect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Header** | Pointer to **bool** | If true, the results will contain a header row | [optional] [default to true]
**Delimiter** | Pointer to **string** | Separator between cells; the default is , | [optional] [default to ","]
**Annotations** | Pointer to **[]string** | Https://www.w3.org/TR/2015/REC-tabular-data-model-20151217/#columns | [optional] 
**CommentPrefix** | Pointer to **string** | Character prefixed to comment strings | [optional] [default to "#"]
**DateTimeFormat** | Pointer to **string** | Format of timestamps | [optional] [default to "RFC3339"]

## Methods

### NewDialect

`func NewDialect() *Dialect`

NewDialect instantiates a new Dialect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDialectWithDefaults

`func NewDialectWithDefaults() *Dialect`

NewDialectWithDefaults instantiates a new Dialect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeader

`func (o *Dialect) GetHeader() bool`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *Dialect) GetHeaderOk() (*bool, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *Dialect) SetHeader(v bool)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *Dialect) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetDelimiter

`func (o *Dialect) GetDelimiter() string`

GetDelimiter returns the Delimiter field if non-nil, zero value otherwise.

### GetDelimiterOk

`func (o *Dialect) GetDelimiterOk() (*string, bool)`

GetDelimiterOk returns a tuple with the Delimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelimiter

`func (o *Dialect) SetDelimiter(v string)`

SetDelimiter sets Delimiter field to given value.

### HasDelimiter

`func (o *Dialect) HasDelimiter() bool`

HasDelimiter returns a boolean if a field has been set.

### GetAnnotations

`func (o *Dialect) GetAnnotations() []string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *Dialect) GetAnnotationsOk() (*[]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *Dialect) SetAnnotations(v []string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *Dialect) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetCommentPrefix

`func (o *Dialect) GetCommentPrefix() string`

GetCommentPrefix returns the CommentPrefix field if non-nil, zero value otherwise.

### GetCommentPrefixOk

`func (o *Dialect) GetCommentPrefixOk() (*string, bool)`

GetCommentPrefixOk returns a tuple with the CommentPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentPrefix

`func (o *Dialect) SetCommentPrefix(v string)`

SetCommentPrefix sets CommentPrefix field to given value.

### HasCommentPrefix

`func (o *Dialect) HasCommentPrefix() bool`

HasCommentPrefix returns a boolean if a field has been set.

### GetDateTimeFormat

`func (o *Dialect) GetDateTimeFormat() string`

GetDateTimeFormat returns the DateTimeFormat field if non-nil, zero value otherwise.

### GetDateTimeFormatOk

`func (o *Dialect) GetDateTimeFormatOk() (*string, bool)`

GetDateTimeFormatOk returns a tuple with the DateTimeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeFormat

`func (o *Dialect) SetDateTimeFormat(v string)`

SetDateTimeFormat sets DateTimeFormat field to given value.

### HasDateTimeFormat

`func (o *Dialect) HasDateTimeFormat() bool`

HasDateTimeFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


