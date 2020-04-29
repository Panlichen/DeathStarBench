---
note:
    createdAt: 2020-03-30T07:46:25.854Z
    modifiedAt: 2020-04-13T12:00:41.084Z
    tags: []
---
# Deploy jaeger

Refer to [here](https://www.jaegertracing.io/docs/1.17/operator/).

## Deploy jaeger operator:

```
kubectl apply -f /home1/root/DeathStarBench/deploy-jaeger/crds/jaegertracing.io_jaegers_crd.yaml
kubectl apply -f /home1/root/DeathStarBench/deploy-jaeger
```

Configurations:

1. Customize the `operator.yaml`, setting the env var `WATCH_NAMESPACES` to have an empty value, so that it can watch for instances across all namespaces.

The operator creates a Kubernetes [`ingress`](https://kubernetes.io/docs/concepts/services-networking/ingress/) route, which is the Kubernetes’ standard for exposing a service to the outside world. Once Ingress is enabled, the address for the Jaeger UI can be found by querying the Ingress object:

```
kubectl get ingress
```



## Deploy jaeger instances

### All-in-one, in memory

```
kubectl apply -f /home1/root/DeathStarBench/deploy-jaeger/crds/jaeger-all-in-one-inmemory.yaml
```

### Configurations

- [Auto-injecting Jaeger Agent Sidecars](https://www.jaegertracing.io/docs/1.17/operator/#auto-injecting-jaeger-agent-sidecars)
- [Deployment Strategies](https://www.jaegertracing.io/docs/1.17/operator/#deployment-strategies)
- [Storage options](https://www.jaegertracing.io/docs/1.17/operator/#storage-options)

# Deploy ingress--not necessary

Refer to [here](https://github.com/kubernetes/ingress-nginx/blob/master/docs/deploy/index.md#bare-metal).

1. Deploy Mandatory stuffs:

   `kubectl apply -f deploy-ingress/static/mandatory.yaml`

2. Deploy nodePort service:

   `kubectl apply -f deploy-ingress/baremetal/service-nodeport.yaml`

   > Before this, may need to configure `spec.ports[0].nodePort` in `deploy-ingress/baremetal/service-nodeport.yaml`

   > Need to configure `spec.selector` or define `Endpoint` object.(refer to [here](https://kubernetes.io/zh/docs/concepts/services-networking/service/#services-without-selectors))
