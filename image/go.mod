module github.com/easygithdev/gip/image

go 1.17

replace github.com/easygithdev/gip/pixel => ../pixel

replace github.com/easygithdev/gip/histogram => ../histogram

require (
	github.com/easygithdev/gip/histogram v0.0.0-00010101000000-000000000000
	github.com/easygithdev/gip/pixel v0.0.0-00010101000000-000000000000
)
