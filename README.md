[![Build Status](https://travis-ci.org/filiptc/store.svg?branch=master)](https://travis-ci.org/filiptc/store)
[![codecov.io](https://codecov.io/github/filiptc/store/coverage.svg?branch=master)](https://codecov.io/github/filiptc/store?branch=master)


Table of contents
=================

  * [How to run](#how-to-run)
  * [Assumptions & reasoning](#assumptions--reasoning)
  * [Design & Code structure](#design--code-structure)
  * [Notes](#notes)


How to run
============

1. Install go ([documentation](https://golang.org/doc/install))
2. Get project `go get github.com/filiptc/store/...`
3. Run tests `go test github.com/filiptc/store/implementation_test.go`

From a different package, the **usage** would be as follows:

```go
package main

import (
    "fmt"

	"github.com/filiptc/store/service"
	"github.com/filiptc/store/service/rules"
)

func main() {
    // see service/rules/rules.go:39
    pricingRules := []rules.Rule{}

    co := service.NewCheckout(pricingRules)
    co.Scan("VOUCHER")
    co.Scan("VOUCHER")
    co.Scan("TSHIRT")
    price := co.GetTotal()
    fmt.Println("Price is %.2f", price)

}
```


Assumptions & reasoning
=====
There was some confusion on my part about the "buy 2 get 1 free" discount logic. It seems a
classical "2 for 1" but the following parens threw me off "buy two of the same product, get an
additional one free" so I initially implemented it adding a free third voucher to the cart. However,
when examining the examples section once I was ready to do the final implementation, the case
"VOUCHER, TSHIRT, VOUCHER" only results in 25€ when using the first interpretation. This is why I
finally went for the first interpretation (2 for 1). But as you will see, the rules are pretty easy
to change with my implementation.

I decided to implement the store as two layers, one being the models (item and cart), the other one
being the service layer that exposes the required API.

The service layer houses the rules package, which has the Rules interface that specifies the methods
IsApplicable and Apply, which rules need to implement. IsApplicable is called first to check if
conditions for rule apply. Apply gets called for modifying the cart passed by reference.

On the concurrency subject, the assignment requires "the Checkout scanning method to be thread safe"
but does not say anything about the GetTotal method. If the scanning, applying of rules and
recalculating the cart is asynchronous, there are two options for the GetTotal method:
1) Let it return at moment it is executed, "non-blockingly"
2) Have it be blocking and wait until all scan operations are complete

I chose the second approach, because it makes most sense to take all queued scan operations into
account before calculationg the total, and also as to facilitate evaluating tests scenarios.

For this reason I finally decided in favor of a scan channel and a wait-group for the total.


Design & code structure
=====

All files containing methods and functions (except `implementation_test.go`) include a `_test` file
with unit tests.

* [implementation_test.go](implementation_test.go): Main implementation, proposed as test. It tests the
examples scenario on the assignment sheet
* `service/`: Contains the `service` package files.
  * [notification.go](service/checkout.go): Library entry point. Defines `Checkout` service.
  * `rules/`: Contains the `rules` package files.
    * [rules.go](service/rules/rules.go): Rule interface specification.
    * [tshirt.go](service/rules/tshirt.go): T-Shirt Rule implementation.
    * [voucher.go](service/rules/voucher.go): Voucher Rule implementation.
* `model/`: Contains the `model` package files.
  * [cart.go](model/cart.go): Represents the cart model as a struct wrapper around a slice of Items,
 a constructor and methods for adding, getting the price-sum and counting by product-code.
  * [item.go](model/item.go): Represents the item model with a constructor and data of available
  items.


Notes
=====

* More rules can be easily added in the [rules](service/rules) directory.