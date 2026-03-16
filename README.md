# cmoli.es-deploy

## Introduction

Project to automate the creation of the [cmoli.es](https://cmoli.es) website content:

- Download the website content (MD, CSS, JS, media, etc).
- Convert MD files to HTML.
- Copy the images and videos to the required paths.
- Etc.

## Configuration

### Media content

The media content (images, videos, etc.) must be in the `$HOME/Software/cmoli-media-content` folder using the same paths as the markdown web files. This is required because the media content will be copied from this path to the web content path with the `cp -r` command.

### Required software

- [Git](https://git-scm.com/).
- [Go](https://go.dev/).

```bash
make run
```
