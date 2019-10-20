---
title: Add RSS to GitHub Pages and link it to Dev.to
layout: post
tags: Guide
published: true
---

I've been posting to [Dev.to](https://dev.to) for a little while now. I've always posted
to my website first, then more or less copy and pasted the markdown into the
Dev.to editor. I thought there must be a better way, and luckily stumbled upon a
setting that allows you to automatically link blog posts from GitHub Pages (or
any RSS feed) to Dev.to! 

We'll go through the steps of enabling RSS on a GitHub pages site, and then
we'll link it up to a Dev.to blog.

# Adding RSS To GitHub Pages
Adding RSS support to a GitHub Pages is very simple. All we have to do is add
the following line the `_config.yml` file in the root directory of the project.

```yaml
plugins:
  - jekyll-feed
```

To test locally, we can add the following entry to the project's
`Gemfile`. (If you haven't yet set up a local environment for your GitHub pages
site, please follow [this guide.](https://help.github.com/en/articles/testing-your-github-pages-site-locally-with-jekyll)

```ruby
gem 'jekyll-feed'
```

We can do a `bundle install`, which will install the required gem. 

When we run the site using `bundle exec jekyll serve`, go to the page
`localhost:4000/feed.xml`. You should see something like this:

```xml
<feed>
    <generator uri="https://jekyllrb.com/" version="3.8.5">Jekyll</generator>
    <link href="https://yoursite.com/feed.xml" rel="self" type="application/atom+xml"/>
    <link href="https://yoursite.com/" rel="alternate" type="text/html"/>```
    <!-- Full xml omitted -->
```

Side note: Technically, this is an [Atom feed](https://en.wikipedia.org/wiki/Atom_(Web_standard)) 
but this will work fine for our use case.

Once we have confirmed it works, we can push it to GitHub via the following
command:
`git commit -am "Add jekyll-feed gem"; git push`

# Linking RSS on Dev.to
Firstly, we need to go to
[https://dev.to/settings/publishing-from-rss](https://dev.to/settings/publishing-from-rss).

We then just have to paste in 'https://yoursite.com/feed.xml' into the box, and
press 'update'. 

![RSS Setting page on Dev.to](/media/DevToRss.png)

I did enable the 'Mark the RSS source as canonical URL by default'. This means
it will mark my blog posts on my website as the original source, which helps my
website appear in Google.

After a small wait, posts from GitHub pages should appear in the Dashboard. They
initially come in as drafts and posts can be published individually.

## Possible Improvements

While the RSS integration works really well, there a couple of small issues I
encountered.

When I initially set up RSS on the Dev.to side, it said it fetched my feed but I didn't 
see any entries in my dashboard. It wasn't until I checked a bit later that posts were 
sent across from my blog. I certainly don't have a problem with waiting, 
but I thought I did something wrong as there was no clear messaging that it take a little while.

Another thing is that I have some entries from my blog that I don't want to publish on Dev.to.
However, if I delete the entries from my dashboard, they just come right back.
It'd be good if there was some way to 'ignore' a blog post from an RSS feed.


# Conclusion

To reiterate the steps:
1. Add the `- jekyll-feed` to the plugin section in `_config.yml`
2. Go to https://dev.to/settings/publishing-from-rss
3. Enter https://yoursite/feed.xml as the RSS Feed URL
4. Press update.
5. Check your Dashboard to see your posts! (may take a little while)
