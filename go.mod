module github.com/jsign/vkt-proof-bench

go 1.20

require (
	github.com/crate-crypto/go-ipa v0.0.0-20230410135559-ce4a96995014
	golang.org/x/sync v0.1.0
)

require golang.org/x/sys v0.0.0-20211020174200-9d6173849985 // indirect

replace github.com/crate-crypto/go-ipa => ../go-ipa
