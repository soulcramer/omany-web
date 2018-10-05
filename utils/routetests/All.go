package routetests

var routeTests = map[string][]string{
	// User
	"/": {
		"/+Scott",
	},
	"/api/analytics/:id": {
		"/api/analytics/4J6qpK1ve",
	},
}

// All returns which specific routes to test for a given generic route.
func All() map[string][]string {
	return routeTests
}
