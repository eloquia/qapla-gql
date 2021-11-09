# Docker Commands

Over time Docker may consume too much space, either through unused layers, dangling containers, etc.

List volumes

```bash
docker volume ls
```

List dangling volumes

```bash
docker volume ls -qf dangling=true
```

Remove dangling volumes

```bash
docker volume rm $(docker volume ls -qf dangling=true)
```
