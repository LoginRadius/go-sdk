# LoginRadius Go SDK

![Home Image](http://docs.lrcontent.com/resources/github/banner-1544x500.png)

## Introduction 

LoginRadius Customer Registration wrapper provides access to LoginRadius Identity Management Platform API.

LoginRadius is an Identity Management Platform that simplifies user registration while securing data. LoginRadius Platform simplifies and secures your user registration process, increases conversion with Social Login that combines 30 major social platforms, and offers a full solution with Traditional Customer Registration. You can gather a wealth of user profile data from Social Login or Traditional Customer Registration. 

LoginRadius centralizes it all in one place, making it easy to manage and access. Easily integrate LoginRadius with all of your third-party applications, like MailChimp, Google Analytics, Livefyre and many more, making it easy to utilize the data you are capturing.

LoginRadius helps businesses boost user engagement on their web/mobile platform, manage online identities, utilize social media for marketing, capture accurate consumer data, and get unique social insight into their customer base.

Please visit [here](http://www.loginradius.com/) for more information.


## Contents

* [Demo](https://github.com/LoginRadius/go-sdk/tree/master/demo) - contains a demo of LoginRadius user auth workflow.


## Documentation

* [Configuration](https://docs.loginradius.comhttps://www.loginradius.com/docs/api/v2/deployment/sdk-libraries/golang-library) - Everything you need to begin using the LoginRadius SDK.


## Installation

To install, run:
`go get github.com/loginradius/go-sdk`

Import the package:

`import "github.com/loginradius/go-sdk"`

Install all package dependencies by running `go get ./...` in the root folder of this SDK.  


## Usage

Take a peek:

Before making any API calls, the LoginRadius API client must be initialized with your Loginradius API key and API secret, This information can be found in your LoginRadius account as described [here](https://www.loginradius.com/docshttps://www.loginradius.com/docs/api/v2/admin-console/platform-security/api-key-and-secret/#api-key-and-secret)
Sample code:

```go
cfg := lr.Config{
    ApiKey:    <your API key>,
    ApiSecret: <your API secret>,
}

lrclient, err := lr.NewLoginradius(&cfg)

if err != nil {
    // handle error
}
```

Many API calls included in this SDK must be completed with an access token, which can be obtained after calling the Authentication Login API and reading the token from the response or from generating an access token through the UID in the Accounts API.

For APIs that require the user's credentials to function properly, the access token must be passed in the `Authorization: Bearer` header; this is handled by the SDK. For APIs that require the user's access token, initialize the LoginRadius client like so:

```go
lrclient, err := lr.NewLoginradius(&cfg, map[string]string{"token": <access token>})

if err != nil {
    // handle error
}
```

Alternatively an already-initalized client can be reused like so:

```go
lrclient.Context.Token = <access token>
```

Please be aware of the dangers of using global variables to store individual user's access token if you choose to reuse an already-initalized client.

## Calling an API provided by the LoginRadius Golang SDK

### Calling an API

API calls are separated into separate packages. Each package contains a struct holding the LoginRadius client object as an embedded struct, e.g.:

```go
package mfa

import (
	lr "github.com/LoginRadius/go-sdk"
)

// Embed Loginradius struct
type Loginradius struct {
	Client *lr.Loginradius
}
```

This allows for the individual API calls to be defined as methods of the Loginradius API client struct.

Require the package containing the API to be called, and type assert the initialized LoginRadius client into the specific package's client struct when calling the API:

```go
res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).PostMFAEmailLogin(login)

```

### Passing Body Parameters

The SDK includes the package `lrbody`, which contains structs for various API endpoints. This package is provided for convenience only, and does not contain every single struct needed to fulfill API requirements. It is useful for endpoints requiring key values to be submitted as nested objects. Alternatively, anonymous structs could be used as well.

Additionally, all API calls included in this SDK requiring body parameters expect `interface{}` as input, allowing for structs or `[]byte({})` as the body parameter.

Passing an anonymous struct initialized with value as body:

```go
body := struct{
  Username string
}{"newname"}

res, err := lraccount.Loginradius(lraccount.Loginradius{lrclient}).PutManageAccountUpdate(
    uid,
    body,
  )
```

Passing `[]byte({})` as body:

```go
res, err := lraccount.Loginradius(lraccount.Loginradius{lrclient}).PutManageAccountUpdate(
  uid,
  []byte(`{"Username":"newname"}`),
)
```

Both of these would work.

For more information on this package and the structs it contains: https://godoc.org/github.com/LoginRadius/go-sdk/lrbody

### Passing query parameters

Some APIs mandate the submission of requests with certain query parameters, for these endpoints the SDK methods expect a `map[string]string` containing the key value pairs for query parameters. Some APIs can take optional query parameters but have no required query parameters, these can be called with or without passing queries.

Passing query parameter:

```go
response, err := tokenmanagement.Loginradius(tokenmanagement.Loginradius{lrclient}).GetRefreshToken(
    map[string]string{"expiresin": expiresIn}, // this is the query parameter
)
```

Alternatively, calling this end point without passing the optional query parameter:

```go
response, err := tokenmanagement.Loginradius(tokenmanagement.Loginradius{lrclient}).GetRefreshToken()
```

## Handling the Response Returned by an API Method

All APIs included in the LoginRadius Golang SDK return `httprutils.Response` and `Error`. For additional information about the `httprutils.Response` struct, please see [Handling the response](#Handling-the-response).

This SDK includes a package called `lrjson`, which contains the method `DynamicUnmarshall`. To access the fields in `response.Body`, call the method like so:

```go
res, err = mfa.Loginradius(mfa.Loginradius{lrclient}).PostMFAUsernameLogin(
  map[string]string{"username": username, "password": password},
  map[string]string{"emailtemplate": "hello"},
)

if err != nil {
  // handle error
}

session, err = lrjson.DynamicUnmarshal(res.Body)
if err != nil {
  // handle error
}

token, ok := session["access_token"].(string); ok{
  // use the token
}
```

Some end points return an array rather than an object, here is a code snippet for handling the returned body from these endpoints:

```go
resp, err := lrsocial.Loginradius(lrsocial.Loginradius{lrclient}).GetSocialAlbum()
data := []interface{}{}
err = json.Unmarshal([]byte(resp.Body), &data)
if err != nil {
  // handle error
}

firstAlbum, ok := data[0].(map[string]interface{}); ok {
  id := firstAlbum["ID"].(string); ok = {
    // use the id
  }
}
```

Some endpoints return nested objects. Here is a code snippet for retrieving data from a nested array:

```go
resp, err := customobject.Loginradius(customobject.Loginradius{lrclient}).GetCustomObjectByToken(map[string]string{"objectname": <object})
if err != nil {
  // handle error
}
obj, err := lrjson.DynamicUnmarshal(resp.Body)
if err != nil {
  // handle error
}
data, ok := obj["data"].([]interface{}); ok {
  // Only get the ID of the first item of the slice
  id, ok := data[0].(map[string]interface{})["Id"].(string); ok {
    // do something with the ID
  }
}
```

A detailed list of LoginRadius datapoints and their types can be found [here.](https://www.loginradius.com/docs/api/v2/getting-started/data-points/detailed-data-points)

An alternative to dynamically unmarshalling the body is to create a struct based on the expected return, and unmarshal the body into the struct; this would be marginally faster, but will not be able to handle all future changes in the API returns. The returned `httprutils.Response` struct contains two fields: string `Body` and []byte `OrigBody` to allow more flexibility in your implementation of the SDK.

## Making Requests and Handling Responses with the httprutils Package

The Loginradius Golang SDK comes with the package `httprutils`(named for HTTP Response/request Utils, as well as to avoid namesquatting the useful package name `httputils`), which contains structs and functions for making REST requests.

This includes utilities for holding and building the response and request struct, as well as for making the request.

For more information on the package `httprutils` please see https://godoc.org/github.com/LoginRadius/go-sdk/httprutils.

### Encoding Body

The package offers a function for encoding body parameters for PUT, POST, and DELETE requests.

Sample code:

```go
encodedBody, error := httprutils.EncodeBody(body)
if error != nil {
  // handle error
}
```

### Building request using the httprutils package

Sample code:

```go
request := &httprutils.Request{
  Method: httprutils.Post,
  URL:    <path>,
  Headers: map[string]string{
    "content-Type":  "application/x-www-form-urlencoded",
    "Authorization": "Bearer <access token>"
  },
  QueryParams: map[string]string{
    "apiKey": <api key>,
  },
  Body: <encodedBody>,
}
```

### Making request using the httprutils package

Sample code:

```go
res, err := httprutils.TimeoutClient.Send(*request)
if err!= nil {
  // handle error
}
```

### Handling the response

The response returned from the previous code snippet will be a struct like so
(defined in httprutils package):

```go
type Response struct {
  StatusCode int
  Body       string
  Headers    map[string][]string
  OrigBody   []byte
}
```

The response body, status code and the headers can be accessed through this struct.

For complete documentation on this package, please refer to https://godoc.org/github.com/LoginRadius/go-sdk/httprutils

## Error package

The LoginRadius Golang SDK also includes an error package, `lrerror`. This package provides the API error interface accessors for the SDK.

Example:

```go
output, err := loginradius.GetAuthVerifyEmail()
if err != nil {
    if lrError, ok := err.(lrError.Error); ok {
        // Get error details
        log.Println("Error:", lrError.Code(), lrError.Message())
        // Prints out full error message, including original error if there was one.
        log.Println("Error:", lrError.Error())
        // Get original error
        if origErr := lrError.OrigErr(); origErr != nil {
            // operate on original error.
        }
    } else {
        ...
    }
}
```

For complete documentation on this package, please refer to https://godoc.org/github.com/LoginRadius/go-sdk

## SOTT Generation

SOTT is a secure one-time token that can be created using the API key, API secret, and a timestamp ( start time and end time ). You can manually create a SOTT using the following util function.

```go

// You can pass the start and end time interval and the SOTT will be valid for this time duration. 

startTime:="2021-01-10 07:10:42"  // Valid Start Date with Date and time

endTime:="2023-01-15 07:20:42" // Valid End Date with Date and time
								
//do not pass the time difference if you are passing startTime & endTime.						
								
timeDifference:="20" // (Optional) The time difference will be used to set the expiration time of SOTT, If you do not pass time difference then the default expiration time of SOTT is 10 minutes.

apiKey:="" //LoginRadius Api Key.
apiSecret:="" //LoginRadius Api Secret (Only Primary Api Secret is used to generate the SOTT manually).
		

sott := sott.Generate(apiKey, apiSecret,timeDifference,startTime,endTime)						
				
```


## Tests

All APIs in the LoginRadius Go SDK are covered by tests; these tests are placed in the lrtest directory, divided into the `lrintegrationtest` and `lrunittest` packages.

The integration tests in the `lrintegrationtest` package must be run with an internet connection as they make calls to the LoginRadius cloud. To run the tests, create /lrtest/lrintegrationtest/config/secret.env based on the content provided by secret.env.example, then run `go test`.

Some tests are skipped by default, uncomment `t.SkipNow()` and manually set the required variable values to run these tests.

## Static Analysis

To analyze the code, [golangci-lint](https://github.com/golangci/golangci-lint) was used. `gometalinter` contains a number of static analysis tools for Go code, including a tool to analyze security issues within the code. Follow the instructions on Github to install and run.

## Demo

A demo project utilizing the SDK is included in the LoginRadius Go SDK. You can find the demo here: https://github.com/LoginRadius/go-sdk/tree/master/demo.
This is a simple project containing a backend server written in Go that exposes endpoints serving JSON returns to a Javascript and HTML frontend. The demo project contains basic functionalities for demonstration purposes; these functionalities include user registration, authentication, forgot-password, multifactor authentication workflows, roles, custom object. A single page application called "Login Screen" is also included in the project for reference purposes.

To configure and run the server, follow these steps:

1. Follow the Installation and Quick Start guide to set up the SDK.
2. Create /demo/config/secret.env based on the example provided by /demo/config/secret.env.example
3. Configure /demo/config.public.env with the base url of your server if needed; the default is localhost:3000
4. Create /demo/ui/assets/js/options.js based on the example provided by /demo/ui/assets/js/options.sample.js
5. Configure /demo/ui/assets/js/loginScreen.js to handle proper link redirection when using the LoginScreen application.
6. Run the server by running `go run main.go` from /demo/cmd/app

Inside the demo project, a playground has also been added for you to try out API calls. You can add your code in /demo/cmd/playground/main.go and run it with `go run main.go` to explore the API calls during development.

## APIs

Please note that that before APIs can be called, the LoginRadius client struct must be initalized with your API key and secret. Please jump to [Initializing the LoginRadius Client](#Intializing-the-LoginRadius-Client) for instructions on how to do so. View [Calling an API](#Calling-an-API) to see general instructions regarding how methods should be invoked. Read on for details about individual API methods.

### Authentication APIs

The Authentication (Auth) APIs allow changes to the account once some form of authentication has been performed. For this reason, they are considered to be user facing client-side/front-end API calls.

To call an Authentication API, import the authentication package like so:

```go
import (
 "github.com/LoginRadius/go-sdk/api/role"
)
```

**List of APIs in this Section:**

- [POST: Auth Add Email](#auth-add-email)
- [POST: Auth Forgot Password](#forgot-password)
- [POST: Auth User Registration by Email](#auth-user-registration-by-email)
- [POST: Auth Login by Email](#auth-login-by-email)
- [POST: Auth Login by Username](#auth-login-by-username)
- [GET: Auth Email Availability](#auth-email-availability)
- [GET: Auth Username Availability](#auth-username-availability)
- [GET: Auth Read Profiles By Token](#auth-read-profiles-by-token)
- [GET: Auth Privacy Policy Accept](#auth-privacy-policy-accept)
- [GET: Auth Send Welcome Email](#auth-send-welcome-email)
- [GET: Auth Social Identity](#auth-social-identity)
- [GET: Auth Validate Access Token](#auth-validate-access-token)
- [GET: Auth Verify Email](#auth-verify-email)
- [GET: Auth Delete Account](#auth-delete-account)
- [GET: Auth Invalidate Access Token](#auth-invalidate-access-token)
- [GET: Security Questions By Token](#security-questions-by-token)
- [GET: Security Questions By Email](#security-questions-by-email)
- [GET: Security Questions By User Name](#security-questions-by-user-name)
- [GET: Security Questions By Phone](#security-questions-by-phone)
- [PUT: Auth Verify Email By OTP](#auth-verify-email-by-otp)
- [PUT: Auth Change Password](#auth-change-password)
- [PUT: Auth Link Social Identities](#auth-link-social-identities)
- [PUT: Auth Resend Email Verification](#auth-resend-email-verification)
- [PUT: Auth Reset Password By Reset Token](#auth-reset-password-by-reset-token)
- [PUT: Auth Reset Password By Email](#auth-reset-password-by-email)
- [PUT: Auth Reset Password By Phone](#auth-reset-password-by-phone)
- [PUT: Auth Reset Password By Username](#auth-reset-password-by-username)
- [PUT: Auth Set or Change Username](#auth-set-or-change-username)
- [PUT: Auth Update Profile By Token](#auth-update-profile-by-token)
- [PUT: Auth Update Security Question By Access Token](#auth-update-security-question-by-access-token)
- [DELETE: Auth Delete Account with Email Confirmation](#auth-delete-account-with-email-confirmation)
- [DELETE: Auth Remove Email](#auth-remove-email)
- [DELETE: Auth Unlink Social Identities](#auth-unlink-social-identities)

##### Auth Add Email

This API is used to add emails to an existing account.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-add-email)

Example:

```go
// use an anonymous struct, alternatively []byte could be passed
body := struct {
  Email string
  Type string
}{
  "example@example.com",
  "Primary", //This can be a value of your designation
}

res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}) .
  PostAuthAddEmail(
    body,
  )

if err != nil {
    // handle error
}
```

##### Auth Forgot Password

This API is used to initiate the Forgot Password workflow.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-forgot-password)

Example:

```go
// use an anonymous struct, alternatively []byte could be passed
email := struct {
  Email string
}{
  "example@example.com,
}

res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).
    PostAuthForgotPassword(
      email,
      map[string]string{"resetpasswordurl": "example.com/password/reset"},
    )
if err != nil {
    // handle error
}
```

##### Auth User Registration by Email

This API is used to register a user with the authentication API.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-user-registration-by-email)

Example:

```go
// Use struct provided by lrbody package to construct body
// alternatively you could initialize your own struct
// []byte could also be passed in lieu of a struct
user := lrbody.RegistrationUser{
    Email: []lrbody.AuthEmail{
      lrbody.AuthEmail{
        Type:  "Primary", //This can be any value of your designation
        Value: "example@example.com",
      },
    },
    Password: "password",
}
sott:=""//(Required) Pass SOTT
res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).
    PostAuthUserRegistrationByEmail(sott,user,)

if err != nil {
    // handle error
}
```

##### Auth Login by Email

This API is used to create an access token for the account used, along with fetching profile data.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-login-by-email)

Example:

```go
// use an anonymous struct, alternatively []byte could be passed in lieu of struct
body := struct {
  Emailstring
  Password string
}{
  userName,
  email, // uses generated email as password
}

res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient})    .PostAuthLoginByEmail(body)

if err != nil {
    // handle error
}
```

##### Auth Login by Username

This API is used to create an access token for the account used, along with fetching profile data. Uses a username instead of email.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-login-by-username)

Example:

```go
// use an anonymous struct, alternatively []byte could be passed in lieu of struct
body := struct {
  Username string
  Password string
}{
  userName,
  email, // uses generated email as password
}

res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).PostAuthLoginByUsername(body)

if err != nil {
  // handle error
}
```

##### Auth Email Availability

This API is used to check used to check whether an email exists or not on your site.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-email-availability)

Example:

```go
res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{loginradius}).
  GetAuthCheckEmailAvailability(map[string]string{"email": "example@example.com"})

if err != nil {
  // handle error
}
```

##### Auth Username Availability

This API is used to check used to check whether a username exists or not on your site.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-username-availability)

Example:

```go
res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{loginradius}).
  GetAuthCheckUsernameAvailability(map[string]string{"username": "exampleusername"})
if err != nil {
  // handle error
}
```

##### Auth Read Profile By Token

This API is used to get the profile data of a user with an access token.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-read-profiles-by-token)

Please note this API requires the access token to be passed in the `Authorization Bearer` header. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

Example:

```go
response, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).
  GetAuthReadProfilesByToken()

if err != nil {
  // handle error
}
```

##### Auth Privacy Policy Accept

This API is used to update the privacy policy of a user using their access token.

Please note this API requires the access token to be passed in the `Authorization Bearer` header. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-privacy-policy-accept)

Example:

```go
res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).
  GetAuthPrivatePolicyAccept()

if err != nil {
  //handle error
}

```

##### Auth Send Welcome Email

This API is used to send a welcome email to the user associated with the access token.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-send-welcome-email)

Example:

```go
res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).GetAuthSendWelcomeEmail(map[string]string{"welcomeemailtemplate": "customer-email-template"})

if err != nil {
  //handle error
}

```

##### Auth Social Identity

This API is called before the account linking API to prevent a second profile from being created.

Please note this API requires the access token to be passed in the `Authorization Bearer` header. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-social-identity)

Example:

```go
//optional queries could also be passed
res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).
  GetAuthSendWelcomeEmail()

if err != nil {
  //handle error
}
```

##### Auth Validate Access Token

This API is used to validate the access token passed in.

Please note this API requires the access token to be passed in the `Authorization Bearer` header. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-validate-access-token)

Example:

```go
res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).
  GetAuthValidateAccessToken()

if err != nil {
  //handle error
}
```

##### Auth Verify Email

This API is used to verify the account using the verification token sent in the verification email.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-verify-email)

Example:

```go
res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).GetAuthVerifyEmail(
  map[string]string{"verificationtoken": <verificationToken>},
)

if err != nil {
  //handle error
}
```

##### Auth Delete Account

This API is used to delete an account using a delete token sent to the user's email.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-delete-account)

Example:

```go
res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).GetAuthDeleteAccount(
  map[string]string{"deletetoken": "<delete token>"}
)

if err != nil {
  // handle error
}
```

##### Auth Invalidate Access Token

This API is used to invalidate the access token passed in.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-invalidate-access-token)

Example:

```go
res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).GetAuthInvalidateAccessToken()

if err != nil {
  // handle error
}

```

##### Security Questions By Access Token

This API is used to retrieve the list of security questions configured for a user using the access token.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/security-questions-by-access-token)

For additional information on security questions, please refer the [documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/security-questions-by-access-token/)

Please note this API requires the access token to be passed in the `Authorization Bearer` header. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

Example:

```go
res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).GetAuthSecurityQuestionByAccessToken()

if err != nil {
  // handle error
}
```

##### Security Questions By Email

This API is used to retrieve the list of security questions configured for a customer using their email.

[Documentation](https://docs.loginradius.comhttps://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/security-questions-by-email)

Example:

```go
res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).GetAuthSecurityQuestionByEmail(
  map[string]string{"email": <email>},
)

if err != nil {
  // handle error
}
```

##### Security Questions By User Name

This API is used to retrieve the list of security questions configured for a user using their username.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/security-questions-by-user-name)

Example:

```go
res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).GetAuthSecurityQuestionByUsername(
  map[string]string{"username": <username>},
)
if err!= nil {
  // handle error
}
```

##### Security Questions By Phone

This API is used to retrieve the list of security questions configured for a user using their phone number.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/security-questions-by-phone)

Example:

```go
res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).GetAuthSecurityQuestionByPhone(
  map[string]string{"phone": <phone>},
)

if err != nil {
  // handle error
}
```

##### Auth Verify Email By OTP

This API is used to verify an account with an OTP. OTP workflow must be enabled in the customer account.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-verify-email-by-otp)

Example:

```go
res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).PutAuthVerifyEmailByOtp(
  map[string]string{
    "email": <email>,
    "otp": <otp>,
  },
)

if err != nil {
  // handle error
}
```

##### Auth Change Password

This API is used to change the user's password.

Please note this API requires the access token to be passed in the `Authorization Bearer` header. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-change-password)

Example:

```go
res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).PutAuthChangePassword(
  map[string]string{
    "oldpassword": <oldpassword>,
    "newpassword": <newpassword>,
  }
)

if err != nil {
  // handle error
}
```

##### Auth Link Social Identities

This API is used to link the user's account with a social account.

Please note this API requires the access token to be passed in the `Authorization Bearer` header. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-link-social-identities)

Example:

```go
res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).PutAuthLinkSocialIdentities(
  map[string]string{
    "candidatetoken": <candidatetoken>,
  }
)

if err != nil {
  // handle error
}
```

##### Auth Resend Email Verification

This API is used to resend the email verification to the user's email.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-resend-email-verification)

Example:

```go
res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).PutResendEmailVerification(
  map[string]string{"email": <email>},
)

if err != nil {
  //handle error
}
```

##### Auth Reset Password By Reset Token

This API is used to reset the password for a user using a reset token received from the user's email.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-reset-password-by-reset-token)

Example:

```go
response, err = lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).PutAuthResetPasswordByResetToken(
  map[string]string{
    "resettoken":<resettoken>,
    "password":<password>,
    // add optional body parameters as needed
  }
)

if err != nil {
  // handle error
}
```

##### Auth Reset Password By OTP

This API is used to reset the password for a user using an OTP received from the user's email. OTP workflow must be enabled for the customer for them to receive an OTP by email.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-reset-password-by-otp)

Example:

```go
response, err = lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).PutAuthResetPasswordByResetToken(
    map[string]string{
    "resettoken":<resettoken>,
    "password":<password>,
    "otp":<otp>,
    // add optional body parameters as needed
  }
)

if err != nil {
  // handle error
}
```

##### Auth Reset Password By Email

This API is used to initiate reset password using a security answer and the user's email.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-reset-password-by-email)

For additional information on security questions, please refer the [documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-reset-password-by-email/)

Example:

```go
// Initialize struct to use for body
// Alternatively, the helper struct lrbody.ResetPwSecurityQuestionEmail from the lrbody package
// could be used, or []byte could be passed as body
type ResetPasswordBySecurityAnswerAndEmailStruct struct {
    SecurityAnswer SecurityQandA `json:"securityanswer"`
    Email string `json:"email"`
    Password string `json:"password"`
}

type SecurityQandA struct {
    SecurityQuestion string `json:<id of security question>`
    // Use your secret question IDs in the json field
    // this would be a random string
    // for more information on this see the security question documentation
}

securityQuestion := SecurityQuestion{"Answer"}
body := ResetPasswordBySecurityAnswerAndEmailStruct{
  securityQuestion,
  <email>,
  <password>
}

response, err = lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).PutAuthResetPasswordBySecurityAnswerAndEmail(
  body,
)

if err != nil {
  // handle error
}
```

##### Auth Reset Password By Phone

This API is used to initiate reset password using a security answer and the user's phone number.

For additional information on security questions, please refer the [documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-reset-password-by-phone/)

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-reset-password-by-phone)

Example:

```go
// Initialize struct to use for body
// Alternatively, the helper struct lrbody.ResetPwSecurityQuestionUsername from the lrbody package
// could be used, or []byte could be passed as body
type ResetPasswordBySecurityAnswerAndUsernameStruct struct {
    SecurityAnswer SecurityQandA `json:"securityanswer"`
    Username string `json:"username"`
    Password string `json:"password"`
}

type SecurityQandA struct {
    SecurityQuestion string `json:<id of security question>`
    // Use your secret question IDs in the json field
    // this would be a random string
    // for more information on this see the security question documentation
}

securityQuestion := SecurityQuestion{"Answer"}
body := ResetPasswordBySecurityAnswerAndUsernameStruct{
  securityQuestion,
  <username>,
  <password>
}

response, err = lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).PutAuthResetPasswordBySecurityAnswerAndUsername(
  body,
)
if err != nil {
  // handle error
}
```

##### Auth Reset Password By Username

This API is used to initiate reset password using a security answer and the username.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-reset-password-by-username)

For additional information on security questions, please refer the [documentation](https://www.loginradius.com/docs/api/v2/admin-console/platform-security/password-policy/#password-policy)

Example:

```go
// Initialize struct to use for body
// Alternatively, the helper struct lrbody.ResetPwSecurityQuestionPhone from the lrbody package
// could be used, or []byte could be passed as body
type ResetPasswordBySecurityAnswerAndPhoneStruct struct {
    SecurityAnswer SecurityQandA `json:"securityanswer"`
    Phone string `json:"phone"`
    Password string `json:"password"`
}

type SecurityQandA struct {
    SecurityQuestion string `json:<id of security question>`
    // Use your secret question IDs in the json field
    // this would be a random string
    // for more information on this see the security question documentation
}

securityQuestion := SecurityQuestion{"Answer"}
body := ResetPasswordBySecurityAnswerAndPhoneStruct{
  securityQuestion,
  <phone>,
  <password>
}

response, err = lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).PutAuthResetPasswordBySecurityAnswerAndPhone(
  body,
)
if err != nil {
  // handle error
}
```

##### Auth Set or Change User Name

This API is used to add a username to an account, or to update the current username.

Please note this API requires the access token to be passed in the `Authorization Bearer` header. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-set-or-change-user-name)

Example:

```go
_, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).PutAuthSetOrChangeUsername(
  map[string]string{"username":<new username>},
)

if err != nil {
  // handle error
}
```

##### Auth Update Profile By Token

This API is used to update the profile of a user using an access token associated with their account.

Please note this API requires the access token to be passed in the `Authorization Bearer` header. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-update-profile-by-token)

Example:

```go
_, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).PutAuthUpdateProfileByToken(
  map[string]string{
    "Username":<new username>,
    "Suffix": <new suffix>,
    // add fields to be updated for the user to the body as needed
  }
)
if err != nil {
  // handle error
}
```

##### Auth Update Security Question By Access Token

This API is used to update the security questions using a user's access token.

Please note this API requires the access token to be passed in the `Authorization Bearer` header. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

For additional information on security questions, please refer the [documentation](https://www.loginradius.com/docs/api/v2/admin-console/platform-security/security-question/#password-policy)

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-update-security-question-by-access-token)

Example:

```go
// Initialize struct to use for body
// Alternatively, the helper struct lrbody.ResetPwSecurityQuestionUsername from the lrbody package
// could be used, or []byte could be passed as body
type PutAuthUpdateSecurityQuestionByAccessTokenStruct struct {
    SecurityAnswer SecurityQandA `json:"securityanswer"`
}

type SecurityQandA struct {
  SecurityQuestion string `json:<id of security question>`
  // Use your secret question IDs in the json field
  // this would be a random string
  // for more information on this see the security question documentation
}

securityQuestion := SecurityQuestion{"Answer"}
body := PutAuthUpdateSecurityQuestionByAccessTokenStruct{
  securityQuestion,
}

response, err = lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).PutAuthUpdateSecurityQuestionByAccessToken(
  body,
)

if err != nil {
  // handle error
}
```

##### Auth Delete Account with Email Confirmation

This API is used to delete a user account by passing in their access token.

Please note this API requires the access token to be passed in the `Authorization Bearer` header. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-delete-account-with-email-confirmation)

Example:

```go
resp, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).DeleteAuthDeleteAccountEmailConfirmation()

if err != nil {
  // handle error
}
```

##### Auth Remove Email

This API is used to remove emails from an account. An account should
have at least one email at all times when used as an identifier.

Please note this API requires the access token to be passed in the `Authorization Bearer` header. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-remove-email)

Example:

```go
resp, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).DeleteAuthDeleteAccountEmailConfirmation()

if err != nil {
  // handle error
}
```

##### Auth Unlink Social Identities

This API is used to unlink a social provider account from the user associated with the access token.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-unlink-social-identities)

Please note this API requires the access token to be passed in the `Authorization Bearer` header. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

Example:

```go
response, err = lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).DeleteAuthUnlinkSocialIdentities(
  map[string]string{
    "provider": <provider>,
    "providerid": <providerid>,
  },
)

if err != nil {
  // handle error
}

```

### Account APIs

The Account Management APIs are used to manage a user's account. These calls require the API Key and API Secret and often the User's Account UID (Unified Identifier) to perform an operation. For this reason, these APIs are considered to be for back-end purposes.

To call an Account API, import the account package like so:

```go
import (
	lraccount "github.com/LoginRadius/go-sdk/api/account"
)
```

**List of APIs in this Section:**

- [POST: Account Create](#account-create)
- [POST: Get Email Verification Token](#get-email-verification-token)
- [POST: Get Forgot Password Token](#get-forgot-password-token)
- [GET: Account Identities By Email](#account-identities-by-email)
- [GET: Account Impersonation API](#account-impersonation-api)
- [GET: Account Password](#account-password)
- [GET: Account Profiles By Email](#account-profiles-by-email)
- [GET: Account Profiles By Username](#account-profiles-by-username)
- [GET: Account Profiles By Phone ID](#account-profiles-by-phone-id)
- [GET: Account Profiles By UID](#account-profiles-by-uid)
- [GET: Refresh Access Token By RefreshToken API](#refresh-access-token-by-refreshtoken-api)
- [GET: Revoke Refresh Token API](#revoke-refresh-token-api)
- [PUT: Account Set Password](#account-set-password)
- [PUT: Account Update](#account-update)
- [PUT: Account Update Security Question Config](#account-update-security-question-config)
- [PUT: Account Invalidate Verification Email](#account-invalidate-verification-email)
- [DELETE: Account Email Delete](#account-email-delete)
- [DELETE: Account Delete](#account-delete)

##### Account Create

This API is used to create an account which bypasses email verification.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/account/account-create)

Example:

```go
// Use struct provided by lrbody package to construct body
// alternatively you could initialize your own struct
// []byte could also be passed in lieu of a struct
user := lrbody.AccountCreate{
    Email: []lrbody.EmailArray{
      lrbody.EmailArray{
        Type:  "Primary", //This can be any value of your designation
        Value: "example@example.com",
      },
    },
    Password: "password",
    // add more profile fields as needed
}

response, err := lraccount.Loginradius(loginradius).PostManageAccountCreate(user)
if err != nil {
  // handle error
}
```

##### Get Email Verification Token

This API is used to generate an email verification token.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/account/get-email-verification-token)

Example:

```go
response, err := lraccount.Loginradius(lraccount.Loginradius{loginradius}).PostManageEmailVerificationToken(
  map[string]string{"email": <email>},
)

if err != nil {
  // handle error
}
```

##### Get Forgot Password Token

This API is used to generate a token to reset the user's password.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/account/get-forgot-password-token)

Example:

```go
response, err = lraccount.Loginradius(lraccount.Loginradius{loginradius}).PostManageForgotPasswordToken(
  map[string]string{"email":<email>},
  map[string]string{"sendemail": "true"}), //queries are optional
)

if err != nil {
  // handle error
}
```

##### Account Identities By Email

This API is used to fetch the account identities associated with an email.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/account/account-identities-by-email)

Example:

```go
response, err := lraccount.Loginradius(lraccount.Loginradius{lrclient}).GetManageAccountIdentitiesByEmail(
  map[string]string{"email": <email>},
)

if err != nil {
  // handle error
}

// This end point returns data in an array, the response needs to be handled like so:
// (please note this is a preliminary example, you may wish to do something
// different with the returned profiles)
body, _ := lrjson.DynamicUnmarshal(response.Body) // unmarshals body
profiles := body["Data"].([]interface{}) // type assertion
profile := profiles[0].(map[string]interface{}) // get first profile
uid := profile["Uid"].(string) // get id of first profile
```

##### Account Impersonation API

This API is used to generate a token by passing in a User ID.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/account/account-impersonation-api)

Example:

```go
response, err := lraccount.Loginradius(lraccount.Loginradius{lrclient}).GetManageAccountProfilesByUid(<uid>)
if err != nil {
  // handle error
}
```

##### Account Password

This API is used to get the hashed password for an account.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/account/account-password)

Example:

```go
response, err := lraccount.Loginradius(lraccount.Loginradius{loginradius}).GetManageAccountPassword(uid)
if err != nil {
  // handle error
}
```

##### Account Profiles By Email

This API is used to get the profile of an account associated with the passed in email.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/account/account-profiles-by-email)

Example:

```go
response, err := lraccount.Loginradius(lraccount.Loginradius{lrclient}).GetManageAccountProfilesByEmail(
  map[string]string{"email": <email>},
)
if err != nil {
  // handle error
}
```

##### Account Profiles By Username

This API is used to get the profile of an account associated with the passed in username.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/account/account-profiles-by-user-name/)

Example:

```go
response, err := lraccount.Loginradius(lraccount.Loginradius{lrclient}).GetManageAccountProfilesByUsername(
  map[string]string{"username": <username>},
)
if err != nil {
  // handle error
}
```

##### Account Profiles By Phone ID

This API is used to get the profile of an account associated with the passed in phone number.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/account/account-profiles-by-phone-id)

Example:

```go
response, err := lraccount.Loginradius(lraccount.Loginradius{lrclient}).GetManageAccountProfilesByPhone(
  map[string]string{"phone": <phone>},
)
if err != nil {
  // handle error
}
```

##### Account Profiles By UID

This API is used to get the profile of an account associated with the passed in user ID.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/account/account-profiles-by-uid)

Example:

```go
response, err := lraccount.Loginradius(lraccount.Loginradius{lrclient}).GetManageAccountProfilesByUid(<uid>)
if err != nil {
    // handle error
}
```


##### Refresh Access Token By RefreshToken API

This API will be used to Refresh Access token using the refresh token API.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/refresh-token/refresh-access-token-by-refresh-token)

Example:

```go
response, err := lraccount.Loginradius(lraccount.Loginradius{lrclient}).GetRefreshAccessTokenByRefreshToken(
  map[string]string{"refresh_token": <Referesh_Token>},
)
if err != nil {
  // handle error
}
```


##### Revoke Refresh Token API

This API will be used to Revoke the Refresh token.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/refresh-token/revoke-refresh-token)

Example:

```go
response, err := lraccount.Loginradius(lraccount.Loginradius{lrclient}).GetRevokeRefreshToken(
  map[string]string{"refresh_token": <Referesh_Token>},
)
if err != nil {
  // handle error
}
```


##### Account Set Password

This API is used to set the password for a user.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/account/account-set-password)

Example:

```go
response, err := lraccount.Loginradius(lraccount.Loginradius{lrclient}).PutManageAccountSetPassword(
  <uid>,
  map[string]string{"password":<new password>},
)

if err != nil {
  // handle error
}
```

##### Account Update

This API is used to update the profile of a user.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/account/account-update)

Example:

```go
// This example passes a []byte as body
// Alternatively a struct containing fields to be updated
// can also be passed in the body
_, err := lraccount.Loginradius(lraccount.Loginradius{lrclient}).PutManageAccountUpdate(
  uid,
  []byte(`{"Username":"newname"}`), // add profile fields as needed
)
if err != nil {
  // handle error
}
```

##### Account Update Security Questiong

This API is used to update security question configurations for an account.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/account/account-update-security-question)

For additional information on security questions, please refer the [documentation](https://www.loginradius.com/docs/api/v2/admin-console/platform-security/security-question/#password-policy)

Example:

```go
// Initialize struct to use for body
// Alternatively, the helper struct lrbody.ResetPwSecurityQuestionEmail from the lrbody package
// could be used, or []byte could be passed as body
type AccountUpdateSecurityQuestionConfigStruct struct {
    SecurityAnswer SecurityQandA `json:"securityanswer"`
}

type SecurityQandA struct {
    SecurityQuestion string `json:<id of security question>`
    // Use your secret question IDs in the json field
    // this would be a random string
    // for more information on this see the security question documentation
}

securityQuestion := SecurityQuestion{"Answer"}
body := ResetPasswordBySecurityAnswerAndEmailStruct{
  securityQuestion,
  <email>,
  <password>
}

response, err = lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).PutManageAccountUpdateSecurityQuestionConfig(
  <uid>,
  body,
)
if err != nil {
  // handle error
}
```

##### Account Invalidate Verification Email

This API is used to invalidate the verification email status on an account.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/account/account-invalidate-verification-email)

Example:

```go
response, err := lraccount.Loginradius(lraccount.Loginradius{lrclient}).PutManageAccountInvalidateVerificationEmail(<uid>)
if err != nil {
  // handle error
}
```

##### Account Email Delete

This API is used to delete an email off of an account. When emails are set as identifiers, an account should always have at least one email.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/account/account-email-delete)

Example:

```go
response, err := lraccount.Loginradius(lraccount.Loginradius{lrclient}).DeleteManageAccountEmail(
  uid,
  map[string]string{"email":<email to delete>},
)
if err != nil {
  // handle error
}
```

##### Account Delete

This API is used to delete an account.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/account/account-delete)

Example:

```go
_, err := lraccount.Loginradius(lraccount.Loginradius{lrclient}).DeleteManageAccount(<uid>)
if err != nil {
  // handle error
}
```


### Roles API

The Roles APIs are used to manage the creation and assignment of user roles for a customer account.

To call a Roles API, import the role package like so:

```go
import (
	"github.com/LoginRadius/go-sdk/api/role"
)
```

**List of APIs in this Section:**

- [POST: Roles Create](#roles-create)
- [GET: Get Context](#get-context)
- [GET: Roles List](#roles-list)
- [GET: Get Roles By UID](#get-roles-by-uid)
- [PUT: Add Permissions To Role](#add-permissions-to-role)
- [PUT: Assign Roles By UID](#assign-roles-by-uid)
- [PUT: Upsert Context](#upsert-context)
- [DELETE: Delete Role](#delete-role)
- [DELETE: Unassign Role By UID](#unassign-role-by-uid)
- [DELETE: Remove Permissions](#remove-permissions)
- [DELETE: Delete Context](#delete-context)
- [DELETE: Delete Role From Context](#delete-role-from-context)
- [DELETE: Delete Permissions From Context](#delete-permissions-from-context)

##### Roles Create

This API is used to create a role.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/roles-management/roles-create)

Example:

```go
// uses struct offered by lrbody package
// alternatively an anonymous struct could be used, or a []byte
role := lrbody.Role{
  Name: <rolename>,
  Permissions: map[string]bool{
    <permission name>:         true,
    <permission name>:         true,
  },
}
roles := lrbody.Roles{[]lrbody.Role{role}}

res, err = role.Loginradius(role.Loginradius{lrclient}).PostRolesCreate(roles)
if err != nil {
  // handle error
}
```

##### Get Context

This API is used to get a context associated with an account.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/roles-management/get-context)

Example:

```go
res, err := role.Loginradius(role.Loginradius{lrclient}).GetContextRolesPermissions(<uid>)

if err != nil {
  // handle error
}
```

##### Roles List

This API is used to get a list of roles within the customer account.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/roles-management/roles-list)

Example:

```go
res, err := role.Loginradius(role.Loginradius{lrclient}).GetRolesList()

if err != nil {
  // handle error
}
```

##### Get Roles By UID

This API is used to retrieve all the assigned roles for a particular user.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/roles-management/get-roles-by-uid)

Example:

```go
res, err := role.Loginradius(role.Loginradius{lrclient}).GetRolesByUID(<uid>)
if err != nil {
  // handle error
}
```

##### Add Permissions To Role

This API is used to add permissions to an existing role.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/roles-management/add-permissions-to-role)

Example:

```go
// This example uses a helper struct from the lrbody package
// Alternatively you could initialize your own struct, or pass a []byte as body
res, err := role.Loginradius(role.Loginradius{lrclient}).PutAccountAddPermissionsToRole(
  <role name>,
  lrbody.PermissionList{[]string{<permission name}},
)

if err != nil {
  // handle error
}
```

##### Assign Roles By UID

This API is used to assign your desired roles to a given user.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/roles-management/assign-roles-by-uid)

Example:

```go
res, err := role.Loginradius(role.Loginradius{lrclient}).PutRolesAssignToUser(
  uid,
  // this example uses a []byte as body, alternatively a struct or map[string]string // could be passed as well
  []byte(`{"roles": ["`+<role name>+`"]}`),
)

if err != nil {
  // handle error
}
```

##### Upsert Context

This API is used to create a context with a set of rules.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/roles-management/upsert-context)

Example:

```go
// This example uses a []byte as body, alternatively a struct or map[string]string // could be passed as well
body := []byte(`{"rolecontext":[{"context":<context name>, "roles":["<role name>"], "additionalpermissions":[<permission>, <permission>]}]}`)

res, err := role.Loginradius(role.Loginradius{lrclient}).PutRolesUpsertContext(
  <uid>,
  body,
)

if err != nil {
  // handle error
}
```

##### Delete Role

This API is used to delete a role.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/roles-management/delete-role)

Example:

```go
res, err := role.Loginradius(role.Loginradius{lrclient}).DeleteAccountRole(<rolename>)

if err != nil {
  // handle error
}
```

##### Unassign Role By UID

This API is used to unassign a role.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/roles-management/unassign-roles-by-uid)

Example:

```go
// This example uses a []byte as body, alternatively a struct or map[string]string // could be passed as well
res, err := role.Loginradius(role.Loginradius{lrclient}).PutRolesAssignToUser(
  uid,
  []byte(`{"roles": ["example_role_name"]}`),
)

if err != nil {
  // handle error
}
```

##### Remove Permissions

This API is used to remove permissions from a role.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/roles-management/remove-permissions)

Example:

```go
// This example uses a helper struct offered by the lrbody package as body, alternatively a []byte could be used as well
res, err := role.Loginradius(role.Loginradius{lrclient}).DeleteRolesAccountRemovePermissions(
  <rolename>,
  lrbody.PermissionList{[]string{<permissionName>}},
)

if err != nil {
  // handle error
}
```

##### Delete Context

This API deletes the passed in role context.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/roles-management/delete-context)

Example:

```go
res, err = role.Loginradius(role.Loginradius{lrclient}).DeleteContextFromRole(
  uid,
  <example_context>,
)

if err != nil {
  // handle error
}
```

##### Delete Role From Context

This API deletes the specified role from a context.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/roles-management/delete-role-from-context)

Example:

```go
// This example uses a []byte as body, alternatively a struct or map[string]string
// could be passed as well
res, err = role.Loginradius(role.Loginradius{lrclient}).DeleteRoleFromContext(
  uid,
  <example_context>,
  []byte(`{"roles":["<rolename>"]}`),
)

if err != nil {
  // handle error
}
```

##### Delete Permissions From Context

This API deletes additional permissions from a context.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/roles-management/delete-permissions-from-context)

Example:

```go
// This example uses a []byte as body, alternatively a struct or map[string]string
// could be passed as well
res, err = role.Loginradius(role.Loginradius{lrclient}).DeleteAdditionalPermissionFromContext(
  <uid>,
  <context name>,
  []byte(`{"additionalpermissions":["<permission name>"]}`),
)

if err != nil {
  // handle error
}
```

### Multi-Factor Authentication APIs

The Multi-Factor Authentication APIs are used to handle Multi-Factor Authentication for users. Multi-Factor Authentication should be enabled on the admin console before utilizing these APIs. Take note, the access tokens used for this section are different from the access tokens obtained from standard login and is obtained from calling the validate Backup Code/OTP/Google Auth Code API.

To call a Multi-Factor Authentication API, import the `mfa` package like so:

```go
import (
	"github.com/LoginRadius/go-sdk/api/mfa"
)
```

**List of APIs in this Section:**

- [POST: MFA Email Login](#mfa-email-login)
- [POST: MFA User Name Login](#mfa-user-name-login)
- [POST: MFA Phone Login](#mfa-phone-login)
- [GET: MFA Validate Access Token](#mfa-validate-access-token)
- [GET: MFA Backup Code By Access Token](#mfa-backup-code-by-access-token)
- [GET: MFA Reset Backup Code By Access Token](#mfa-reset-backup-code-by-access-token)
- [GET: MFA Backup Code By UID](#mfa-backup-code-by-uid)
- [GET: MFA Reset Backup Code By UID](#mfa-reset-backup-code-by-uid)
- [PUT: MFA Validate Backup Code](#mfa-validate-backup-code)
- [PUT: Validate MFA by OTP](#validate-mfa-by-otp)
- [PUT: Validate MFA by Google Authenticator Code](#validate-mfa-by-google-authenticator-code)
- [PUT: MFA Update Phone Number](#mfa-update-phone-number)
- [PUT: MFA Update Phone Number By Token](#mfa-update-phone-number-by-token)
- [PUT: Update MFA By Access Token](#update-mfa-by-access-token)
- [PUT: Update MFA Settings](#update-mfa-settings)
- [DELETE: MFA Reset Google Authenticator By Token](#mfa-reset-google-authenticator-by-token)
- [DELETE: MFA Reset SMS Authenticator By Token](#mfa-reset-sms-authenticator-by-token)
- [DELETE: MFA Reset Google Authenticator By UID](#mfa-reset-google-authenticator-by-uid)
- [DELETE: MFA Reset SMS Authenticator By UID](#mfa-reset-sms-authenticator-by-uid)

##### MFA Email Login

This API uses the multi-factor process to perform an authentication using email.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/mfa-email-login)

Example:

```go
// This example uses a map[string]string as body, alternatively []byte or
// struct could be passed as body as well
res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).PostMFAEmailLogin(
  map[string]string{"email":<email>, "password":<password>},
)

if err != nil {
  // handle error
}
```

##### MFA User Name Login

This API uses the multi-factor process to perform an authentication using username.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/mfa-user-name-login)

Example:

```go
// This example uses a map[string]string as body, alternatively []byte or
// struct could be passed as body as well
res, err = mfa.Loginradius(mfa.Loginradius{lrclient}).PostMFAUsernameLogin(
  map[string]string{"username": username, "password": password},
  map[string]string{"emailtemplate": "hello"},
)

if err != nil {
  // handle error
}
```

##### MFA Phone Login

This API uses the multi-factor process to perform an authentication using phone.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/mfa-phone-login)

Example:

```go
// This example uses a map[string]string as body, alternatively []byte or
// struct could be passed as body as well
res, err = mfa.Loginradius(mfa.Loginradius{lrclient}).PostMFAPhoneLogin(
  map[string]string{"phone": phone, "password": password},
  map[string]string{"emailtemplate": "hello"},
)

if err != nil {
  // handle error
}
```

##### MFA Validate Access Token

This API validates the access token after logging in with an optional MFA setting. Take note, the access tokens used for this section are different from the access tokens obtained from standard login and is obtained from calling the validate Backup Code/OTP/Google Auth Code API.

Please note this API requires the access token to be passed in the `Authorization Bearer` header. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/mfa-validate-access-token)

Example:

```go
// This example uses a map[string]string as body, alternatively []byte or
// struct could be passed as body as well
res, err = mfa.Loginradius(mfa.Loginradius{lrclient}).GetMFAValidateAccessToken(
  map[string]string{"smstemplate2fa": "hello"}, // this is an optional query parameter
)

if err != nil {
  // handle error
}
```

##### MFA Backup Code By Access Token

This API creates backup codes for user login after passing in a valid access token. Take note, the access tokens used for this section are different from the access tokens obtained from standard login and is obtained from calling the validate Backup Code/OTP/Google Auth Code API.

Please note this API requires the access token to be passed in the `Authorization Bearer` header. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/mfa-backup-code-by-access-token)

Example:

```go
res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).GetMFABackUpCodeByAccessToken()

if err != nil {
  // handle error
}
```

##### MFA Reset Backup Code By Access Token

This API resets any backup codes created for the account when passed in an access token. Take note, the access tokens used for this section are different from the access tokens obtained from standard login and is obtained from calling the validate Backup Code/OTP/Google Auth Code API.

Please note this API requires the access token to be passed in the `Authorization Bearer` header. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/mfa-reset-backup-code-by-access-token)

Example:

```go
res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).GetMFAResetBackUpCodeByAccessToken()

if err != nil {
  // handle error
}
```

##### MFA Backup Code By UID

This API creates backup codes for a user after passing in a UID. The UID passed in must have an account with a valid Multi-Factor Authentication setup.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/mfa-backup-code-by-uid)

Example:

```go
res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).GetMFABackUpCodeByUID(<uid>)

if err != nil {
  // handle error
}
```

##### MFA Reset Backup Code By UID

This API resets any backup codes created for the account when passed in a UID. The UID passed in must have an account with a valid Multi-Factor Authentication setup.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/mfa-reset-backup-code-by-uid)

Example:

```go
res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).GetMFAResetBackUpCodeByUID(<uid>)

if err != nil {
  // handle error
}
```

##### MFA Validate Backup Code

This API validates the passed in backup code and returns a multi-factor access token. For more information on the Multi-Factor Authentication token, read the overview document for Multi-Factor Authentication. It can be obtained from calling one of the multi-factor login APIs.

Please note this API requires the access token to be passed in the `Authorization Bearer` header. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/mfa-validate-backup-code)

Example:

```go
res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).GetMFAResetBackUpCodeByAccessToken()

if err != nil {
  // handle error
}
```

##### MFA Validate OTP

This API validates the OTP sent for Multi-Factor Authentication. For more information on the Multi-Factor Authentication token, read the overview document for Multi-Factor Authentication. It can be obtained from calling one of the multi-factor login APIs.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/sms-authenticator/mfa-validate-otp)

Example:

```go
// This example uses a map[string]string as body, alternatively a struct or a
// []byte could be passed as body as well.
res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).PostMFAEmailLogin(
  map[string]string{"email": <email>, "password": <password>},
)

if err != nil {
  // handle error
}
```

##### MFA Validate Google Auth Code

This API validates the Google Auth Code generated for Multi-Factor Authentication. For more information on the Multi-Factor Authentication token, read the overview document for Multi-Factor Authentication. It can be obtained from calling one of the multi-factor login APIs.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/google-authenticator/mfa-validate-google-auth-code/)

Example:

```go
// This example uses a map[string]string as body, alternatively a struct or a
// []byte could be passed as body as well.
res, err = mfa.Loginradius(mfa.Loginradius{lrclient}).PutMFAValidateGoogleAuthCode(
  map[string]string{"secondfactorauthenticationtoken": <token>},
  map[string]string{"googleauthenticatorcode": <google authenticator code>},
)

if err != nil {
  // handle error
}
```

##### MFA Update Phone Number

This API updates the phone number associated with MFA OTPs if configured for the account. For more information on the Multi-factor Authentication token, read the overview document for Multi-Factor Authentication. It can be obtained from calling one of the multi-factor login APIs.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/sms-authenticator/mfa-update-phone-number)

Example:

```go
// This example uses a map[string]string as body, alternatively a struct or a
// []byte could be passed as body as well.
res, err = mfa.Loginradius(mfa.Loginradius{lrclient}).PutMFAUpdatePhoneNumber(
  map[string]string{"secondfactorauthenticationtoken": <token>},
  map[string]string{"phoneno2fa": <phone number>},
)

if err != nil {
  // handle error
}
```

##### MFA Update Phone Number By Token

This API updates the phone number associated with MFA OTPs if configured for the account. This API uses the access token to update instead of second factor auth token.

Please note this API requires the access token to be passed in the `Authorization Bearer` header. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/sms-authenticator/mfa-update-phone-number-by-token)

Example:

```go
// This example uses a map[string]string as body, alternatively a struct or a
// []byte could be passed as body as well.
res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).PutMFAUpdatePhoneNumberByToken(
  map[string]string{"phoneno2fa": <phone number>},
)

if err != nil {
  // handle error
}
```

##### Update MFA By Access Token

This API enables MFA authentication through Google Authenticator Codes when the user is already logged in.

Please note this API requires the access token to be passed in the `Authorization Bearer` header. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/google-authenticator/update-mfa-by-access-token)

Example:

```go
res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).PutMFAUpdateByToken(<google authenticator code>)

if err != nil {
  // handle error
}
```

##### Update MFA Settings

This API enables MFA authentication through OTP after login.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/sms-authenticator/update-mfa-setting)

Please note this API requires the access token to be passed in the `Authorization Bearer` header. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

Example:

```go
// This example uses a map[string]string as body, alternatively a struct or a
// []byte could be passed as body as well.
res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).PutMFAUpdateSettings(
  map[string]string{"otp":<otp>}, // add optional body parameters in map if any
)

if err != nil {
  // handle error
}
```

##### MFA Reset Google Authenticator By Token

This API resets the MFA configuration for a Google Authenticator associated with a user's account.

Please note this API requires the access token to be passed in the `Authorization Bearer` header. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/google-authenticator/mfa-reset-google-authenticator-by-token)

Example:

```go
res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).DeleteMFAResetGoogleAuthenticatorByToken()

if err != nil {
  // handle error
}
```

##### MFA Reset SMS Authenticator By Token

This API resets the MFA configuration for a phone device associated with a user's account.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/sms-authenticator/mfa-reset-sms-authenticator-by-token)

Please note this API requires the access token to be passed in the `Authorization Bearer` header. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

Example:

```go
res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).DeleteMFAResetSMSAuthenticatorByToken()

if err != nil {
  // handle error
}
```

##### MFA Reset Google Authenticator By UID

This API resets the MFA configuration for a Google Authenticator associated with a user's account.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/google-authenticator/mfa-reset-google-authenticator-by-uid)

Example:

```go
res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).DeleteMFAResetGoogleAuthenticatorByUid(<uid>)

if err != nil {
  // handle error
}
```

##### MFA Reset SMS Authenticator By UID

This API resets the MFA configuration for a phone device associated with a user's account.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/sms-authenticator/mfa-reset-sms-authenticator-by-uid)

Example:

```go
res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).DeleteMFAResetSMSAuthenticatorByUid(<uid>)

if err != nil {
  // handle error
}
```

##### MFA Re-authenticate

This API is used to trigger the Multi-Factor Autentication workflow for the provided access_token

Please note this API requires the access token to be passed in the `Authorization Bearer` header. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/re-authentication/mfa-re-authenticate)

Example:

```go
res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).GetMFAReAuthenticate()
if err != nil {
  // handle error
}
```

##### Validate MFA by Google Authenticator Code

This API is used to re-authenticate via Multi-Factor Authentication by passing the Google Authenticator code.

Please note this API requires the access token to be passed in the `Authorization Bearer` header. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/re-authentication/mfa/re-auth-by-google-authenticator-code/)

Example:

```go
// This example uses a map[string]string as body, alternatively a []byte or
// a struct could also be passed as body.
res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).PutMFAReauthenticateByGoogleAuthenticator(
  map[string]string{"googleauthenticatorcode": <google authenticator code>},
)

if err != nil {
  // handle error
}
```

##### Validate MFA by OTP

This API is used to re-authenticate via Multi-Factor Authentication by passing the One-Time Passcode received via SMS authentication.

Please note this API requires the access token to be passed in the `Authorization Bearer` header. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/re-authentication/mfa/re-auth-by-otp)

Example:

```go
// This example uses a map[string]string as body, alternatively a []byte or
// a struct could also be passed as body.
res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).PutMFAReauthenticateByOTP(
  map[string]string{"otp": <otp?},
)

if err != nil {
  // handle error
}
```

##### Validate MFA by Backup Code

This API is used to re-authenticate by set of backup codes via access_token on the site that has Multi-Factor Authentication enabled in re-authentication for the user that does not have the device.

Please note this API requires the access token to be passed in the `Authorization Bearer` header. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/re-authentication/validate-mfa-by-backup-code)

Example:

```go
// This example uses a map[string]string as body, alternatively a []byte or
// a struct could also be passed as body.
res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).PutMFAReauthenticateByBackupCode(
  map[string]string{"backupcode": <backup code>},
)

if err != nil {
  // handle error
}
```

##### Validate MFA by Password

This API is used to re-authenticate via Multi-Factor Authentication by passing in the password.

Please note this API requires the access token to be passed in the `Authorization Bearer` header. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

[Documentation](https://www.loginradius.com/docshttps://www.loginradius.com/docs/api/v2/customer-identity-api/re-authentication/re-auth-validate-password)

Example:

```go
// This example uses a map[string]string as body, alternatively a []byte or
// a struct could also be passed as body.
res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).PutMFAReauthenticateByPassword(
  map[string]string{"password": "password"},
)
if err != nil {
  // handle error
}
```

### Social APIs

The Social APIs are used to fetch user profile and other data from providers linked to the user accounts. The access tokens in this section are obtained after validating an access token using a social provider. Look at Access Token via Facebook, Access Token via Twitter, Access Token via VKontakte to get these access tokens.

To call a Social API, import the social package like so:

```go
import (
  lrsocial "github.com/LoginRadius/go-sdk/api/social"
)
```

**List of APIs in this Section:**

- [POST: Post Message API](#post-message-api)
- [POST: Trackable Status Posting](#trackable-status-posting)
- [GET: Access Token](#access-token)
- [GET: Validate Access Token](#validate-access-token)
- [GET: Invalidate Access Token](#invalidate-access-token)
- [GET: Album](#album)
- [GET: Audio](#audio)
- [GET: Check-in](#check-in)
- [GET: Company](#company)
- [GET: Contact](#contact)
- [GET: Event](#event)
- [GET: Following](#following)
- [GET: Group](#group)
- [GET: Like](#like)
- [GET: Mention](#mention)
- [GET: Get Message API](#get-message-api)
- [GET: Page](#page)
- [GET: Photo](#photo)
- [GET: Post](#post)
- [GET: Status Fetching](#status-fetching)
- [GET: Status Posting](#status-posting)
- [GET: Video](#video)

##### Post Message API

This API is used to send messages through the user's provider account.

Please note this API requires the access token to be passed in. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/social-login/advanced-social-api/post-message-api)

Example:

```go
// This example uses a map[string]string as body, alternatively a []byte or
// a struct could also be passed as body.
resp, err := lrsocial.Loginradius(lrsocial.Loginradius{lrclient}).PostSocialMessageAPI(
  map[string]string{
    "to":      <receipient's social provider id>,
    "subject": <message subject>,
    "message": <message content>,
  },
)

if err != nil {
  // handle error
}
```

##### Trackable Status Posting

This API is used to post a new status for the user through their provider account.

Please note this API requires the access token to be passed in. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/social-login/advanced-social-api/trackable-status-posting)

Example:

```go
// This example uses a map[string]string as body, alternatively a []byte or
// a struct could also be passed as body.
queries := map[string]string{
  "url":         <url>,
  "title":       <title>,
  "imageurl":    <img url>,
  "status":      <status>,
  "caption":     <caption>,
  "description": <description>,
}

resp, err := lrsocial.Loginradius(lrsocial.Loginradius{lrclient}).PostSocialStatusPost(queries)

if err != nil {
  // handle error
}
```

##### Access Token

This API is used to translate a LoginRadius Request Token into an Access Token that can be used with all APIs.

For more information on LoginRadius Request Tokens, see [this documentation](/infrastructure-and-security/loginradius-tokens#loginradius-request-token-expiration-15-mins-).

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/social-login/access-token)

Example:

```go
resp, err := lrsocial.Loginradius(lrsocial.Loginradius{lrclient}).GetSocialAccessToken(<request token>)

if err != nil {
  // handle error
}
```

##### Validate Access Token

This API validates the passed in access token for social API usage.

Please note this API requires the access token to be passed in. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/social-login/validate-access-token)

Example:

```go
resp, err := lrsocial.Loginradius(lrsocial.Loginradius{lrclient}).GetSocialTokenValidate()

if err != nil {
  // handle error
}
```

##### Invalidate Access Token

This API invalidates the access token passed in.

Please note this API requires the access token to be passed in. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/social-login/invalidate-access-token)

Example:

```go
resp, err := lrsocial.Loginradius(lrsocial.Loginradius{lrclient}).GetSocialTokenInvalidate()

if err != nil {
  // handle error
}
```

##### Album

This API retrieves Album data associated with the user's provider.

Please note this API requires the access token to be passed in. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/social-login/advanced-social-api/album)

Example:

```go
resp, err := lrsocial.Loginradius(lrsocial.Loginradius{lrclient}).GetSocialAlbum()

if err!= nil {
  // handle error
}
```

##### Audio

This API retrieves Audio associated with the user's provider.

Please note this API requires the access token to be passed in. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/social-login/advanced-social-api/audio)

Example:

```go
resp, err := lrsocial.Loginradius(lrsocial.Loginradius{lrclient}).GetSocialAudio()

if err != nil {
  // handle error
}
```

##### Check-in

This API retrieves Check-in data associated with the user's provider.

Please note this API requires the access token to be passed in. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/social-login/advanced-social-api/check-in)

Example:

```go
res, err = lrsocial.Loginradius(lrsocial.Loginradius{lrclient}).GetSocialCheckin()
if err != nil {
  // handle error
}
```

##### Company

This API retrieves Company data associated with the user's provider.

Please note this API requires the access token to be passed in. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/social-login/advanced-social-api/company)

Example:

```go
res, err = lrsocial.Loginradius(lrsocial.Loginradius{lrclient}).GetSocialCompany()
if err != nil {
  // handle error
}
```

##### Contact

This API retrieves Contact data associated with the user's provider.

Please note this API requires the access token to be passed in. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/social-login/advanced-social-api/contact)

Example:

```go
res, err = lrsocial.Loginradius(lrsocial.Loginradius{lrclient}).GetSocialContact()

if err != nil {
  // handle error
}
```

##### Event

This API retrieves Event data associated with the user's provider.

Please note this API requires the access token to be passed in. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/social-login/advanced-social-api/event)

Example:

```go
res, err = lrsocial.Loginradius(lrsocial.Loginradius{lrclient}).GetSocialEvent()

if err != nil {
  // handle error
}
```

##### Following

This API retrieves follower data associated with the user's provider.

Please note this API requires the access token to be passed in. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/social-login/advanced-social-api/following)

Example:

```go
res, err = lrsocial.Loginradius(lrsocial.Loginradius{lrclient}).GetSocialFollowing()
if err != nil {
  // handle error
}
```

##### Group

This API retrieves Group data associated with the user's provider.

Please note this API requires the access token to be passed in. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/social-login/advanced-social-api/group)

Example:

```go
res, err = lrsocial.Loginradius(lrsocial.Loginradius{lrclient}).GetSocialGroup()
if err != nil {
  // handle error
}
```

##### Like

This API retrieves Like data associated with the user's provider.

Please note this API requires the access token to be passed in. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/social-login/advanced-social-api/like)

Example:

```go
res, err = lrsocial.Loginradius(lrsocial.Loginradius{lrclient}).GetSocialLike()
if err != nil {
  // handle error
}
```

##### Mention

This API retrieves Mention data associated with the user's provider.

Please note this API requires the access token to be passed in. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/social-login/advanced-social-api/mention)

Example:

```go
res, err = lrsocial.Loginradius(lrsocial.Loginradius{lrclient}).GetSocialMention()
if err != nil {
  // handle error
}
```

##### Get Message API

The Message API is used to post messages to the user’s contacts.

Supported Providers: LinkedIn, Twitter

The Message API is used to post messages to the user’s contacts. This is one of the APIs that makes up the LoginRadius Friend Invite System. After using the Contact API, you can send messages to the retrieved contacts. This API requires setting permissions in your LoginRadius Admin Console.

GET and POST Message APIs work the same way except the API method is different.

Please note this API requires the access token to be passed in. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/social-login/advanced-social-api/get-message-api)

Example:

```go
// This example uses a map[string]string as body, alternatively a []byte or
// a struct could also be passed as body.
queries := map[string]string{
  "url":         <url>,
  "title":       <title>,
  "imageurl":    <img url>,
  "status":      <status>,
  "caption":     <caption>,
  "description": <description>,
}

resp, err := lrsocial.Loginradius(lrsocial.Loginradius{lrclient}).GetSocialStatusPost(queries)

if err != nil {
  // handle error
}
```

##### Page

This API retrieves Page data for the page passed in. The passed in page typically is an ID designated by the provider.

Please note this API requires the access token to be passed in. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/social-login/advanced-social-api/page)

Example:

```go
_, err = lrsocial.Loginradius(lrsocial.Loginradius{lrclient}).GetSocialPage(<pagename>)
if err != nil {
  // handle error
}
```

##### Photo

This API retrieves Photo data associated with a passed in Album ID.

Please note this API requires the access token to be passed in. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/social-login/advanced-social-api/photo)

Example:

```
// Takes string as argument
res, err = lrsocial.Loginradius(lrsocial.Loginradius{lrclient}).GetSocialPhoto(<albumid>)

if err != nil {
  // handle error
}
```

##### Post

This API retrieves Post data associated with a user's account.

Please note this API requires the access token to be passed in. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/social-login/advanced-social-api/post)

Example:

```go
res, err = lrsocial.Loginradius(lrsocial.Loginradius{lrclient}).GetSocialPost()

if err != nil {
  // handle error
}
```

##### Status Fetching

This API retrieves Status data associated with a user's account.

Please note this API requires the access token to be passed in. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/social-login/advanced-social-api/status-fetching)

Example:

```go
res, err = lrsocial.Loginradius(lrsocial.Loginradius{lrclient}).GetSocialStatus()
if err != nil {
  // handle error
}
```

##### Status Posting

This API is used to post a new status for the user through their provider account.

Please note this API requires the access token to be passed in. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/social-login/advanced-social-api/status-posting)

Example:

```go
func StatusPostingGet() {
    // Set LoginRadius Environment Here
    loginradius.SetLoginRadiusEnv("MyAPIKey", "MyAPISecret", "https://api.loginradius.com")

    // Call the API
    response, err := loginradius.GetSocialStatusPost("AccessToken", "StatusTitle", "StatusURL", "StatusImageURL", "StatusMessage", "StatusCaption", "StatusDescription")
    if(err != nil) {
        // handle the error
    }

    // Use the returned object
    fmt.Printf("%+v\n", response)
}
```

##### Video

This API retrieves Video data associated with a user's account.

Please note this API requires the access token to be passed in. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/social-login/advanced-social-api/video)
Example:

```go
res, err = lrsocial.Loginradius(lrsocial.Loginradius{lrclient}).GetSocialVideo()

if err != nil {
  // handle error
}
```

### Phone Authentication APIs

The Phone Authentication APIs are used similar to the regular Login Authentication APIs, but consist of a flow using an SMS device instead. Phone OTPs are used for verification instead of email verification for the APIs that need verification.

To call a Phone Authentication API, import the `phoneauthentication` package like so:

```go
import (
  "github.com/LoginRadius/go-sdk/api/phoneauthentication"
)
```

**List of APIs in this Section:**

- [POST: Phone Login](#phone-login)
- [POST: Phone Forgot Password By OTP](#phone-forgot-password-by-otp)
- [POST: Phone Resend OTP](#phone-resend-otp)
- [POST: Phone Resend OTP By Token](#phone-otp-by-token)
- [POST: Phone User Registration By SMS](#phone-user-registration-by-sms)
- [GET: Passwordless Login By Phone](#passwordless-login-by-phone)
- [GET: Phone Number Availability](#phone-number-availability)
- [GET: Passwordless Login Phone Verification](#passwordless-login-phone-verification)
- [PUT: Phone Number Update](#phone-number-update)
- [PUT: Phone Reset Password By OTP](#phone-reset-password-by-otp)
- [PUT: Phone Verify OTP](#phone-verify-otp)
- [PUT: Phone Verify OTP By Token](#phone-verify-otp-by-token)
- [PUT: Reset Phone ID Verification](#reset-phone-id-verification)
- [DELETE: Remove Phone ID By Access Token](#remove-phone-id-by-access-token)

##### Phone Login

This API logs in a user using their PhoneId and returns an access token and profile data.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/phone-authentication/phone-login)

Example:

```go
// This example uses a map[string]string as body, alternatively a struct or
// []byte could be passed as well
resp, err := phoneauthentication.Loginradius(phoneauthentication.Loginradius{lrclient}).PostPhoneLogin(
  map[string]string{"phone": <phone id>, "password": <password>},
)

if err != nil {
  // handle error
}
```

##### Phone Forgot Password By OTP

This API starts the forgot password flow for a phone user by sending a forgot password OTP to their phone.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/phone-authentication/phone-forgot-password-by-otp)

Example:

```go
// This example uses a map[string]string as body, alternatively a struct or
// []byte could be passed as well
resp, err := phoneauthentication.Loginradius(phoneauthentication.Loginradius{lrclient}).PostPhoneForgotPasswordByOTP(
  map[string]string{"phone": <phone number>},
)

if err != nil {
  // handle error
}
```

##### Phone Resend OTP

This API resends the account verification OTP to the PhoneId associated with the user account.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/phone-authentication/phone-resend-otp)

Example:

```go
// This example uses a map[string]string as body, alternatively a struct or
// []byte could be passed as well
resp, err := phoneauthentication.Loginradius(phoneauthentication.Loginradius{lrclient}).PostPhoneResendVerificationOTP(
  map[string]string{"phone": <phone number>},
)

if err != nil {
  // handle error
}
```

##### Phone Resend OTP By Token

This API resends the account verification OTP to the PhoneId associated with the user account in cases where an active token exists.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/phone-authentication/phone-resend-otp-by-token)

Please note this API requires the access token to be passed in the `Authorization Bearer` header. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

Example:

```go
// This example uses a map[string]string as body, alternatively a struct or
// []byte could be passed as well
resp, err := phoneauthentication.Loginradius(phoneauthentication.Loginradius{lrclient}).PostPhoneResendVerificationOTPByToken(
  map[string]string{"phone": <phone number>},
)

if err != nil {
  // handle error
}
```

##### Phone User Registration By SMS

This API registers a user with the profile data provided in the body.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/phone-authentication/phone-user-registration-by-sms)

Example:

```go
// This example uses initialized struct, alternatively a helper struct from lrbody
// package could be used, an anonymous struct or []byte would also work.
type Email struct {
	Type  string `json:"Type"`
	Value string `json:"Value"`
}

type User struct {
	Email    []Email `json:"Email"`
	Password string  `json:"Password"`
  PhoneId string  `json:"PhoneId"`

}

user:= User{
  Email: []Email{
    Email{
      Type:  "Primary",
      Value: <email>,
    },
  },
  Password: "password",
  PhoneId:"phone_number",
}
sott:="" //(Required) Sott
res, err:= phoneauthentication.Loginradius(phoneauthentication.Loginradius{lrclient}).PostPhoneUserRegistrationBySMS(sott,user,)
if err != nil {
  // handle error
}
```

##### Passwordless Login By Phone

This API sends an OTP to the phone. This API is part of the Passwordless Login module.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/passwordless-login/passwordless-login-by-phone)

Example:

```go
// This example uses a map[string]string as body, alternatively a struct or
// []byte could be passed as well
res, err := phoneauthentication.Loginradius(phoneauthentication.Loginradius{lrclient}).PutPhoneLoginUsingOTP(
  map[string]string{"phone": "", "otp": "871962"},
)

if err != nil {
  // handle error
}
```

##### Phone Number Availability

This API checks whether a phone number is available for use within the database.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/phone-authentication/phone-number-availability)

Example:

```go
// This example uses a map[string]string as body, alternatively a struct or
// []byte could be passed as well
resp, err := phoneauthentication.Loginradius(phoneauthentication.Loginradius{lrclient}).GetPhoneNumberAvailability(
  map[string]string{"phone": <phone number>},
)

if err != nil {
  // handle error
}
```

##### Passwordlass Login Phone Verification

This API logs in a user using an OTP. This API is part of the Passwordless Login module.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/passwordless-login/passwordless-login-phone-verification)

Example:

```go
// This example uses a map[string]string as body, alternatively a struct or
// []byte could be passed as well
res, err := phoneauthentication.Loginradius(phoneauthentication.Loginradius{lrclient}).PutPhoneLoginUsingOTP(
  map[string]string{"phone": "", "otp": <otp>},
)

if err != nil {
  // handle error
}
```

##### Phone Number Update

This API updates a user's SMS PhoneId.

Please note this API requires the access token to be passed in the `Authorization Bearer` header. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/phone-authentication/phone-number-update)

Example:

```go
// This example uses a map[string]string as body, alternatively a struct or
// []byte could be passed as well
res, err := phoneauthentication.Loginradius(phoneauthentication.Loginradius{lrclient}).PutPhoneNumberUpdate(
  map[string]string{"phone": <phone number>},
)

if err != nil {
  // handle error
}
```

##### Phone Reset Password By OTP

This API resets the user's password by passing in an OTP.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/phone-authentication/phone-reset-password-by-otp)

Example:

```go
// This example uses a map[string]string as body, alternatively a struct or
// []byte could be passed as well
resp, err := phoneauthentication.Loginradius(phoneauthentication.Loginradius{lrclient}).PutPhoneResetPasswordByOTP(
  map[string]string{"phone": <phone number>, "password": <password>, "otp": <otp>},
)

if err != nil {
  // handle error
}
```

##### Phone Verify OTP

This API verifies a user account by validating an OTP sent to the user's device. Take note that OTP is a query parameter in this function.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/phone-authentication/phone-verify-otp)

Example:

```go
// This example uses a map[string]string as body, alternatively a struct or
// []byte could be passed as well
res, err := phoneauthentication.Loginradius(phoneauthentication.Loginradius{lrclient}).PutPhoneVerificationByOTP(
  map[string]string{"otp": <otp>},
  map[string]string{"phone": <phone number>},
)

if err != nil {
  // handle error
}
```

##### Phone Verfiy OTP By Token

This API verifies a user account that is already logged in using an OTP.

Please note this API requires the access token to be passed in the `Authorization Bearer` header. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/phone-authentication/phone-verify-otp-by-token)

Example:

```go
// This example uses a map[string]string as body, alternatively a struct or
// []byte could be passed as well
res, err := phoneauthentication.Loginradius(phoneauthentication.Loginradius{lrclient}).PutPhoneVerificationByOTPByToken(
  map[string]string{"otp": <otp>},
)

if err != nil {
  // handle error
}
```

##### Reset Phone ID Verification

This API resets the verified status of a PhoneId on a user's account.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/phone-authentication/reset-phone-id-verification)

Example:

```go
// Takes a string argument
res, err := phoneauthentication.Loginradius(phoneauthentication.Loginradius{lrclient}).PutResetPhoneIDVerification(<uid>)
if err != nil {
  // handle error
}
```

##### Remove Phone ID By Access Token

This API deletes the PhoneId on a user's account.

Please note this API requires the access token to be passed in the `Authorization Bearer` header. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/phone-authentication/remove-phone-id-by-access-token)

Example:

```go
res, err := phoneauthentication.Loginradius(phoneauthentication.Loginradius{lrclient}).DeleteRemovePhoneIDByAccessToken()

if err != nil {
  // handle error
}
```

### Custom Object APIs

The Custom Object APIs are used to interact with custom objects that are set on the customer account. For more information on how to set up custom objects for an account, contact support.

To call a Custom Object API, import the `customobject` package like so:

```go
import (
	"github.com/LoginRadius/go-sdk/api/customobject"
)
```

**List of APIs in this Section:**

- [POST: Create Custom Object By UID](#create-custom-object-by-uid)
- [POST: Create Custom Object By Token](#create-custom-object-by-token)
- [GET: Custom Object By ObjectRecordID and UID](#custom-object-by-objectrecordid-and-uid)
- [GET: Custom Object By ObjectRecordID and Token](#custom-object-by-objectrecordid-and-token)
- [GET: Custom Object By UID](#custom-object-by-uid)
- [GET: Custom Object By Token](#custom-object-by-token)
- [PUT: Custom Object Update By ObjectRecordID and UID](#custom-object-update-by-objectrecordid-and-uid)
- [PUT: Custom Object Update By ObjectRecordID and Token](#custom-object-update-by-objectrecordid-and-token)
- [DELETE: Custom Object Delete By ObjectRecordID and UID](#custom-object-delete-by-objectrecordid-and-uid)
- [DELETE: Custom Object Delete By ObjectRecordID and Token](#custom-object-delete-by-objectrecordid-and-token)

##### Create Custom Object By UID

This API creates a custom object for a user using the UID.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/custom-object/create-custom-object-by-uid)

Example:

```go
// custom object is sent in the body, it can be any data you wish to store
// in the LoginRadius cloud.
customObj := map[string]string{
  <key>: <value>,
  <key>: <value>,
}

resp, err := customobject.Loginradius(customobject.Loginradius{lrclient}).PostCustomObjectCreateByUID(
  <uid>, // this is a string
  map[string]string{"objectname": <object name>}, // this is the query parameter
  customObj
)

if err != nil {
  // handle error
}
```

##### Create Custom Object By Token

This API creates a custom object for a user using the token.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/custom-object/create-custom-object-by-token)

Please note this API requires the access token to be passed in the `Authorization Bearer` header. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

Example:

```go
customObj := map[string]string{
  <key>: <value>,
  <key>: <value>,
}

resp, err := customobject.Loginradius(customobject.Loginradius{lrclient}).PostCustomObjectCreateByToken(
  map[string]string{"objectname": <object name>}, // this is the query parameter
  customObj
)

if err!= nil {
  // handle error
}
```

##### Custom Object By ObjectRecordID and UID

This API gets the custom objects for a user with an object record ID and their UID.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/custom-object/custom-object-by-objectrecordid-and-uid)

Example:

```go
resp, err := customobject.Loginradius(customobject.Loginradius{lrclient}).GetCustomObjectByObjectRecordIDAndUID(
  <uid>,  // string
  <object id>, // string
  map[string]string{"objectname": <object},
)

if err != nil {
  // handle error
}
```

##### Custom Object By ObjectRecordID and Token

This API gets the custom objects for a user with an object record ID and their access token.

Please note this API requires the access token to be passed in the `Authorization Bearer` header. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/custom-object/custom-object-by-objectrecordid-and-token)

Example:

```go
resp, err := customobject.Loginradius(customobject.Loginradius{lrclient}).GetCustomObjectByObjectRecordIDAndToken(
  <object id>, // string
  map[string]string{"objectname": <object name>},
)
if err != nil {
  // handle error
}
```

##### Custom Object By UID

This API gets the custom objects for a user with their UID.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/custom-object/custom-object-by-uid)

Example:

```go
resp, err := customobject.Loginradius(customobject.Loginradius{lrclient}).GetCustomObjectByUID(
  <uid>, // uid is a string
  map[string]string{"objectname": <object},
)

if err != nil {
  // handle error
}
```

##### Custom Object By Token

This API gets the custom objects for a user with their access token.

Please note this API requires the access token to be passed in the `Authorization Bearer` header. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/custom-object/custom-object-by-token)

Example:

```go
resp, err := customobject.Loginradius(customobject.Loginradius{lrclient}).GetCustomObjectByToken(
  map[string]string{"objectname": <object name>},
)

if err != nil {
  // handle error
}
```

##### Custom Object Update By ObjectRecordID and UID

This API updates a custom object associated with a user UID and a Object Record ID. There are different update types - please refer the [documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/custom-object/custom-object-update-by-objectrecordid-and-uid) for details.

Example:

```go
// The custom object to be sent
customObj := map[string]string{
  <key>: <value>,
  <key>: <value>,
}

resp, err := customobject.Loginradius(customobject.Loginradius{lrclient}).PutCustomObjectUpdateByUID(
  <uid>, // string
  <objectId>, // string
  map[string]string{"objectname": <object, "updatetype": "replace"}, // query parameters
  customObj, // string
)

if err != nil {
  // handle error
}
```

##### Custom Object Update By ObjectRecordID and Token

This API updates a custom object associated with a user access token and a Object Record ID. There are different update types that can be found [here]().

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/custom-object/custom-object-update-by-objectrecordid-and-token)

Please note this API requires the access token to be passed in the `Authorization Bearer` header. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).
Example:

```go
// The custom object to be sent
customObj := map[string]string{
  <key>: <value>,
  <key>: <value>,
}

resp, err := customobject.Loginradius(customobject.Loginradius{lrclient}).PutCustomObjectUpdateByToken(
  <objectId>, // string
  map[string]string{"objectname": <object, "updatetype": "replace"}, // query parameters
  customObj, // string
)

if err != nil {
  // handle error
}
```

##### Custom Object Delete By ObjectRecordID and UID

This API deletes a custom object associated with a user UID and an ObjectRecordID.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/custom-object/custom-object-delete-by-objectrecordid-and-uid)

Example:

```go
resp, err := customobject.Loginradius(customobject.Loginradius{lrclient}).PostCustomObjectCreateByUID(
  <uid>, // a string
  map[string]string{"objectname": <object},
  customObj,
)

if err != nil {
  // handle error
}
```

##### Custom Object Delete By ObjectRecordID and Token

This API deletes a custom object associated with a user access token and an ObjectRecordID.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/custom-object/custom-object-delete-by-objectrecordid-and-token)

Example:

```go
resp, err = customobject.Loginradius(customobject.Loginradius{lrclient}).DeleteCustomObjectByObjectRecordIDAndToken(
  <id>, // string
  map[string]string{"objectname": <object},
)

if err != nil {
  //handle err
}
```

### Smart Login APIs

The Smart Login APIs are logins that allow a user to login through a unique client GUID. This is focused on creating an easy method for Smart devices to access the LoginRadius authentication features. The client GUID is a unique identifier that can only be used once per login.

To call a Smart Login API, import the `smartlogin` package like so:

```go
import (
	"github.com/LoginRadius/go-sdk/api/smartlogin"
)
```

**List of APIs in this Section:**

- [GET: Smart Login By Email](#smart-login-by-email)
- [GET: Smart Login By Username](#smart-login-by-username)
- [GET: Smart Login Ping](#smart-login-ping)
- [GET: Smart Login Verify Token](#smart-login-verify-token)

##### Smart Login By Email

This API sends a Smart Login link to the user's email to sign in.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/smart-login/smart-login-by-email)

Example:

```go
res, err := smartlogin.Loginradius(smartlogin.Loginradius{lrclient}).GetSmartLoginByEmail(
  map[string]string{
    "email": <email>,
    "clientguid": <guid>,
  },
)
if err != nil {
  // handle error
}
```

##### Smart Login By Username

This API sends a Smart Login link to the user's email to sign in.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/smart-login/smart-login-by-username)

Example:

```go
res, err := smartlogin.Loginradius(smartlogin.Loginradius{lrclient}).GetSmartLoginByUsername(
    map[string]string{
    "username": <username>,
    "clientguid": <guid>,
  },
)
if err != nil {
  // handle error
}
```

##### Smart Login Ping

This API checks if the Smart Login link for a client GUID has been accessed and verified.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/smart-login/smart-login-ping)

Example:

```go
res, err = smartlogin.Loginradius(smartlogin.Loginradius{lrclient}).GetSmartLoginPing(
  map[string]string{"clientguid": <guid>},
)

if err != nil {
  // handle error
}
```

##### Smart Login Verify Token

This API validates the token generated by a Smart Login link.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/smart-login/smart-login-verify-token)

Example:

```go
res, err := smartlogin.Loginradius(smartlogin.Loginradius{lrclient}).GetSmartLoginVerifyToken(
  map[string]string{"verificationtoken": <verification token>},
)
if err != nil {
  // handle error
}
```

### One Touch Login APIs

The One Touch Login APIs use email and phone verification to create links that allow the user to login.

To call a One Touch Login API, import the `onetouchlogin` package like so:

```go
import (
	"github.com/LoginRadius/go-sdk/api/onetouchlogin"
)
```

**List of APIs in this Section:**

- [GET: One Touch Login By Email Captcha](#one-touch-login-by-email-captcha)
- [GET: One Touch Login By Phone Captcha](#one-touch-login-by-phone-captcha)
- [PUT: One Touch OTP Verification](#one-touch-otp-verification)

##### One Touch Login By Email

This API sends a link to an email to start the One Touch Login workflow.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/one-touch-login/one-touch-login-by-email-captcha)

Example:

```go
res, err := onetouchlogin.Loginradius(onetouchlogin.Loginradius{lrclient}).PostOneTouchLoginByEmail(
  map[string]string{
    "clientguid": <guid>,
    "email": <email>,
    "g-recaptcha-response": <google captcha response>,
  },
)

if err!= nil {
  // handle error
}
```

##### One Touch Login By Phone Captcha

This API sends an OTP to an SMS device to start the One Touch Login workflow.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/one-touch-login/one-touch-login-by-phone-captcha)

Example:

```go
res, err := onetouchlogin.Loginradius(onetouchlogin.Loginradius{lrclient}).PostOneTouchLoginByEmail(
  map[string]string{
    "clientguid": <guid>,
    "phone": <phone>,
    "g-recaptcha-response": <google captcha response>,
  },
)

if err!= nil {
  // handle error
}
```

##### One Touch OTP Verification

This API verifies an OTP sent to a user SMS and returns a user profile and access token if valid.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/one-touch-login/one-touch-otp-verification)

Example:

```go
res, err := onetouchlogin.Loginradius(onetouchlogin.Loginradius{lrclient}).PutOneTouchOTPVerification(
  map[string]string{"otp": <otp>},
  map[string]string{"phone": <phone number>},
)

if err!= nil {
  // handle error
}
```

### Configuration and Infrastructure APIs

The Configuration and Infrastructure APIs are used to view configurations and information around the customer account.

To call a Configuration and Infrastructure API, import the `lrconfiguration` package like so:

```go
import (
	lrconfiguration "github.com/LoginRadius/go-sdk/api/configuration"
)
```

**List of APIs in this Section:**

- [GET: Get Configurations](#get-configuration)
- [GET: Get Server Time](#get-server-time)
- [GET: Generate SOTT Token](#generate-sott-token)
- [GET: Get Active Session Details](#get-active-session-details)

##### Configuration

This API gets the configuration for the customer admin console.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/configuration/get-configurations)

Example:

```go
res, err := lrconfiguration.Loginradius(lrconfiguration.Loginradius{lrclient}).GetConfiguration()

if err != nil {
  // handle error
}
```

##### Get Server Time

This API queries the LoginRadius account for server information and time.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/configuration/get-server-time)

Example:

```go
res, err := lrconfiguration.Loginradius(lrconfiguration.Loginradius{lrclient}).GetServerTime()

if err != nil {
  // handle error
}
```

##### Generate SOTT Token

This API generates a SOTT which can be used for account creation.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/session/generate-sott-token)

Example:

```go
res, err := lrconfiguration.Loginradius(lrconfiguration.Loginradius{lrclient}).GetGenerateSottAPI()

if err != nil {
  // handle error
}
```

##### Get Active Session Details

This API gets all the active sessions that exist with the access token.

Please note this API requires the access token to be passed in the `Authorization Bearer` header. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/social-login/get-active-session-details)

Example:

```go
res, err := lrconfiguration.Loginradius(lrconfiguration.Loginradius{lrclient}).GetActiveSessionDetails()

if err != nil {
  // handle error
}
```

### Token Management APIs

The Token Management APIs allow management of access tokens and generation tokens usable by the social APIs.

To call a Token Management API, import the `tokenmanagement` package like so:

```go
import (
	"github.com/LoginRadius/go-sdk/api/tokenmanagement"
)
```

**List of APIs in this Section:**

- [GET: Access Token Via Facebook Token](#access-token-via-facebook-token)
- [GET: Access Token Via Twitter Token](#access-token-via-twitter-token)
- [GET: Access Token Via Vkontakte Token](#access-token-via-vkontakte-token)
- [GET: Refresh User Profile](#refresh-user-profile)
- [GET: Refresh Token](#refresh-token)

##### Access Token Via Facebook Token

This API generates a token that can be used with Facebook compatible Social APIs.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/social-login/native-social-login-api/access-token-via-facebook-token)

Example:

```go
res, err := tokenmanagement.Loginradius(tokenmanagement.Loginradius{lrclient}).GetAccessTokenViaFacebook(
  map[string]string{"fb_access_token": <fb access token>},
)

if err != nil {
  // handle error
}
```

##### Access Token Via Twitter Token

This API generates a token that can be used with Twitter compatible Social APIs.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/social-login/native-social-login-api/access-token-via-twitter-token)

Example:

```go
res, err := tokenmanagement.Loginradius(tokenmanagement.Loginradius{lrclient}).GetAccessTokenViaTwitter(
  map[string]string{
    "tw_access_token": <twitter access token>,
    "tw_token_secret": <twitter token secret>,
  },
)

if err != nil {
  // handle error
}
```

##### Access Token Via VKontakte Token

This API generates a token that can be used with VKontakte compatible Social APIs.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/social-login/native-social-login-api/access-token-via-vkontakte-token)

Example:

```go
res, err := tokenmanagement.Loginradius(tokenmanagement.Loginradius{lrclient}).GetAccessTokenViaVkontakte(
  map[string]string{"vk_access_token": <vk access token>},
)

if err != nil {
  // handle error
}
```

##### Refresh User Profile

Refreshes the user profile on the LoginRadius account by passing in a Social access token.

Please note this API requires the access token to be passed in the `Authorization Bearer` header. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/refresh-token/refresh-user-profile)

Example:

```go
res, err := tokenmanagement.Loginradius(tokenmanagement.Loginradius{lrclient}).GetRefreshUserProfile()

if err != nil {
  // handle error
}
```

##### Refresh Token

Refreshes and increases the lifetime of the token to up to 60 days.

Please note this API requires the access token to be passed in the `Authorization Bearer` header. The LoginRadius API client struct [must be initialized with a token](#Intializing-the-LoginRadius-Client).

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/refresh-token/refresh-token)

Example:

```go
res, err := tokenmanagement.Loginradius(tokenmanagement.Loginradius{lrclient}).GetRefreshToken()

if err != nil {
  // handle error
}
```

### Web Hooks APIs

WebHooks allow you to build or set up integrations which subscribe to certain events on LoginRadius. When one of those events is triggered, we'll send an HTTP POST payload to the WebHook's configured URL. WebHooks can be used to update an external tracker or update a backup mirror.

Each WebHook can be configured on LoginRadius or a specific LoginRadius site. Once configured, they will be triggered each time one or more subscribed events occur on that LoginRadius site.

For additional information, see [here](https://www.loginradius.com/docs/api/v2/integrations/webhooks/overview).

To call a Web Hooks API, import the `webhook` package like so:

```go
import (
	"github.com/LoginRadius/go-sdk/api/webhook"
)
```

**List of APIs in this Section:**

- [POST: Webhook Subscribe](#webhook-subscribe)
- [GET: Webhook Test](#webhook-test)
- [GET: Webhook Subscribed URLs](#webhook-subscribed-urls)
- [DELETE: Webhook Unsubscribe](#webhook-unsubscribe)

##### Webhook Subscribe

Configures a webhook onto a website where data will be sent when an event is triggered.

[Documentation](https://www.loginradius.com/docs/api/v2/integrations/webhooks/webhook-subscribe)

Example:

```go
// This example users a map[string]string, but a struct or []byte could be passed as body as well
res, err := webhook.Loginradius(webhook.Loginradius{lrclient}).PostWebhookSubscribe(
  map[string]string{
    "TargetUrl": <target url>,
    "Event":     <event>,
  },
)

if err != nil {
  // handle error
}
```

##### Webhook Test

This API is used to test a subscribed webhook.

[Documentation](https://www.loginradius.com/docs/api/v2/integrations/webhooks/webhook-test)

Example:

```go
res, err := webhook.Loginradius(webhook.Loginradius{lrclient}).GetWebhookTest()

if err != nil {
  // handle error
}
```

##### Webhook Subscribed URLs

This API is used to fetch all subscribed URLs for a particular event.

[Documentation](https://www.loginradius.com/docs/api/v2/integrations/webhooks/webhook-subscribed-urls)

Example:

```go
res, err := webhook.Loginradius(webhook.Loginradius{lrclient}).GetWebhookSubscribedURLs(
  map[string]string{"event": <event>},
)

if err != nil {
  // handle error
}
```

##### Webhook Unsubscribe

This API is used to fetch all subscribed URLs for a particular event.

[Documentation](https://www.loginradius.com/docs/api/v2/integrations/webhooks/webhook-unsubscribe)

Example:

```go
// This example users a map[string]string, but a struct or []byte could be passed as body as well
res, err := webhook.Loginradius(webhook.Loginradius{lrclient}).DeleteWebhookUnsubscribe(
  map[string]string{
    "targeturl": <target url>,
    "event":     <event>,
  },
)

if err != nil {
  // handle error
}
```

### Passwordless Login APIs

The Passwordless Login APIs are used to login to LoginRadius systems with an email link. Phone authentication also contains some information on passwordless logins.

To call a Passwordless Login API, import the `lrauthentication` package like so:

```go
import (
	lrauthentication "github.com/LoginRadius/go-sdk/api/authentication"
)
```

**List of APIs in this Section:**

- [GET: Passwordless Login By Email](#passwordless-login-by-email)
- [GET: Passwordless Login By Username](#passwordless-login-by-username)
- [GET: Passwordless Login Verification](#passwordless-login-verification)

##### Passwordless Login By Email

This API is used to send an email containing a link to start the passwordless login flow.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/passwordless-login/passwordless-login-by-email)

Example:

```go
res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).GetPasswordlessLoginByEmail(
  map[string]string{"email": <email>},
)

if err != nil {
  // handle error
}
```

##### Passwordless Login By Username

This API is used to send an email containing a link to start the passwordless login flow.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/passwordless-login/passwordless-login-by-username)

Example:

```go
res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).GetPasswordlessLoginByUsername(
  map[string]string{"username": <username>},
)
if err != nil {
  // handle error
}
}
```

##### Passwordless Login Verification

This API is to verify the passwordless login token and returns the user profile and access token.

[Documentation](https://www.loginradius.com/docs/api/v2/customer-identity-api/passwordless-login/passwordless-login-verification)

Example:

```go
res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).GetPasswordlessLoginVerification(
  map[string]string{"verificationtoken": <verification token>},
)

if err != nil {
  // handle error
}
```

