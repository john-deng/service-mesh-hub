<!-- Code generated by protoc-gen-solo-kit. DO NOT EDIT. -->

## Package:
supergloo.solo.io

## Source File:
install.proto 

## Description:  

## Contents:
- Messages:  
	- [Install](#Install)  
	- [HelmChartLocator](#HelmChartLocator)  
	- [HelmChartPath](#HelmChartPath)

---
  
### <a name="Install">Install</a>

Description: 

```yaml
"status": .core.solo.io.Status
"metadata": .core.solo.io.Metadata
"istio": .supergloo.solo.io.Istio
"linkerd2": .supergloo.solo.io.Linkerd2
"consul": .supergloo.solo.io.Consul
"chartLocator": .supergloo.solo.io.HelmChartLocator
"encryption": .supergloo.solo.io.Encryption
"enabled": .google.protobuf.BoolValue

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| status | [.core.solo.io.Status](install.proto.sk.md#Install) | Status indicates the validation status of this resource. Status is read-only by clients, and set by gloo during validation |  |
| metadata | [.core.solo.io.Metadata](install.proto.sk.md#Install) | Metadata contains the object metadata for this resource |  |
| istio | [.supergloo.solo.io.Istio](install.proto.sk.md#Install) |  |  |
| linkerd2 | [.supergloo.solo.io.Linkerd2](install.proto.sk.md#Install) |  |  |
| consul | [.supergloo.solo.io.Consul](install.proto.sk.md#Install) |  |  |
| chartLocator | [.supergloo.solo.io.HelmChartLocator](install.proto.sk.md#Install) |  |  |
| encryption | [.supergloo.solo.io.Encryption](install.proto.sk.md#Install) |  |  |
| enabled | [.google.protobuf.BoolValue](install.proto.sk.md#Install) | whether or not this install should be enabled if disabled, corresponding resources will be uninstalled defaults to true |  |
  
### <a name="HelmChartLocator">HelmChartLocator</a>

Description: 

```yaml
"chartPath": .supergloo.solo.io.HelmChartPath

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| chartPath | [.supergloo.solo.io.HelmChartPath](install.proto.sk.md#HelmChartLocator) | path to a local directory, local tar.gz, or url tar.gz |  |
  
### <a name="HelmChartPath">HelmChartPath</a>

Description: 

```yaml
"path": string

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| path | string |  |  |

