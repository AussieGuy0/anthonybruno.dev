---
title: Bit shifting
layout: post
tags: programming
---
We have two main bit shifting operators:
- `>>` right shift
- `<<` left shift

## Left Shift (<<)
Essentially moves the bits N to the left. 

For instance, let's say I have the number `6`. This can represented by the following binary (representing a byte)

```
00000110
```

If we were to shift this number by `1`, we would get the following:
```
6       # 00000110
6 << 1  # 00001100
```

Shifting left is actually equivalent to multiplying by powers of two so:
```
6 << 1 == 6 * 2^1 == 6 * 2 == 12
6 << 2 == 6 * 2^2 == 6 * 4 == 24
6 << 3 == 6 * 2^3 == 6 * 8 == 48
```

## Right Shift (>>)
Opposite of the left shift, moves the bits to the right. Let's look at the 6 example again!

```
6      # 00000110 
6 >> 1 # 00000011 (3)
6 >> 2 # 00000001 (1)
6 >> 3 # 00000000 (0)
```

Shifting right is actually equivalent to  dividing by powers of two and rounding down:
```
6 >> 1 == 6 / 2^1 == 6 / 2 == 3
6 >> 2 == 6 / 2^2 == 6 / 4 == 1 (1.5 rounded down)
6 >> 3 == 6 / 2^3 == 6 / 8 == 0 (0.75 rounded down)
```

## Application
Bit shifting can be used to extract certain specific values from a 'bitfield', where data is encoded at the bit level for a value.

For instance, in 8-bit color, we have 3 colors that we can assign to the 8 bits (1 byte). We could assign the 1 byte to the colors like this:

```
Bit    7  6  5  4  3  2  1  0
Data   R  R  R  G  G  G  B  B
```
This gives us 4 possible blue values, 8 possible green values and 8 possible red values.

Let's say we want to extract the green value, we first need to create a mask. The mask will remove all the data we don't want.

```
Data       R  R  R  G  G  G  B  B
Color      1  1  0  1  0  1  1  1 (215)
Green Mask 0  0  0  1  1  1  0  0 (28)

masked = color & greenMask (& ensures only bits of 1 is kept)

masked     0  0  0  1  0  1  0  0 (20)
```

We then just need to get the correct value by shifting the bits

```
Data       R  R  R  G  G  G  B  B
masked     0  0  0  1  0  1  0  0 (20)

greenValue = masked >> 2 = 5
```

And there we have it, we have decoded that our color byte has the green value of '5'.

## Sources
https://stackoverflow.com/questions/141525/what-are-bitwise-shift-bit-shift-operators-and-how-do-they-work