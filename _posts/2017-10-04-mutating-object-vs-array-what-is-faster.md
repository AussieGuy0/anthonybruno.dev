---
layout: post
title: Mutating object vs single element array in java, what is faster?
published: false
---
Recently a coworker asked what he should do to do pass an `int` into a method,
change it and get the changed variable after the method ended. Ignoring the
anti-pattern of programming via side-effects, I suggested wrapping the variable
in an `Object` with a getter and setter. Another coworker suggested creating a
single element array BLA.

A *somewhat* heated discussion occurred as I debated it's more semantically
correct to use an object, while my coworker suggested that creating and using an
array is faster. 

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

To test the array, we have this code

{% highlight Java %}
private void mutateWithArrayTest() {
	int toMutate = 5;
	for (int i = 0; i < testsNum; i++) {
 		int[] array = {toMutate};
        mutate(array);
        int mutated = array[0];
    }
}

private void mutate(int[] toMutate) {
   toMutate[0] = generateRandomNumber();
}
{% endhighlight %}

Finally, to test the object we have this

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
