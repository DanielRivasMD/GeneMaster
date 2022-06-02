####################################################################################################

_default:
  @just --list

####################################################################################################

# print justfile
@show:
  bat .justfile --language make

####################################################################################################

# edit justfile
@edit:
  micro .justfile

####################################################################################################

# aliases

####################################################################################################

# build for OSX
osx OUT:
  #!/bin/bash
  set -euo pipefail

  echo "Building..."
  go build -v -o {{OUT}}

####################################################################################################

# build for linux
linux OUT:
  #!/bin/bash
  set -euo pipefail

  echo "Building..."
  env GOOS=linux GOARCH=amd64 go build -v -o {{OUT}}

####################################################################################################

# install locally
install:
  #!/bin/bash
  set -euo pipefail

  echo "Install..."
  go install
  mv -v "${HOME}/.go/bin/GeneMaster" "${HOME}/.go/bin/gene"

####################################################################################################

# deliver gene binary
hermesUppmax:
  #!/bin/bash
  set -euo pipefail

  # declarations
  source .just.sh

  echo "Deploying to Uppmax..."
  rsync -azvhP "${gene}/excalibur/gene" "${uppmaxID}:${uppmaxBin}"

####################################################################################################
