# Working with Private GitHub Repositories

In order to work with private github repositories you need to create a [deploy key](https://developer.github.com/v3/guides/managing-deploy-keys/).

Lets start by creating an ssh key

```sh
ssh-keygen -t rsa -b 4096 -C "your_email@example.com"
```

Save the key in any location. Now lets upload the key as a secret

```sh
kci secret <Location of Key>
```

Now we can create a [private repo](https://help.github.com/en/articles/setting-repository-visibility) in GitHub. After the repository is created navigate to Settings -> Deploy Keys and paste the value of your public key. Make sure to check the "Allow Push Access" checkbox.

We can now create the build file which will clone the repo and run steps

```yaml
repository: git@github.com:example/privatexample.git
repositorySSHKeySecret: private.key 
steps:
- image: alpine
  args:
    - echo
    - Hello
```

The **repositorySSHKeySecret** field specifies the name of the private key file we uploaded as a secret.

We can now run the build

```sh
kci <path to build file>
```