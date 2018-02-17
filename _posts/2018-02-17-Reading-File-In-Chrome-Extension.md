---
layout: post
title: Reading files in Chrome Extension
published: false
---
Often, it is desirable for a Chrome Extension to be bundled with files that
should be read. This files may contain data or configuration information to help
the extension function. This short guide will show you how to include and read
files in a Chrome Extension.

## Add file path(s) to manifest.json
Firstly, you must add the file paths to the `web_accessible_resources` property
in `manifest.json`. The file paths are relative to the extension's root (where
the `manifest.json` is located). For instance, if I wanted to include a file
called `info.json` that is located in a folder `data`, it would look like:

```
"web_accessible_resources": [
    "data/info.json"
]
```

A cool feature is that these paths support wildcards. For example:


```
"web_accessible_resources": [
    "data/*"
]
```

will allow access to any file in the data folder.


## Read from file
The next step is to read the data from the file. To do this, we need to get the
url of the file and make a request to it.

To get the url of the file we can use `chrome.runtime.getUrl('path/to/file')`.

The last step is to make a request to the url. In this example, we will use the ES6 feature Fetch (BLA link) but methods such as XmlHttpRequest will also work.

```
const url = chrome.runtime.getUrl('path/to/file');

fetch(url)
.then((response) => {response.json()})
    .then((json) => doSomething(json);
```
