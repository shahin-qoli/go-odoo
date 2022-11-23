package odoo

import (
	"fmt"
)

// IrActionsActWindowClose represents ir.actions.act_window_close model.
type IrActionsActWindowClose struct {
	LastUpdate     *Time      `xmlrpc:"__last_update,omptempty"`
	BindingModelId *Many2One  `xmlrpc:"binding_model_id,omptempty"`
	BindingType    *Selection `xmlrpc:"binding_type,omptempty"`
	CreateDate     *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName    *String    `xmlrpc:"display_name,omptempty"`
	Help           *String    `xmlrpc:"help,omptempty"`
	Id             *Int       `xmlrpc:"id,omptempty"`
	Name           *String    `xmlrpc:"name,omptempty"`
	Type           *String    `xmlrpc:"type,omptempty"`
	WriteDate      *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One  `xmlrpc:"write_uid,omptempty"`
	XmlId          *String    `xmlrpc:"xml_id,omptempty"`
}

// IrActionsActWindowCloses represents array of ir.actions.act_window_close model.
type IrActionsActWindowCloses []IrActionsActWindowClose

// IrActionsActWindowCloseModel is the odoo model name.
const IrActionsActWindowCloseModel = "ir.actions.act_window_close"

// Many2One convert IrActionsActWindowClose to *Many2One.
func (iaa *IrActionsActWindowClose) Many2One() *Many2One {
	return NewMany2One(iaa.Id.Get(), "")
}

// CreateIrActionsActWindowClose creates a new ir.actions.act_window_close model and returns its id.
func (c *Client) CreateIrActionsActWindowClose(iaa *IrActionsActWindowClose) (int64, error) {
	return c.Create(IrActionsActWindowCloseModel, iaa)
}

// UpdateIrActionsActWindowClose updates an existing ir.actions.act_window_close record.
func (c *Client) UpdateIrActionsActWindowClose(iaa *IrActionsActWindowClose) error {
	return c.UpdateIrActionsActWindowCloses([]int64{iaa.Id.Get()}, iaa)
}

// UpdateIrActionsActWindowCloses updates existing ir.actions.act_window_close records.
// All records (represented by ids) will be updated by iaa values.
func (c *Client) UpdateIrActionsActWindowCloses(ids []int64, iaa *IrActionsActWindowClose) error {
	return c.Update(IrActionsActWindowCloseModel, ids, iaa)
}

// DeleteIrActionsActWindowClose deletes an existing ir.actions.act_window_close record.
func (c *Client) DeleteIrActionsActWindowClose(id int64) error {
	return c.DeleteIrActionsActWindowCloses([]int64{id})
}

// DeleteIrActionsActWindowCloses deletes existing ir.actions.act_window_close records.
func (c *Client) DeleteIrActionsActWindowCloses(ids []int64) error {
	return c.Delete(IrActionsActWindowCloseModel, ids)
}

// GetIrActionsActWindowClose gets ir.actions.act_window_close existing record.
func (c *Client) GetIrActionsActWindowClose(id int64) (*IrActionsActWindowClose, error) {
	iaas, err := c.GetIrActionsActWindowCloses([]int64{id})
	if err != nil {
		return nil, err
	}
	if iaas != nil && len(*iaas) > 0 {
		return &((*iaas)[0]), nil
	}
	return nil, fmt.Errorf("id %v of ir.actions.act_window_close not found", id)
}

// GetIrActionsActWindowCloses gets ir.actions.act_window_close existing records.
func (c *Client) GetIrActionsActWindowCloses(ids []int64) (*IrActionsActWindowCloses, error) {
	iaas := &IrActionsActWindowCloses{}
	if err := c.Read(IrActionsActWindowCloseModel, ids, nil, iaas); err != nil {
		return nil, err
	}
	return iaas, nil
}

// FindIrActionsActWindowClose finds ir.actions.act_window_close record by querying it with criteria.
func (c *Client) FindIrActionsActWindowClose(criteria *Criteria) (*IrActionsActWindowClose, error) {
	iaas := &IrActionsActWindowCloses{}
	if err := c.SearchRead(IrActionsActWindowCloseModel, criteria, NewOptions().Limit(1), iaas); err != nil {
		return nil, err
	}
	if iaas != nil && len(*iaas) > 0 {
		return &((*iaas)[0]), nil
	}
	return nil, fmt.Errorf("ir.actions.act_window_close was not found with criteria %v", criteria)
}

// FindIrActionsActWindowCloses finds ir.actions.act_window_close records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrActionsActWindowCloses(criteria *Criteria, options *Options) (*IrActionsActWindowCloses, error) {
	iaas := &IrActionsActWindowCloses{}
	if err := c.SearchRead(IrActionsActWindowCloseModel, criteria, options, iaas); err != nil {
		return nil, err
	}
	return iaas, nil
}

// FindIrActionsActWindowCloseIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrActionsActWindowCloseIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(IrActionsActWindowCloseModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindIrActionsActWindowCloseId finds record id by querying it with criteria.
func (c *Client) FindIrActionsActWindowCloseId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrActionsActWindowCloseModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("ir.actions.act_window_close was not found with criteria %v and options %v", criteria, options)
}
