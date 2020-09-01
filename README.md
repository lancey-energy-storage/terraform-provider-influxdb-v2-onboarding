# InfluxDB v2 onboarding provider

This provider is only used to setup the initial configuration on influxdbv2 startup.

The InfluxDB V2 provider allows Terraform to setup
[InfluxDB v2](https://www.influxdata.com/products/influxdb-overview/).

## How to use
### Download the provider 

#### Terraform 0.12.x 

Download the release and extract it to (on linux):
`~/.terraform.d/plugins/linux_amd64/terraform-provider-influxdbv2-onboarding_v0.1.0`


#### Terraform 0.13.x

Add this snippet to your code:

```hcl
terraform {
    required_providers {
        influxdbv2-onboarding {
            source = "lancey-energy-storage/influxdbv2-onboarding"
            version = "0.1.0"
        }
    }
}
```

Until the provider is available on registry.terraform.io, you need to manually download the release and extract it to, eg (on linux):
`~/.terraform.d/plugins/registry.terraform.io/lancey-energy-storage/influxdbv2-onboarding/0.1.0/linux_amd64/terraform-provider-influxdbv2-onboarding_v0.1.0`

### Initialize the provider 

```hcl
provider "influxdbv2-onboarding" {
  url = "http://influxdb.example.com:9999"
  username = "influxdbUsername"
  password = "InfluxdbPassword"
}
```
The provider configuration block accepts the following arguments:

* ``url`` (Optional) The root URL of a Influxdb V2 server. May alternativly be set via the INFLUXDB_V2_URL environment variable. Default to `http://localhost:9999`.
* ``username`` (Optional) The username that will be created as administrator. Defaults to `administrator`
* ``password`` (Optional) The password that will be set for the initial user. Defaults to `Administrator1.`

At first, you need to start a new Influxdb V2 instance. To do so, you can follow the official documentation [here](https://v2.docs.influxdata.com/v2.0/get-started/#start-with-influxdb-oss)

### Available functionalities

Documentation is available in [website/docs/](website/docs/).
Influxdb v2 api documentation is available [here](https://v2.docs.influxdata.com/v2.0/api/).

### Resources

* **setup** to setup initial user, bucket and organization, documentation [here](website/docs/r/setup.html.md)

### Examples
Find examples in `examples/`. To run them:

```bash
terraform init
terraform apply
```

## Dev

In case you need to update the influx client, run `go get github.com/lancey-energy-storage/influxdb-client-go@<commit sha>`.  
Also don't forget to run `go mod tidy` from time to time to remove useless dependencies.

## Test

First execute this command to check fmt requirements:
 
```bash
make fmt
```

Then execute this command to run provider unit tests: 

```bash
make test
```

And finally to run acceptance test, run this command: 

```bash
make testacc
```

## Build

```bash
go build -o terraform-provider-influxdb-v2-onboarding
```
