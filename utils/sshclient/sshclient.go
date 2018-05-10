package sshclient

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"github.com/fatih/color"
	"golang.org/x/crypto/ssh"
	"io"
	"io/ioutil"
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
		return nil
	}
	return session
}

func (this *CONClient) Execute(module, cmd string) error {
	switch module {
	case "cmd":
		res, err := this.Run(cmd)
		if err != nil {
			color.Cyan("\n------ %s [%s] ------", this.Hostname, this.Addr)
			color.Red("Running command [%s] failed: %s\n", cmd, err)
			color.Red(string(res))
		} else {
			color.Cyan("\n------ %s [%s] ------\n", this.Hostname, this.Addr)
			color.Green(string(res))
		}
		return nil
	case "copy":
		act := NewCopyFile()
		err := act.Parse(cmd)
		if err != nil {
			return err
		}
		err = this.Copy(act.Src, act.Dest, act.Mode)
		return err
	}
	return nil
}

func (this *CONClient) Run(cmd string) ([]byte, error) {
	session := this.session()
	defer session.Close()
	res, err := session.CombinedOutput(cmd)
	return res, err
}

func (c *CONClient) Copy(src, dest string, mode string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	directory := filepath.Dir(dest)
	if c.checkDirNotExists(directory) {
		//err = c.mkRemoteDir(directory)
		//if err != nil {
		return errors.New(fmt.Sprintf("Destination directory %s does not exist", directory))
		//}
	}
	err = c.CopyFile(srcFile, dest, mode)
	color.Green("Copy file %s to remote server [%s] finished!", src, c.Hostname)
	return nil
}

func (c *CONClient) CopyFile(fileReader io.Reader, remotePath string, permissions string) error {
	session := c.session()

	contents, _ := ioutil.ReadAll(fileReader)
	reader := bytes.NewReader(contents)

	size := int64(len(contents))
	filename := filepath.Base(remotePath)
	directory := filepath.Dir(remotePath)
	go func() {
		w, _ := session.StdinPipe()
		defer w.Close()
		fmt.Fprintln(w, "C"+permissions, size, filename)
		io.Copy(w, reader)
		fmt.Fprintln(w, "\x00")
	}()
	err := session.Run("/usr/bin/scp -qrt " + directory)
	return err

}

func (c *CONClient) checkDirNotExists(dir string) bool {
	session := c.session()
	err := session.Run("ls -d " + dir)
	if err != nil {
		return true
	}
	return false
}

func (c *CONClient) mkRemoteDir(dir string) error {
	session := c.session()
	err := session.Run("mkdir -p " + dir)
	return err
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
	//hostKey, err := getHostKey(ip)
	if err != nil {
		return nil, err
	}
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	clientConfig = &ssh.ClientConfig{
		User:    user,
		Auth:    auth,
		Timeout: 30 * time.Second,
		//HostKeyCallback: ssh.FixedHostKey(hostKey),
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", ip, port)
	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}
	return client, nil
}

type CopyFile struct {
	Src    string
	Dest   string
	Mode   string
	Force  bool
	Backup bool
	User   string
	Group  string
}

func NewCopyFile() *CopyFile {
	return &CopyFile{}
}

func (c *CopyFile) Parse(str string) error {
	if len(str) == 0 {
		return errors.New("copy params empty")
	}
	strlist := strings.Split(str, " ")
	for _, v := range strlist {
		keys := strings.Split(v, "=")
		switch keys[0] {
		case "src":
			c.Src = keys[1]
		case "dest":
			c.Dest = keys[1]
		case "mode":
			c.Mode = keys[1]
		case "force":
			if keys[1] == "no" {
				c.Force = false
			} else {
				c.Force = true
			}
		case "backup":
			if keys[1] == "yes" {
				c.Backup = true
			} else {
				c.Backup = false
			}
		case "user":
			c.User = keys[1]
		case "group":
			c.Group = keys[1]
		}
	}
	return nil
}
