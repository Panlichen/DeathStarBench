helm template hubble \
    --namespace kube-system \
    --set metrics.enabled="{dns:query;ignoreAAAA,tcp,flow:destinationContext=pod;sourceContext=pod,port-distribution:destinationContext=pod;sourceContext=pod,icmp,http}" \
    > install-hubble.yaml