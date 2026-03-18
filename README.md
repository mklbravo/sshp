# sshp

A DevOps utility for managing SSH connections and operations.

## Features
- Manage SSH hosts and credentials
- Automate SSH tasks
- Easy configuration and extensibility

## Installation
### Install via script (recommended):

The script will install latest released version.

```ssh
   sh -c "$(curl -fsSL https://raw.githubusercontent.com/mklbravo/sshp/main/tools/install.sh)"
```

>[!NOTE]
> Install script will download the binary into `$HOME/.local/bin`. Don't forget to add this folder to your `PATH`

### Build and install locally:

With this method the latest main version will be installed.

1. Clone the repository:
```ssh 
git clone https://github.com/mklbravo/sshp.git 
```

2. `cd` into repository folder
```ssh 
cd sshp
```

3. Execute make job:
```ssh 
make install
```

>[!NOTE]
> Make job will install the binary into `$HOME/.local/bin`. Don't forget to add this folder to your `PATH`


## Getting Started
1. Clone the repository:
   ```sh
   git clone <repo-url>
   ```
2. Install dependencies as required.
3. Run the application or scripts as described in the documentation.

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

