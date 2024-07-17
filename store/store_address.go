package store

import (
	"time"

	"github.com/Conflux-Chain/go-conflux-util/store/mysql"
	"github.com/pkg/errors"
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

func (as *AddressStore) Count(startTime, endTime time.Time) (uint64, error) {
	db := as.DB.Model(&Address{})
	nilTime := time.Time{}
	if startTime != nilTime && endTime != nilTime {
		db = db.Where("block_time >= ? and block_time < ?", startTime, endTime)
	}
	if startTime == nilTime && endTime != nilTime {
		db = db.Where("block_time < ?", endTime)
	}

	var count int64
	err := db.Count(&count).Error
	if err != nil {
		return 0, err
	}

	return uint64(count), nil
}

type AddressStat struct {
	ID       uint64    `json:"-"`
	StatType string    `gorm:"size:4;not null;uniqueIndex:idx_statType_statTime,priority:1" json:"-"`
	StatTime time.Time `gorm:"not null;uniqueIndex:idx_statType_statTime,priority:2" json:"statTime"`

	AddrCount  uint64 `gorm:"not null;default:0" json:"addrCount"`  // Number of address in a specific time interval
	AddrActive uint64 `gorm:"not null;default:0" json:"addrActive"` // Number of active address in a specific time interval
	AddrTotal  uint64 `gorm:"not null;default:0" json:"addrTotal"`  // Total number of address by a certain time
}

func (AddressStat) TableName() string {
	return "address_stats"
}

type AddressStatStore struct {
	*mysql.Store
}

func newAddressStatStore(db *gorm.DB) *AddressStatStore {
	return &AddressStatStore{
		Store: mysql.NewStore(db),
	}
}

func (t *AddressStatStore) LastByType(statType string) (*AddressStat, error) {
	var addressStat AddressStat
	err := t.Store.DB.Where("stat_type = ?", statType).Order("stat_time desc").Last(&addressStat).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &addressStat, nil
}

func (t *AddressStatStore) Add(dbTx *gorm.DB, addressStats []*AddressStat) error {
	return dbTx.CreateInBatches(addressStats, batchSizeInsert).Error
}

func (t *AddressStatStore) Del(dbTx *gorm.DB, addressStat *AddressStat) error {
	return dbTx.Where("stat_type = ? and stat_time = ?", addressStat.StatType, addressStat.StatTime).Delete(&AddressStat{}).Error
}

func (t *AddressStatStore) List(intervalType *string, minTimestamp, maxTimestamp *int, desc bool, skip, limit int) (int64,
	[]AddressStat, error) {
	var conds []func(db *gorm.DB) *gorm.DB

	if intervalType != nil {
		intervalType := IntervalTypes[*intervalType]
		conds = append(conds, StatType(intervalType))
	}

	if minTimestamp != nil {
		conds = append(conds, MinTimestamp(*minTimestamp))
	}

	if maxTimestamp != nil {
		conds = append(conds, MaxTimestamp(*maxTimestamp))
	}

	dbRaw := t.DB.Model(&AddressStat{})
	dbRaw.Scopes(conds...)

	list := new([]AddressStat)
	total, err := t.Store.List(dbRaw, desc, skip, limit, list)
	if err != nil {
		return 0, nil, err
	}

	return total, *list, nil
}

type Miner struct {
	ID              uint64
	FirstMiningTime time.Time `gorm:"not null"`
}

func (Miner) TableName() string {
	return "miners"
}

type MinerStore struct {
	*mysql.Store
}

func newMinerStore(db *gorm.DB) *MinerStore {
	return &MinerStore{
		Store: mysql.NewStore(db),
	}
}

func (ms *MinerStore) Add(id uint64, firstMiningTime time.Time) (uint64, error) {
	var miner Miner
	existed, err := ms.Store.Exists(&miner, "id = ?", id)
	if err != nil {
		return 0, err
	}
	if existed {
		return miner.ID, nil
	}

	miner = Miner{
		ID:              id,
		FirstMiningTime: firstMiningTime,
	}

	if err := ms.DB.Create(&miner).Error; err != nil {
		return 0, err
	}

	return miner.ID, nil
}

func (ms *MinerStore) Count(startTime, endTime time.Time) (uint64, error) {
	db := ms.DB.Model(&Miner{})
	nilTime := time.Time{}
	if startTime != nilTime && endTime != nilTime {
		db = db.Where("first_mining_time >= ? and first_mining_time < ?", startTime, endTime)
	}
	if startTime == nilTime && endTime != nilTime {
		db = db.Where("first_mining_time < ?", endTime)
	}

	var count int64
	err := db.Count(&count).Error
	if err != nil {
		return 0, err
	}

	return uint64(count), nil
}

type MinerStat struct {
	ID       uint64    `json:"-"`
	StatType string    `gorm:"size:4;not null;uniqueIndex:idx_statType_statTime,priority:1" json:"-"`
	StatTime time.Time `gorm:"not null;uniqueIndex:idx_statType_statTime,priority:2" json:"statTime"`

	MinerCount  uint64 `gorm:"not null;default:0" json:"minerCount"`  // Number of miner in a specific time interval
	MinerActive uint64 `gorm:"not null;default:0" json:"minerActive"` // Number of active miner in a specific time interval
	MinerTotal  uint64 `gorm:"not null;default:0" json:"minerTotal"`  // Total number of miner by a certain time
}

func (MinerStat) TableName() string {
	return "miner_stats"
}

type MinerStatStore struct {
	*mysql.Store
}

func newMinerStatStore(db *gorm.DB) *MinerStatStore {
	return &MinerStatStore{
		Store: mysql.NewStore(db),
	}
}

func (t *MinerStatStore) LastByType(statType string) (*MinerStat, error) {
	var minerStat MinerStat
	err := t.Store.DB.Where("stat_type = ?", statType).Order("stat_time desc").Last(&minerStat).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &minerStat, nil
}

func (t *MinerStatStore) Add(dbTx *gorm.DB, minerStats []*MinerStat) error {
	return dbTx.CreateInBatches(minerStats, batchSizeInsert).Error
}

func (t *MinerStatStore) Del(dbTx *gorm.DB, minerStat *MinerStat) error {
	return dbTx.Where("stat_type = ? and stat_time = ?", minerStat.StatType, minerStat.StatTime).Delete(&MinerStat{}).Error
}

func (t *MinerStatStore) List(intervalType *string, minTimestamp, maxTimestamp *int, desc bool, skip, limit int) (int64,
	[]MinerStat, error) {
	var conds []func(db *gorm.DB) *gorm.DB

	if intervalType != nil {
		intervalType := IntervalTypes[*intervalType]
		conds = append(conds, StatType(intervalType))
	}

	if minTimestamp != nil {
		conds = append(conds, MinTimestamp(*minTimestamp))
	}

	if maxTimestamp != nil {
		conds = append(conds, MaxTimestamp(*maxTimestamp))
	}

	dbRaw := t.DB.Model(&MinerStat{})
	dbRaw.Scopes(conds...)

	list := new([]MinerStat)
	total, err := t.Store.List(dbRaw, desc, skip, limit, list)
	if err != nil {
		return 0, nil, err
	}

	return total, *list, nil
}
