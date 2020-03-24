# Social Network Microservices on Kubernetes

## Pre-requirements

- A running kubernetes cluster is needed.
- This repo should be put in the same path on every node of the cluster for volume mount (here we use `/root/DeathStarBench/`).
- Pre-requirements mentioned [here](https://github.com/delimitrou/DeathStarBench/blob/master/socialNetwork/README.md) should be met.

## Running the social network application on kubernetes

> NOTE: you need to modify the ip:port of nginx service in commands and files used below. Files you should pay attention to are:
> - ` <path-of-repo>/socialNetwork/scripts/init_social_graph.py`
> - `<path-of-repo>/socialNetwork/wrk2/scripts/social-network/compose-post.lua`
> - `<path-of-repo>/socialNetwork/wrk2/scripts/social-network/read-home-timeline.lua`
> - `<path-of-repo>/socialNetwork/wrk2/scripts/social-network/read-user-timeline.lua`

### Before you start

Set the resolver to be the FQDN of the core-dns or kube-dns  service of your cluster in `<path-of-repo>/socialNetwork/nginx-web-server/conf/nginx-k8s.conf` and `<path-of-repo>/socialNetwork/media-frontend/conf/nginx-k8s.conf`. If you do not deploy core-dns or kube-dns, I am not sure whether all things below still work.

### Deploy services

Run `kubectl apply -f <path-of-repo>/socialNetwork/k8s-yaml/` and wait `kubectl -n social-network get pod` to show all pods with status `Running`.

### Register users and construct social graphs

Register users and construct social graph by running `python3 <path-of-repo>/socialNetwork/scripts/init_social_graph.py`. This will initialize a social graph based on [Reed98 Facebook Networks](http://networkrepository.com/socfb-Reed98.php), with 962 users and 18.8K social graph edges. 

### Running HTTP workload generator

#### Compose posts

```bash
cd <path-of-repo>/socialNetwork/wrk2
./wrk -D exp -t <num-threads> -c <num-conns> -d <duration> -L -s ./scripts/social-network/compose-post.lua http://10.0.0.252:31000/wrk2-api/post/compose -R <reqs-per-sec>
```

#### Read home timelines

```bash
cd <path-of-repo>/socialNetwork/wrk2
./wrk -D exp -t <num-threads> -c <num-conns> -d <duration> -L -s ./scripts/social-network/read-home-timeline.lua http://10.0.0.252:31000/wrk2-api/home-timeline/read -R <reqs-per-sec>
```

#### Read user timelines

```bash
cd <path-of-repo>/socialNetwork/wrk2
./wrk -D exp -t <num-threads> -c <num-conns> -d <duration> -L -s ./scripts/social-network/read-user-timeline.lua http://10.0.0.252:31000/wrk2-api/user-timeline/read -R <reqs-per-sec>
```

#### View Jaeger traces

 View Jaeger traces by accessing `http://10.0.0.252:31010` 
 > The ip and port depends on how you deploy your jaeger.

## TODO

We need a more configurable way to config service ip and port. I am not familiar with lua, and I just leave it there. You are really welcome to submit a pull request!

## Questions and contact

If you have questions or advices about the k8s-yamls, contact me at myles.l.pan@gmail.com