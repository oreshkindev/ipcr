# Image-Processor

![CLI](https://img.shields.io/badge/CLI-blue.svg?logo=cli&logoColor=white&style=for-the-badge)

Image-Processor is a CLI microservice that extends the capabilities of image processing. Use this tool to automatically decode / compress PNG and JPEG files and then convert them to WEBP and JPG formats.

You can simply build the ipcr CLI with the provided Makefile.

```
$ ❯  make build
```

## Usage

```
Usage:
  ipcr [command]

Available Commands:
  post-process Use this command to watch your images directory and post-process the added files.
  pre-process  Use this command with the image directory path to pre-process existing files.

Flags:
  -h, --help          help for ipcr
  -q, --quality int   compressed image quality (default 60)
  -v, --version       version for ipcr
  -w, --workers int   determine the number of wokers (default 4)
```

### For example

```
ipcr input output
-
ipcr uncompresseddir compresseddir -w 10 -q 40
-
ipcr images images -q 100 (if we want change dir)
```

### Simple systemd service configuration

```
[Unit]

Description=Image processor service

Documentation=https://api.example.com/doc

[Service]

Type=simple

User=root

Group=root

TimeoutStartSec=0

Restart=on-failure

RestartSec=30s

ExecStart=/var/www/api.example.com/bin/ipcr post-process '/var/www/api.example.com/in/' '/var/www/api.example.com/out/'

[Install]

WantedBy=multi-user.target
```

## Author

[Yo :)](https://t.me/oreshkindev)

## License

See LICENSE for the full license text.

## Security Contact

To report a security vulnerability, please [push me](https://t.me/oreshkindev).
