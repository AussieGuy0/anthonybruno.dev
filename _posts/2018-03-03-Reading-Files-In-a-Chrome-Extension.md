---
layout: post
title: Reading files in a Chrome Extension
published: true
---
Often, it is desirable for a Chrome extension to be bundled with files that
need to be read. These files may contain data or configuration information to help
the extension function. This short guide will show you how you can set up your Chrome extension to read files.

## Add file path(s) to manifest.json
Firstly, you must add the file paths to the `web_accessible_resources` property
in the `manifest.json` file. The file paths are relative to the extension's root (where
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
    "data/*.json"
]
```

will allow access to any json file in the data folder.


## Read from the file
The next step is to read the data from the file. To do this, we need to get the
URL of the file and make a request to it.

To get the URL of the file we can use `chrome.runtime.getUrl('path/to/file')`.

Then we make a GET request to the URL. In this example, we will use the ES6 feature [Fetch](https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API) but methods such as XmlHttpRequest will also work.

```javascript
const url = chrome.runtime.getUrl('path/to/file');

fetch(url)
    .then((response) => {response.json()}) //assuming file contains json
    .then((json) => doSomething(json));    
```

And there we have it!

To reiterate the steps simply:

1. Add file path to the `web_accessible_resources` property in the `manifest.json` file
2. Get the URL of the file using `chrome.runtime.getUrl('path/to/file')`
3. Make a GET request to the URL
