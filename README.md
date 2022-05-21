# GithubStatsGenerator

Generate github stats in SVG using Go. They generate 4 SVGs:

![my toplangs](https://github.com/mrkresnofatih/mrkresnofatih/blob/main/generated/myTopLanguages.svg)
![my recentrepos](https://github.com/mrkresnofatih/mrkresnofatih/blob/main/generated/myRecentRepositories.svg)
![my recentlangs](https://github.com/mrkresnofatih/mrkresnofatih/blob/main/generated/myRecentLanguages.svg)
![my contributions](https://github.com/mrkresnofatih/mrkresnofatih/blob/main/generated/myRecentContributions.svg)

## How To Get Started

1. Create Github Token [here](https://github.com/settings/tokens). Not sure how big the scopes should be, I just select all scopes to make sure the token allows me to do anything. 
2. Create a target repo if you don't have one already. If the repo doesn't exist, it won't work.
3. Using docker, inject them as env. variables like so:

```docker
docker run --name githubstatsgenerator -e GH_TOKEN=someGithubToken -e REPO=existingRepoName ghcr.io/mrkresnofatih/ghcr.io/mrkresnofatih/ghstatsgenerator:latest
```

4. You probably want this to be ran in some cron tiggers. I'm currently testing this with `Github Actions`.

## Releases

<img src="https://img.shields.io/badge/image-latest-green">