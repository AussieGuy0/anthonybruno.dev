---
layout: post
title: The non-broken promise of static typing
published: true
tags: Thoughts
---
A while back I read an article about how static typing does not prevent
bugs being added to software. The article is appropriately named: [The broken promise of static
typing](https://dev.to/danlebrero/the-broken-promise-of-static-typing). 
The author conducted research by generating and comparing 'bug density' scores for GitHub repositories.
The bug density score was determined by getting the average
number of issues labelled 'bug' per repository. The results showed that there
were not any less bugs in statically typed languages vs dynamically typed languages. The author concludes on the results:

> "the lack of evidence in the charts that more advanced type languages are going to save us from writing bugs is very disturbing."

While this article brings up good points and makes an effort at original
research, I've always felt that the claims made were wrong. I strongly
believe less bugs will occur when a statically typed language is used.
However, I've never had any proper evidence to back up my claims...**until now!**

Enter: [The Morning Paper](https://blog.acolyer.org), a blog that summarises
tech white papers. It recently released an article talking about the same subject called: 
[To type or not to type: quantifying detectable bugs in
JavaScript](https://blog.acolyer.org/2017/09/19/to-type-or-not-to-type-quantifying-detectable-bugs-in-javascript/).

The article covers a study of the same name. In it, researchers looked at 400
fixed bugs in JavaScript projects hosted on GitHub. For each bug, the
researchers tried to see if adding type annotations (using TypeScript and Flow)
would detect the bug. The results? A substantial **15%** of bugs could be
detected using type annotations. With this reduction in bugs, it's hard to deny
the value of static typing.

While these results show a benefit from using static typing, people will continue
to prefer a specific type system. So, let's hear from you! What type system do you
prefer, and why?
