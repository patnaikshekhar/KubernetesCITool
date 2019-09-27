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
export ACR_NAME=<unique acr name>

az acr create -n $ACR_NAME -g kciacr --sku basic -l eastus
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

We now need to set permissions for MIC. The MIC uses the service principal credentials stored in the cluster to access Azure resources. This service principal needs Microsoft.ManagedIdentity/userAssignedIdentities/\*/assign/action permission on the identity to work with user-assigned MSI.

Lets get the service principal ID first

```sh
export SP_ID=$(az aks show -n <Your AKS Cluster Name> -g <Your AKS Cluster Group> \
  --query=servicePrincipalProfile.clientId -o tsv)
```
