---
layout: post
title: Add RSS to Github Pages site and link it to Dev.to
published: true
---
I've been posting to [Dev.to](https://dev.to) for a little while now. I've always posted
to my website first, then more or less copy and pasted the markdown into the
Dev.to editor. I thought there must be a better way, and luckily stumbled upon a
setting that allows you to automatically link blog posts from Github Pages (or
any RSS feed, really) to Dev.to! 

We'll go through the steps of enabling RSS on a Github pages site, and then
we'll link it up to a Dev.to blog


# Adding RSS To Github Pages
Adding RSS support to a Github Pages is very simple. All we have to do is add
the following line to the `Gemfile`. 

```ruby
gem 'jekyll-feed'
```

If we want to test locally, we can do a `bundle install`, which will install the
required gem. 

When we run the site using `bundle exec jekyll serve`, go to the page
`localhost:4000/feed.xml`. You should see something like this:

```xml
<feed>
    <generator uri="https://jekyllrb.com/" version="3.8.5">Jekyll</generator>
    <link href="https://yoursite.com/feed.xml" rel="self" type="application/atom+xml"/>
    <link href="https://yoursite.com/" rel="alternate" type="text/html"/>```
    <!-- Full xml omitted -->
```

# Linking RSS on Dev.to
Firstly, we need to go to https://dev.to/settings/publishing-from-rss

We then just have to paste in 'https://yoursite.com/feed.xml' into the box, and
press 'update'. 


# Possible Improvements
I encountered a few things that could of been done beter.

## Delay on first fetch
Dev.to said "Last fetched" but I didn't have any entries in my dashboard. It
wasn't until I checked later in the day that posts were sent across from my
blog.

## Be able to ignore/remove specific feed entries
I have some entries from my blog that I don't want to publish on Dev.to.
However, if I delete the entries from my dashboard, they just come right back.
It'd be good if there was someway to 'ignore' a blog post.


# Conclusion
To reiterate the steps:
1. Add the `gem 'jekyll-feed'` to your github pages site
2. Go to https://dev.to/settings/publishing-from-rss
3. Enter https://yoursite/feed.xml as the RSS Feed URL
4. Press update.
5. Check your Dashboard to see your posts! (may take a little while)
