package store

import (
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"

	nhContract "github.com/0glabs/0g-storage-scan/contract"

	"github.com/openweb3/web3go/types"

	"github.com/Conflux-Chain/go-conflux-util/store/mysql"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Address struct {
	ID           uint64
	Address      string          `gorm:"size:64;unique"`
	BlockTime    time.Time       `gorm:"not null"`
	DataSize     uint64          `gorm:"not null;default:0"`                  // Size of storage data
	StorageFee   decimal.Decimal `gorm:"type:decimal(65);not null;default:0"` // The base fee for storage
	Txs          uint64          `gorm:"not null;default:0"`                  // Number of layer1 transaction
	Files        uint64          `gorm:"not null;default:0"`                  // Number of files/layer2 transaction
	ExpiredFiles uint64          `gorm:"not null;default:0"`                  // Number of expired files
	PrunedFiles  uint64          `gorm:"not null;default:0"`                  // Number of pruned files
	UpdatedAt    time.Time       `gorm:"not null"`
}

func (Address) TableName() string {
	return "addresses"
}

type AddressStore struct {
	*mysql.Store
	runBatchIncrStat bool
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
		UpdatedAt: blockTime,
	}

	if err := as.DB.Create(&addr).Error; err != nil {
		return 0, err
	}

	return addr.ID, nil
}

func (as *AddressStore) Get(address string) (Address, bool, error) {
	var addr Address
	exist, err := as.Store.Exists(&addr, "address = ?", address)
	return addr, exist, err
}

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

func (as *AddressStore) DeltaUpdate(dbTx *gorm.DB, a *Address) error {
	db := as.DB
	if dbTx != nil {
		db = dbTx
	}

	u := map[string]interface{}{
		"data_size":   gorm.Expr("data_size + ?", a.DataSize),
		"storage_fee": gorm.Expr("storage_fee + ?", a.StorageFee),
		"txs":         gorm.Expr("txs + ?", a.Txs),
		"files":       gorm.Expr("files + ?", a.Files),
		"updated_at":  a.UpdatedAt,
	}
	if err := db.Model(&a).Where("id=?", a.ID).Updates(u).Error; err != nil {
		return err
	}

	return nil
}

func (as *AddressStore) BatchDeltaUpsertExpiredFiles(dbTx *gorm.DB, addresses []Address) error {
	db := as.DB
	if dbTx != nil {
		db = dbTx
	}

	var placeholders string
	var params []interface{}
	size := len(addresses)
	for i, a := range addresses {
		placeholders += "(?,?,?,?)"
		if i != size-1 {
			placeholders += ",\n\t\t\t"
		}
		params = append(params, []interface{}{a.ID, a.ExpiredFiles, time.Now(), time.Now()}...)
	}

	sql := fmt.Sprintf(`
		insert into 
    		addresses(id, expired_files, updated_at, block_time)
		values
			%s
		on duplicate key update
			expired_files = expired_files + values(expired_files)
	`, placeholders)

	if err := db.Exec(sql, params...).Error; err != nil {
		return err
	}

	return nil
}

func (as *AddressStore) BatchDeltaUpsertPrunedFiles(dbTx *gorm.DB, addresses []Address) error {
	db := as.DB
	if dbTx != nil {
		db = dbTx
	}

	var placeholders string
	var params []interface{}
	size := len(addresses)
	for i, a := range addresses {
		placeholders += "(?,?,?,?)"
		if i != size-1 {
			placeholders += ",\n\t\t\t"
		}
		params = append(params, []interface{}{a.ID, a.PrunedFiles, time.Now(), time.Now()}...)
	}

	sql := fmt.Sprintf(`
		insert into 
    		addresses(id, pruned_files, updated_at, block_time)
		values
			%s
		on duplicate key update
			pruned_files = pruned_files + values(pruned_files)
	`, placeholders)

	if err := db.Exec(sql, params...).Error; err != nil {
		return err
	}

	return nil
}

func (as *AddressStore) BatchUpsert(dbTx *gorm.DB, addresses []Address) error {
	db := as.DB
	if dbTx != nil {
		db = dbTx
	}

	var placeholders string
	var params []interface{}
	size := len(addresses)
	for i, a := range addresses {
		placeholders += "(?,?,?,?,?,?,?)"
		if i != size-1 {
			placeholders += ",\n\t\t\t"
		}
		params = append(params, []interface{}{a.ID, a.DataSize, a.StorageFee, a.Txs, a.Files, a.UpdatedAt, time.Now()}...)
	}

	sql := fmt.Sprintf(`
		insert into 
    		addresses(id, data_size, storage_fee, txs, files, updated_at, block_time)
		values
			%s
		on duplicate key update
			data_size = values(data_size),
			storage_fee = values(storage_fee),
			txs = values(txs),
			files = values(files),
			updated_at=values(updated_at)
	`, placeholders)

	if err := db.Exec(sql, params...).Error; err != nil {
		return err
	}

	return nil
}

func (as *AddressStore) Topn(field string, duration time.Duration, limit int) ([]Address, error) {
	db := as.DB.Model(&Address{})

	if duration != 0 {
		db = db.Where("updated_at >= ?", time.Now().Add(-duration))
	}

	list := new([]Address)
	if err := db.Order(fmt.Sprintf("%s DESC", field)).Limit(limit).Find(list).Error; err != nil {
		return nil, err
	}

	return *list, nil
}

type AddressStat struct {
	ID       uint64    `json:"-"`
	StatType string    `gorm:"size:4;not null;uniqueIndex:idx_statType_statTime,priority:1" json:"-"`
	StatTime time.Time `gorm:"not null;uniqueIndex:idx_statType_statTime,priority:2" json:"statTime"`

	AddrNew    uint64 `gorm:"not null;default:0" json:"addrNew"`    // Number of newly increased address in a specific time interval
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
	FirstMiningTime time.Time       `gorm:"not null"`
	Amount          decimal.Decimal `gorm:"type:decimal(65);not null;index:idx_amount,sort:desc"`
	UpdatedAt       time.Time       `gorm:"not null;index:idx_updatedAt,sort:desc"`
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

func (ms *MinerStore) Add(id uint64, firstMiningTime time.Time, amount decimal.Decimal) (uint64, error) {
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
		Amount:          amount,
		UpdatedAt:       firstMiningTime,
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

func (ms *MinerStore) BatchGetMiners(minerIDs []uint64) (map[uint64]Miner, error) {
	miners := new([]Miner)
	err := ms.DB.Raw("select * from miners where id in ?", minerIDs).Scan(miners).Error
	if err != nil {
		return nil, err
	}

	m := make(map[uint64]Miner)
	for _, miner := range *miners {
		m[miner.ID] = miner
	}

	return m, nil
}

func (ms *MinerStore) BatchUpsert(dbTx *gorm.DB, miners []Miner) error {
	db := ms.DB
	if dbTx != nil {
		db = dbTx
	}

	var placeholders string
	var params []interface{}
	size := len(miners)
	for i, m := range miners {
		placeholders += "(?,?,?,?)"
		if i != size-1 {
			placeholders += ",\n\t\t\t"
		}
		params = append(params, []interface{}{m.ID, m.Amount, m.UpdatedAt, time.Now()}...)
	}

	sql := fmt.Sprintf(`
		insert into 
    		miners(id, amount, updated_at, first_mining_time)
		values
			%s
		on duplicate key update
			amount = values(amount),
			updated_at = values(updated_at)
	`, placeholders)

	if err := db.Exec(sql, params...).Error; err != nil {
		return err
	}

	return nil
}

func (ms *MinerStore) Topn(duration time.Duration, limit int) ([]Miner, error) {
	db := ms.DB.Model(&Miner{})

	if duration != 0 {
		db = db.Where("updated_at >= ?", time.Now().Add(-duration))
	}

	list := new([]Miner)
	if err := db.Order("amount DESC").Limit(limit).Find(list).Error; err != nil {
		return nil, err
	}

	return *list, nil
}

type MinerStat struct {
	ID       uint64    `json:"-"`
	StatType string    `gorm:"size:4;not null;uniqueIndex:idx_statType_statTime,priority:1" json:"-"`
	StatTime time.Time `gorm:"not null;uniqueIndex:idx_statType_statTime,priority:2" json:"statTime"`

	MinerNew    uint64 `gorm:"not null;default:0" json:"minerNew"`    // Number of newly increased miner in a specific time interval
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

type MinerRegister struct {
	ID              uint64
	RegisterMinerID string `gorm:"size:66;not null"`
	Address         string `gorm:"-"`
	AddressID       uint64 `gorm:"not null"`
	PreAddress      string `gorm:"-"`
	PreID           uint64 `gorm:"not null;default:0"`

	BlockNumber uint64    `gorm:"not null;index:idx_bn"`
	BlockTime   time.Time `gorm:"not null;index:idx_bt"`
	TxHash      string    `gorm:"size:66;not null;index:idx_txHash,length:10"`
}

func NewMinerRegister(blockTime time.Time, log types.Log, filter *nhContract.PoraMineFilterer) (*MinerRegister,
	error) {
	minerRegister, err := filter.ParseNewMinerId(*log.ToEthLog())
	if err != nil {
		return nil, err
	}

	register := &MinerRegister{
		RegisterMinerID: common.BytesToHash(minerRegister.MinerId[:]).String(),
		Address:         minerRegister.Beneficiary.String(),

		BlockNumber: log.BlockNumber,
		BlockTime:   blockTime,
		TxHash:      log.TxHash.String(),
	}

	return register, nil
}

func NewMinerUpdate(blockTime time.Time, log types.Log, filter *nhContract.PoraMineFilterer) (*MinerRegister,
	error) {
	minerUpdate, err := filter.ParseUpdateMinerId(*log.ToEthLog())
	if err != nil {
		return nil, err
	}

	register := &MinerRegister{
		RegisterMinerID: common.BytesToHash(minerUpdate.MinerId[:]).String(),
		Address:         minerUpdate.To.String(),
		PreAddress:      minerUpdate.From.String(),

		BlockNumber: log.BlockNumber,
		BlockTime:   blockTime,
		TxHash:      log.TxHash.String(),
	}

	return register, nil
}

func (MinerRegister) TableName() string {
	return "miner_registers"
}

type MinerRegisterStore struct {
	*mysql.Store
}

func newMinerRegisterStore(db *gorm.DB) *MinerRegisterStore {
	return &MinerRegisterStore{
		Store: mysql.NewStore(db),
	}
}

func (mrs *MinerRegisterStore) Add(dbTx *gorm.DB, registers []MinerRegister) error {
	return dbTx.CreateInBatches(registers, batchSizeInsert).Error
}

func (mrs *MinerRegisterStore) Pop(dbTx *gorm.DB, block uint64) error {
	return dbTx.Where("block_number >= ?", block).Delete(&MinerRegister{}).Error
}
