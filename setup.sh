mkdir -p $HOME/docker/volumes/crudapipg && \
docker network create crudapi-net && \
docker run --name crudapipg \
    -e POSTGRES_PASSWORD=hyher43edasdt45y34rweryu \
    -e POSTGRES_USER=crudapi \
    -e POSTGRES_DB=crudapi \
    --restart unless-stopped \
    -v $HOME/docker/volumes/crudapipg:/var/lib/postgresql/data \
    -p 5432:5432 \
    --net crudapi-net \
    -d postgres