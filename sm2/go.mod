module github.com/kingstenzzz/sm2-improvement/sm2

require github.com/kingstenzzz/sm2-improvement v0.0.0

//replace github.com/kingstenzzz/sm2-improvement/sm3 => ../sm3
replace (
	sm2-improvement => github.com/kingstenzzz/sm2-improvement v0.0.0
)

go 1.16
