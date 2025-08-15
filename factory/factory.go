// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

/*
 * PCF Configuration Factory
 */

package factory

import (
	"fmt"
	"os"

	"github.com/omec-project/pcf/logger"
	"gopkg.in/yaml.v2"
)

var PcfConfig Config

// TODO: Support configuration update from REST api
func InitConfigFactory(f string) error {
	if content, err := os.ReadFile(f); err != nil {
		return err
	} else {
		PcfConfig = Config{}

		if yamlErr := yaml.Unmarshal(content, &PcfConfig); yamlErr != nil {
			return yamlErr
		}
		if PcfConfig.Configuration.WebuiUri == "" {
			PcfConfig.Configuration.WebuiUri = "webui:9876"
		}
	}

	if PcfConfig.Configuration.ManualConfigs != nil {
		logger.CfgLog.Infof("Manual Configuration provided for network functions")
		for nfType, nfs := range PcfConfig.Configuration.ManualConfigs.NFs {
			for _, nf := range nfs {
				logger.CfgLog.Debugf("Manual Configuration - NF Type: %s, Name: %s, URL: %s, Port: %d", nfType, nf.NfInstanceName, nf.NfServices)
			}
		}
	} else {
		logger.CfgLog.Infof("No manual configuration provided for network functions")
	}

	return nil
}

func CheckConfigVersion() error {
	currentVersion := PcfConfig.GetVersion()

	if currentVersion != PCF_EXPECTED_CONFIG_VERSION {
		return fmt.Errorf("config version is [%s], but expected is [%s]",
			currentVersion, PCF_EXPECTED_CONFIG_VERSION)
	}

	logger.CfgLog.Infof("config version [%s]", currentVersion)

	return nil
}
