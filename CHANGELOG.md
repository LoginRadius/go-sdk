
> **LoginRadius  Golang  SDK Change Log** provides information regarding what has changed, more specifically what changes, improvements and bug fix has been made to the SDK. For more details please refer to the [LoginRadius API Documention](https://www.loginradius.com/docs/libraries/sdk-libraries/golang-library/)

# Version 11.4.0
-  We have introduced connection pooling in the Go Default Http client to keep fewer connections open and it will support more requests with minimal server resources.
- Enhancement in README.md file.

## Added new multiple APIs for better user experience
- GetRevokeRefreshToken
- GetRefreshAccessTokenByRefreshToken

## Breaking Changes

For developers migrating from v11.3.0, there will be some minor breaking changes in terms of SDK implementation as mentioned below.

- In this version, we have added additional param `timeDifference`, `startTime` and `endTime` into manual SOTT generate method `Generate()`.

- We have standardize `PostAuthUserRegistrationByEmail` and `PostPhoneUserRegistrationBySMS` Api and now SOTT will be passed explicity as an function parameter.




# Version 11.3.0
- Added Contribution Guideline file.
- Updated jquery version 3.3.1 with latest version 3.6.0

# Version 11.3.0-beta
- Added ``Licence.md`` file



See the documentation [here](https://www.loginradius.com/docs/libraries/sdk-libraries/golang-library/)
