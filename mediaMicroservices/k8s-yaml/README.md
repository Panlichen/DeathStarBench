# Media Microservices on Kubernetes

## Pre-requirements

- A running kubernetes cluster is needed.
- This repo should be put in the same path on every node of the cluster for volume mount (here we use `/root/DeathStarBench/`).
- Pre-requirements mentioned [here](https://github.com/delimitrou/DeathStarBench/blob/master/mediaMicroservices/README.md) should be met.

## Running the media service application on kubernetes

> NOTE: you need to modify the ip:port of nginx service in commands and files used below. Files you should pay attention to are:
> - `<path-of-repo>/mediaMicroservices/scripts/write_movie_info.py`
> - `<path-of-repo>/mediaMicroservices/scripts/register_users.sh`
> - `<path-of-repo>/mediaMicroservices/wrk2/scripts/media-microservices/compose-review-tiny.lua`

### Before you start

Set the resolver to be the FQDN of the core-dns or kube-dns  service of your cluster in `<path-of-repo>/mediaMicroservices/nginx-web-server/conf/nginx-k8s.conf`. If you do not deploy core-dns or kube-dns, I am not sure whether all things below still work.

### Deploy services

Run `kubectl apply -f <path-of-repo>/mediaMicroservices/k8s-yaml/` and wait `kubectl -n media-microsvc get pod` to show all pods with status `Running`.

### Register users and movie information

```
python3 <path-of-repo>/mediaMicroservices/scripts/write_movie_info.py
bash <path-of-repo>/mediaMicroservices/scripts/register_users.sh
 ```

### Running HTTP workload generator

#### Compose reviews

```bash
cd <path-of-repo>/mediaMicroservices/wrk2
./wrk -D exp -t <num-threads> -c <num-conns> -d <duration> -L -s ./scripts/media-microservices/compose-review-tiny.lua http://10.0.0.252:32000/wrk2-api/review/compose -R <reqs-per-sec>
```

- `compose-review-tiny.lua` uses a subset of movies, excluding problematic movies in the original dataset.

#### View Jaeger traces

 View Jaeger traces by accessing `http://10.0.0.252:31010` 
 > The ip and port depends on how you deploy your jaeger.