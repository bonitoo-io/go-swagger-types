# \AuthorizationsApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAuthorizationsID**](AuthorizationsApi.md#DeleteAuthorizationsID) | **Delete** /authorizations/{authID} | Delete a authorization
[**GetAuthorizations**](AuthorizationsApi.md#GetAuthorizations) | **Get** /authorizations | List all authorizations
[**GetAuthorizationsID**](AuthorizationsApi.md#GetAuthorizationsID) | **Get** /authorizations/{authID} | Retrieve an authorization
[**PatchAuthorizationsID**](AuthorizationsApi.md#PatchAuthorizationsID) | **Patch** /authorizations/{authID} | Update an authorization to be active or inactive
[**PostAuthorizations**](AuthorizationsApi.md#PostAuthorizations) | **Post** /authorizations | Create an authorization



## DeleteAuthorizationsID

> DeleteAuthorizationsID(ctx, authID).ZapTraceSpan(zapTraceSpan).Execute()

Delete a authorization

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
    authID := "authID_example" // string | The ID of the authorization to delete.
    zapTraceSpan := "{"trace_id":"1","span_id":"1","baggage":{"key":"value"}}" // string | OpenTracing span context (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthorizationsApi.DeleteAuthorizationsID(context.Background(), authID).ZapTraceSpan(zapTraceSpan).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationsApi.DeleteAuthorizationsID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authID** | **string** | The ID of the authorization to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthorizationsIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **zapTraceSpan** | **string** | OpenTracing span context | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthorizations

> Authorizations GetAuthorizations(ctx).ZapTraceSpan(zapTraceSpan).UserID(userID).User(user).OrgID(orgID).Org(org).Execute()

List all authorizations

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
    userID := "userID_example" // string | Only show authorizations that belong to a user ID. (optional)
    user := "user_example" // string | Only show authorizations that belong to a user name. (optional)
    orgID := "orgID_example" // string | Only show authorizations that belong to an organization ID. (optional)
    org := "org_example" // string | Only show authorizations that belong to a organization name. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthorizationsApi.GetAuthorizations(context.Background()).ZapTraceSpan(zapTraceSpan).UserID(userID).User(user).OrgID(orgID).Org(org).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationsApi.GetAuthorizations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthorizations`: Authorizations
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationsApi.GetAuthorizations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthorizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zapTraceSpan** | **string** | OpenTracing span context | 
 **userID** | **string** | Only show authorizations that belong to a user ID. | 
 **user** | **string** | Only show authorizations that belong to a user name. | 
 **orgID** | **string** | Only show authorizations that belong to an organization ID. | 
 **org** | **string** | Only show authorizations that belong to a organization name. | 

### Return type

[**Authorizations**](Authorizations.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthorizationsID

> Authorization GetAuthorizationsID(ctx, authID).ZapTraceSpan(zapTraceSpan).Execute()

Retrieve an authorization

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
    authID := "authID_example" // string | The ID of the authorization to get.
    zapTraceSpan := "{"trace_id":"1","span_id":"1","baggage":{"key":"value"}}" // string | OpenTracing span context (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthorizationsApi.GetAuthorizationsID(context.Background(), authID).ZapTraceSpan(zapTraceSpan).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationsApi.GetAuthorizationsID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthorizationsID`: Authorization
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationsApi.GetAuthorizationsID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authID** | **string** | The ID of the authorization to get. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthorizationsIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **zapTraceSpan** | **string** | OpenTracing span context | 

### Return type

[**Authorization**](Authorization.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAuthorizationsID

> Authorization PatchAuthorizationsID(ctx, authID).AuthorizationUpdateRequest(authorizationUpdateRequest).ZapTraceSpan(zapTraceSpan).Execute()

Update an authorization to be active or inactive

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
    authID := "authID_example" // string | The ID of the authorization to update.
    authorizationUpdateRequest := *openapiclient.NewAuthorizationUpdateRequest() // AuthorizationUpdateRequest | Authorization to update
    zapTraceSpan := "{"trace_id":"1","span_id":"1","baggage":{"key":"value"}}" // string | OpenTracing span context (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthorizationsApi.PatchAuthorizationsID(context.Background(), authID).AuthorizationUpdateRequest(authorizationUpdateRequest).ZapTraceSpan(zapTraceSpan).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationsApi.PatchAuthorizationsID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchAuthorizationsID`: Authorization
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationsApi.PatchAuthorizationsID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authID** | **string** | The ID of the authorization to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAuthorizationsIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorizationUpdateRequest** | [**AuthorizationUpdateRequest**](AuthorizationUpdateRequest.md) | Authorization to update | 
 **zapTraceSpan** | **string** | OpenTracing span context | 

### Return type

[**Authorization**](Authorization.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAuthorizations

> Authorization PostAuthorizations(ctx).Authorization(authorization).ZapTraceSpan(zapTraceSpan).Execute()

Create an authorization

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
    authorization := *openapiclient.NewAuthorization("OrgID_example", []openapiclient.Permission{*openapiclient.NewPermission("Action_example", *openapiclient.NewResource("Type_example"))}) // Authorization | Authorization to create
    zapTraceSpan := "{"trace_id":"1","span_id":"1","baggage":{"key":"value"}}" // string | OpenTracing span context (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthorizationsApi.PostAuthorizations(context.Background()).Authorization(authorization).ZapTraceSpan(zapTraceSpan).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationsApi.PostAuthorizations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostAuthorizations`: Authorization
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationsApi.PostAuthorizations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthorizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | [**Authorization**](Authorization.md) | Authorization to create | 
 **zapTraceSpan** | **string** | OpenTracing span context | 

### Return type

[**Authorization**](Authorization.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

