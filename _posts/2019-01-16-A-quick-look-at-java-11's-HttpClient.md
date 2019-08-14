---
layout: post
title: A quick look at Java 11's HttpClient
published: true 
tags: Java
---
Java 11 was released in September 2018 and is the first Long-Term-Support version 
after Java 8. One of its features is [HttpClient](https://docs.oracle.com/en/java/javase/11/docs/api/java.net.http/java/net/http/HttpClient.html), 
a new way to make HTTP requests. This post will give a quick overview of  HttpClient, and how it's a much-needed replacement 
for [URLConnection](https://docs.oracle.com/javase/8/docs/api/java/net/URLConnection.html)!

## The task
I've created a page on my website [http://anthonybruno.dev/last-update](http://anthonybruno.dev/last-update)  that
simply has a Unix timestamp of the last time the site was built.

The task is simple, create some code that requests this page and returns the
timestamp!

## Using URLConnection
Below is the code that uses `URLConnection`. 
```java
// 1. Create URL object
URL updateUrl = new URL("http://anthonybruno.com.au/last-update");

// 2. We use openConnection() on the url to get a HttpURLConnection, 
//    that we have to cast(?!). Also, this doesn't actually make a 
//    network request(?!?!?)
HttpURLConnection connection = (HttpURLConnection) updateUrl.openConnection();

// 3. We can then set things like set request methods, headers.
connection.setRequestMethod("GET");

// 4. Then we actually connect! Note: connect doesn't return anything, it
//    mutates the connection object!
connection.connect();
int statusCode = connection.getResponseCode();
if (statusCode != 200) {
    throw new RuntimeException("Got non 200 response code! " + statusCode);
}
// 5. Content is returned in an InputStream (Don't forget to close it!)
InputStream content = connection.getInputStream()

Instant timestamp = processIntoInstant(content)

// 6. Remember to disconnect! Note: HttpURLConnnection is not autoclosable!
connection.disconnect()
```

After creating the `URL` object, things quickly go awry. It's extremely
counter-intuitive to use a method called `openConnection()`, that doesn't
actually open a connection! Having to cast the returned `URLConnection` object to
`HttpURLConnection` to access methods like `setRequestMethod` and `disconnect`
is plain silly. Finally, calling `connect()` (which actually makes a network
request!) doesn't return anything, instead, you have to get response information 
from the `connection` object itself.

## Using HttpClient
Below is the code that uses `HttpClient`. You'll see a big difference.
```java
// 1. Create HttpClient object
HttpClient httpClient = HttpClient.newHttpClient();

// 2. Create URI object
URI uri = URI.create(updateUrl);

// 3. Build a request
HttpRequest request = HttpRequest.newBuilder(uri).GET().build();

// 4. Send the request and get a HttpResponse object back!
//    Note: HttpResponse.BodyHandlers.ofString() just parses the response body
//          as a String
HttpResponse<String> response = httpClient.send(request, HttpResponse.BodyHandlers.ofString());
int statusCode = response.statusCode();
if (statusCode != 200) {
    throw new RuntimeException("Got non 200 response code! " + statusCode);
}
Instant timestamp = processIntoInstant(response.body())
```

Now, isn't that much nicer than the `URlConnection` code we saw before? We first
set up a `HttpClient` object, which will send our requests. We then instantiate a `HttpRequest` object,
which holds the request method, headers, etc. We send the `HttpRequest`, using
the previously created `HttpClient`, giving us a nice `HttpResponse` object
back.

The second parameter in `httpClient.send` is a `BodyHandler`, which is
responsible for parsing the response body into the format you want. Java
provides a bunch of default ones in `BodyHandlers`, that covers common use cases
like parsing to `String`, `File` and `InputStream`. Of course, it's possible to
create your own, which deserves an article by itself.

The idea of creating a client, creating requests and receiving responses is
quite a bit more intuitive than using `URlConnection`! `HttpClient` also supports
asynchronous requests, HTTP/2 and websockets. It's an enticing reason to migrate
from 8 to 11!

*Code used in this article can be found
[here](https://github.com/AussieGuy0/trash-heap/blob/master/code/java/http/src/main/java/au/com/anthonybruno/http/Main.java)*
