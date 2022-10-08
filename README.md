# urlscan
<div align="center">
A <a href="https://urlscan.io">urlscan.io</a> API library for Go

![urlscan logo](assets/logo.png)
</div>

## Features
* Scanning URLs with urlscan
* Getting scan results
* Searching for scan results

## Usage

### Initializing the Client
To get an API Key, [sign up for urlscan](https://urlscan.io/user/signup/) and go to the "Settings & API" profile menu
```go
    client := urlscan.NewClient("MY-API-KEY")
```

### Scanning
```go
scan, err := client.Scan("https://example.com", urlscan.ScanOptions{
		Tags:    []string{"test", "urlscan-go"},
		Country: "ca",
	})
```

### Getting a result
```go
result, err := client.GetResult("result-uuid-here")
```

### Searching
```go
result, _ := client.Search("domain:(example.com OR example.net)", 100)

for _, result := range result.Results {
	// Do something here
}
```

