# props

props is a command line interface to interact with [Java .properties files](https://en.wikipedia.org/wiki/.properties).
It utilizes the [magiconair/properties](https://github.com/magiconair/properties) library and provides a thin command line interface wrapper.

## Features

* Extract a property value from a file
* Replace a property value in a file
* Create the properties file if it doesn't exist
* ISO-8859-1 or UTF-8 encoding
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
