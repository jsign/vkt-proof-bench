module github.com/jsign/vkt-proof-bench

go 1.20

require (
	github.com/crate-crypto/go-ipa v0.0.0-20230601143733-1eef4aed89d9
	github.com/gballet/go-verkle v0.0.0-20230511144752-ab52d15a8a34
	golang.org/x/sync v0.1.0
)

require golang.org/x/sys v0.2.0 // indirect

replace github.com/gballet/go-verkle => github.com/jsign/go-verkle v0.0.0-20230515162510-171d74d383cc
