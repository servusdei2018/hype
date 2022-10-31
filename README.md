# hype

Hype is a <b>minimal</b> command-line tool for creating and serving plaintext blogs.

```shell
Usage:

	hype <command>
			
The commands are:

	init		initialize empty Hype repository
	post		create Hype post
	serve		serve Hype repository
	version		print Hype version
```

With all the bloat and clutter in modern CMS systems, it's easy to get lost in the details. By deliberately stripping down functionality to what matters (most), hype allows you to focus on writing.

Hype is not a replacement for a full CMS. Hype focuses on minimalism. Sometimes simple is beautiful.

## Features

 - Everything is plaintext
 - Edit and serve with a single utility

## Installation

Prebuilt binaries for Mac, Windows, and Linux are available on the [Releases](https://github.com/servusdei2018/hype/releases) page.

### Manual Installation

Alternately, if you have `go` installed, you may run the following command to install hype from source:

```
go install github.com/servusdei2018/hype
```

## Getting Started

 1. Navigate to a fresh folder and initialize a new hype repository there with `hype init`. This creates a configuration file (`hype.env`) as well as an initial directory structure (`www`). By default, a tiny HTML homepage is created at `www/index.html`. This is the first page visitors to your site will see. If you want, you can delete that and it will default to a directory listing.
 2. To create a new post, run `hype post`.
 3. To serve your hype repository, run `hype serve`, and visit `localhost`.

## Configuration

Configuration may be achieved via environment variables, which are loaded from `hype.env`.

| Option | Description |
|-|-|
| HYPE_EDITOR | Editor with which to edit posts (default: vi) |
| HYPE_PORT | Port on which to serve HTTP (default: 80) |

## License

Hype is free and open-source software, and is distributed under the MIT License. See `LICENSE`.
