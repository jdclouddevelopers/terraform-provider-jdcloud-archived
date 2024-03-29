Terraform Provider for JDCloud
==================

[![Build Status](https://travis-ci.com/jdclouddevelopers/terraform-provider-jdcloud.svg?branch=master)](https://travis-ci.com/jdclouddevelopers/terraform-provider-jdcloud)

# *Reminders

Hi 👋! JDCloud is now one of the Terraform official providers. All files and documents 
has been moved to the [Our New Repository](https://github.com/terraform-providers/terraform-provider-jdcloud). 
This project has been deprecated. Any bugs or issues, please report it to [Our New Repository](https://github.com/terraform-providers/terraform-provider-jdcloud) 

If you have any questions, feel free to let us know 😊

- JDCloud-Team <ark@jd.com>
- Email <liangxiaohan@jd.com>


# Requirements

* Terraform 0.10+
* Go 1.12(to build the provider plugin)

# Using the provider 

* Prepare a folder for Terraform 
```bash
mkdir terraform && cd terraform 
```
* Prepare your `jdcloud.tf`. This is where Terraform begins 
```bash
touch jdcloud.tf
```
* Download [Terraform](https://www.terraform.io/downloads.html) 
* Download [Terraform-Jdcloud-Plugin](https://github.com/jdclouddevelopers/terraform-provider-jdcloud)
* Launch!
```bash
terraform init
``` 
Terraform is now start working, it will manage your resources according to your `jdcloud.tf`
We would recommend our users begin with some simple resource, say `VPC` and `ElasticIP`

How to write `jdcloud.tf` ? [Check Here](https://github.com/XiaohanLiang/terraform-provider-jdcloud/blob/master/example/main.tf)

# Developing the Provider

Contributions and advices to this plugin is very welcomed. In order to get start with, you 
need to do the following steps.

### 1.Prepare Golang Environment

First you will need to have [Golang1.12](https://golang.org/dl/) installed on your machine. Besides, 
You will need to correctly set up $GOPATH, as well as adding `$GOPATH/bin` to your `$PATH`

### 2.Compile project


After you have modified the code you can compile this plugin by `make build`. 
Remember to format your code by using `go fmt`. If it works fine. Plugin will be in your $GOPATH/bin
``` go
$ make build
==> Checking that code complies with gofmt requirements...
go install

$ ls $GOPATH/bin | grep jdcloud
terraform-provider-jdcloud
```   

### 3.Running Acceptance Test


Acceptance test can be an important part of developing process. Basically, it will first create a resource,
validate its attributes and see if it works as expected. Update this resource if applied and eventually delete this resource.
Acceptance tests are files in `jdcloud` with suffix `_test.go`. You can invoke an acceptance test by 
```bash
make testacc
```

_Note_ 
* Acceptance creates real resources, it will probably cost some money.
* Process usually takes 20~30 minutes depends on your network condition.

## License

Mozilla Public License 2.0

