# Bitbucket CLI

A command line tool to help access the Bitbucket API. It brings Pipelines, pull requests and other Bitbucket features to the terminal. All which come with command completion

# Downloads 

[Releases](https://github.com/ysph-tech/bitbucket-cli/releases)


# Features

## Getting started

After installation you need to add your username (not email) and bitbucket generated password in your config file `%HOMEDIR%/.bitbucketcmd.json`

[Create password](https://bitbucket.org/account/settings/app-passwords/)

```json
{
  "password": "",
  "user": ""
}
```
or via repo access token
[https://bitbucket.org/$REPO/admin/access-tokens](https://bitbucket.org/$REPO/admin/access-tokens)
```json
{
    "token":""
}
```
Permission should be to have `READ/WRITE` for pipelines, pull-requests

https://developer.atlassian.com/cloud/bitbucket/rest/intro/#authentication


## Overview

All commands that run on repos will automaticlly asign the repo to current directory remote origin. Other wise you can use the --repo flag to set a repo

## Pipelines

### List Pipelines

List pipelines 

**Usage**

```bash
bitbucket pipelines
bitbucket pipelines -p 2
```
**Flags**

| Name | Description | Default |
| ---- | ------------ | ------- |
| `--page` `-p` |  Current page pagination | 1 |

### Pipeline details

Fetches pipeline details
Usage

```bash
bitbucket pipelines $PIPELINE_ID
```
**Example**

```bash
bitbucket pipelines 12 
bitbucket pipelines 12 -d #For detailed view including script 
```

You can use the pipeline uuid or build number

| Name | Description | Default |
| ---- | ------------ | ------- |
| `--detailed` `-d` | Includes pipeline scripts steps | 1 |

### Run pipeline

You can run a pipeline by specifying the pipeline using the `--pipeline` `-p` flag and which branch (`--branch` `-b`) or which commit (`--commit` `-c`) to run it on 

**Example**

```bash
bitbucket pipelines run -p deploy-to-staging -b master
bitbucket pipelines run -p custom:deploy-to-staging -c cba0e8d21da448f1264351ba2ebe5545958aa2ab
```

If the pipeline is custom it can be named without a prefix
If it is triggered by a tag or branch you can specify it by adding the tirgger as prefix 

```bash
bitbucket pipelines run -p branches:deploy-from-master -b master
bitbucket pipelines run -p tags:release-to-prod -c cba0e8d21da448f1264351ba2ebe5545958aa2ab
```


| Name | Description | Example |
| ---- | ------------ | ------ |
| `--pipeline` `-p` | Pipeline name | deploy-to-staging |
| `--branch` `-b` | Targeted branch name | master |
| `--commit` `-c` | Targeted commit name | cba0e8d21da448f1264351ba2ebe5545958aa2ab  |
| `--variables` `-v` | Targeted commit name | [{ "key": "var1key",  "value": "var1value", "secured": true}] |

### Pipeline Step

Gets pipeline's step log

**Usage**

```bash
bitbucket pipelines step -p $PIPELINE_ID -s $STEP_ID 
```
**Flags**

| Name | Description | Example |
| ---- | ------------ | ------- |
| `--pipeline` `-p` | Pipeline Id or build number | 1 |
| `--step` `-s` | step uuid | {1791efee-9e20-4c60-8a6a-bc1b071a15cc} |


### Stop Pipeline

Stop pipeline

**Usage**

```bash
bitbucket pipelines stop $PIPELINE_ID
```

## Pull requests

### List Pull requests

Lists pull requests 

**Usage**

```bash
bitbucket pr
bitbucket pr -p $PAGE_NUMBER
bitbucket pr -s MERGED
```
**Flags**

| Name | Description | Default |
| ---- | ------------ | ------- |
| `--page` `-p` |  Current page pagination | 1 |
| `--state` `-s` | PR state | OPEN |

### Pull request details

Fetches pull request details

**Usage**

```bash
bitbucket pr $PR_ID
```
**Example**

```bash
bitbucket pr 12 
bitbucket pr 12 -d #For detailed view including script 
```

### Pull request Merge

Merges pull request

**Usage**

```bash
bitbucket pr merge $PR_ID
bitbucket pr merge $PR_ID -m $MERGE_MESSAGE
bitbucket pr merge $PR_ID -s $MERGE_STRATEGY #specifying merge strategy [fast_forward, merge_commit, squash]
bitbucket pr merge $PR_ID -c #closes source
```
**Example**

```bash
bitbucket pr merge 12 
bitbucket pr merge 12 -c 
bitbucket pr merge 12 -m "Closing message" 
```
**Flags**

| Name | Description | Default |
| ---- | ------------ | ------- |
| `--message` `-m` |  Merge message | '' |
| `--strategy` `-s` | Merge strategy | OPEN |
| `--close-source` `-c` | Close source | false |

### Pull request decline

Declines a pull request

**Usage**

```bash
bitbucket pr decline $PR_ID
```
**Example**

```bash
bitbucket pr decline 12 
```

## Environments

### List Environments

Lists Environments

**Usage**

```bash
bitbucket envs
```




