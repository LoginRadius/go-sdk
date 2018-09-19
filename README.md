# LoginRadius Go SDK


![Home Image](http://docs.lrcontent.com/resources/github/banner-1544x500.png)

## Introduction ##

LoginRadius Customer Registration wrapper provides access to LoginRadius Identity Management Platform API.

LoginRadius is an Identity Management Platform that simplifies user registration while securing data. LoginRadius Platform simplifies and secures your user registration process, increases conversion with Social Login that combines 30 major social platforms, and offers a full solution with Traditional Customer Registration. You can gather a wealth of user profile data from Social Login or Traditional Customer Registration. 

LoginRadius centralizes it all in one place, making it easy to manage and access. Easily integrate LoginRadius with all of your third-party applications, like MailChimp, Google Analytics, Livefyre and many more, making it easy to utilize the data you are capturing.

LoginRadius helps businesses boost user engagement on their web/mobile platform, manage online identities, utilize social media for marketing, capture accurate consumer data, and get unique social insight into their customer base.

Please visit [here](http://www.loginradius.com/) for more information.


## Contents ##

* [Demo](https://github.com/LoginRadius/go-sdk/tree/master/demo): It contains a demo of LoginRadius Flow.


## Documentation

* [Configuration](https://docs.loginradius.com/api/v2/sdk-libraries/golang) - Everything you need to begin using the LoginRadius SDK.



# LoginRadius

Go wrapper for the LoginRadius API. Get social graph information and send messages using LoginRadius' many social network clients!

## Installation

Run the following command:

``` 
go get -u github.com/loginradius/go-sdk
```

And then on any code that uses the sdk, include the following line as an import:
```
import "github.com/loginradius/loginradiusgo"
```

This package also uses pbkdf2 encryption from the Golang crypto libraries. To add them, run the following command:
```
go get -u golang.org/x/crypto/pbkdf2
```

## Usage

Take a peek:
```
type LoginStruct struct {
    Email string
    Password string
}

func GetAccessToken() string {
    loginradius.SetLoginRadiusEnv("MyAPIKey", "MyAPISecret", "https://api.loginradius.com")
    login := LoginStruct{"example@example.com", "password"}
    sessionObj, err := loginradius.PostAuthLoginByEmail("", "", "", "", "", login)
    if err != nil {
        // handle the error
    }
    return sessionObj.AccessToken
}
```
		