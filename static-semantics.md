Static Semantics of Go
======================

* packages
  * identified by /-separated names, `package "a/b"`
  * imported with `import "a/b"`, which introduces a `b` that can be
    used to access exported elements using dot notation `b.Fun`. _This
    syntax seems to be ambiguous with field access._
  * visibility is controlled by identified capitalization (lowercase
    private, uppercase public)
  * packages can be split over multiple files

* in several places contracted notation for type annotations is
  allowed, as `var a, b int` where both variables are of type `int`

* multiple variable (not tuples!)
  * functions can return multiple values `f(x, y int) (int, int) { ... }`
  * assignment could be `var a, b = f(1,2)`, or with literals `var a, b = 1, 2`
  * an n-ary return value can also be passed to an n-ary function,
    such as `f(f(1,2))`.

* named return values
  * variable names can be named in the signature, as in

        func f() (x, y int) {
            x = 1
            return
        }

    the return statement than does not require values, the named
    variables are used (although explicit return values are still
    allowed)

* unused variables are considered errors

* type conversions are mostly explicit `uint(f)`,
  * _syntax is the same as function application_
  * numerical types have some implicit conversions

        v := 1 + 1 // v int
        v := 1 + 1.0 // v float
        v := 1 + 1i // v complex

* locally scoped statements in loops, conditions, switches

      if v := math.Sqrt(7); v < math.Pi { ... }

  here `v` is only visible in the condition and the body

* structs
  * constructed either by providing all fields positionally, or by providing
    some fields by name
  * structurally equal structs with a different name, (same fields and field
    types) can be converted by `Struct2(struct1)`
  * struct members are accessed using dot-notation `s.Fld`. _ambiguous with
    package element access_
  * _accessing a field via a struct pointer does not require explicit
    dereferencing of the pointer_
  * anonymous fields
    * unqualified type name is used as the field name
    * anonymous fields cause methods of the anonymous fields be promoted to
      methods of the struct

* types
  * new named types are introduced using `type Name TypeExpr`
  * _the rules of implicit conversion when naming primitive types are
    not yet clear to me_

* arrays and slices
  * the size is part of the array type `[16]int`
  * array initialization checks that no more than size elements are
    provided (less is okay)
  * array access is statically checked for some indices (`1`, `1 + 4`)
  * array and slice builtin functions are polymorphic in the type its
    elements
  * slices have no static size `ys []int`

* maps
  * _a literal map initialization can omit the struct name for the
    entry values_, so these are the same
    
        var m1 = map[string]Struct{
            "key": Struct{ ... }, // explicit struct init
            "key": { ... },       // special form for map init
        }

    Both positional and named initialization forms for structs are
    allowed in both variants.

* `make` is a builtin to create slices, maps, or channels. It has
  several variants, some with different number of arguments, all with
  a type expression as the first argument.

* methods
  * methods are functions with a special receiver type, e.g. a method
    for type `T`:

        func (x T) m(...) { ... }

  * methods must be defined in the same package (not the same file!)
    as the type (also built-ins are disallowed)
  * methods implicitly convert between pointers and values for both
    receivers and arguments (functions do not do this)

* interfaces
  * a set of methods that needs to be implemented by a type
  * adherence to an interface is structural
  * there is no implicit conversion between pointers and values, so if
    the interface requires a pointer argument to a function, and the
    type has a method that accepts a value, it will not match the
    interface

* type switch
  * in each of the cases, the match variable has a different type,
    corresponding to the type expression for that case

* dependency analysis for package initializers
  * there is a static check there exists an initialization order for
    package variables

* operators are overloaded (plus for different numeric types, and
  strings, for example)

---

