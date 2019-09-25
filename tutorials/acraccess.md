# ACR Access

This tutorial walks through the steps of accessing ACR during the build to push images. You can also use this method to access any other azure services (via Pod Identity)

The first thing we will do is install [AAD Pod Identity](https://github.com/Azure/aad-pod-identity)

```sh
kubectl apply -f https://raw.githubusercontent.com/Azure/aad-pod-identity/master/deploy/infra/deployment-rbac.yaml
```

Next we will create a resource group and an azure identity

```sh

export SUBSCRIPTION_ID=$(az account show --query="id" -o tsv)

az group create -n kciacr -l eastus

export CLIENT_ID=$(az identity create -g kciacr -n kcibuild -o tsv --query="clientId")
```

We will also create our ACR Instance

```sh

```

Next we will create the Azure Identity resource

```sh
cat <<EOF | kubectl apply -f -
apiVersion: "aadpodidentity.k8s.io/v1"
kind: AzureIdentity
metadata:
  name: kci-az-identity
spec:
  type: 0
  ResourceID: /subscriptions/${SUBSCRIPTION_ID}/resourcegroups/kciacr/providers/Microsoft.ManagedIdentity/userAssignedIdentities/kcibuild
  ClientID: ${CLIENT_ID}
EOF
```

Now we will create the AzureIdentityBinding object. Which will bind the identity to our build agent

```sh
cat <<EOF | kubectl apply -f -
apiVersion: "aadpodidentity.k8s.io/v1"
kind: AzureIdentityBinding
metadata:
  name: kci-azure-identity-binding
spec:
  AzureIdentity: kci-az-identity
  Selector: kcibuildpod
EOF
```