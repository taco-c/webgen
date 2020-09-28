# webgen

Generate an entire website from XML and a template HTML file!

## How to run

```sh
$ webgen <file.xml> [output_directory]
```

`output_directory` is optional, defaults to current directory.

## XML file format
```xml
<website>
    <template-file> <!-- Path to template file, relative to this XML file --> </template-file>
    <stylesheet> <!-- Path to stylesheet file, relative to this XML file --> </stylesheet>
    
    <page>
        <path> <!-- Relative path --> </path>
        <filename> <!-- Name of file --> </filename>
        <style> <!-- CSS --> </style>
        <title1> <!-- Title and H1 --> </title1>
        <title2> <!-- H2 --> </title2>
        <title3> <!-- H3 --> </title3>
        <content> <!-- HTML --> </content>
    </page>

    <page>
        ...
    </page>
</website>
```