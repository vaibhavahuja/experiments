module github.com/vaibhavahuja/projects/experiments/rabbitmq

go 1.19

replace github.com/ldclabs/cose => ../cose

require (
	github.com/fxamacker/cbor/v2 v2.7.0
	//github.com/ldclabs/cose v1.3.1
	github.com/stretchr/testify v1.9.0
)

require github.com/ldclabs/cose v0.0.0-00010101000000-000000000000

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/x448/float16 v0.8.4 // indirect
	golang.org/x/crypto v0.25.0 // indirect
	golang.org/x/sys v0.22.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
