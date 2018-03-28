package sshclient

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type CONClient struct {
	SshClient *ssh.Client
	Hostname  string
	Addr      string
}

func New(ip, user, pass string, port int, hostname string) *CONClient {
	sshclient, err := connect(ip, user, pass, port)
	if err != nil {
		log.Println(err)
		return nil
	}
	var conclient = &CONClient{}
	conclient.SshClient = sshclient
	conclient.Hostname = hostname
	conclient.Addr = ip
	return conclient
}

func (this *CONClient) session() *ssh.Session {
	session, err := this.SshClient.NewSession()
	if err != nil {
		log.Println(err)
		return nil
	}
	return session
}

func (this *CONClient) Execute(module, cmd string) error {
	if module == "cmd" {
		res, err := this.Exec(cmd)
		if err != nil {
			color.Cyan("\n------ %s [%s] ------", this.Hostname, this.Addr)
			color.Red("Command [%s] of host %s: %s\n", cmd, this.Addr, err)
			color.Green(string(res))
		} else {
			color.Cyan("\n------ %s [%s] ------\n", this.Hostname, this.Addr)
			color.Green(string(res))
		}
		return nil
	}
	filepath := strings.Split(cmd, " ")
	src := filepath[0]
	workdir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	if strings.HasPrefix(src, "workdir://") {
		src = strings.Replace(src, "workdir:/", workdir, 1)
	}
	dst := filepath[1]
	if strings.HasPrefix(dst, "workdir://") {
		dst = strings.Replace(dst, "workdir:/", workdir, 1)
	}
	if module == "sendfile" {
		err := this.SendFile(src, dst)
		if err != nil {
			return err
		}
	}
	if module == "getfile" {
		fmt.Println(dst)
		err := this.GetFile(src, dst)
		if err != nil {
			return err
		}
	}
	defer this.Close()
	return nil
}

func (this *CONClient) Exec(cmd string) ([]byte, error) {
	session := this.session()
	defer session.Close()
	res, err := session.CombinedOutput(cmd)
	return res, err
}

func (this *CONClient) addSftpClient() *sftp.Client {
	sftpClient, err := sftp.NewClient(this.SshClient)
	if err != nil {
		log.Fatal(err)
	}
	return sftpClient
}

func (this *CONClient) SendFile(src, dst string) error {
	sftpClient := this.addSftpClient()
	defer sftpClient.Close()
	srcFile, err := os.Open(src)
	if err != nil {
		log.Printf("open the dst file [%s] error\n", src)
		return err
	}
	defer srcFile.Close()

	dstFile, err := sftpClient.Create(dst)
	if err != nil {
		log.Printf("create the dst file [%s] error\n", dst)
		return err
	}
	defer dstFile.Close()

	buf := make([]byte, 1024)
	for {
		n, _ := srcFile.Read(buf)
		if n == 0 {
			break
		}
		dstFile.Write(buf)
	}

	color.Green("Send file %s to remote server [%s] finished!", src, this.Hostname)
	return nil
}

func (this *CONClient) GetFile(src, dst string) error {
	sftpClient := this.addSftpClient()
	defer sftpClient.Close()
	srcFile, err := sftpClient.Open(src)
	if err != nil {
		log.Printf("Try to open file [%s] error\n", src)
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		log.Printf("Try to create file [%s] error\n", dst)
		return err
	}
	defer dstFile.Close()

	if _, err = srcFile.WriteTo(dstFile); err != nil {
		log.Println("writeto dstfile error")
		return err
	}

	color.Green("copy file %s from remote server [%s] finished!", src, this.Hostname)

	return nil
}

func (this *CONClient) Close() {
	this.SshClient.Close()
}

func getHostKey(host string) (ssh.PublicKey, error) {
	file, err := os.Open(filepath.Join(os.Getenv("HOME"), ".ssh", "known_hosts"))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var hostKey ssh.PublicKey
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")
		if len(fields) != 3 {
			continue
		}
		if strings.Contains(fields[0], host) {
			var err error
			hostKey, _, _, _, err = ssh.ParseAuthorizedKey(scanner.Bytes())
			if err != nil {
				return nil, errors.New(fmt.Sprintf("error parsing %q: %v", fields[2], err))
			}
			break
		}
	}

	if hostKey == nil {
		return nil, errors.New(fmt.Sprintf("no hostkey for %s", host))
	}
	return hostKey, nil
}

func connect(ip, user, password string, port int) (*ssh.Client, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		err          error
	)
	hostKey, err := getHostKey(ip)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	clientConfig = &ssh.ClientConfig{
		User:            user,
		Auth:            auth,
		Timeout:         30 * time.Second,
		HostKeyCallback: ssh.FixedHostKey(hostKey),
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", ip, port)
	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}
	return client, nil
}
