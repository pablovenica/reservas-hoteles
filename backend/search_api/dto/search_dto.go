package dto

type SearchRequestDTO struct {
	HotelName string `json:"hotel_name"`
	City      string `json:"city"`
	CheckIn   string `json:"check_in"`
	CheckOut  string `json:"check_out"`
	Guests    int    `json:"guests"`
}

type SearchResponseDTO struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	HotelName string `json:"hotel_name"`
	City      string `json:"city"`
	CheckIn   string `json:"check_in"`
	CheckOut  string `json:"check_out"`
	Guests    int    `json:"guests"`
	Timestamp string `json:"timestamp"`
	Status    string `json:"status"`
}

type SearchHistoryDTO struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	HotelName string `json:"hotel_name"`
	City      string `json:"city"`
	Timestamp string `json:"timestamp"`
}
