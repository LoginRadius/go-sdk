# LoginRadius Go SDK


![Home Image](http://docs.lrcontent.com/resources/github/banner-1544x500.png)

## Introduction ##

LoginRadius Customer Registration wrapper provides access to LoginRadius Identity Management Platform API.

LoginRadius is an Identity Management Platform that simplifies user registration while securing data. LoginRadius Platform simplifies and secures your user registration process, increases conversion with Social Login that combines 30 major social platforms, and offers a full solution with Traditional Customer Registration. You can gather a wealth of user profile data from Social Login or Traditional Customer Registration. 

LoginRadius centralizes it all in one place, making it easy to you ,manage and access. Easily integrate LoginRadius with all of your third-party applications, like MailChimp, Google Analytics, Livefyre and many more, making it easy to utilize the data you are capturing.

LoginRadius helps businesses boost user engagement on their web/mobile platform, manage online identities, utilize social media for marketing, capture accurate consumer data, and get unique social insight into their customer base.

Please visit [here](http://www.loginradius.com/) for more information.


## Contents ##

* [Demo](https://github.com/LoginRadius/go-sdk/tree/master/demo) - A demo of LoginRadius user auth workflow.

## Documentation

* [Configuration](https://docs.loginradius.com/api/v2/deployment/sdk-libraries/golang-library) - Everything you need to Know to begin using the LoginRadius SDK.

## Installation

To install, run:
`go get github.com/loginradius/go-sdk`

Import the package:

`import "github.com/loginradius/go-sdk"`

Install all package dependencies by running `go get ./...` in the root folder of this SDK.  

## Usage

Take a peek:

Before making any API calls, the LoginRadius API client must be initialized with your Loginradius API key and API secret.

Sample code:

```
cfg := lr.Config{
    ApiKey:    <your API key>,
    ApiSecret: <your API secret>,
}

lrclient, err := lr.NewLoginradius(&cfg)

if err != nil {
    // handle error
}
```
