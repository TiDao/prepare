package cryptogen

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"chainmaker.org/chainmaker-go/common/crypto"
	"github.com/spf13/cobra"
)

func ExtendCmd() *cobra.Command {
	extendCmd := &cobra.Command{
		Use:   "extend",
		Short: "Extend existing network",
		Long:  "Extend existing network",
		RunE: func(cmd *cobra.Command, args []string) error {
			return extend()
		},
	}
	extendCmd.Flags().StringVarP(&OutputDir, "output", "o", "crypto-config", "specify the output directory in which to place artifacts")
	return extendCmd
}

func extend() error {
	cryptoGenConfig := GetCryptoGenConfig()

	for _, item := range cryptoGenConfig.Item {
		for i := 0; i < int(item.Count); i++ {
			orgName := fmt.Sprintf("%s%d.%s", item.HostName, i+1, item.Domain)
			if item.Count == 1 {
				orgName = fmt.Sprintf("%s.%s", item.HostName, item.Domain)
			}
			keyType := crypto.AsymAlgoMap[strings.ToUpper(item.PKAlgo)]
			hashType := crypto.HashAlgoMap[strings.ToUpper(item.SKIHash)]

			//caPath := filepath.Join(OutputDir, orgName, "ca")
			//caKeyPath := filepath.Join(caPath, "ca.key")
			//caCertPath := filepath.Join(caPath, "ca.crt")
			//userPath := filepath.Join(OutputDir, orgName, "user")
			caPath := filepath.Join(OutputDir,"ca")
			caKey := fmt.Sprintf("ca%d.key",i+1)
			caCrt := fmt.Sprintf("ca%d.crt",i+1)
			caKeyPath := filepath.Join(caPath, caKey)
			caCertPath := filepath.Join(caPath, caCrt)
			userPath := filepath.Join(OutputDir, orgName, "user")
			nodePath := filepath.Join(OutputDir, orgName, "node")


			if _, err := os.Stat(caPath); os.IsNotExist(err) {
				caCN := fmt.Sprintf("ca.%s", orgName)
				caSANS := append(item.CA.Specs.SANS, caCN)
				if err := generateCA(caPath,caKey,caCrt,
					item.CA.Location.Country, item.CA.Location.Locality, item.CA.Location.Province, "root-cert", orgName, caCN,
					item.CA.Specs.ExpireYear, caSANS, keyType, hashType); err != nil {
					return err
				}
			}
			for _, user := range item.User {
				for j := 0; j < int(user.Count); j++ {
					name := fmt.Sprintf("%s%d", user.Type, j+1)
					path := filepath.Join(userPath, name)
					if _, err := os.Stat(path); os.IsNotExist(err) {
						err := generateUser(path, name, caKeyPath, caCertPath,
							user.Location.Country, user.Location.Locality, user.Location.Province, orgName, user.Type,
							user.ExpireYear, keyType, hashType)
						if err != nil {
							return err
						}
					}
				}
			}

			for _, node := range item.Node {
				for j := 0; j < int(node.Count); j++ {
					name := fmt.Sprintf("%s%d", node.Type, j+1)
					path := filepath.Join(nodePath, name)
					if _, err := os.Stat(path); os.IsNotExist(err) {
						err := generateNode(path, name, caKeyPath, caCertPath,
							node.Location.Country, node.Location.Locality, node.Location.Province, orgName, node.Type,
							node.Specs.ExpireYear, node.Specs.SANS, keyType, hashType)
						if err != nil {
							return err
						}
					}
				}
			}
		}
	}

	return nil
}
