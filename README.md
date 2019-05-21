# Golang package for signing Aliyun API gateway requests

Same algorithm as the following 'official' demos in other language. 

- https://github.com/aliyun/api-gateway-demo-sign-python
- https://github.com/aliyun/api-gateway-demo-sign-net
- https://github.com/aliyun/api-gateway-demo-sign-php

Compliance is not fully guaranteed, since I have not found the official documentation on this algorithm. Use at your own risk.

## Installation

```
go get -u github.com/HuanTeng/go-aliyun-sign
```

## Usage

```golang
import "github.com/HuanTeng/go-aliyun-sign"

// Prepare a Request
req, _ := http.NewRequest("POST", url, body)
// Set headers
req.Header.Set("Content-Type", "application/json")
req.Header.Set("Accept", "application/json")
// Sign the request
if err := sign.Sign(req, appKey, appSecret); err != nil {
    // Handle error
}
// Do the request round-trip
resp, err := http.DefaultClient.Do(req)
```
