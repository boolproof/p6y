# p6y - duration string parsing in GO
Package provides parsing of duration information encoded in string compliant with format described by **ISO 8601** (https://en.wikipedia.org/wiki/ISO_8601#Durations).

## Current limitations
Only `PnYnMnDTnHnMnS` and `PnW` formats are parsed by the package (no suppport for both `PYYYYMMDDThhmmss` and `P[YYYY]-[MM]-[DD]T[hh]:[mm]:[ss]`).

Package can not parse values with fractions (only integer values).

If parsing is successfull (no error), use following result struct methods to access intrger values of duration components:
- `Years()`
- `Months()`
- `Days()`
- `Weeks()`
- `Hours()`
- `Minutes()`
- `Seconds()`

## Example usage
```go
package main

import (
	"fmt"

	"github.com/boolproof/p6y"
)

func main() {
	d, err := p6y.NewDuration("P7Y5DT12H")
	if err == nil {
		fmt.Printf("This will take %d years, %d months,"+
			"%d days and %d hours\n",
			d.Years(), d.Months(), d.Days(), d.Hours())

		//outputs: "This will take 7 years, 0 months, 5 days and 12 hours"
	} else {
		fmt.Println(err.Error())
	}
}

```