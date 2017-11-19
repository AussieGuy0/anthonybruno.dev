---
layout: post
title: Mutating an object vs a single element array in Java, which is faster?
published: true
---
Recently a coworker asked what he should do to pass an `int` into a method,
change it, and then get the changed variable after the method ends. Apparently simply returning it was not an option! I suggested wrapping 
the variable in an `Object` with a getter and setter. Another coworker suggested 
creating a single element int array containing the value. 

A *somewhat* heated discussion occurred as I debated it's more semantically
correct to use an object, while my coworker suggested that creating and using an
array is faster. 

But we're programmers, let's write some code to figure out which is faster!

## The Code
Here's the testing code that simply runs the test, records how long it takes 
and prints the results.


{% highlight Java %}
private void timeTests(Runnable test, String type) {
	long timeStarted = System.currentTimeMillis();
	test.run();
	long timeEnded = System.currentTimeMillis();
	long totalTime = timeEnded - timeStarted;
	System.out.println("Using " + type + " - avg: " + totalTime/testsNum + "ms. Total time: " + totalTime + "ms");
}
{% endhighlight %}

To test mutating the variable using an array, we have this code:

{% highlight Java %}
private void mutateWithArrayTest() {
	int toMutate = 5;
	for (int i = 0; i < testsNum; i++) {
 		int[] array = { toMutate };
        mutate(array);
        int mutated = array[0];
    }
}

private void mutate(int[] toMutate) {
   toMutate[0] = generateRandomNumber();
}
{% endhighlight %}

Finally, to test mutating the variable with an object, we have this:

{% highlight Java %}
private void mutateWithObjectTest() {
	int toMutate = 5;
    for (int i = 0; i < testsNum; i++) {
        IntWrapper intWrapper = new IntWrapper(toMutate);
        mutate(intWrapper);
        int mutated = intWrapper.getWrapped();
    }
}

private void mutate(IntWrapper intWrapper) {
    intWrapper.setWrapped(generateRandomNumber());
}

private static class IntWrapper {

    private int wrapped;

    public IntWrapper(int wrapped) {
        this.wrapped = wrapped;
    }

    public int getWrapped() {
        return wrapped;
    }

    public void setWrapped(int wrapped) {
        this.wrapped = wrapped;
    }
}
{% endhighlight %}

## The Results

![Graph](/media/mutateIntGraph.jpg)
    
After running each test a billion times, 10 times each, the results are surprising.
There is practically **no** difference in speed when using an array vs an
object. This shows how efficient object creation is in Java, demonstrating that java programmers shouldn't be scared to create objects. 

It also shows that I was right! ;)
