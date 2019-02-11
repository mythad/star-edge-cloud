package exchage

import (
	"container/list"
	"regexp"
)

var routeCollection *list.List

// Route -
type Route struct {
	Rule    string
	Address string
}

// RegisterRegexpRoute -
func RegisterRegexpRoute(rule string, addr string) {
	if routeCollection == nil {
		routeCollection = list.New()
	}

	routeCollection.PushBack(Route{Rule: rule, Address: addr})
}

// GetRegexpRoutes -
func GetRegexpRoutes(target string) []Route {
	var routes []Route
	for e := routeCollection.Front(); e != nil; e = e.Next() {
		r := e.Value.(Route)
		if isMatch(r.Rule, target) {
			routes = append(routes, r)
		}
	}

	return routes
}

func isMatch(rule, target string) bool {
	match, _ := regexp.MatchString(rule, target)
	return match
}
