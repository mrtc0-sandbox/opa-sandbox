package example.authz

default allow := false

allow {
    input.method == "GET"
    input.path == ["salary", input.subject.user]
}

allow {
    is_admin
}

is_admin {
    input.subject.groups[_] = "admin"
}
