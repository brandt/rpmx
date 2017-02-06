# rpmx

A simple tool to extract RPMs.


## About

`rpmx` extracts RPM files to a new directory.

Unlike `rpm2cpio.pl *.src.rpm | cpio -idmv`, it doesn't litter the working directory with files.

Thanks to Go's cross-compilation support, it works on both Linux and macOS.


## Usage

Usage is simple:

    rpmx <src_file> [dest_dir]

If a destination isn't specified, the contents will be extracted to a directory with the name listed in the RPM's header.

### Examples

    rpmx ./foo-1.0.5-1.el7.src.rpm

If an output directory isn't specified, `rpmx` will lookup the name specified in the RPM's header and create a new directory with that name.

    rpmx ./foo.rpm /tmp/output

If an output directory is specified, `rpmx` will create the directory and put all the files into it.


## Installing

    go get -u github.com/brandt/rpmx


## Author

- J. Brandt Buckley
