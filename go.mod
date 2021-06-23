module main
require (
	github.com/kingstenzzz/sm2-improvement v0.0.0
	github.com/kingstenzzz/sm2-improvement/sm2 v0.0.0
	github.com/kingstenzzz/sm2-improvement/sm3 v0.0.0

)
replace (
	github.com/kingstenzzz/sm2-improvement/sm2 =>  ./sm2
	github.com/kingstenzzz/sm2-improvement/sm3 =>  ./sm3

)
go 1.16
