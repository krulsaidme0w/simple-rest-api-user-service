# 1st golang pet 
#### _simple user service_

Install and run (you can change user storage)
```sh
docker-compose -f docker-compose.yml 
```

My service has 2 different user storages (u can change them in docker-compose.yml):
* s3 storage (minio)
* file storage (store file for every user in local fs)

Also my service works nice to get user with username or name, for this i use cache. It helps to get user with logN complexity. Its filling up from storage when server is starting.

Technologies i was working with:

api:
* swagger

container:
* docker (Dockerfile, docker-compose)

go:
* Hexagonal Architecture
* dig (uber-go/dig)
* fasthttp
* s3 minio
* read/write from/to local files
