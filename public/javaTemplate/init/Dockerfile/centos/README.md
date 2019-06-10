```
docker build -t my/centos:latest .
docker run --rm  my/centos
```
#### docker网络设置
```bash
vi /etc/docker/daemon.json
{
  "dns" : [
    "223.5.5.5",
    "223.6.6.6",
    "8.8.8.8",
    "8.8.4.4"
  ]
}
```
