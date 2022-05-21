# GithubStatsGenerator

Generate github stats in SVG using Go. 

## How To Get Started

1. Create Github Token [here](https://github.com/settings/tokens). Not sure how big the scopes should be, I just select all scopes to make sure the token allows me to do anything. 
2. Create a target repo if you don't have one already. If the repo doesn't exist, it won't work.
3. Using docker, do:

```docker
docker run --name githubstatsgenerator -e GH_TOKEN=someGithubToken -e REPO=mrkresnofatih ghcr.io/mrkresnofatih/ghcr.io/mrkresnofatih/ghstatsgenerator:latest
```

## Releases

<img src="https://img.shields.io/badge/image-latest-green">