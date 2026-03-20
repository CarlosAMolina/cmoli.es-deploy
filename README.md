# cmoli.es-deploy

## Introduction

Project to automate the creation of the [cmoli.es](https://cmoli.es) website content:

- Download the website content (MD, CSS, JS, media, etc).
- Convert MD files to HTML.
- Copy the images and videos to the required paths.
- Etc.

## Configuration

### Media content

Multimedia content (images, videos, etc.) must be located in the `$HOME/Software/cmoli.es-media` folder on the computer where the web content will be generated.

The `cmoli.es-media` folder must use the same paths as the markdown web files; otherwise, the symlinks won't be created correctly.

### Create the directories

We need to create and configure the following paths in:

- The computer that will run this script to create the web content.
- The VPS where the web content will be sent.

```
# HTML
sudo mkdir -p /var/www/cmoli.es
sudo chown nonroot:nonroot /var/www/cmoli.es
sudo chmod 755 /var/www/cmoli.es

# Media
sudo mkdir -p /var/www/cmoli.es-media
sudo chown nonroot:nonroot /var/www/cmoli.es-media
sudo chmod 755 /var/www/cmoli.es-media
# Copy media
cp -r ~/Software/cmoli.es-media/* /var/www/cmoli.es-media
```

### VPS connection

- Configure in your local pc the `~/.ssh/config` file with the VPS information.
- Set the host name in `~/.ssh/config` as the key `vps_alias` in the `config.json` file.

### Required software

- [Git](https://git-scm.com/).
- [Go](https://go.dev/).
- rsync. Installation example: `sudo apt install rsync`.

## Run

```bash
make run
```
