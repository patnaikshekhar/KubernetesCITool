# Kubernetes CI (KCI)

KCI is a simple build tool that allows you to run builds in your Kubernetes
cluster using a simple build file format.


Here is an example of a sample build file:


```yaml
repository: https://github.com/patnaikshekhar/KubernetesCIToolExample
steps:
- image: golang:1.13.0-alpine3.10
  args: ["go", "test"]
- image: golang:1.13.0-alpine3.10
  args: ["go", "build", "-o", "app"]
```

The build file consists of a series of steps that will execute sequentially. Each step
consists of an image and a list of args. In the above example, code is
downloaded from the repository mentioned in the repository parameter. The golang
image is downloaded and then `go test` is executed. Then in step 2 again the
golang image is used and go build is executed. 

## Usage

KCI can be installed in your organization by running

```sh

git clone https://github.com/patnaikshekhar/KubernetesCITool

kubectl apply -f k8s/kci-server.yaml
```

You then need to download the CLI from ???

The CLI needs to be configured with the following commands:

```sh
KCI_IP=$(kubectl get svc kci-server -n kci -o jsonpath="{.status.loadBalancer.ingress[0].ip}")

./kci config url KCI_IP:10000
```

Now you can run a build by running

```sh
./kci build ./examples/simple.yaml
```

You can also integrate KCI will GitHub so that builds are started automatically when you commit your code. All you need to do is add a Web Hook to github that points to the IP/Hostname of the KCI service running in Kubernetes.

## Coming Soon
- Pod Identity Support - In Progress
- User login / namespacing
