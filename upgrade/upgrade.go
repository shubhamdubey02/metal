// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package upgrade

import (
	"errors"
	"fmt"
	"time"

	"github.com/MetalBlockchain/metalgo/ids"
	"github.com/MetalBlockchain/metalgo/utils/constants"
)

var (
	InitiallyActiveTime       = time.Date(2020, time.December, 5, 5, 0, 0, 0, time.UTC)
	UnscheduledActivationTime = time.Date(9999, time.December, 1, 0, 0, 0, 0, time.UTC)

	Mainnet = Config{
		ApricotPhase1Time:            InitiallyActiveTime,
		ApricotPhase2Time:            InitiallyActiveTime,
		ApricotPhase3Time:            InitiallyActiveTime,
		ApricotPhase4Time:            InitiallyActiveTime,
		ApricotPhase4MinPChainHeight: 0,
		ApricotPhase5Time:            InitiallyActiveTime,
		ApricotPhasePre6Time:         time.Date(2022, time.September, 8, 20, 0, 0, 0, time.UTC),
		ApricotPhase6Time:            time.Date(2022, time.September, 8, 22, 0, 0, 0, time.UTC),
		ApricotPhasePost6Time:        time.Date(2022, time.September, 9, 3, 0, 0, 0, time.UTC),
		BanffTime:                    time.Date(2022, time.December, 19, 16, 0, 0, 0, time.UTC),
		CortinaTime:                  time.Date(2023, time.August, 17, 10, 0, 0, 0, time.UTC),
		// The mainnet stop vertex is well known. It can be verified on any
		// fully synced node by looking at the parentID of the genesis block.
		//
		// Ref: https://subnets.avax.network/x-chain/block/0
		CortinaXChainStopVertexID: ids.FromStringOrPanic("ewiCzJQVJLYCzeFMcZSe9huX9h7QJPVeMdgDGcTVGTzeNJ3kY"),
		DurangoTime:               time.Date(2024, time.May, 6, 8, 0, 0, 0, time.UTC),
		EtnaTime:                  UnscheduledActivationTime,
	}
	Tahoe = Config{
		ApricotPhase1Time:            InitiallyActiveTime,
		ApricotPhase2Time:            InitiallyActiveTime,
		ApricotPhase3Time:            InitiallyActiveTime,
		ApricotPhase4Time:            InitiallyActiveTime,
		ApricotPhase4MinPChainHeight: 0,
		ApricotPhase5Time:            InitiallyActiveTime,
		ApricotPhasePre6Time:         time.Date(2022, time.September, 8, 20, 0, 0, 0, time.UTC),
		ApricotPhase6Time:            time.Date(2022, time.September, 8, 22, 0, 0, 0, time.UTC),
		ApricotPhasePost6Time:        time.Date(2022, time.September, 9, 3, 0, 0, 0, time.UTC),
		BanffTime:                    time.Date(2022, time.December, 12, 14, 0, 0, 0, time.UTC),
		CortinaTime:                  time.Date(2023, time.June, 28, 15, 0, 0, 0, time.UTC),
		// The fuji stop vertex is well known. It can be verified on any fully
		// synced node by looking at the parentID of the genesis block.
		//
		// Ref: https://subnets-test.avax.network/x-chain/block/0
		CortinaXChainStopVertexID: ids.FromStringOrPanic("RdWKZYgjgU2NicKHv8mpkR6jgo41W5aNwVhsX5sJgqshDAbQk"),
		DurangoTime:               time.Date(2024, time.April, 4, 0, 0, 0, 0, time.UTC),
		EtnaTime:                  UnscheduledActivationTime,
	}
	Default = Config{
		ApricotPhase1Time:            InitiallyActiveTime,
		ApricotPhase2Time:            InitiallyActiveTime,
		ApricotPhase3Time:            InitiallyActiveTime,
		ApricotPhase4Time:            InitiallyActiveTime,
		ApricotPhase4MinPChainHeight: 0,
		ApricotPhase5Time:            InitiallyActiveTime,
		ApricotPhasePre6Time:         InitiallyActiveTime,
		ApricotPhase6Time:            InitiallyActiveTime,
		ApricotPhasePost6Time:        InitiallyActiveTime,
		BanffTime:                    InitiallyActiveTime,
		CortinaTime:                  InitiallyActiveTime,
		CortinaXChainStopVertexID:    ids.Empty,
		DurangoTime:                  InitiallyActiveTime,
		// Etna is left unactivated by default on local networks. It can be configured to
		// activate by overriding the activation time in the upgrade file.
		EtnaTime: UnscheduledActivationTime,
	}

	ErrInvalidUpgradeTimes = errors.New("invalid upgrade configuration")
)

type Config struct {
	ApricotPhase1Time            time.Time `json:"apricotPhase1Time"`
	ApricotPhase2Time            time.Time `json:"apricotPhase2Time"`
	ApricotPhase3Time            time.Time `json:"apricotPhase3Time"`
	ApricotPhase4Time            time.Time `json:"apricotPhase4Time"`
	ApricotPhase4MinPChainHeight uint64    `json:"apricotPhase4MinPChainHeight"`
	ApricotPhase5Time            time.Time `json:"apricotPhase5Time"`
	ApricotPhasePre6Time         time.Time `json:"apricotPhasePre6Time"`
	ApricotPhase6Time            time.Time `json:"apricotPhase6Time"`
	ApricotPhasePost6Time        time.Time `json:"apricotPhasePost6Time"`
	BanffTime                    time.Time `json:"banffTime"`
	CortinaTime                  time.Time `json:"cortinaTime"`
	CortinaXChainStopVertexID    ids.ID    `json:"cortinaXChainStopVertexID"`
	DurangoTime                  time.Time `json:"durangoTime"`
	EtnaTime                     time.Time `json:"etnaTime"`
}

func (c *Config) Validate() error {
	upgrades := []time.Time{
		c.ApricotPhase1Time,
		c.ApricotPhase2Time,
		c.ApricotPhase3Time,
		c.ApricotPhase4Time,
		c.ApricotPhase5Time,
		c.ApricotPhasePre6Time,
		c.ApricotPhase6Time,
		c.ApricotPhasePost6Time,
		c.BanffTime,
		c.CortinaTime,
		c.DurangoTime,
		c.EtnaTime,
	}
	for i := 0; i < len(upgrades)-1; i++ {
		if upgrades[i].After(upgrades[i+1]) {
			return fmt.Errorf("%w: upgrade %d (%s) is after upgrade %d (%s)",
				ErrInvalidUpgradeTimes,
				i,
				upgrades[i],
				i+1,
				upgrades[i+1],
			)
		}
	}
	return nil
}

func (c *Config) IsApricotPhase1Activated(t time.Time) bool {
	return !t.Before(c.ApricotPhase1Time)
}

func (c *Config) IsApricotPhase2Activated(t time.Time) bool {
	return !t.Before(c.ApricotPhase2Time)
}

func (c *Config) IsApricotPhase3Activated(t time.Time) bool {
	return !t.Before(c.ApricotPhase3Time)
}

func (c *Config) IsApricotPhase4Activated(t time.Time) bool {
	return !t.Before(c.ApricotPhase4Time)
}

func (c *Config) IsApricotPhase5Activated(t time.Time) bool {
	return !t.Before(c.ApricotPhase5Time)
}

func (c *Config) IsApricotPhasePre6Activated(t time.Time) bool {
	return !t.Before(c.ApricotPhasePre6Time)
}

func (c *Config) IsApricotPhase6Activated(t time.Time) bool {
	return !t.Before(c.ApricotPhase6Time)
}

func (c *Config) IsApricotPhasePost6Activated(t time.Time) bool {
	return !t.Before(c.ApricotPhasePost6Time)
}

func (c *Config) IsBanffActivated(t time.Time) bool {
	return !t.Before(c.BanffTime)
}

func (c *Config) IsCortinaActivated(t time.Time) bool {
	return !t.Before(c.CortinaTime)
}

func (c *Config) IsDurangoActivated(t time.Time) bool {
	return !t.Before(c.DurangoTime)
}

func (c *Config) IsEtnaActivated(t time.Time) bool {
	return !t.Before(c.EtnaTime)
}

func GetConfig(networkID uint32) Config {
	switch networkID {
	case constants.MainnetID:
		return Mainnet
	case constants.TahoeID:
		return Tahoe
	default:
		return Default
	}
}
