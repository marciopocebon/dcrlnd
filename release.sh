#!/bin/bash

# Simple bash script to build basic lnd tools for all the platforms
# we support with the golang cross-compiler.
#
# Copyright (c) 2016 Company 0, LLC.
# Use of this source code is governed by the ISC
# license.

# If no tag specified, use date + version otherwise use tag.
if [[ $1x = x ]]; then
    DATE=`date +%Y%m%d`
    VERSION="01"
    TAG=$DATE-$VERSION
else
    TAG=$1
fi

PACKAGE=dcrlnd
MAINDIR=$PACKAGE-$TAG

mkdir -p releases/$MAINDIR
cd releases/$MAINDIR

# Enable exit on error.
set -e

# If LNDBUILDSYS is set the default list is ignored. Useful to release
# for a subset of systems/architectures.
SYS=${LNDBUILDSYS:-"windows-386 windows-amd64 openbsd-386 openbsd-amd64 linux-386 linux-amd64 linux-armv6 linux-armv7 linux-arm64 darwin-386 darwin-amd64 dragonfly-amd64 freebsd-386 freebsd-amd64 freebsd-arm netbsd-386 netbsd-amd64 linux-mips64 linux-mips64le linux-ppc64"}

# Use the first element of $GOPATH in the case where GOPATH is a list
# (something that is totally allowed).
PKG="github.com/decred/dcrlnd"
COMMIT=$(git describe --abbrev=40 --dirty)
COMMITFLAGS="-X $PKG/build.Commit=$COMMIT"

for i in $SYS; do
    OS=$(echo $i | cut -f1 -d-)
    ARCH=$(echo $i | cut -f2 -d-)
    ARM=
    if [[ $ARCH = "armv6" ]]; then
      ARCH=arm
      ARM=6
    elif [[ $ARCH = "armv7" ]]; then
      ARCH=arm
      ARM=7
    fi
    mkdir $PACKAGE-$i-$TAG
    cd $PACKAGE-$i-$TAG
    echo "Building:" $OS $ARCH $ARM
    env GOOS=$OS GOARCH=$ARCH GOARM=$ARM go build -v -ldflags "$COMMITFLAGS" github.com/decred/dcrlnd
    env GOOS=$OS GOARCH=$ARCH GOARM=$ARM go build -v -ldflags "$COMMITFLAGS" github.com/decred/dcrlnd/cmd/dcrlncli
    cd ..
    if [[ $OS = "windows" ]]; then
	zip -r $PACKAGE-$i-$TAG.zip $PACKAGE-$i-$TAG
    else
	tar -cvzf $PACKAGE-$i-$TAG.tar.gz $PACKAGE-$i-$TAG
    fi
    rm -r $PACKAGE-$i-$TAG
done

shasum -a 256 * > manifest-$PACKAGE-$TAG.txt
