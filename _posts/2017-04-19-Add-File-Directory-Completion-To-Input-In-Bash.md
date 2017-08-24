---
layout: post
title: Add file/directory auto-completion to user input in Bash
published: true
excerpt: A small guide on how to add path auto-completion to Bash scripts
---

The most common way to get input from a user in a Bash script is to use the read function like so:

{% highlight bash %}
#!/bin/bash
echo "Type a word and press enter"
read INPUT
echo "Your word was $INPUT"
{% endhighlight  %}

However, this is not the best method to use when you want the user to input a directory path.
This is because pressing `tab` enters a literal tab instead of auto-completing the directory path
like what happens in the terminal.

To add directory auto-completion to your bash script, all you have to do is add the `-e` flag to the
read function like so:

{% highlight bash %}
#!/bin/bash
echo "Type a file or directory (use tab to autocomplete path!) and press enter"
read -e  INPUT
echo "Your path was $INPUT"
{% endhighlight  %}
