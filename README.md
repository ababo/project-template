# Template for dockerized project

## Prerequisites

1. Install Docker on your host machine.

## Build

To build debug images run:

```sh
docker-compose build
```

**Note**: After this step the source files are still not compiled. The sources
are compiled on the *run*-step.

To build release images run:

```sh
docker-compose build -f docker-compose.release.yml
```

**Note**: Unlike building the debug images the source files are compiled during
this step.

## Run

To perform unit-tests and launch the application services in debug mode run:

```sh
docker-compose up
```

To launch release version run:

```sh
docker-compose -f docker-compose.release.yml up
```

Read the Docker documentation for further information.
