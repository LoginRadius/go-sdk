# LoginRadius Go SDK Demo

![Home Image](http://docs.lrcontent.com/resources/github/banner-1544x500.png)

## Installation

Add LoginRadius API Credentials in the following file, You can obtain credentials from [here](https://www.loginradius.com/docs/api/v2/admin-console/platform-security/api-key-and-secret/#api-key-and-secret)

 - Add the API credentials in `secret.env.example` under `demo/config`  and rename the file name to `secret.env`
 - Add the API credentials in `options.sample.js` under `demo/ui/assets/js` and rename the file name to `options.js`


In `demo/cmd/app/` directory 
``
Run the command 
```
"go mod init demo/cmd/app"
```

To install, run: 
```
go get github.com/loginradius/go-sdk
```

Install all package dependencies by running `go get ./...`.

After installation of all dependencies run `go run main.go` in the ./demo/cmd/app/ directory

You can view the demo on the [http://localhost:3000/](http://localhost:3000/)