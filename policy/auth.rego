package auth

import future.keywords.if
import future.keywords.in

# By default, deny requests.
default allow := false

allow if {
	some uri in data.uri[input.role]
	uri.path == input.path
}

allow if {
	some uri in data.uri.ROLE_TYPE_GUEST
	uri.path == input.path
}