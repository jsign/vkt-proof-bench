module github.com/jsign/vkt-proof-bench

go 1.20

require (
	github.com/crate-crypto/go-ipa v0.0.0-20230410135559-ce4a96995014
	github.com/gballet/go-verkle v0.0.0-20230511144752-ab52d15a8a34
	golang.org/x/sync v0.1.0
)

require golang.org/x/sys v0.0.0-20220919091848-fb04ddd9f9c8 // indirect

replace github.com/crate-crypto/go-ipa => github.com/jsign/go-ipa v0.0.0-20230515134231-b3476889b67b

//replace github.com/crate-crypto/go-ipa => ../go-ipa

replace github.com/gballet/go-verkle => github.com/jsign/go-verkle v0.0.0-20230515162510-171d74d383cc

//replace github.com/gballet/go-verkle => ../go-verkle
