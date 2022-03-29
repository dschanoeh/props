# props

![Tests](https://github.com/dschanoeh/props/actions/workflows/test.yaml/badge.svg)
![License](https://img.shields.io/github/license/dschanoeh/props)
[![Docker Image Size (latest semver)](https://img.shields.io/docker/image-size/dschanoeh/props)](https://hub.docker.com/r/dschanoeh/props)

props is a command line interface to interact with [Java .properties files](https://en.wikipedia.org/wiki/.properties).
These files aren't as easy to parse as it seems at first sight (see the [format section](https://en.wikipedia.org/wiki/.properties#Format) for an example).

Therefore, props utilizes the [magiconair/properties](https://github.com/magiconair/properties) library for proper parsing and provides a thin command line interface wrapper.

There are also the following GitHub Actions that wrap props for use in GitHub workflows:
* [get-property](https://github.com/dschanoeh/get-property)
* [change-property](https://github.com/dschanoeh/change-property)

## Features

* Extract a property value from a file
* Replace a property value in a file
* Create the properties file if it doesn't exist
* Support for UTF-8 (default) or ISO-8859-1/Latin1 encoding. Note that
  ISO-8859-1 is used by Java < 9 and UTF-8 is used by default in Java >= 9
* File-based processing or STDIN/STDOUT

## Example Usage
Reading a property:
```bash
% cat my.properties
foo=bar
% props read foo --file my.properties
bar
%
```

Writing a property
```bash
% cat my.properties
foo=bar
foo2=bar
% props write foo2 baz --file my.properties
% cat my.properties
foo = bar
foo2 = baz
``` 
