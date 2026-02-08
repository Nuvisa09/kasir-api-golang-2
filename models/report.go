package models

type BestSeller struct {
	Name string `json:"nama"`
	Qty  int    `json:"qty_terjual"`
}

type ReportResponse struct {
	TotalRevenue   int        `json:"total_revenue"`
	TotalTransaksi int        `json:"total_transaksi"`
	ProdukTerlaris BestSeller `json:"produk_terlaris"`
}