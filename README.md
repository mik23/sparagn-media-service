### Requirements
- Install go1.15.1 darwin/amd64 ( see https://golang.org/doc/install)
- Make sure the file is present under the directory resources/GCP/credentials/google.json

#### How to test
go test ./...

go clean
go build app.go
./app (or "go run app.go")

### Configuring keys to deploy from CircleCi to Cloud

Generate a private key with (user sparagn):

```
ssh-keygen -m pem
```

Copy the private ssh key

```
less ~/.ssh/id_rsa | pbcopy
```

Note that if you have the private key is not in "PEM Format" use the following command:

```
ssh-keygen -p -N "" -m pem -f /path/to/key
```

Paste SSH private key copied in memory before, to the circleCI project in "Project Settings" -> "SSH Keys" -> Additional SSH Keys
Add the sparagn public key into the machine as ssh-trusted in GCP VM machine console.


