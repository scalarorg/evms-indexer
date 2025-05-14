package evm

type BTCFeeOpts int32

const (
	MinimumFee  BTCFeeOpts = 0
	EconomyFee  BTCFeeOpts = 1
	HourFee     BTCFeeOpts = 2
	HalfHourFee BTCFeeOpts = 3
	FastestFee  BTCFeeOpts = 4
)

var BTCFeeOpts_name = map[int32]string{
	0: "MINIMUM_FEE",
	1: "ECONOMY_FEE",
	2: "HOUR_FEE",
	3: "HALF_HOUR_FEE",
	4: "FASTEST_FEE",
}

var BTCFeeOpts_value = map[string]int32{
	"MINIMUM_FEE":   0,
	"ECONOMY_FEE":   1,
	"HOUR_FEE":      2,
	"HALF_HOUR_FEE": 3,
	"FASTEST_FEE":   4,
}

func (x BTCFeeOpts) String() string {
	return BTCFeeOpts_name[int32(x)]
}
