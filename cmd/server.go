// Copyright © 2019 The Vultr-cli Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/vultr/vultr-cli/cmd/printer"
)

// Server represents the server command
func Server() *cobra.Command {
	serverCmd := &cobra.Command{
		Use:   "server",
		Short: "A brief description of your command",
		Long:  ``,
	}

	serverCmd.AddCommand(serverStart, serverStop, serverRestart, serverReinstall, serverTag, serverDelete, serverLabel, serverBandwidth, serverIPV4List, serverIPV6List, serverList, serverInfo, updateFwgGroup)

	serverTag.Flags().StringP("tag", "t", "", "tag you want to set for a given instance")
	serverTag.MarkFlagRequired("tag")

	serverLabel.Flags().StringP("label", "l", "", "label you want to set for a given instance")
	serverLabel.MarkFlagRequired("label")

	serverIPV4List.Flags().StringP("public", "p", "", "include information about the public network adapter : True or False")

	updateFwgGroup.Flags().StringP("instance-id", "i", "", "instance id of the server you want to use")
	updateFwgGroup.Flags().StringP("firewall-group-id", "f", "", "firewall group id that you want to assign. 0 Value will unset the firewall-group")
	updateFwgGroup.MarkFlagRequired("instance-id")
	updateFwgGroup.MarkFlagRequired("firewall-group-id")

	osCmd := &cobra.Command{
		Use:   "os",
		Short: "operating system commands ",
		Long:  ``,
	}

	osCmd.AddCommand(osList, osUpdate)
	osUpdate.Flags().StringP("os", "o", "", "operating system ID you wish to use")
	osUpdate.MarkFlagRequired("os")
	// Sub commands for OS
	serverCmd.AddCommand(osCmd)

	return serverCmd
}

var serverStart = &cobra.Command{
	Use:   "start <instanceID>",
	Short: "starts a server",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("please provide an instanceID")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]

		err := client.Server.Start(context.TODO(), id)

		if err != nil {
			fmt.Printf("error starting server : %v", err)
			os.Exit(1)
		}

		fmt.Println("Started up server")
	},
}

var serverStop = &cobra.Command{
	Use:   "stop <instanceID>",
	Short: "stops a server",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("please provide an instanceID")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]

		err := client.Server.Halt(context.TODO(), id)

		if err != nil {
			fmt.Printf("error stopping server : %v", err)
			os.Exit(1)
		}

		fmt.Println("Stopped the server")
	},
}

var serverRestart = &cobra.Command{
	Use:   "restart <instanceID>",
	Short: "restart a server",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("please provide an instanceID")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]

		err := client.Server.Reboot(context.TODO(), id)

		if err != nil {
			fmt.Printf("error rebooting server : %v", err)
			os.Exit(1)
		}

		fmt.Println("Rebooted server")
	},
}

var serverReinstall = &cobra.Command{
	Use:   "reinstall <instanceID>",
	Short: "reinstall a server",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("please provide an instanceID")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]

		err := client.Server.Reinstall(context.TODO(), id)

		if err != nil {
			fmt.Printf("error reinstalling server : %v", err)
			os.Exit(1)
		}

		fmt.Println("Reinstalled server")
	},
}

var serverTag = &cobra.Command{
	Use:   "tag <instanceID>",
	Short: "add/modify tag on server",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("please provide an instanceID")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		tag, _ := cmd.Flags().GetString("tag")
		err := client.Server.SetTag(context.TODO(), id, tag)

		if err != nil {
			fmt.Printf("error adding tag to server : %v", err)
			os.Exit(1)
		}

		fmt.Printf("Tagged server with : %s", tag)
	},
}

var serverDelete = &cobra.Command{
	Use:   "delete <instanceID>",
	Short: "delete a server",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("please provide an instanceID")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		err := client.Server.Delete(context.TODO(), id)

		if err != nil {
			fmt.Printf("error deleting server : %v", err)
			os.Exit(1)
		}

		fmt.Println("Deleted server")
	},
}

var serverLabel = &cobra.Command{
	Use:   "label <instanceID>",
	Short: "label a server",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("please provide an instanceID")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		label, _ := cmd.Flags().GetString("label")
		err := client.Server.SetLabel(context.TODO(), id, label)

		if err != nil {
			fmt.Printf("error labeling server : %v", err)
			os.Exit(1)
		}

		fmt.Printf("Labeled server with : %s", label)
	},
}

var serverBandwidth = &cobra.Command{
	Use:   "bandwidth <instanceID>",
	Short: "bandwidth for server",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("please provide an instanceID")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		bw, err := client.Server.Bandwidth(context.TODO(), id)

		if err != nil {
			fmt.Printf("error getting bandwidth for server : %v", err)
			os.Exit(1)
		}

		printer.ServerBandwidth(bw)
	},
}

var serverIPV4List = &cobra.Command{
	Use:     "list-ipv4 <instanceID>",
	Aliases: []string{"v4"},
	Short:   "list ipv4 for a server",
	Long:    ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("please provide an instanceID")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		public, _ := cmd.Flags().GetString("public")

		pub := false
		if strings.ToLower(public) == "true" {
			pub = true
		}

		v4, err := client.Server.IPV4Info(context.TODO(), id, pub)

		if err != nil {
			fmt.Printf("error getting ipv4 info : %v", err)
			os.Exit(1)
		}

		printer.ServerIPV4(v4)
	},
}

var serverIPV6List = &cobra.Command{
	Use:     "list-ipv6 <instanceID>",
	Aliases: []string{"v6"},
	Short:   "list ipv6 for a server",
	Long:    ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("please provide an instanceID")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]

		v6, err := client.Server.IPV6Info(context.TODO(), id)

		if err != nil {
			fmt.Printf("error getting ipv6 info : %v", err)
			os.Exit(1)
		}

		printer.ServerIPV6(v6)
	},
}

var serverList = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "list all available servers",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		s, err := client.Server.List(context.TODO())

		if err != nil {
			fmt.Printf("error getting list of servers : %v", err)
			os.Exit(1)
		}

		printer.ServerList(s)
	},
}

var serverInfo = &cobra.Command{
	Use:   "info <instanceID>",
	Short: "info about a specific server",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("please provide an instanceID")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		s, err := client.Server.GetServer(context.TODO(), id)

		if err != nil {
			fmt.Printf("error getting server : %v", err)
			os.Exit(1)
		}

		printer.ServerInfo(s)
	},
}

var updateFwgGroup = &cobra.Command{
	Use:   "update-firewall-group",
	Short: "assign a firewall group to server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("instance-id")
		fwgID, _ := cmd.Flags().GetString("firewall-group-id")

		err := client.Server.SetFirewallGroup(context.TODO(), id, fwgID)

		if err != nil {
			fmt.Printf("error setting firewall group : %v", err)
			os.Exit(1)
		}

		fmt.Println("Added firewall-group")
	},
}

var osList = &cobra.Command{
	Use:   "list <instanceID>",
	Short: "list available operating systems this instance can be changed to",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("please provide an instanceID")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		o, err := client.Server.ListOS(context.TODO(), id)

		if err != nil {
			fmt.Printf("error getting os list : %v", err)
			os.Exit(1)
		}

		printer.OsList(o)
	},
}

var osUpdate = &cobra.Command{
	Use:   "change <instanceID>",
	Short: "changes operating system",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("please provide an instanceID")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		osID, _ := cmd.Flags().GetString("os")

		err := client.Server.ChangeOS(context.TODO(), id, osID)

		if err != nil {
			fmt.Printf("error updating os : %v", err)
			os.Exit(1)
		}

		fmt.Println("Updated OS")
	},
}
//backup                 get and set backup schedules
//create                 create a new virtual machine
//os                     show and change OS on a virtual machine
//app                    show and change application on a virtual machine
//iso                    attach/detach ISO of a virtual machine
//restore                restore from backup/snapshot
//create-ipv4            add a new IPv4 address to a virtual machine
//delete-ipv4            remove IPv4 address from a virtual machine
//reverse-dns            modify reverse DNS entries
//upgrade-plan           upgrade plan of a virtual machine
