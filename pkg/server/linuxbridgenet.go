package server

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"

	"github.com/tangfeixiong/go-to-docker/pb"
	"github.com/tangfeixiong/go-to-docker/pb/moby"
	"github.com/tangfeixiong/go-to-docker/pkg/util"
	"github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/kopit/pkg/netutils"
)

func (m *myService) snoopBridgedNetworkLandscape(req *pb.BridgedNetworkingData) (*pb.BridgedNetworkingData, error) {
	resp := new(pb.BridgedNetworkingData)
	resp.LinuxBridges = make([]*pb.BridgedNetworkingData_LinuxBridgeInfo, 0)
	resp.ContainersNetworking = &pb.BridgedNetworkingData_ContainersNetworkingInfo{
		NetworkResources: make([]*moby.NetworkResource, 0),
		ContainersInfo:   make([]*moby.ContainerJSON, 0),
	}
	resp.VethPairs = make(map[string]string)

	if req == nil || len(req.LinuxBridges) == 0 {
		return resp, fmt.Errorf("At leasst one bridge required")
	}

	brs, err := netutils.Execute_brctl_show()
	if err != nil {
		return resp, fmt.Errorf("Failed to gain Linux Bridges: %s", err.Error())
	}

	for _, v := range brs {
		for _, v1 := range req.LinuxBridges {
			if v.Id == v1.Id || v.Name == v1.Name {
				resp.LinuxBridges = append(resp.LinuxBridges, &pb.BridgedNetworkingData_LinuxBridgeInfo{
					Id:              v.Id,
					Name:            v.Name,
					StpEnabled:      v.STPEnabled,
					Interfaces:      v.Interfaces,
					MacInfo:         make([]*pb.BridgedNetworkingData_LinuxBridgeLearnedMac, 0),
					IpAddressesInfo: make([]*pb.BridgedNetworkingData_IPAddressInfo, 0),
				})
				macs, err := netutils.Execute_brctl_showmacs(v.Name)
				if err != nil {
					return resp, fmt.Errorf("Failed to gain MACs in Linux Bridges %s: %s", v.Name, err.Error())
				}
				for _, v2 := range macs {
					l := len(resp.LinuxBridges)
					resp.LinuxBridges[l-1].MacInfo = append(resp.LinuxBridges[l-1].MacInfo, &pb.BridgedNetworkingData_LinuxBridgeLearnedMac{
						PortNo:      v2.PortNo,
						MacAddr:     v2.MacAddr,
						IsLocal:     v2.IsLocal,
						AgeingTimer: v2.AgeingTimer,
					})
				}
			}
		}
	}

	ips, err := netutils.Execute_ip_addr_show()
	if err != nil {
		return resp, fmt.Errorf("Failed to gain IP Addresses: %s", err.Error())
	}

	for _, v := range ips {
		for _, lbr := range resp.LinuxBridges {
			if v.Name == lbr.Name {
				lbr.IpAddressesInfo = append([]*pb.BridgedNetworkingData_IPAddressInfo{
					&pb.BridgedNetworkingData_IPAddressInfo{
						LinkInfo: &pb.BridgedNetworkingData_LinkLayerInfo{
							Index:            v.Id,
							Name:             v.Name,
							DataLinkStatus:   v.DataLinkStatus,
							DataLinkConf:     v.DataLinkConf,
							DataLinkFrame:    v.DataLinkFrame,
							DataLinkEtherMac: v.DataLinkEtherMAC,
							DataLinkEtherBrd: v.DataLinkEtherBRD,
							DataLinkNetnsId:  v.DataLinkNetnsID,
						},
						Ipv4:       v.IPv4,
						V4Mask:     v.V4Mask,
						V4Info:     v.V4Info,
						V4Lifetime: v.V4Lifetime,
						Ipv6:       v.IPv6,
						V6Mask:     v.V6Mask,
						V6Info:     v.V6Info,
						V6Lifetime: v.V6Lifetime,
					},
				}, lbr.IpAddressesInfo...)
			} else {
				for _, veth := range lbr.Interfaces {
					if strings.HasPrefix(v.Name, veth) {
						lbr.IpAddressesInfo = append(lbr.IpAddressesInfo, &pb.BridgedNetworkingData_IPAddressInfo{
							LinkInfo: &pb.BridgedNetworkingData_LinkLayerInfo{
								Index:            v.Id,
								Name:             v.Name,
								DataLinkStatus:   v.DataLinkStatus,
								DataLinkConf:     v.DataLinkConf,
								DataLinkFrame:    v.DataLinkFrame,
								DataLinkEtherMac: v.DataLinkEtherMAC,
								DataLinkEtherBrd: v.DataLinkEtherBRD,
								DataLinkNetnsId:  v.DataLinkNetnsID,
							},
							Ipv4:       v.IPv4,
							V4Mask:     v.V4Mask,
							V4Info:     v.V4Info,
							V4Lifetime: v.V4Lifetime,
							Ipv6:       v.IPv6,
							V6Mask:     v.V6Mask,
							V6Info:     v.V6Info,
							V6Lifetime: v.V6Lifetime,
						})
					}
				}
			}
		}
	}

	resultnetworking, err := m.reapDockerNetworking(&pb.DockerNetworkData{})
	if err != nil {
		util.Logger.Println(err)
		return resp, err
	}
	if resultnetworking.StateCode != 0 {
		util.Logger.Println(resultnetworking.StateMessage)
		return resp, err
	}

	for _, v := range resultnetworking.NetworkResources {
		for _, br := range resp.LinuxBridges {
			if v.Name == "bridge" && br.Name == "docker0" || strings.HasSuffix(br.Name, v.Id) {
				resp.ContainersNetworking.NetworkResources = append(resp.ContainersNetworking.NetworkResources, v)
				for key, _ := range v.Containers {
					resultinspection, err := m.inspectContainer(&pb.DockerContainerInspection{
						ContainerInfo: &moby.ContainerJSON{
							ContainerJsonBase: &moby.ContainerJSONBase{
								Id: key,
							},
						},
					})
					if err != nil {
						util.Logger.Println(err)
						return resp, err
					}
					resp.ContainersNetworking.ContainersInfo = append(resp.ContainersNetworking.ContainersInfo, resultinspection.ContainerInfo)
					if resultinspection.ContainerInfo == nil || resultinspection.ContainerInfo.ContainerJsonBase == nil || resultinspection.ContainerInfo.ContainerJsonBase.State == nil || resultinspection.ContainerInfo.ContainerJsonBase.State.Pid < 1 {
						util.Logger.Println("container PID required")
						return resp, fmt.Errorf("Container PID required for ns")
					}

					if v.Ipam == nil || len(v.Ipam.Config) == 0 || v.Ipam.Config[0].Gateway == "" {
						util.Logger.Println("gateway required")
						return resp, fmt.Errorf("Gateway required to find container ns")
					}
					resultips, err := nsenterContainerNetworkingViaSSH(v.Ipam.Config[0].Gateway, int(resultinspection.ContainerInfo.ContainerJsonBase.State.Pid))
					for _, v := range resultips {
						if v.Id != "1" && v.Name != "lo" && !strings.HasPrefix(v.Name, "loop") {
							resp.ContainersNetworking.AddressesInfo = append(resp.ContainersNetworking.AddressesInfo, &pb.BridgedNetworkingData_IPAddressInfo{
								LinkInfo: &pb.BridgedNetworkingData_LinkLayerInfo{
									Index:            v.Id,
									Name:             v.Name,
									DataLinkStatus:   v.DataLinkStatus,
									DataLinkConf:     v.DataLinkConf,
									DataLinkFrame:    v.DataLinkFrame,
									DataLinkEtherMac: v.DataLinkEtherMAC,
									DataLinkEtherBrd: v.DataLinkEtherBRD,
									DataLinkNetnsId:  v.DataLinkNetnsID,
								},
								Ipv4:       v.IPv4,
								V4Mask:     v.V4Mask,
								V4Info:     v.V4Info,
								V4Lifetime: v.V4Lifetime,
								Ipv6:       v.IPv6,
								V6Mask:     v.V6Mask,
								V6Info:     v.V6Info,
								V6Lifetime: v.V6Lifetime,
							})
							for i := 1; i < len(br.IpAddressesInfo); i++ {
								if strings.HasSuffix(br.IpAddressesInfo[i].LinkInfo.Name, v.Id) && strings.HasSuffix(v.Name, br.IpAddressesInfo[i].LinkInfo.Index) {
									resp.VethPairs[br.IpAddressesInfo[i].LinkInfo.Name] = v.Name
									continue
								}
							}
							fmt.Println("veth orphan:", v.Name)
						}
					}
				}
			}
		}
	}

	return resp, nil
}

const (
	datalink_Line1_regexp = `([0-9]+): ([\w-@]+): <([\w-,]+)> ([\w ]+)`                                         // 2: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP group default qlen 1000
	datalink_line2_regexp = `\s+link/(\w+) (([0-9a-f]{2}:?){6}) brd (([0-9a-f]{2}:?){6})( link-netnsid (\w+))?` //     link/ether 08:00:27:46:54:e7 brd ff:ff:ff:ff:ff:ff
	//     link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
	//     link/ether ce:68:9a:f1:bf:7c brd ff:ff:ff:ff:ff:ff link-netnsid 1
	net_ipv4_regexp = `\s+inet ((\d{1,3}\.?){4})/(\d\d) ((\w ?)+)` //     inet 172.17.0.1/22 scope global docker0
	net_ipv6_regexp = `\s+inet6 ([0-9a-f:]+)/(\d\d) ([\w ]+)`      //     inet6 fe80::42:1ff:fe74:cc7e/64 scope link
	net_lft_regexp  = `\s+([\w ]+)`                                //        valid_lft forever preferred_lft forever

)

func nsenterContainerNetworkingViaSSH(host string, pid int) ([]*netutils.IPAddress, error) {
	scc := gainSSHConnectionConfig()
	sshConfig := &ssh.ClientConfig{
		User: scc.SSHUser.Name,
		Auth: []ssh.AuthMethod{
		// ssh.Password("your_password"),
		// PublicKeyFile("/path/to/your/pub/certificate/key")
		// SSHAgent(),
		},
	}
	if scc.SSHAgent {
		sshConfig.Auth = append(sshConfig.Auth, util.SSHAgent())
	}
	if len(scc.SSHUser.Password) != 0 {
		sshConfig.Auth = append(sshConfig.Auth, ssh.Password(scc.SSHUser.Password))
	} else {
		if len(scc.SSHUser.RSAKeyContent) != 0 {
			sshConfig.Auth = append(sshConfig.Auth, util.PublicKeys(scc.SSHUser.RSAKeyContent))
		}
	}

	client := &util.SSHClient{
		Config: sshConfig,
		Host:   host,
		Port:   22,
	}

	inbuf := bytes.Buffer{}
	outbuf := new(bytes.Buffer)
	errbuf := new(bytes.Buffer)

	path := strings.Join([]string{
		"sudo",
		"nsenter", "-t", strconv.Itoa(pid), "-n",
		"--",
		"ip", "addr", "show",
	}, " ")

	cmd := &util.SSHCommand{
		Path:   path,
		Env:    []string{},
		Stdin:  &inbuf, // os.Stdin,
		Stdout: outbuf, // os.Stdout,
		Stderr: errbuf, // os.Stderr,
	}

	data := make([]*netutils.IPAddress, 0)
	util.Logger.Println("Running command:", cmd.Path)
	if err := client.RunCommand(cmd); err != nil {
		fmt.Fprintf(os.Stderr, "command run error: %s\n", err)
		return data, fmt.Errorf("Failed to call via SSH: %v", err)
	}
	time.Sleep(time.Second * 5)

	resultlines := []string{}
	if errbuf.Len() > 0 {
		scanner := bufio.NewScanner(errbuf)
		for scanner.Scan() {
			resultlines = append(resultlines, scanner.Text())
		}
		if err := scanner.Err(); err != nil && err != io.EOF {
			fmt.Fprintln(os.Stderr, "reading input:", err)
		} else {
			fmt.Fprintln(os.Stderr, "Execution failed!", resultlines)
		}
		return data, fmt.Errorf(strings.Join(resultlines, "\n"))
	}
	if outbuf.Len() > 0 {
		var resp *netutils.IPAddress
		var lftv4 bool
		scanner := bufio.NewScanner(outbuf)
		for i := 0; scanner.Scan(); i++ {
			if nil == scanner.Err() && 0 != len(scanner.Text()) {
				resultlines = append(resultlines, scanner.Text())
				line := scanner.Text()
				switch {
				case strings.HasPrefix(line, "    inet "):
					lftv4 = true
					re := regexp.MustCompile(net_ipv4_regexp)
					result := re.FindAllStringSubmatch(line, -1)
					if nil != result {
						resp.IPv4 = result[0][1]
						resp.V4Mask = result[0][3]
						resp.V4Info = strings.Split(result[0][4], " ")
					}
				case strings.HasPrefix(line, "    inet6 "):
					lftv4 = false
					re := regexp.MustCompile(net_ipv6_regexp)
					result := re.FindAllStringSubmatch(line, -1)
					if nil != result {
						resp.IPv6 = result[0][1]
						resp.V6Mask = result[0][2]
						resp.V6Info = strings.Split(result[0][3], " ")
					}
				case strings.HasPrefix(line, "       "):
					if lftv4 {
						resp.V4Lifetime = strings.Split(strings.TrimLeft(line, " "), " ")
					} else {
						resp.V6Lifetime = strings.Split(strings.TrimLeft(line, " "), " ")
					}
				default:
					if 0 < i {
						data = append(data, resp)
					}
					resp = new(netutils.IPAddress)
					re := regexp.MustCompile(datalink_Line1_regexp)
					result := re.FindAllStringSubmatch(line, -1)
					if nil != result {
						resp.Id = result[0][1]
						resp.Name = result[0][2]
						resp.DataLinkStatus = strings.Split(result[0][3], ",")
						//				resp.DataLinkConf = make(map[string]string)
						conf := strings.Split(result[0][4], " ")
						//				for i := 0; i < len(conf)-1; i += 2 {
						//					resp.DataLinkConf[conf[i]] = conf[i+1]
						//				}
						resp.DataLinkConf = conf
						if scanner.Scan() && nil == scanner.Err() {
							line = scanner.Text()
							re = regexp.MustCompile(datalink_line2_regexp)
							result = re.FindAllStringSubmatch(line, -1)
							resp.DataLinkFrame = result[0][1]
							if "ether" == resp.DataLinkFrame || "loopback" == resp.DataLinkFrame {
								resp.DataLinkEtherMAC = result[0][2]
								resp.DataLinkEtherBRD = result[0][4]
								if len(result[0]) == 8 {
									resp.DataLinkNetnsID = result[0][7]
								}
							}
						}
					}
				}
			} else if io.EOF == scanner.Err() {
				break
			} else {
				fmt.Fprintln(os.Stderr, "reading input:", scanner.Err())
				return data, fmt.Errorf("Reading input error: %v", scanner.Err())
			}
		}
		data = append(data, resp)
	}
	fmt.Fprintln(os.Stdout, resultlines)
	return data, nil
}

func gainSSHConnectionConfig() util.SSHConnectionConfig {
	username := os.Getenv("NSH_USERNAME")
	password := os.Getenv("NSH_PASSWORD")
	b64key := os.Getenv("NSH_B64KEY")

	scc := util.SSHConnectionConfig{}
	if 0 != len(username) && username != "fake user" {
		scc.SSHUser.Name = username
	}
	if 0 != len(password) && password != "fake secret" {
		scc.SSHUser.Password = password
	}
	if 0 != len(b64key) {
		b, err := base64.StdEncoding.DecodeString(b64key)
		if err != nil {
			fmt.Println("Failed to decode RSA private key:", err.Error())
		} else {
			scc.SSHUser.RSAKeyContent = b
		}
	}
	if 0 == len(scc.SSHUser.Name) {
		scc.SSHUser.Name = "root"
	}
	if 0 == len(scc.SSHUser.RSAKeyContent) {
		scc.SSHUser.RSAKeyContent = []byte(util.Key_vagrant)
		scc.SSHUser.RSAPubContent = []byte(util.Pub_vagrant)
	}
	return scc
}
