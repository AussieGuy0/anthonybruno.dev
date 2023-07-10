---
layout: post
title: "Full stack web development in a single Java file: An intro to Javalin and htmx"
tags: Java Htmx
---

TODO: Explain HTMX.

Here's the list of dependencies that you can chuck in you `pom.xml` file.
Don't worry, this blog post will explain what each of them does!

```xml
<!-- Place the below in the <dependencies> section of your pom.xml -->
<dependency>
    <groupId>com.j2html</groupId>
    <artifactId>j2html</artifactId>
    <version>1.6.0</version>
</dependency>
<dependency>
    <groupId>io.javalin</groupId>
    <artifactId>javalin</artifactId>
    <version>5.5.0</version>
</dependency>
<dependency>
    <groupId>org.webjars.npm</groupId>
    <artifactId>htmx.org</artifactId>
    <version>1.9.2</version>
</dependency>
<!-- Defining a logger as Javalin will complain if there is no SL4J logger -->
<dependency>
    <groupId>org.slf4j</groupId>
    <artifactId>slf4j-api</artifactId>
    <version>2.0.5</version>
</dependency>
<dependency>
    <groupId>org.slf4j</groupId>
    <artifactId>slf4j-reload4j</artifactId>
    <version>2.0.5</version>
</dependency>
```

## Server side HTML with j2html
In a typical architecture of a Single Page Application, the server only communicates in JSON.
It is then up to frontend framework, like React, to consume that JSON and translate that to HTML.

In HTMX, however we want the server to respond with HTML.
HTMX will then simply replace the relevant element on the page with the server returned HTML.

A way to generate HTML in Java is to use a templating engine like
[Thymeleaf](https://www.thymeleaf.org/).
However, these means the template files are outside Java, and as the blog title suggests, we want
everything in Java!

A nice alternative to a templating engine is **j2html**.

j2html is a Java library that allows us to build HTML in a fluent and typesafe manner. If you know HTML, you
know j2html!

```java
import static j2html.TagCreator.*;

public class Main
  public static void main(String[] args) {
    // Fluently build a HTML structure.
    var content = html(
      body(
        h1("Hello World!")
      )
    );

    // Let's render that HTML!
    System.out.println(content.render());
  }
}
```

Output:

```html
<html><body><h1>Hello world</h1></body></html>
```

Now, let's serve that html through a sever. We can do that using **Javalin**.

## Handling requests with Javalin
According to <https://javalin.io/>, Javalin is a simple web framework for Java and Kotlin. I agree!

I feel like it is very similar to the [Express](https://expressjs.com/) framework for Node.js.
For instance, the following Express code:

```js
// Using Express as an example.
const express = require('express')
const app = express()
const port = 3000

app.get('/', (req, res) => {
  res.send('Hello World!')
})
app.listen()
```

Is the same as the following Javalin code:

```java
import io.javalin.Javalin;

public class Main {
  public static void main(String[] args) {
    var app = Javalin.create()
            .get("/", ctx -> ctx.result("Hello World"))
            .start();
    }
}
```

Now let's combine j2html and Javalin.
We can use Javalin's `Context#html` to allow the server to respond with HTML.

```java
import io.javalin.Javalin;
import static j2html.TagCreator.*;

public static void main(String[] args) {
  // Create a new Javalin instance.
  // By default it will run on port 8080, but can be changed.
  var javalin = Javalin.create();

  // Handle a GET request to the path "/".
  javalin.get("/", ctx -> {
    // Build the HTML, and render it.
    var content = html(
            body(
              h1("Hello world")
            )
    );
    // We prefix with the doctype to ensure the browser does not use quirks mode.
    // https://developer.mozilla.org/en-US/docs/Web/HTML/Quirks_Mode_and_Standards_Mode
    var rendered = "<!DOCTYPE html>\n" + content.render();

    // Return the html in the response.
    ctx.html(rendered);
  });

  javalin.start();
}
```

If we run this, we will see some output in the console.

```
 INFO [main] (JavalinLogger.kt:16) - Listening on http://localhost:8080/
 INFO [main] (JavalinLogger.kt:16) - You are running Javalin 5.5.0 (released May 1, 2023).
 INFO [main] (JavalinLogger.kt:16) - Javalin started in 273ms \o/
```

Once you see the "Javalin" started command, we can put in the url `http://localhost:8080` into a
browser. This should show something like this:

![](/media/htmx-1/hello-world.png)

Isn't it beautiful?

## Interactivity with htmx
Lastly, lets add some interactivity with htmx!

To add htmx to our project, we can add it as a [WebJar](https://www.webjars.org/).
A WebJar is basically just a way to manage Javascript dependencies using standard Java build tools,
like Maven or Gradle.

```xml
<!-- Was in the dependencies at the top of this blog post, but putting
     here in case you forgot! -->
<dependency>
    <groupId>org.webjars.npm</groupId>
    <artifactId>htmx.org</artifactId>
    <version>1.9.2</version>
</dependency>
```

To use this, we need to:
1. Enable WebJar support in Javalin.
2. Include it in our `<head>` element of our HTML.

```java
public static void main(String[] args) {
  var javalin = Javalin.create(config -> {
    // Enable WebJar support.
    config.staticFiles.enableWebJars();
  });

  javalin.get("/", ctx -> {
    var content = html(
            head(
              // The WebJars path follows the format:
              // /webjars/<artifactId>/<version>/<path-to-file>
              // We can find the <path-to-file> via npmjs.com 'Code' view.
              // For instance: https://www.npmjs.com/package/htmx.org?activeTab=code
              script().withSrc("/webjars/htmx.org/1.9.2/dist/htmx.min.js")
            ),
            body(
              h1("Hello world")
            )
        );
    var rendered = "<!DOCTYPE html>\n" + content.render();

    ctx.html(rendered);
  });

  javalin.start();
}
```

Now, let's start actually using htmx!
For this example, we'll create a simple counter app.
It will show a number, that can be incremented by pressing a button.

Firstly, let's create a method that produces the HTML for the count.
This will make a bit more sense later on.

```java
  // Create a H2 tag with an id.
  private H2Tag createCounter(int count) {
    return h2("count: " + count)
        .withId("counter");
  }
```

Next, we need a way to store the count on the server.
The easiest way is to use an `AtomicInteger`.

We can use this count and the `createCounter` method we created before in the handler we have:

```java
  public static void main() {
    var count = new AtomicInteger();
    javalin.get("/", ctx -> {
      var content = html(
          head(
              script().withSrc("/webjars/htmx.org/1.9.2/dist/htmx.min.js")
          ), (
              body(
                  h1("Hello world"),
                  createCounter(count.get())
              )
          )
      );
      var rendered = "<!DOCTYPE html>\n" + content.render();
      ctx.html(rendered);
    });

  }
```

Lastly, and most importantly, we can use htmx by creating a button with specific attributes.

```java
button("Increment")
  .attr("hx-post", "/increment")
  .attr("hx-target", "#counter")
  .attr("hx-swap", "outerHTML")
```

To explain each attribute:
- [hx-post](https://htmx.org/attributes/hx-post/): When clicked, make a `POST` request to the `/increment` path.
- [hx-target](https://htmx.org/attributes/hx-target/): Swap the element with the id `counter`.
- [hx-swap](https://htmx.org/attributes/hx-swap/): Using `outerHTML` means htmx will swap the entire html of the counter element.



```java
  public static void main() {
    var javalin = Javalin.create(config -> {
          config.staticFiles.enableWebjars();
        }
    );

    var counter = new AtomicInteger();
    javalin.get("/", ctx -> {
      var content = html(
          head(
              script().withSrc("/webjars/htmx.org/1.9.2/dist/htmx.min.js")
          ), (
              body(
                  h1("Hello world"),
                  createCounter(counter.get()),
                  button("Increment")
                      .attr("hx-post", "/increment")
                      .attr("hx-swap", "outerHTML")
                      .attr("hx-target", "#counter")
              )
          )
      );
      var rendered = "<!DOCTYPE html>\n" + content.render();
      ctx.html(rendered);
    });

    javalin.post("/increment", ctx -> {
      var newCounter = createCounter(counter.incrementAndGet());
      ctx.html(newCounter.render());
    });
  }
```




