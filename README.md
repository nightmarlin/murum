# murum

> A wallpaper engine for your music listening history!

murum generates wallpapers for your machine!

## Usage

### CLI

- Generation

   ```shell
   $ murum generate [...]
   ```

### Daemon

- Installation

   ```shell
   $ murum install
   ```

   This will install the murum daemon and guide you through set up.

- Configuration

   ```shell
   $ murum config [...]
   ```

   You can use the config wizard to go through and edit your murum
   configuration. Alternatively you can edit `~/.murum.yaml`

## How it works

To put it simply, murum gets the album art for your recent listening history,
figures out some canvas sizes, works out (through the power of
[fancy maths](https://en.wikipedia.org/wiki/Natural_neighbor_interpolation))
where the art needs to go and then sets your desktop to the result. Easy! Right?

## Architecture

The murum daemon/cli requires an internet connection to work.
