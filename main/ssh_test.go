package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"net"
	"time"
)

type ShellCli struct {
	IP         string       //IP地址
	Username   string       //用户名
	Password   string       //密码
	Port       int          //端口号
	session    *ssh.Session //sshSession
	LastResult string       //最近一次Run的结果
}

func sshConnect(user, password, host string, port int) (*ssh.Session, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		session      *ssh.Session
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	clientConfig = &ssh.ClientConfig{
		User:    user,
		Auth:    auth,
		Timeout: 30 * time.Second,
		//需要验证服务端，不做验证返回nil就可以，点击HostKeyCallback看源码就知道了
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	// connet to ssh
	addr = fmt.Sprintf("%s:%d", host, port)
	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}
	// create session
	if session, err = client.NewSession(); err != nil {
		return nil, err
	}
	return session, nil
}

//执行shell
func (c ShellCli) RunShell(shell string) (string, error) {
	c.session, _ = sshConnect(c.Username, c.Password, c.IP, c.Port)
	buf, err := c.session.CombinedOutput(shell)
	c.LastResult = string(buf)
	return c.LastResult, err
}

func NewShell(ip string, username string, password string, port ...int) *ShellCli {
	cli := new(ShellCli)
	cli.IP = ip
	cli.Username = username
	cli.Password = password
	if len(port) <= 0 {
		cli.Port = 22
	} else {
		cli.Port = port[0]
	}
	return cli
}
