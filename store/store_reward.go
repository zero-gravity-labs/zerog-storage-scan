package store

import (
	"fmt"
	"time"

	nhContract "github.com/0glabs/0g-storage-scan/contract"
	"github.com/Conflux-Chain/go-conflux-util/store/mysql"
	"github.com/openweb3/web3go/types"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Reward struct {
	BlockNumber  uint64          `gorm:"primaryKey;autoIncrement:false"`
	Miner        string          `gorm:"-"`
	MinerID      uint64          `gorm:"primaryKey;autoIncrement:false"`
	TxHash       string          `gorm:"size:66;primaryKey;autoIncrement:false"`
	BlockTime    time.Time       `gorm:"not null"`
	PricingIndex uint64          `gorm:"not null"`
	Amount       decimal.Decimal `gorm:"type:decimal(65);not null"`
}

func NewReward(blockTime time.Time, log types.Log, filter *nhContract.OnePoolRewardFilterer) (*Reward, error) {
	distributeReward, err := filter.ParseDistributeReward(*log.ToEthLog())
	if err != nil {
		return nil, err
	}

	reward := &Reward{
		PricingIndex: distributeReward.PricingIndex.Uint64(),
		Miner:        distributeReward.Beneficiary.String(),
		Amount:       decimal.NewFromBigInt(distributeReward.Amount, 0),
		BlockNumber:  log.BlockNumber,
		BlockTime:    blockTime,
		TxHash:       log.TxHash.String(),
	}

	return reward, nil
}

func (Reward) TableName() string {
	return "rewards"
}

type RewardStore struct {
	*mysql.Store
}

func newRewardStore(db *gorm.DB) *RewardStore {
	return &RewardStore{
		Store: mysql.NewStore(db),
	}
}

func (rs *RewardStore) Sum(startTime, endTime time.Time) (*decimal.Decimal, error) {
	nilTime := time.Time{}
	if startTime == nilTime && endTime == nilTime {
		return nil, errors.New("At least provide one parameter for startTime and endTime")
	}

	db := rs.DB.Model(&Reward{}).Select(`IFNULL(sum(Amount), 0) as amount`)
	if startTime != nilTime {
		db = db.Where("block_time >= ?", startTime)
	}
	if endTime != nilTime {
		db = db.Where("block_time < ?", endTime)
	}

	var sum struct {
		Amount decimal.Decimal
	}
	err := db.Find(&sum).Error
	if err != nil {
		return nil, err
	}

	return &sum.Amount, nil
}

func (rs *RewardStore) Add(dbTx *gorm.DB, rewards []Reward) error {
	return dbTx.CreateInBatches(rewards, batchSizeInsert).Error
}

func (rs *RewardStore) Pop(dbTx *gorm.DB, block uint64) error {
	return dbTx.Where("block_number >= ?", block).Delete(&Reward{}).Error
}

func (rs *RewardStore) List(idDesc bool, skip, limit int) (int64, []Reward, error) {
	dbRaw := rs.DB.Model(&Reward{})

	var orderBy string
	if idDesc {
		orderBy = "block_number DESC"
	} else {
		orderBy = "block_number ASC"
	}

	list := new([]Reward)
	total, err := rs.Store.ListByOrder(dbRaw, orderBy, skip, limit, list)
	if err != nil {
		return 0, nil, err
	}

	return total, *list, nil
}

func (rs *RewardStore) CountActive(startTime, endTime time.Time) (uint64, error) {
	db := rs.DB.Model(&Reward{})

	nilTime := time.Time{}
	if startTime != nilTime && endTime != nilTime {
		db = db.Where("block_time >= ? and block_time < ?", startTime, endTime)
	}
	if startTime == nilTime && endTime != nilTime {
		db = db.Where("block_time < ?", endTime)
	}

	var countActive int64
	err := db.Select(`count(distinct miner_id) as miner_count`).
		Find(&countActive).Error
	if err != nil {
		return 0, err
	}

	return uint64(countActive), nil
}

type GroupedReward struct {
	MinerID   uint64
	Amount    decimal.Decimal
	UpdatedAt time.Time
}

func (rs *RewardStore) GroupByMiner(minBn, maxBn uint64) ([]GroupedReward, error) {
	groupedRewards := new([]GroupedReward)
	err := rs.DB.Model(&Reward{}).
		Select(`miner_id, IFNULL(sum(Amount), 0) amount, max(block_time) updated_at`).
		Where("block_number between ? and ?", minBn, maxBn).
		Group("miner_id").
		Scan(groupedRewards).Error

	if err != nil {
		return nil, err
	}

	return *groupedRewards, nil
}

func (rs *RewardStore) GroupByMinerByTime(startBlockTime, endBlockTime time.Time) ([]GroupedReward, error) {
	groupedRewards := new([]GroupedReward)
	err := rs.DB.Model(&Reward{}).
		Select(`miner_id, IFNULL(sum(Amount), 0) amount, max(block_time) updated_at`).
		Where("block_time >= ? and block_time < ?", startBlockTime, endBlockTime).
		Group("miner_id").
		Scan(groupedRewards).Error

	if err != nil {
		return nil, err
	}

	return *groupedRewards, nil
}

func (rs *RewardStore) AvgRewardRecently(duration time.Duration) (*decimal.Decimal, error) {
	var stat struct {
		DistinctMiners int64
		TotalAmount    decimal.Decimal
	}

	err := rs.DB.Model(&Reward{}).
		Select(`COUNT(DISTINCT miner_id) distinct_miners, IFNULL(SUM(amount), 0) total_amount`).
		Where("block_time >= ?", time.Now().Add(-duration)).
		Find(&stat).Error
	if err != nil {
		return nil, err
	}

	if stat.DistinctMiners == 0 {
		return &decimal.Zero, nil
	}

	miners, err := decimal.NewFromString(fmt.Sprintf("%d", stat.DistinctMiners))
	if err != nil {
		return nil, err
	}

	avgReward := stat.TotalAmount.DivRound(miners, 0)

	return &avgReward, nil
}

type RewardStat struct {
	ID          uint64          `json:"-"`
	StatType    string          `gorm:"size:4;not null;uniqueIndex:idx_statType_statTime,priority:1" json:"-"`
	StatTime    time.Time       `gorm:"not null;uniqueIndex:idx_statType_statTime,priority:2" json:"statTime"`
	RewardNew   decimal.Decimal `gorm:"type:decimal(65);not null;default:0" json:"rewardNew"`   // New reward
	RewardTotal decimal.Decimal `gorm:"type:decimal(65);not null;default:0" json:"rewardTotal"` // Total reward
}

func (RewardStat) TableName() string {
	return "reward_stats"
}

type RewardStatStore struct {
	*mysql.Store
}

func newRewardStatStore(db *gorm.DB) *RewardStatStore {
	return &RewardStatStore{
		Store: mysql.NewStore(db),
	}
}

func (t *RewardStatStore) LastByType(statType string) (*RewardStat, error) {
	var rewardStat RewardStat
	err := t.Store.DB.Where("stat_type = ?", statType).Order("stat_time desc").Last(&rewardStat).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &rewardStat, nil
}

func (t *RewardStatStore) Add(dbTx *gorm.DB, rewardStats []*RewardStat) error {
	return dbTx.CreateInBatches(rewardStats, batchSizeInsert).Error
}

func (t *RewardStatStore) Del(dbTx *gorm.DB, rewardStat *RewardStat) error {
	return dbTx.Where("stat_type = ? and stat_time = ?", rewardStat.StatType, rewardStat.StatTime).Delete(&RewardStat{}).Error
}

func (t *RewardStatStore) List(intervalType *string, minTimestamp, maxTimestamp *int, desc bool, skip, limit int) (int64,
	[]RewardStat, error) {
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

	dbRaw := t.DB.Model(&RewardStat{})
	dbRaw.Scopes(conds...)

	list := new([]RewardStat)
	total, err := t.Store.List(dbRaw, desc, skip, limit, list)
	if err != nil {
		return 0, nil, err
	}

	return total, *list, nil
}

type RewardTopnStat struct {
	ID        uint64
	StatTime  time.Time       `gorm:"not null;uniqueIndex:idx_statTime_addressId,priority:1"`
	AddressID uint64          `gorm:"not null;uniqueIndex:idx_statTime_addressId,priority:2"`
	Amount    decimal.Decimal `gorm:"type:decimal(65);not null"`
}

func (RewardTopnStat) TableName() string {
	return "reward_topn_stats"
}

type RewardTopnStatStore struct {
	*mysql.Store
}

func newRewardTopnStatStore(db *gorm.DB) *RewardTopnStatStore {
	return &RewardTopnStatStore{
		Store: mysql.NewStore(db),
	}
}

func (t *RewardTopnStatStore) BatchDeltaUpsert(dbTx *gorm.DB, rewards []RewardTopnStat) error {
	db := t.DB
	if dbTx != nil {
		db = dbTx
	}

	var placeholders string
	var params []interface{}
	size := len(rewards)
	for i, r := range rewards {
		placeholders += "(?,?,?)"
		if i != size-1 {
			placeholders += ",\n\t\t\t"
		}
		params = append(params, []interface{}{r.StatTime, r.AddressID, r.Amount}...)
	}

	sqlString := fmt.Sprintf(`
		insert into 
    		reward_topn_stats(stat_time, address_id, amount)
		values
			%s
		on duplicate key update
			stat_time = values(stat_time),
			address_id = values(address_id),                
			amount = amount + values(amount)
	`, placeholders)

	if err := db.Exec(sqlString, params...).Error; err != nil {
		return err
	}

	return nil
}

type TopnMiner struct {
	Address string
	Amount  decimal.Decimal
}

func (t *RewardTopnStatStore) Topn(duration time.Duration, limit int) ([]TopnMiner, error) {
	miners := new([]TopnMiner)

	db := t.DB.Model(&RewardTopnStat{}).
		Select(`addresses.address address, IFNULL(sum(reward_topn_stats.amount), 0) amount`).
		Joins("left join addresses on addresses.id = reward_topn_stats.address_id")

	if duration != 0 {
		db = db.Where("reward_topn_stats.stat_time >= ?", time.Now().Add(-duration))
	}

	if err := db.Group("reward_topn_stats.address_id").
		Order("amount DESC").
		Limit(limit).
		Scan(miners).Error; err != nil {
		return nil, err
	}

	return *miners, nil
}
