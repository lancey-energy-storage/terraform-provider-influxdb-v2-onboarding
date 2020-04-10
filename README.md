# terraform-provider-influxdb-v2-onboarding
A terraform provider for influxdb v2 onboarding step specifically

This provider is only used to setup the initial configuration on influxdbv2 startup.

The InfluxDB V2 provider allows Terraform to setup
[InfluxDB v2](https://www.influxdata.com/products/influxdb-overview/).

The provider configuration block accepts the following arguments:

* ``url`` (Optional) The root URL of a Influxdb V2 server. May alternativly be set via the INFLUXDB_V2_URL environment variable. Default to `http://localhost:9999`.
* ``username`` (Optional) The username that will be created as administrator. Defaults to `administrator`
* ``password`` (Optional) The password that will be set for the initial user. Defaults to `Administrator1.`

## Build

```bash
go build -o terraform-provider-influxdb-v2-onboarding
```

Don't forget to copy `terraform-provider-influxdbv2` to your terraform plugin directory (eg. `~/.terraform.d/plugins/linux_amd64` on linux).

## Test

```bash
go test ./influxdbv2-onboarding
```

## How to use

### Initialize the provider
```hcl
provider "influxdbv2-onboarding" {
  url = "http://influxdb.example.com:9999"
  username = "influxdbUsername"
  password = "influxdbPassword"
}
 ```

### Available functionalities

* **setup** to setup initial user, bucket and organization, documentation [here](website/docs/r/setup.html.md)

### Examples file
Find more examples in `examples/`. To run them:
```bash
terraform init
terraform apply
```

## Dev

In case you need to update the influx client, run `go get github.com/lancey-energy-storage/influxdb-client-go@<commit sha>`.  
Also don't forget to run `go mod tidy` from time to time to remove useless dependencies.