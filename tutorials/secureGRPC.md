# Adding TLS to the Build Server

This tutorial walks you through the process of adding TLS support to the build server. In this example we create TLS certificates first and then create a secret which is mounted to the application.

Lets create certificates first.

```sh
# Create Private Key
openssl genrsa -out server.key 2048

# Create Certificate
openssl req -new -x509 -sha256 -key server.key \
    -out server.crt -days 3650
```

Now, we can add these certs to a secret.

```sh
kubectl create secret generic kci-server-secret \
    --from-file=server.key \
    --from-file=server.crt
```

