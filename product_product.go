package odoo

import (
	"fmt"
)

// ProductProduct represents product.product model.
type ProductProduct struct {
	LastUpdate            *Time      `xmlrpc:"__last_update,omptempty"`
	Active                *Bool      `xmlrpc:"active,omptempty"`
	ActivityDateDeadline  *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityIds           *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState         *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary       *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeId        *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId        *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	AttributeLineIds      *Relation  `xmlrpc:"attribute_line_ids,omptempty"`
	AttributeValueIds     *Relation  `xmlrpc:"attribute_value_ids,omptempty"`
	Barcode               *String    `xmlrpc:"barcode,omptempty"`
	CategId               *Many2One  `xmlrpc:"categ_id,omptempty"`
	Code                  *String    `xmlrpc:"code,omptempty"`
	Color                 *Int       `xmlrpc:"color,omptempty"`
	CompanyId             *Many2One  `xmlrpc:"company_id,omptempty"`
	CostCurrencyId        *Many2One  `xmlrpc:"cost_currency_id,omptempty"`
	CostMethod            *String    `xmlrpc:"cost_method,omptempty"`
	CreateDate            *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid             *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId            *Many2One  `xmlrpc:"currency_id,omptempty"`
	DefaultCode           *String    `xmlrpc:"default_code,omptempty"`
	Description           *String    `xmlrpc:"description,omptempty"`
	DescriptionPicking    *String    `xmlrpc:"description_picking,omptempty"`
	DescriptionPickingin  *String    `xmlrpc:"description_pickingin,omptempty"`
	DescriptionPickingout *String    `xmlrpc:"description_pickingout,omptempty"`
	DescriptionPurchase   *String    `xmlrpc:"description_purchase,omptempty"`
	DescriptionSale       *String    `xmlrpc:"description_sale,omptempty"`
	DisplayName           *String    `xmlrpc:"display_name,omptempty"`
	ExpensePolicy         *Selection `xmlrpc:"expense_policy,omptempty"`
	Id                    *Int       `xmlrpc:"id,omptempty"`
	Image                 *String    `xmlrpc:"image,omptempty"`
	ImageMedium           *String    `xmlrpc:"image_medium,omptempty"`
	ImageSmall            *String    `xmlrpc:"image_small,omptempty"`
	ImageVariant          *String    `xmlrpc:"image_variant,omptempty"`
	IncomingQty           *Float     `xmlrpc:"incoming_qty,omptempty"`
	InvoicePolicy         *Selection `xmlrpc:"invoice_policy,omptempty"`
	IsProductVariant      *Bool      `xmlrpc:"is_product_variant,omptempty"`
	ItemIds               *Relation  `xmlrpc:"item_ids,omptempty"`
	ListPrice             *Float     `xmlrpc:"list_price,omptempty"`
	LocationId            *Many2One  `xmlrpc:"location_id,omptempty"`
	LstPrice              *Float     `xmlrpc:"lst_price,omptempty"`
}

// ProductProducts represents array of product.product model.
type ProductProducts []ProductProduct

// ProductProductModel is the odoo model name.
const ProductProductModel = "product.product"

// Many2One convert ProductProduct to *Many2One.
func (pp *ProductProduct) Many2One() *Many2One {
	return NewMany2One(pp.Id.Get(), "")
}

// CreateProductProduct creates a new product.product model and returns its id.
func (c *Client) CreateProductProduct(pp *ProductProduct) (int64, error) {
	ids, err := c.CreateProductProducts([]*ProductProduct{pp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProductProduct creates a new product.product model and returns its id.
func (c *Client) CreateProductProducts(pps []*ProductProduct) ([]int64, error) {
	var vv []interface{}
	for _, v := range pps {
		vv = append(vv, v)
	}
	return c.Create(ProductProductModel, vv)
}

// UpdateProductProduct updates an existing product.product record.
func (c *Client) UpdateProductProduct(pp *ProductProduct) error {
	return c.UpdateProductProducts([]int64{pp.Id.Get()}, pp)
}

// UpdateProductProducts updates existing product.product records.
// All records (represented by ids) will be updated by pp values.
func (c *Client) UpdateProductProducts(ids []int64, pp *ProductProduct) error {
	return c.Update(ProductProductModel, ids, pp)
}

// DeleteProductProduct deletes an existing product.product record.
func (c *Client) DeleteProductProduct(id int64) error {
	return c.DeleteProductProducts([]int64{id})
}

// DeleteProductProducts deletes existing product.product records.
func (c *Client) DeleteProductProducts(ids []int64) error {
	return c.Delete(ProductProductModel, ids)
}

// GetProductProduct gets product.product existing record.
func (c *Client) GetProductProduct(id int64) (*ProductProduct, error) {
	pps, err := c.GetProductProducts([]int64{id})
	if err != nil {
		return nil, err
	}
	if pps != nil && len(*pps) > 0 {
		return &((*pps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of product.product not found", id)
}

// GetProductProducts gets product.product existing records.
func (c *Client) GetProductProducts(ids []int64) (*ProductProducts, error) {
	pps := &ProductProducts{}
	if err := c.Read(ProductProductModel, ids, nil, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindProductProduct finds product.product record by querying it with criteria.
func (c *Client) FindProductProduct(criteria *Criteria) (*ProductProduct, error) {
	pps := &ProductProducts{}
	if err := c.SearchRead(ProductProductModel, criteria, NewOptions().Limit(1), pps); err != nil {
		return nil, err
	}
	if pps != nil && len(*pps) > 0 {
		return &((*pps)[0]), nil
	}
	return nil, fmt.Errorf("product.product was not found with criteria %v", criteria)
}

// FindProductProducts finds product.product records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductProducts(criteria *Criteria, options *Options) (*ProductProducts, error) {
	pps := &ProductProducts{}
	if err := c.SearchRead(ProductProductModel, criteria, options, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindProductProductIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductProductIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProductProductModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProductProductId finds record id by querying it with criteria.
func (c *Client) FindProductProductId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductProductModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("product.product was not found with criteria %v and options %v", criteria, options)
}
