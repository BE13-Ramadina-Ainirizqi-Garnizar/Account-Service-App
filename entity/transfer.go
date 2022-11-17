package entity

type Transfer struct {
	IDTransfer     int
	UserIDPengirim int
	UserIDPenerima int
	NamaPengirim   string
	NamaPenerima   string
	SaldoTransfer  int
	CreatedAt      string
}
