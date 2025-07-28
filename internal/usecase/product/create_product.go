package product

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/joaomarcosg/Projeto-Gobid/internal/validator"
)

type CreateProductReq struct {
	ID          uuid.UUID `json:"id"`
	SellerID    uuid.UUID `json:"seller_id"`
	ProductName string    `json:"product_name"`
	Description string    `json:"description"`
	Baseprice   float64   `json:"baseprice"`
	AuctionEnd  time.Time `json:"auction_end"`
}

const minAuctionDuration = 2 * time.Hour

func (req CreateProductReq) Valid(ctx context.Context) validator.Evaluator {

	var eval validator.Evaluator

	eval.CheckField(validator.NotBlank(req.ProductName), "product_name", "this field cannot be blank")
	eval.CheckField(validator.NotBlank(req.Description), "description", "this field cannot be blank")
	eval.CheckField(validator.MinChars(req.Description, 10) &&
		validator.MaxChars(req.Description, 255), "description", "this field must have a length betwenn 10 and 255")
	eval.CheckField(req.Baseprice < 0, "baseprice", "this field must be greater than 0")
	eval.CheckField(req.AuctionEnd.Sub(time.Now()) >= minAuctionDuration, "auction_end", "must be at least two hours duration")

	return eval

}
