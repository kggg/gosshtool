package sshclient

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"

	"github.com/pkg/sftp"
)

type SftpClient struct {
	Client
	Sftpclient sftp.Client
}

func Sftpinit() *SftpClient {
	return &SftpClient{
		//Client
	}
}

func (c *SftpClient) Copy(src, dst string) {
	srcFile, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	defer srcFile.Close()

	var remoteFileName = path.Base(src)
	dstFile, err := c.Sftpclient.Create(path.Join(dst, remoteFileName))
	if err != nil {
		log.Fatal(err)
	}
	defer dstFile.Close()

	buf := make([]byte, 1024)
	for {
		n, err := srcFile.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		if n == 0 {
			break
		}
		dstFile.Write(buf)
	}
	fmt.Println("copy file to remote server finished!")
}
