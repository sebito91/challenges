# Perfect Numbers

Determine if a number is perfect, abundant, or deficient based on Nicomachus' (60 - 120 CE) classification scheme for natural numbers.

The Greek mathematican Nicomachus devised a classification scheme for
natural numbers, identifying each as belonging uniquely to the
categories of _abundant_, _perfect_, or _deficient_.  A perfect number
equals the sum of its positive divisors — the pairs of numbers whose
product yields the target number, excluding the number itself.

- Perfect: Sum of factors = number
- Abundant: Sum of factors > number
- Deficient: Sum of factors < number

The Aliquot sum is defined as the sum of the factors of a number not
including the number itself.

## Examples

- 6 is a perfect number because its divisors are 1, 2, 3 and 6 = 1 + 2 +
  3.
- 28 is perfect number because 28 = 1 + 2 + 4 + 7 + 14.
- Prime numbers 7, 13, etc are Deficient by the Nicomachus
  classification.

To run the tests simply run the command `go test` in the exercise directory.

If the test suite contains benchmarks, you can run these with the `-bench`
flag:

    go test -bench .

For more detailed info about the Go track see the [help
page](http://exercism.io/languages/go).

## Source

Taken from Chapter 2 of Functional Thinking by Neal Ford. [http://shop.oreilly.com/product/0636920029687.do](http://shop.oreilly.com/product/0636920029687.do)

## Submitting Incomplete Problems
It's possible to submit an incomplete solution so you can see how others have completed the exercise.

