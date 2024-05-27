//go:generate goa gen github.com/kormiltsev/proofofwork/api/design -o ./api/

package api

import (
	_ "goa.design/goa/v3/codegen"
	_ "goa.design/goa/v3/codegen/generator"
)
