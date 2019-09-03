---
layout: post
title: Add RSS to Github Pages site and link it to Dev.to
published: false    
---
I've been posting to [Dev.to](dev.to) for a little while now. I've always posted
to my website first, then more or less copy and pasted the markdown into the
Dev.to editor. I thought there must be a better way, and luckily stumbled upon a
setting that allows you to automatically link blog posts from Github Pages (or
any RSS feed, really) to Dev.to! 

# Adding RSS To Github Pages
Adding RSS support to a Github Pages is very simple. All we have to do is add
the following line to the `Gemfile`. 

```ruby
gem 'jekyll-feed'
```

If we want to test locally, we can do a `bundle install`, which will install the
required gem. 

When we run the site `bundle run server` {BLA: check command}, go to the page
`/feed.xml`. You should see something like this:

```xml
{BLA: FILL IN}
```

# Linking RSS on Dev.to
- Go to `/bla`
- enter url
- press button
- wait 


# Possible Improvements
    - Delay on first fetch, better messages would be good
    - Be able to ignore/remove specific feed entries

# Conclusion
