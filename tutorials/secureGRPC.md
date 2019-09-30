# Adding TLS to the Build Server

This tutorial walks you through the process of adding TLS support to the build server. In this example we create TLS certificates first and then create a secret which is mounted to the application.

Lets create certificates first.

```sh
# Create Private Key
openssl genrsa -out server.key 2048

# Create Certificate use the hostname of your server as the CN
openssl req -new -x509 -sha256 -key server.key \
    -out server.crt -days 3650
```

Now, we can add these certs to a secret.

```sh
kubectl create secret generic kci-server-secret \
    --from-file=server.key \
    --from-file=server.crt
```

Next we can mount the secret as a volume to the pods.

```sh
kubectl apply -f k8s/kci-server-with-secrets.yaml
```

Now we need to configure the client to expect the certificate

```sh
kci config cert $(pwd)/server.crt
```

