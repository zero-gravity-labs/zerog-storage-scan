package sync

import (
	"math/big"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/0glabs/0g-storage-client/contract"
	nhContract "github.com/0glabs/0g-storage-scan/contract"
	"github.com/0glabs/0g-storage-scan/store"
	viperUtil "github.com/Conflux-Chain/go-conflux-util/viper"
	"github.com/ethereum/go-ethereum/common"
	"github.com/openweb3/web3go"
	"github.com/openweb3/web3go/types"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type flowConfig struct {
	Address                string
	SubmitEventSignature   string
	NewEpochEventSignature string
}

type rewardConfig struct {
	Address              string
	RewardEventSignature string
}

type daEntranceConfig struct {
	Address                            string
	DataUploadSignature                string
	ErasureCommitmentVerifiedSignature string
	DARewardSignature                  string
}

type daSignersConfig struct {
	Address                string
	NewSignerSignature     string
	SocketUpdatedSignature string
}

type baseSyncer struct {
	conf *SyncConfig
	sdk  *web3go.Client
	db   *store.MysqlStore

	flowAddr        string
	flowSubmitSig   string
	flowNewEpochSig string
	rewardAddr      string
	rewardSig       string

	daEntranceAddr    string
	dataUploadSig     string
	commitVerifiedSig string
	daRewardSig       string
	daSignersAddr     string
	newSignerSig      string
	socketUpdatedSig  string

	addresses []common.Address
	topics    [][]common.Hash

	pricePerSector *big.Int
	currentBlock   uint64
}

func (s *baseSyncer) mustInit() {
	var flow flowConfig
	viperUtil.MustUnmarshalKey("flow", &flow)
	var reward rewardConfig
	viperUtil.MustUnmarshalKey("reward", &reward)
	var daEntrance daEntranceConfig
	viperUtil.MustUnmarshalKey("daEntrance", &daEntrance)
	var daSigners daSignersConfig
	viperUtil.MustUnmarshalKey("daSigners", &daSigners)

	s.flowAddr = flow.Address
	s.flowSubmitSig = flow.SubmitEventSignature
	s.flowNewEpochSig = flow.NewEpochEventSignature
	s.rewardAddr = reward.Address
	s.rewardSig = reward.RewardEventSignature

	s.daEntranceAddr = daEntrance.Address
	s.dataUploadSig = daEntrance.DataUploadSignature
	s.commitVerifiedSig = daEntrance.ErasureCommitmentVerifiedSignature
	s.daRewardSig = daEntrance.DARewardSignature
	s.daSignersAddr = daSigners.Address
	s.newSignerSig = daSigners.NewSignerSignature
	s.socketUpdatedSig = daSigners.SocketUpdatedSignature

	s.addresses = []common.Address{
		common.HexToAddress(flow.Address),
		common.HexToAddress(reward.Address),
		common.HexToAddress(daEntrance.Address),
		common.HexToAddress(daSigners.Address),
	}

	s.topics = [][]common.Hash{{
		common.HexToHash(flow.SubmitEventSignature),
		common.HexToHash(flow.NewEpochEventSignature),
		common.HexToHash(reward.RewardEventSignature),
		common.HexToHash(daEntrance.DataUploadSignature),
		common.HexToHash(daEntrance.ErasureCommitmentVerifiedSignature),
		common.HexToHash(daEntrance.DARewardSignature),
		common.HexToHash(daSigners.NewSignerSignature),
		common.HexToHash(daSigners.SocketUpdatedSignature),
	}}

	s.mustInitPricePerSector()
	s.mustInitExpireInSec()
}

func (s *baseSyncer) mustInitPricePerSector() {
	ethClient, _ := s.sdk.ToClientForContract()

	flowContract, err := contract.NewFlow(common.HexToAddress(s.flowAddr), ethClient)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to instantiate flow contract")
	}
	marketAddress, err := flowContract.Market(nil)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to get market contract address")
	}

	marketContract, err := contract.NewMarket(marketAddress, ethClient)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to instantiate market contract")
	}
	pricePerSector, err := marketContract.PricePerSector(nil)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to get price per sector")
	}

	s.pricePerSector = pricePerSector
}

func (s *baseSyncer) mustInitExpireInSec() {
	value, exist, err := s.db.ConfigStore.Get(store.FileExpireSeconds)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to get file expiration from DB")
	}
	if exist {
		_, success := new(big.Int).SetString(value, 10)
		if !success {
			logrus.WithError(err).Fatal("Failed to parse file expiration from DB")
		}
		return
	}

	releaseSeconds, err := s.initReleaseSeconds()
	if err == nil {
		if err := s.db.ConfigStore.Upsert(nil, store.FileExpireSeconds, releaseSeconds.String()); err != nil {
			logrus.WithError(err).Fatal("Failed to create file expiration config")
		}
		return
	}

	lifetimeInSeconds, err1 := s.initLifetimeInSeconds()
	if err1 == nil {
		if err := s.db.ConfigStore.Upsert(nil, store.FileExpireSeconds, lifetimeInSeconds.String()); err != nil {
			logrus.WithError(err).Fatal("Failed to create file expiration config")
		}
		return
	}

	logrus.WithError(err).WithError(err1).Fatal("Failed to init file expiration config")
}

func (s *baseSyncer) initReleaseSeconds() (*big.Int, error) {
	ethClient, _ := s.sdk.ToClientForContract()

	chunkLinearReward, err := nhContract.NewChunkLinearReward(common.HexToAddress(s.rewardAddr), ethClient)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to instantiate ChunkLinearReward contract")
	}

	releaseSeconds, err := chunkLinearReward.ReleaseSeconds(nil)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to get release seconds")
	}

	return releaseSeconds, nil
}

func (s *baseSyncer) initLifetimeInSeconds() (*big.Int, error) {
	ethClient, _ := s.sdk.ToClientForContract()

	rewardContract, err := nhContract.NewOnePoolReward(common.HexToAddress(s.rewardAddr), ethClient)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to instantiate OnePoolReward contract")
	}

	lifetimeInSeconds, err := rewardContract.LifetimeInSeconds(nil)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to get lifetime in seconds")
	}

	return lifetimeInSeconds, nil
}

func (s *baseSyncer) decodeSubmit(blkTime time.Time, log types.Log) (*store.Submit, error) {
	addr := log.Address.String()
	sig := log.Topics[0].String()
	if !strings.EqualFold(addr, s.flowAddr) || sig != s.flowSubmitSig {
		return nil, nil
	}

	submit, err := store.NewSubmit(s.pricePerSector, blkTime, log, nhContract.DummyFlowFilterer())
	if err != nil {
		return nil, err
	}

	senderID, err := s.db.AddressStore.Add(submit.Sender, blkTime)
	if err != nil {
		return nil, err
	}

	submit.SenderID = senderID

	return submit, nil
}

func (s *baseSyncer) decodeFlowEpoch(blkTime time.Time, log types.Log) (*store.FlowEpoch, error) {
	addr := log.Address.String()
	sig := log.Topics[0].String()
	if !strings.EqualFold(addr, s.flowAddr) || sig != s.flowNewEpochSig {
		return nil, nil
	}

	flowEpoch, err := store.NewFlowEpoch(blkTime, log, nhContract.DummyFlowFilterer())
	if err != nil {
		return nil, err
	}

	senderID, err := s.db.AddressStore.Add(flowEpoch.Sender, blkTime)
	if err != nil {
		return nil, err
	}

	flowEpoch.SenderID = senderID

	return flowEpoch, nil
}

func (s *baseSyncer) decodeReward(blkTime time.Time, log types.Log) (*store.Reward, error) {
	addr := log.Address.String()
	sig := log.Topics[0].String()
	if !strings.EqualFold(addr, s.rewardAddr) || sig != s.rewardSig {
		return nil, nil
	}

	reward, err := store.NewReward(blkTime, log, nhContract.DummyRewardFilterer())
	if err != nil {
		return nil, err
	}

	minerID, err := s.db.AddressStore.Add(reward.Miner, blkTime)
	if err != nil {
		return nil, err
	}

	_, err = s.db.MinerStore.Add(minerID, blkTime, reward.Amount)
	if err != nil {
		return nil, err
	}

	reward.MinerID = minerID

	return reward, nil
}

func (s *baseSyncer) decodeNewSigner(blkTime time.Time, log types.Log) (*store.DASigner, error) {
	addr := log.Address.String()
	sig := log.Topics[0].String()
	if !strings.EqualFold(addr, s.daSignersAddr) || sig != s.newSignerSig {
		return nil, nil
	}

	signer, err := store.NewDASigner(blkTime, log, nhContract.DummyDASignersFilterer())
	if err != nil {
		return nil, err
	}

	signerID, err := s.db.AddressStore.Add(signer.Address, blkTime)
	if err != nil {
		return nil, err
	}

	signer.SignerID = signerID

	return signer, nil
}

func (s *baseSyncer) decodeSocketUpdated(blkTime time.Time, log types.Log) (*store.DASigner, error) {
	addr := log.Address.String()
	sig := log.Topics[0].String()
	if !strings.EqualFold(addr, s.daSignersAddr) || sig != s.socketUpdatedSig {
		return nil, nil
	}

	signer, err := store.NewDASignerSocket(log, nhContract.DummyDASignersFilterer())
	if err != nil {
		return nil, err
	}

	signerID, err := s.db.AddressStore.Add(signer.Address, blkTime)
	if err != nil {
		return nil, err
	}

	signer.SignerID = signerID

	return signer, nil
}

func (s *baseSyncer) decodeDataUpload(blkTime time.Time, log types.Log) (*store.DASubmit, error) {
	addr := log.Address.String()
	sig := log.Topics[0].String()
	if !strings.EqualFold(addr, s.daEntranceAddr) || sig != s.dataUploadSig {
		return nil, nil
	}

	daSubmit, err := store.NewDASubmit(blkTime, log, nhContract.DummyDAEntranceFilterer())
	if err != nil {
		return nil, err
	}

	senderID, err := s.db.AddressStore.Add(daSubmit.Sender, blkTime)
	if err != nil {
		return nil, err
	}

	_, err = s.db.DAClientStore.Add(senderID, blkTime)
	if err != nil {
		return nil, err
	}

	daSubmit.SenderID = senderID

	return daSubmit, nil
}

func (s *baseSyncer) decodeCommitVerified(blkTime time.Time, log types.Log) (*store.DASubmit, error) {
	addr := log.Address.String()
	sig := log.Topics[0].String()
	if !strings.EqualFold(addr, s.daEntranceAddr) || sig != s.commitVerifiedSig {
		return nil, nil
	}

	daSubmit, err := store.NewDASubmitVerified(blkTime, log, nhContract.DummyDAEntranceFilterer())
	if err != nil {
		return nil, err
	}

	return daSubmit, nil
}

func (s *baseSyncer) decodeDAReward(blkTime time.Time, log types.Log) (*store.DAReward, error) {
	addr := log.Address.String()
	sig := log.Topics[0].String()
	if !strings.EqualFold(addr, s.daEntranceAddr) || sig != s.daRewardSig {
		return nil, nil
	}

	daReward, err := store.NewDAReward(blkTime, log, nhContract.DummyDAEntranceFilterer())
	if err != nil {
		return nil, err
	}

	minerID, err := s.db.AddressStore.Add(daReward.Miner, blkTime)
	if err != nil {
		return nil, err
	}

	daReward.MinerID = minerID

	return daReward, nil
}

func (s *baseSyncer) convertLogs(logs []types.Log, bn2TimeMap map[uint64]uint64) (*store.DecodedLogs, error) {
	var decodedLogs store.DecodedLogs

	for _, log := range logs {
		if log.Removed {
			continue
		}

		ts := bn2TimeMap[log.BlockNumber]
		blockTime := time.Unix(int64(ts), 0)

		switch log.Topics[0].String() {
		case s.flowSubmitSig:
			submit, err := s.decodeSubmit(blockTime, log)
			if err != nil {
				return nil, err
			}
			if submit != nil {
				decodedLogs.Submits = append(decodedLogs.Submits, *submit)
			}
		case s.flowNewEpochSig:
			flowEpoch, err := s.decodeFlowEpoch(blockTime, log)
			if err != nil {
				return nil, err
			}
			if flowEpoch != nil {
				decodedLogs.FlowEpochs = append(decodedLogs.FlowEpochs, *flowEpoch)
			}
		case s.rewardSig:
			reward, err := s.decodeReward(blockTime, log)
			if err != nil {
				return nil, err
			}
			if reward != nil {
				decodedLogs.Rewards = append(decodedLogs.Rewards, *reward)
			}
		case s.newSignerSig:
			signer, err := s.decodeNewSigner(blockTime, log)
			if err != nil {
				return nil, err
			}
			if signer != nil {
				decodedLogs.DASigners = append(decodedLogs.DASigners, *signer)
			}
		case s.socketUpdatedSig:
			signer, err := s.decodeSocketUpdated(blockTime, log)
			if err != nil {
				return nil, err
			}
			if signer != nil {
				decodedLogs.DASignersWithSocketUpdated = append(decodedLogs.DASignersWithSocketUpdated, *signer)
			}
		case s.dataUploadSig:
			daSubmit, err := s.decodeDataUpload(blockTime, log)
			if err != nil {
				return nil, err
			}
			if daSubmit != nil {
				decodedLogs.DASubmits = append(decodedLogs.DASubmits, *daSubmit)
			}
		case s.commitVerifiedSig:
			daSubmit, err := s.decodeCommitVerified(blockTime, log)
			if err != nil {
				return nil, err
			}
			if daSubmit != nil {
				decodedLogs.DASubmitsWithVerified = append(decodedLogs.DASubmitsWithVerified, *daSubmit)
			}
		case s.daRewardSig:
			daReward, err := s.decodeDAReward(blockTime, log)
			if err != nil {
				return nil, err
			}
			if daReward != nil {
				decodedLogs.DARewards = append(decodedLogs.DARewards, *daReward)
			}
		default:
			return nil, errors.Errorf("Faild to decode log, sig %v", log.Topics[0].String())
		}
	}

	return &decodedLogs, nil
}

func (s *baseSyncer) findClosedInterval(input string, regStr string) (uint64, uint64, error) {
	reg := regexp.MustCompile(regStr)
	matches := reg.FindStringSubmatch(input)

	if len(matches) < 3 {
		return 0, 0, errors.Errorf("Failed to match by regExp")
	}

	start, err := strconv.ParseUint(matches[1], 10, 64)
	if err != nil {
		return 0, 0, err
	}
	end, err := strconv.ParseUint(matches[2], 10, 64)
	if err != nil {
		return 0, 0, err
	}

	return start, end, nil
}
