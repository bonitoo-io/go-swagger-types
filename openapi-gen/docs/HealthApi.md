# \HealthApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHealth**](HealthApi.md#GetHealth) | **Get** /health | Get the health of an instance



## GetHealth

> HealthCheck GetHealth(ctx).ZapTraceSpan(zapTraceSpan).Execute()

Get the health of an instance

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    zapTraceSpan := "{"trace_id":"1","span_id":"1","baggage":{"key":"value"}}" // string | OpenTracing span context (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HealthApi.GetHealth(context.Background()).ZapTraceSpan(zapTraceSpan).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `HealthApi.GetHealth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHealth`: HealthCheck
    fmt.Fprintf(os.Stdout, "Response from `HealthApi.GetHealth`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHealthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zapTraceSpan** | **string** | OpenTracing span context | 

### Return type

[**HealthCheck**](HealthCheck.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

