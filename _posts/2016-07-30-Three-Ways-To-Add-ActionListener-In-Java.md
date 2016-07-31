---
layout: post
title: Three Ways To Add an ActionListener In Java
published: false
---

When building GUI's, the interface needs a way to listen and respond to events.
One of these types event's in Java Swing is the ActionEvent. A commonly used component that generates a ActionEvent is a JButton. This article serves to explain three different ways an ActionListener can be added to a component.

## Component as ActionListener
This is when the component itself listens for actions. To do this, simply implement the ActionListener interface in the component, for instance:

{% highlight Java %}
public class ButtonExample extends JButton implements ActionListener {

    public ButtonExample() {
        addActionListener(this);
    }

    public void actionPerformed(ActionEvent ae) {
        //handle event here
    }

}
{% endhighlight %}

Components implementing the ActionListener therefore become both the source and listener. You would want to use this method if events are relatively simple and internal (e.g. not connecting multiple components). 

Doing this breaks the [Single responsibility principle](https://en.wikipedia.org/wiki/Single_responsibility_principle) as the component is both responsible for creating events and handling it. 

## Inner ActionListener
This method is when you use the 'new' keyword to
create a new ActionListener for each component. 
{% highlight Java %}
component.addActionListener(new ActionListener() {
        @Override
        public void actionPerformed(ActionEvent ae) {
            //handle event here
        }
});
{% endhighlight %}

This implementation is great for simple listeners and it also separates the listener from the component. Avoid this if your trying to add ActionListeners to sub-components, for example, sub-components of a JPanel. This can result in messy code as every component is creating it's own version of an ActionListener. 


## Separate ActionListener
This technique requires you to create a separate class that implements ActionListener. A component that needs this ActionListener simply creates a new instance of this class and adds it. 

{% highlight Java %}
public class ButtonExample extends JButton implements ActionListener {

    public ButtonExample() {
        addActionListener(new ButtonListener());
    }
}

public ButtonHandler implements ActionListener {
    public void actionPerformed(ActionEvent ae) {
        //handle event here
    }
}
{% endhighlight %}

This is the cleanest way to add an ActionListener. It separates the component from the listener,  creates a reusable listener that can be used across different classes and allows for complex interactions between components. 

