# WoWForge-CLI

WoWForge-CLI is a simple command line utility for Linux that can be used to track and install addons for World of Warcraft from CurseForge. It's extremely simplistic and lightweight requiring very little setup or external dependencies.

## Installation
Currently pre-built binaries are available for download in the repositories [releases](https://github.com/m-triassi/wowforge-cli/releases) page. Simply download a release and follow the guidance below.

### Linux

Pick the correct file for your current distribution and architecture and download it. 

For example if you are on Ubuntu, download the `wowforge-cli_x.x.x_amd64.deb` file where `x.x.x` is the latest release version
and run the following:

```bash
sudo apt install ./wowforge-cli_x.x.x_amd64.deb
```

Once complete you may run `wowforge-cli` globally.

### MacOS 

`wowforge-cli` is distributed via [brew](https://brew.sh/), you can install the latest version of the application by adding the tap, and installing
the application.

```bash
brew tap m-triassi/public
brew install wowforge-cli
```

Once complete you may run `wowforge-cli` globally.

### Set up
Before being able to use the application you **must** set the install path configuration value. This path specifies where 
to install addons to. If you are using Lutris or Wine this path should be mounted somewhere in your file system. 

With Lutris specifically you can right-click on the game and press "Browse Files" and it should open a file browser in 
the root of the wine environment. 

Make sure to get the path all the way down to the `.../interface/addons` directory. 
Once you've found the path use the `set` command to set the path:

```bash
wowforge-cli set --install "/path/to/Games/battlenet/drive_c/Program Files (x86)/World of Warcraft/_retail_/Interface/AddOns/"
```

## Usage
Once `wowforge-cli` is installed and set up you can simply track new addons with the `add` command, and update them with the `update` command.

### Adding a Addon
To add an addon you will need the "project id" of the addon. Thankfully this ID is easy to find and is listed under "Project ID" 
in the description of each addon, under the "About Project" heading on the right-hand side. With this ID in hand you can simply run:

```bash
wowforge-cli add <ID>
```

This will download and install the addon to the path you specified above, as well as save the ID for future updates

### Updating Addons
Updating addons is even easier, since all tracked ID are saved, running the `update` command will update all tracked addons one by one without any intervention.
```bash
wowforge-cli update
```

## Contributing

Please see [CONTRIBUTING](https://github.com/m-triassi/wowforge-cli/blob/main/CONTRIBUTING.md) for details.

## License

The MIT License (MIT). Please see [License File](https://github.com/m-triassi/wowforge-cli/blob/main/LICENSE.md) for more information.
