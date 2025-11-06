package constant

type CtxKey string

const(
	NewRelicTransactionCtx CtxKey = "newRelicTransaction"

	RoleAdmin 	int64 = 1
	RoleUser  	int64 = 2
	RoleSeller 	int64 = 3
)