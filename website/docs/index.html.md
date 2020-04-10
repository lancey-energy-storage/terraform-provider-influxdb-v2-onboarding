---
layout: "influxdbv2-onboarding"
page_title: "Provider: InfluxDB V2 Onboarding"
sidebar_current: "docs-influxdbv2-onboarding-index"
description: |-
  The InfluxDB V2 Onboarding provider configures the onboarding on an InfluxDB V2 server.
---

# InfluxDB V2 Provider Onboarding

The InfluxDB V2 provider allows Terraform to manage
[InfluxDB v2](https://v2.docs.influxdata.com/v2.0/get-started/).

The provider configuration block accepts the following arguments:

* ``url`` (Optional) The root URL of a InfluxDB V2 server. May alternatively be set via the `INFLUXDB_V2_URL` environment variable. Defaults to `http://localhost:9999/`.

* ``username`` (Optional) The username of the Influxdb V2 account. May alternatively be set via the `INFLUXDB_V2_USERNAME` environment variable.

* ``password`` (Optional) The password of the Influxdb V2 account. May alternatively be set via the `INFLUXDB_V2_PASSWORD` environment variable.

## Example Usage

```hcl
provider "influxdbv2" {
  url = "http://influxdb.example.com:9999"
  username = "influxdbUsername"
  password = "influxdbPassword"
}
 ```
