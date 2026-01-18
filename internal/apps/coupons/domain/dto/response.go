package dto

type CouponInfoResp struct {
	Name      string   `json:"name"`
	Amount    int      `json:"amount"`
	RemAmount int      `json:"remaining_amount"`
	Users     []string `json:"claimed_by"`
}
