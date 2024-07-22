// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "golang.org/x/crypto/acme"

// ---------------------------------------------------------
// EXERCISE: Undeclarables
//
//  1. Declare the variables below:
//      3speed
//      !speed
//      spe?ed
//      var
//      func
//      package
//
//  2. Observe the error messages
//
// NOTE
//  The types of the variables are not important.
// ---------------------------------------------------------

func main() {
	var 3speed int
	var !speed string
	var !speed bool
	var var float64
	var func int
	var package string
}
