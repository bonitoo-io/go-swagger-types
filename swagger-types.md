# Generated code analysis
This document describes possible ways how to get server API types and request/response handling code for Go client.

## Swagger generators

There two Open API generators for Go language available:
 - https://github.com/deepmap/oapi-codegen
 - https://github.com/OpenAPITools/openapi-generator
 
### OAPI Codegen
The _oapi-codegen_  was originally used in the version 0 of Go client and it is still used on the current version. 
It generates plain schema types and clear REST call handling code  with a small amount of boilerplate code. It also offers simple Go templates to customize code, 
which is mainly required to reuse the HTTP function of the current Go client. 

It's written in Go. 

As it missed some features, such as (typed) enums generation, and had some flaws (e.g. time serialization to request params), 
so I created [3 pull requests](https://github.com/deepmap/oapi-codegen/pulls?q=+is%3Apr+author%3Avlastahajek+) to that repository.
For now, using this generator requires checkout from my fork: https://github.com/bonitoo-io/oapi-codegen and branch _dev-master_, which has all 3 PRs merged. 

It can generate code from the InfluxDB server swagger without any problem. The generated code is produced that each code type (schema types, client, server) is placed to a single 
file, which leads to source files with thousands lines. It generates inner complex types as anonymous structs.

### OpenAPI Generator
The _openapi-generator_ is used for other InfluxDB clients, e.g. Java, C#, Python. It is written in Java and it supports many backends, including Go. 
It by default produces a complex structure of files allowing easily creating of new git repo.  
Generated sources are quite Java-ish. Domain types have setters and getters, generated REST calls look cumbersome.  
Generated code is organized that each schema type  is placed in the separated file, and a file for code handling communication for each server endpoint.

This generator has problems with _oneOf_ compound schema type, so it cannot generate code from  the original InfluxDB server swagger without workarounds.
It supports features that were necessary to add to OAPI Codegen by PRs, such as Enums, or generated code handles unexpected HTTP errors by default. Templates are also supported, 
in the _Mustache_ format. 
It generates inner complex types as named types.

Summary:
| Property | OAPI Codegen | OpenAPI Generator |
| :---    |  :---:       |   :---:           |
| Templates |   :+1: (Go) | :+1: (Mustache) |
| Typed enums | Added by PR | :+1: |
| oneOf | :+1: | :-1: |
| # generated files | 2 | 19 |
| inner complex type | anonymous | typed |
| external dependence of generated code | yes | no |

# generated files is as generated from testing swagger.yml

### Generated Schema types
Both generators produce schema types as structs with public fields where optional fields are pointers.  
OpenAPI Generator also generates setters and getters, which eases using pointers, but it is not Go idiomatic.

**Authorization by OAPI Codegen**
```go
// Authorization defines model for Authorization.
type Authorization struct {
	// Embedded struct due to allOf(#/components/schemas/AuthorizationUpdateRequest)
	AuthorizationUpdateRequest
	// Embedded fields due to inline allOf schema
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Id        *string    `json:"id,omitempty"`
	Links     *struct {

		// URI of resource.
		Self *Link `json:"self,omitempty"`

		// URI of resource.
		User *Link `json:"user,omitempty"`
	} `json:"links,omitempty"`

	// Name of the org token is scoped to.
	Org *string `json:"org,omitempty"`

	// ID of org that authorization is scoped to.
	OrgID *string `json:"orgID,omitempty"`

	// List of permissions for an auth.  An auth must have at least one Permission.
	Permissions *[]Permission `json:"permissions,omitempty"`

	// Passed via the Authorization Header and Token Authentication type.
	Token     *string    `json:"token,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	// Name of user that created and owns the token.
	User *string `json:"user,omitempty"`

	// ID of user that created and owns the token.
	UserID *string `json:"userID,omitempty"`
}
```
Note that pointer in Permissions declaration. In swagger this is required type, but because of [JSON serialization issue in Go](https://github.com/golang/go/issues/27589), 
it must be marked as _omitempty_ and thus by default pointer. Although, empty arrays/slices have clear meaning this doesn't need to be a pointer.

**Authorization by OpenAPI Generator**
```go
// Authorization struct for Authorization
type Authorization struct {
	// If inactive the token is inactive and requests using the token will be rejected.
	Status *string `json:"status,omitempty"`
	// A description of the token.
	Description *string `json:"description,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// ID of org that authorization is scoped to.
	OrgID string `json:"orgID"`
	// List of permissions for an auth.  An auth must have at least one Permission.
	Permissions []Permission `json:"permissions"`
	Id *string `json:"id,omitempty"`
	// Passed via the Authorization Header and Token Authentication type.
	Token *string `json:"token,omitempty"`
	// ID of user that created and owns the token.
	UserID *string `json:"userID,omitempty"`
	// Name of user that created and owns the token.
	User *string `json:"user,omitempty"`
	// Name of the org token is scoped to.
	Org *string `json:"org,omitempty"`
	Links *AuthorizationAllOfLinks `json:"links,omitempty"`
}

// AuthorizationAllOfLinks struct for AuthorizationAllOfLinks
type AuthorizationAllOfLinks struct {
	// URI of resource.
	Self *string `json:"self,omitempty"`
	// URI of resource.
	User *string `json:"user,omitempty"`
}
```

Main entities, e.g. Bucket, Authorization, Organization, etc,  has basic fields, such as ID (uint64), Name, Description as simple types with meaning what empty/zero value means no valid data, and thus id 
doesn't need a pointer. 

  On the other hand, schema types for updating those entities require distinguishing for missing and empty value and thus have to have pointers.
 
 ## Schema types in the server repo
 InfluxDB server defines main entities in the root _influxdb_ package. Basic fields are non-pointers:
 
 ```go
 // Authorization is an authorization. ??
type Authorization struct {
	ID          ID           `json:"id"`
	Token       string       `json:"token"`
	Status      Status       `json:"status"`
	Description string       `json:"description"`
	OrgID       ID           `json:"orgID"`
	UserID      ID           `json:"userID,omitempty"`
	Permissions []Permission `json:"permissions"`
	CRUDLog
}
 ```
 
 However, status result types, such as results of `/health/` or `/ready` are not defined in the InfluxDB repo.  
 
 ## Conclusion
 
 To have user-friendly structures for server API entities we have 3 basic options:
 1. reuse types from InfluxDB server package and modify generated client code in template to use those entities. 
    - It would require to manually write Go type for `HealthCheck` and `Ready`schema types.
      - It might be be added to the server repo
    - Entities doesn't have properties docs
      - It can be fixed in the server repo
    - It would create circular dependency if the client would be used by influx tool
      - influx tool in separate repo?
 2. Fix a generator to produces non-pointers for non-update related schema types
    - ~2MD work for OAPI Codegen
    - It will be non generalizable hack - it would require custom fork forever
 3. Provide entities by client (copy and fix generated code or from server) and modify generated client code in template to use those entities.
    - It would require to check updates of server API for changes, however API in GA should be stable
 
 For a long term solution option [1] looks best. If not possible, then [3] seems to be a viable solution.

