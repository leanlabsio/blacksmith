# Development

Development environment assumed to be portable and have as less external dependencies as possible.

### Dependencies

The following software is required to start development environment:

- Docker 1.9+
- GNU Make

### Start dev env

To start development environment you first should obtain
GitHub OAuth client id and secret. Then run the following command in project root dir.

```bash
GITHUB_CLIENT_ID=YOUR_CLIENT_ID GITHUB_CLIENT_SECRET=YOUR_CLIENT_SECRET make dev
```

That's it, now you have running dev env.
