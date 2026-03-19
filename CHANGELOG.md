## [1.2.0] - TODO

TODO: git tag -a 1.2.0 -m "Replace Pandoc Docker with the Go project md-to-html-go"

### Added

- Send files to the VPS.

### Changed

- Convert md to html. Replace Pandoc Docker with the Go project [md-to-html-go](https://github.com/CarlosAMolina/md-to-html-go).
- Menu order.
- Test local. Omit git pull.
- Media. Create symlinks instead of copy media.

### Removed

- Pandoc and Docker are not required. Move code to the project [md-to-html-python](https://github.com/CarlosAMolina/md-to-html-python).

## [1.1.0] - 20260211

### Changed

- Deploy. Instead of copying the movie.mp4 file to the Docker volume, use a symlink to reduce storage on the VPS.
- Local test. Deprecate the Nginx Docker image to open the website locally. Open the files directly with Firefox, this avoids errors with the movie.mp4 symlink as it exists on the local host but not in the Nginx container. This error does not happen in the VPS as it does not use the Nginx Docker container to serve the web.

## [1.0.0] - 20260131

### Added

- Complete CLI tool to manage the website:
  - Deploy: create website files.
  - Test locally.
