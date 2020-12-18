# Xendint API go test code

## Credentials

```
cd cmd
cp test.sample.env test.secret.env 
# then edit test.secret.env to include your secret token

. test.secret.env
```

## Test code
1. cmd/disbursement; ID={invoice ID} go run main.go
1. cmd/virtual-account; go run main.go
