# external_endpoint_mocking_example_golang
 GoLang Example of how to mock clients that would call external endpoints during production

# Run (as you would in production)

```shell
go run .
```

This outputs the below to the console.  The `Phoned home!` means the client called the _real_ endpoint.  This should only happen in production.
```json
Phoned home! &{{0x14000112020 app123}}
{'finaloutput':{'externaldata':'27'}}
```

# Test (as you would in development)

```shell
go test --v
```

This outputs the below to the console.  Notice the `Phoned home!` message does not get printed because the mock client does not call the real endpoint, it only returns a response that simulates/mocks what it would return if it really did call the endpoint like the real client does.  This should only happen in production.
```json
=== RUN   Test_main
--- PASS: Test_main (0.00s)
PASS
ok      external_endpoint_mocking_example       0.196s

```