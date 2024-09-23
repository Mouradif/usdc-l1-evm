// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package usdcreceiver

import (
	"github.com/ava-labs/subnet-evm/precompile/precompileconfig"
)

var _ precompileconfig.Config = &Config{}

// Config implements the precompileconfig.Config interface while adding in the
// ContractNativeMinter specific precompile config.
type Config struct {
	precompileconfig.Upgrade
}

// NewConfig returns a config for a network upgrade at [blockTimestamp] that enables
// ContractNativeMinter with the given [admins], [enableds] and [managers] as members of the allowlist.
// Also mints balances according to [initialMint] when the upgrade activates.
func NewConfig(blockTimestamp *uint64) *Config {
	return &Config{
		Upgrade: precompileconfig.Upgrade{BlockTimestamp: blockTimestamp},
	}
}

// NewDisableConfig returns config for a network upgrade at [blockTimestamp]
// that disables ContractNativeMinter.
func NewDisableConfig(blockTimestamp *uint64) *Config {
	return &Config{
		Upgrade: precompileconfig.Upgrade{
			BlockTimestamp: blockTimestamp,
			Disable:        true,
		},
	}
}

// Key returns the key for the ContractNativeMinter precompileconfig.
// This should be the same key as used in the precompile module.
func (*Config) Key() string { return ConfigKey }

// Equal returns true if [cfg] is a [*ContractNativeMinterConfig] and it has been configured identical to [c].
func (c *Config) Equal(s precompileconfig.Config) bool {
	other, ok := (s).(*Config)
	if !ok {
		return false
	}
	// modify this boolean accordingly with your custom Config, to check if [other] and the current [c] are equal
	// if Config contains only Upgrade you can skip modifying it.
	equals := c.Upgrade.Equal(&other.Upgrade)
	return equals
}

func (c *Config) Verify(chainConfig precompileconfig.ChainConfig) error {
	return nil
}
