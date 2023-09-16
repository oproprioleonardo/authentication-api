package entity

import "errors"

type OfflineRecord struct {
	ID         string
	UUID       string
	Name       string
	OnlineMode bool
	Registered bool
}

func NewOfflineRecord(id string, uuid string, name string, onlineMode bool) (*OfflineRecord, error) {
	record := &OfflineRecord{
		ID:         id,
		UUID:       uuid,
		Name:       name,
		OnlineMode: onlineMode,
		Registered: false,
	}

	if err := record.Validate(); err != nil {
		return nil, err
	}

	return record, nil
}

func (o *OfflineRecord) Validate() error {
	if o.ID == "" {
		return errors.New("id is empty")
	}
	if o.UUID == "" {
		return errors.New("uuid is empty")
	}
	if o.Name == "" {
		return errors.New("name is empty")
	}
	return nil
}
