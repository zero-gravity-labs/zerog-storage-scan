package store

import (
	"time"

	"github.com/Conflux-Chain/go-conflux-util/store/mysql"
	"gorm.io/gorm"
)

type Address struct {
	ID        uint64
	Address   string    `gorm:"size:64;unique"`
	BlockTime time.Time `gorm:"not null"`
}

func (Address) TableName() string {
	return "addresses"
}

type AddressStore struct {
	*mysql.Store
}

func newAddressStore(db *gorm.DB) *AddressStore {
	return &AddressStore{
		Store: mysql.NewStore(db),
	}
}

func (as *AddressStore) Add(address string, blockTime time.Time) (uint64, error) {
	var addr Address
	existed, err := as.Store.Exists(&addr, "address = ?", address) //TODO using LRU cache for improving the query performance
	if err != nil {
		return 0, err
	}
	if existed {
		return addr.ID, nil
	}

	addr = Address{
		Address:   address,
		BlockTime: blockTime,
	}

	if err := as.DB.Create(&addr).Error; err != nil {
		return 0, err
	}

	return addr.ID, nil
}

// BatchGetAddresses TODO LRU cache
func (as *AddressStore) BatchGetAddresses(addrIDs []uint64) (map[uint64]Address, error) {
	addresses := new([]Address)
	err := as.DB.Raw("select * from addresses where id in ?", addrIDs).Scan(addresses).Error
	if err != nil {
		return nil, err
	}

	m := make(map[uint64]Address)
	for _, addr := range *addresses {
		m[addr.ID] = addr
	}

	return m, nil
}

func (as *AddressStore) Get(address string) (Address, bool, error) {
	var addr Address
	exist, err := as.Store.Exists(&addr, "address = ?", address)
	return addr, exist, err
}
