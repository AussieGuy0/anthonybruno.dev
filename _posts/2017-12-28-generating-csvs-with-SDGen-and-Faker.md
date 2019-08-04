---
layout: post
title: Generating CSV files with SDGen and Faker
published: true
tags: Java Project Guide
---
The CSV format is a simple but commonly used format for exchanging
data. Many applications support the import and export of information as CSV
files. Due to the popularity of this format, there is a requirement for
developers to generate large amounts of CSV files for testing. This is where my
latest open source project [SDGen](https://github.com/AussieGuy0/SDgen) comes
into play.

[SDGen](https://github.com/AussieGuy0/SDgen) is a Java library that helps
developers generate randomised data files for testing purposes. It supports CSV
and Fixed Width formats, with more formats such as JSON planned for the future.

This guide will show you how to generate a simple CSV file using [SDGen](https://github.com/AussieGuy0/SDgen) and
[Faker](https://github.com/DiUS/java-faker). Faker will be used to assist creating
random values.

## Maven
For a Maven project, we can add the required libraries by inserting the
following xml into the pom.xml file of a project. 
```xml
<dependencies>
    <dependency>
        <groupId>au.com.anthonybruno</groupId>
        <artifactId>SdGen</artifactId>
        <version>0.3.0</version>
    </dependency>
    <dependency>
        <groupId>com.github.javafaker</groupId>
        <artifactId>javafaker</artifactId>
        <version>0.14</version>
    </dependency>
</dependencies>
```

## Instructions
Firstly, we need to get the Faker instance by writing:
```java
    Faker faker = Faker.instance();
```

We can then use the faker instance to generate values such as URLs
`faker.internet().url()` and planet names `faker.space().planet()`.

Next, we'll use `SDGen`'s fluent builder to create the schema for the CSV file
we want to create. To begin, we write:
```java
    Gen.start()
```

We can then add fields (aka columns) using the `addField` method. `addField` takes 2
parameters: A `String` name, which will be to identify the field in the produced
file and a `Generator`. A `Generator` is a simple interface with a single method
`generate`. This is the way that random values are created and added to a field. 

We are going to make a simple CSV file of people. To do that, we will add a
'First Name' and a 'Last Name' column using the corresponding `Faker` methods to
generate values for these:

```java
    Gen.start()
        .addField("First Name", () -> faker.name().firstName())
        .addField("Last Name", () -> faker.name().lastName())
```

*Note: Using lambdas (e.g. `() -> faker.name().firstName()` is the equivalent of
        writing:*
        
```java
new Generator() {
    @Override
    public Object generate() {
        return faker.name().firstName();
    }
}
```

We also want to add an 'Age' field. To do this, we can use `SDGen`'s inbuilt
`IntGenerator`. We can give it a sensible minimum and maximum value to limit the
range of numbers it will generate. `SDGen` provides generators for all primitive types. 


```java
Gen.start()
    .addField("First Name", () -> faker.name().firstName())
    .addField("Last Name", () -> faker.name().lastName())
    .addField("Age", new IntGenerator(18, 80))
```


Next, we specify how many rows to generate by using the `generate` method.  We also want to select the format of the generated data. We
will be using `asCsv` to generate the data in CSV format. `SDGen` also supports
the Fixed Width format and will support other data formats such as JSON in the future.
    
```java
Gen.start()
    .addField("First Name", () -> faker.name().firstName())
    .addField("Last Name", () -> faker.name().lastName())
    .addField("Age", new IntGenerator(18, 80))
    .generate(1000) //1000 rows will be generated
    .asCsv()
```
Finally, we specify how the data will be output. We will use the `toFile` method
to put the information into a file.

```java
Gen.start()
    .addField("First Name", () -> faker.name().firstName())
    .addField("Last Name", () -> faker.name().lastName())
    .addField("Age", new IntGenerator(18, 80))
    .generate(1000)
    .asCsv()
    .toFile("people.csv");
``` 

And that's it! Running the code will produce a CSV file in the project's working
directory. Here is some data that was produced when I ran it:

```csv
First Name,Last Name,Age
Corrine,Berge,78
Gerald,Carter,63
Enid,Padberg,66
Eleanora,Murray,79
Coy,Okuneva,76
Jovan,Reynolds,77
Lane,Haag,48
```

For more information about SDGen, please visit it on
[Github](https://github.com/AussieGuy0/SDgen). 

## Details
- Faker: [Github](https://github.com/DiUS/java-faker)
- SDGen: [Github](https://github.com/AussieGuy0/SDgen)
