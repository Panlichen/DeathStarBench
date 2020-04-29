helm template hubble \
    --namespace kube-system \
    --set metrics.enabled="{dns:query;ignoreAAAA,tcp,flow:destinationContext=pod-short;sourceContext=pod-short,port-distribution:destinationContext=pod-short;sourceContext=pod-short,icmp,http}" \
    > install-hubble.yaml