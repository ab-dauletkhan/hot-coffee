package models

type Sales struct {
	TotalRevenue   float64 `json:"total_revenue"`
	TotalItemsSold int     `json:"total_items_sold"`
	TimeReceived   string  `json:"time_reseived"`
}

type PopularItems struct {
	List []MenuItem `json:"popular_items_list"`
}
