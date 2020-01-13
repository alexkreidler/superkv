package api

type RTransaction interface {
	// When the transaction is closed it will be committed
	Closer
	AGetter
}

type RWTransaction interface {
	Closer
	AGetter
	ASetter
}

type TStore interface {
	RWStore
	RTransact() RTransaction
	RWTransact() RWTransaction
}
