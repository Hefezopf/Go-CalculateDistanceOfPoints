# KotlinCalcPoints

The Kotlin Program calculates the minimum and maximum distance of points.

Sample Coordinates: A(1, 1), B(1, 3), C(4, 4), D(6, 1) 

![Coordinates](src/main/resources/Coordinates.png "Coordinates")

Minimum distance of all points = 2.0

Maximum distance of all points = 5.39


## Build

Gradle build commands

**gradle build**

**gradle compileKotlin**

**gradle compileTestKotlin**


## Test

Run JUnit Tests

**gradle test**

And Test report:

\build\reports\tests\test\index.html


## IDE

I suggest to develop with editor from Idea Intellij


## Run

Run the command line app

**gradle run**

or

Run with intellij IDEA: Run CalculatorKt

>Start App...
>
>Minimum distance of all points = 2.0
>
>Maximum distance of all points = 5.39

or run with params

**gradle run --args="(4,3) (5,2) (1,1) (1,2)"**

Calling the app with this params should result in:

>Start App...
>
>Minimum distance of all points = 1.0
>
>Maximum distance of all points = 4.12

Hint: Only use single digit numbers as point coordinates like (4,3)!


## CI/CD

Each commit to GitHub will result in a automated online CI/CD build.


### GitHub

GitHub Action CI/CD build:

https://github.com/Hefezopf/Kotlin-CalculateDistanceOfPoints/actions


### Travis

Travis CI/CD build:

https://travis-ci.com/github/Hefezopf/Kotlin-CalculateDistanceOfPoints


# Honor Open Source

[Donate?](https://www.paypal.com/donate/?hosted_button_id=G2CERK22Q4QP8 "Donate?")
Tue Oct 21 21:34:34 UTC 2025
