---
layout: post
title: Common Java patterns in Kotlin
tags: Java Kotlin
---

![Java logo with an arrow pointing to the Kotlin logo](/media/kotlin-to-java.png)

Recently, I've been using Kotlin for a project at work. As someone coming from a 
Java background, I've been taking notes on certain patterns that are used in
Java, that can be done nicer in Kotlin. I feel like when using a new
language, it's easy to fall into the trap of writing non-idiomatic code simply
because it's what I would've done in another language.

I wrote this article with the idea that this is what I would've wanted when starting
with Kotlin, and I plan on updating it as I learn more techniques. Anyway, let's
begin!

## Multi-line Strings
Sometimes, it is nice to embed SQL or have HTML snippets in our
application. In these cases, we want to be able to have a `String` that spans
multiple lines. If we had it on a single line, it would be near impossible to
read!

### Java
In Java 11 and below, we have to resort to simple `String` `+` concatenation. Make sure you don't forget
to add new line endings `\n` or else when the String is built, it will all end up
on one line!
```java
// Java 11 and older
String query = "SELECT w.id, w.title, w.date, we.weight, e.name\n" +
        "FROM workouts as w\n" +
        "LEFT JOIN workout_exercises as we ON w.id = we.workout_id\n" +
        "LEFT JOIN exercises as e on we.exercise_id = e.id\n" +
        "WHERE w.user_id = ?";
```

However, starting from Java 13, we have a much nicer way of writing this! 

```java
// Java 13 and above
String query = """
        SELECT w.id, w.title, w.date, we.weight, e.name
        FROM workouts as w
        LEFT JOIN workout_exercises as we ON w.id = we.workout_id
        LEFT JOIN exercises as e on we.exercise_id = e.id
        WHERE w.user_id = ?
        """;
```

This is an example of a [Text Block](https://openjdk.java.net/jeps/378), which is a preview 
feature in Java 13 and 14. At the time of writing, this feature will be finalised in Java 15.


### Kotlin
For Kotlin, the syntax is similar to what is available in Java 13+. 

```kotlin
val query = """
    SELECT w.id, w.title, w.date, we.weight, e.name
    FROM workouts as w
    LEFT JOIN workout_exercises as we ON w.id = we.workout_id
    LEFT JOIN exercises as e on we.exercise_id = e.id
    WHERE w.user_id = ?
""".trimIndent()
```

You might be asking why we have to call the `trimIndent` function at the end.
This is because if we don't, Kotlin will construct the `String` including the
initial whitespace indentation on each line. As we are only inserting that
indentation for readability purposes, we have to call `trimIndent` which will
remove this initial whitespace from each line.

I think this is a case where the Java way of Text Blocks is a bit better, as it
will automatically trim the whitespace for us. However, Kotlin is **here now** and 
(at the time of writing) Java 15 is still a month or so away!

## String Concatenation
Another common thing we like to do with `String`s is to construct them with
variables. For instance, we receive some input from the user and we want to
create a message based on that input.

### Java
Unfortunately, Java does not have a nice, modern way to do this and we have to 
resort to simply using `+` to build our String.

```java
String name = "bob";
int age = 18;
String message = "My name is " + name + " and my age is" + age;
```

We do have other options that are useful in certain situations, such as
[StringBuilder](https://docs.oracle.com/en/java/javase/11/docs/api/java.base/java/lang/StringBuilder.html)
and
[String#format](https://docs.oracle.com/en/java/javase/11/docs/api/java.base/java/lang/String.html#format(java.lang.String,java.lang.Object...)). But for simple, straight forward
situations, we can only use `+`.

### Kotlin
In Kotlin, we can use template expressions inside a `String`. This greatly improves
readability when a `String` is concatenated with multiple values.
```kotlin
val name = "bob";
val age = 18;
val message = "My name is $name and my age is $age"

// Can also use ${expression} for more complex usages
val message = "My name is $name and my age in 10 years will be ${age + 10}"

```

## Static Utility Methods
When we want utility methods that we want to use without having create a new Object, 
it is common practice to add them as static methods on a class.

Frequently, 'Util' classes exist that only contain static methods. This is a
pattern employed by popular libraries like [Google
Guava](https://github.com/google/guava) and [Apache
Commons](https://github.com/apache/commons-lang).

### Java
```java
public class StringUtils {

    public static char getLastLetter(String str) {
        if (str.isEmpty()) {
            throw new IllegalArgumentException("Provided string must not be empty");
        }
        return str.charAt(str.length() - 1);
    }

}

StringUtils.getLastLetter("abc") // => c
```

### Kotlin

In Kotlin, we have a few options. 

Firstly, we can have a top-level function. That is, a function that is not
attached to a class.  This language feature does not exist in Java. 

We can call the function from anywhere we want, without needing to create a new
object, or prefix a class name in front of the function call.

```kotlin
// Using a top-level function

fun getLastLetter(str: String): Char {
    if (str.isEmpty()) {
        throw IllegalArgumentException("Provided string must not be empty")
    }
    return str[str.length - 1]
}

getLastLetter("abc") // => c
```

Next, we can add a method to an `object`. This is roughly equivalent to the static
method way we saw in Java, and the calling code is the same.

```kotlin
// Using an Object
// https://kotlinlang.org/docs/tutorials/kotlin-for-py/objects-and-companion-objects.html
object StringUtils {

    fun getLastLetter(str: String): Char {
        if (str.isEmpty()) {
            throw IllegalArgumentException("Provided string must not be empty")
        }
        return str[str.length - 1]
    }

}
StringUtils.getLastLetter("abc")  // => c
```

Finally, we have an extension function. This is a cool feature that allows us to
attach a method to an existing class, without the need of creating a subclass or
a new wrapped type!

```kotlin
// Using an extension function
// https://kotlinlang.org/docs/reference/extensions.html
// Preferred!
fun String.getLastLetter(): Char {
    if (this.isEmpty()) {
        throw IllegalArgumentException("Provided string must not be empty")
    }
    return this[this.length - 1]
}

"abc".getLastLetter() // => c
```

This extension function feature was one the initial features I saw when
looking at Kotiln where I thought that it was really useful!

## Singletons
In some cases, we want to be able to define and use
[singleton's](https://en.wikipedia.org/wiki/Singleton_pattern) in our code. 

### Java
With Java, we can do this by creating a class with a private constructor. The class
then provides a public static field that can be accessed from calling code.

```java
public class BeanFactory {

    public static final BeanFactory INSTANCE = new BeanFactory();

    private BeanFactory() {
        // Prevent outside world creating a new BeanFactory
    }

    public Beans createBakedBeans() {
        return new BakedBeans();
    }
}
```

### Kotlin
In Kotlin, we greatly reduce boilerplate code by simply defining a new `object`
type. 

```kotlin
object BeanFactory {

    fun createBakedBeans(): Beans {
        return BakedBeans()
    }

}
```

For more information about `object`, please read [the relevant
docs](https://kotlinlang.org/docs/reference/object-declarations.html).

## Conclusion
I've covered how to write common Java techniques in idiomatic Kotlin. Often, the
equivalent Kotlin code is shorter and more readable. 

My advice for Java developers starting with Kotlin: **always be learning!** If you find
yourself encountering a new problem in Kotlin and you solve it by writing
Java-ish code, slow down and see if there's a better way to do it in Kotlin.
Have a search on Google or talk to your favourite coworker about it.
This way, you will continuously learn best practices while still being
productive!

