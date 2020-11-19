package cookies

import (
	_ "github.com/zellyn/kooky/allbrowsers" // register cookie store finders!
)

type Cookies struct {
	Name string `json:"name"`
	Domain string `json:"domain"`
}
