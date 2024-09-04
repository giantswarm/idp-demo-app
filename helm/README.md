# [[image-name]]

Chart version 0.1.0, app version 0.1.0

A Helm chart to deploy the sample golang based web service project.

**Homepage:** <https://github.com/[[repo-owner]]/[[image-name]]/>

## Requirements

Kubernetes: `>=1.25.0`

## Values

Use the values below to configure the chart's values.
| Key | Type | Default | Description |
|-----|------|---------|-------------|
| autoscaling.enabled | bool | `true` | Turn on Pod replicas number autoscaling instead of setting a constant value. your cluster must support [ Horizontal Pod Autoscaling ](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/). |
| autoscaling.maxReplicas | int | `10` | Max number of Pods autoscaler can deploy. |
| autoscaling.minReplicas | int | `3` | Min number of Pods autoscaler can deploy. |
| autoscaling.targetCPUUtilizationPercentage | int | `80` | Pod scale up critieria based on CPU usage. |
| autoscaling.targetMemoryUtilizationPercentage | int | `80` | Pod scale up critieria based on Memory usage. |
| fullnameOverride | string | `""` | Override the default name generated for this specific chart Release. |
| image.pullPolicy | string | `"IfNotPresent"` | Configure image pull policy. |
| image.registry | string | `"[[registry-domain]]"` | Set the domain of your container images registry. |
| image.repository | string | `"[[registry-name]]/[[image-name]]"` | Set the name of the repository within the registry. |
| imagePullSecrets | list | `[]` | Configure login secrets for the container images registry. |
| ingress.annotations | object | `{}` | Optional annotations for the Ingress definition. If your cluster has "CertManager" operator running, you can use "cert-manager.io/cluster-issuer" annotation to [automatically generate a certificate for it](https://cert-manager.io/docs/usage/). |
| ingress.enabled | bool | `true` | Should the Service be accessible through an Ingress. This needs an Ingress controller to be configured already on your cluster. |
| ingress.host | string | `"chart-example.local"` | HTTP host that you want to use for your service. |
| ingress.tls | list | `[]` | Optional TLS certificate configuration. You can use it with "CertManager" or provide your own certificate. |
| monitoring.serviceMonitor | object | `{"enabled":true}` | If your cluster supports prometheus-operator configuration of metrics data, enable this to have metrics from your application automatically ingested by prometheus. |
| nameOverride | string | `""` | Override the default name generated for the chart objects. |
| nodeSelector | object | `{}` | Optional node delector to limit the nodes where pods of the chart can be deployed. |
| pdb | object | `{"enabled":true}` | Should the chart deploy a [PodDisruptionBudget](https://kubernetes.io/docs/tasks/run-application/configure-pdb/) to limit disruptions based on administrative tasks. |
| podAnnotations | object | `{}` | Set additional annotations for the pods created. |
| podListenPort | int | `8080` | Configure the TCP port on which your pods will listen for connections. |
| replicaCount | int | `3` | Number of Pod replicas to deploy. Used only if 'autoscaling.enabled' is 'false'. |
| resources | object | `{"limits":{"cpu":"500m","memory":"512Mi"},"requests":{"cpu":"100m","memory":"128Mi"}}` | Configure [Pod resources](https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/). |
| service.annotations | object | `{}` | Optional annotations for the Service definition. If your cluster has "ExternalDNS" operator running, you can use "external-dns.alpha.kubernetes.io/hostname" annotation to [automatically register DNS name for your service](https://github.com/kubernetes-sigs/external-dns). |
| service.port | int | `80` | TCP port that the service will be exposed on. |
| service.type | string | `"ClusterIP"` | The type of [ Service ](https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types). |
| serviceAccount.annotations | object | `{}` | Annotations to add to the service account. |
| serviceAccount.create | bool | `true` | Specifies whether a service account should be created. |
| serviceAccount.name | string | `""` | The name of the service account to use. If not set and create is true, a name is generated using the fullname template |

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| [[repo-owner]] | <noemail@nothing.com> |  |

