# emailvalidator
* endpoint: /email/validate
* listen on PORT environment variable
* that accepts json `{"email":"xxx@yyy.zzz"}`
* returns json
```json
{
  "valid": false,
  "validators": {
    "regexp": {
      "valid": true
    },
    "domain": {
      "valid": false,
      "reason": "INVALID_TLD"
    },
    "smtp": {
      "valid": false,
      "reason": "UNABLE_TO_CONNECT"
    }
  }
}
```

* image on dockerhub
* should be able to run with 
```
docker run -t -p 127.0.0.1:8080:8080-e PORT=8080 ksaritek/emailvalidator
``` 

example request:
curl -XPOST -d '{"email":"xxx@yyy.zzz"}' http://localhost:8080/email/validate