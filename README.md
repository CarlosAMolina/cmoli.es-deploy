# cmoli.es-deploy

## Introduction

Project to automate the creation of the [cmoli.es](https://cmoli.es) website content:

- Download the website content (MD, CSS, JS, media, etc).
- Convert MD files to HTML.
- Copy the images and videos to the required paths.
- Etc.

## Configuration

### VPS connection

- Configure in your local pc the `~/.ssh/config` file with the VPS information.
- Set the host name in `~/.ssh/config` as the key `vps_alias` in the `config.json` file.

### Media content

Multimedia content (images, videos, etc.) must be located in the `$HOME/Software/cmoli-media-content` folder on both the computer where the web content will be generated and the VPS server from which that content will be served. Otherwise, the symbolic links generated will not work.

The `cmoli-media-content` folder must use the same paths as the markdown web files; otherwise, the symlinks won't be created correctly.

### Required software

- [Git](https://git-scm.com/).
- [Go](https://go.dev/).

```bash
make run
```
