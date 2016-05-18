# Development

There are two ways to run the development environment - based on docker, or install required software on your machine.

### Running with docker

The following software is required to start development environment:

- Docker 1.9+
- GNU Make

To start development environment you 

1. first should obtain GitHub OAuth client id and secret. 

2. Then run the following command in project root dir:

```bash
GITHUB_CLIENT_ID=YOUR_CLIENT_ID GITHUB_CLIENT_SECRET=YOUR_CLIENT_SECRET make dev
```

That's it, now you have running dev env and could start writing code.

### Running without docker

The following software is required to start development environment:

- Node.js + npm
- Golang 1.6 

To start development environment you 

1. first should obtain GitHub OAuth client id and secret. 

2. Start Redis server instance.

2. Then run the following command in project root dir:

```
npm run build

go get -u github.com/jteeuwen/go-bindata/...

go-bindata -debug -pkg=templates -o templates/templates.go templates/...

go-bindata -debug -pkg=web -o web/web.go web/...

go run -v main.go daemon --redis-addr REDIS_SERVER_IP:REDIS_SERVER_PORT 
--docker-host unix:///var/run/docker.sock --github-client-id YOUR_CLIENT_ID 
--github-client-secret YOUR_CLIENT_SECRET

npm run watch
```

That's it, now you should have running dev env and could start writing code.
