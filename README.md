# Blacksmith CD server

Blacksmith is a powerful continuous integration/delivery server.

### Why another CI/CD server and how it's different?

Blacksmith is based on simple idea - we have containers in production, in development, in testing,
we have data containers, application containers, executable containers, containers, containers, containers...

So why just not have builder containers? Why just not have full featured build pipelines in containers?

What if you could just pull docker container, containing the whole amount of instructions to build project of
any complexity and just run you build? Run you build locally, on cloud? Anywhere where docker runs?

Here comes Blacksmith. Blacksmith stores your build configurations, and triggers builds on incoming webhooks,
it is simple.

The full power comes with builder containers, check our [Docker builder](https://github.com/vasiliy-t/blacksmith-docker-builder),
it allows to build and publish docker containers to docker hub, and you'll get the idea.

### Usage

Blacksmith dependes on Redis to store users account information, build configuration and logs.

With docker:

```bash
docker run -d \
		-p 80:80 \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-e REDIS_ADDR=redis:6379 \
		-e GITHUB_CLIENT_ID=GITHUB_OAUTH_CLIENT_ID) \
		-e GITHUB_CLIENT_SECRET=GITHUB_OAUTH_CLIENT_SECRET \
		-e DOCKER_HOST=unix:///var/run/docker.sock \
		leanlabs/blacksmith:latest
```
