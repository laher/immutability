# immutabillity

Immutable Lists and Maps, using generics and higher order functions.

 * Easy to create an immutability collection from a builtin slice or map.
 * Easy to operate on these using Higher Order Functions.
 * No fluent interfaces (keeping itself simple).

[![PkgGoDev](https://pkg.go.dev/badge/github.com/laher/immutability)](https://pkg.go.dev/github.com/laher/immutability)
[![Go Report Card](https://goreportcard.com/badge/github.com/laher/immutability)](https://goreportcard.com/report/github.com/laher/immutability)


I wrote this after working a bit on the `benbjohnson/immutable` package (see below). It's a great package with some useful performance optimisations - especially where you're writing over and over.

But, I realised I wanted a simpler library instead.

This library, `immutability`, is more about conveniently exposing collections of data without exposing the backing slice or map (which would be modifialble).

## Status

NOTE: this is just a very early version.

 * TODO `Set`s; better testing.


## Similar projects

### benbjohnson/immutable

 * See https://github.com/benbjohnson/immutable
 * uses a more sophisticated internal representation
 * offers a "fluent interface" API - 'modify by returning a modified collection'.
 * It is optimised for performance during transformations
 * NOTE: I wrote a few PRs to help to upgrade the benbjohnson package for use with generics, and added Sets.

It's a great package with some useful performance optimisations, but I realised I wanted a simpler library instead.

### myitcv

https://pkg.go.dev/myitcv.io/immutable uses code generation. It doesn't use generics yet.

