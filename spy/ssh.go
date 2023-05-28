package spy

import (
	"bufio"
	"fmt"
	"github.com/a-dakani/logSpy/configs"
	"github.com/a-dakani/logSpy/logger"
	"github.com/jcmturner/gokrb5/v8/config"
	"golang.org/x/crypto/ssh"
	"io"
	"os"
	"sync"
)

type Spy struct {
	Service  configs.Service
	Client   *ssh.Client
	Sessions []*ssh.Session
}

func (spy *Spy) CreateClient() error {
	var err error

	//hostKeyCallback, err := knownhosts.New(filepath.Join(os.Getenv("HOME"), ".ssh/known_hosts"))
	//if err != nil {
	//	logger.Fatal(err.Error())
	//}

	conf := &ssh.ClientConfig{
		User: spy.Service.User,
		//TODO replace InsecureIgnoreHostKey with hostKeyCallback for production
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth:            []ssh.AuthMethod{},
	}

	//If private key path is provided, use it
	if spy.Service.PrivateKeyPath != "" {
		pemBytes, err := os.ReadFile(spy.Service.PrivateKeyPath)
		if err != nil {
			logger.Warning("Private key file does not exist")
			return err
		}
		// create signer
		signer, err := ssh.ParsePrivateKey(pemBytes)
		if err != nil {
			logger.Warning("Private key file is not valid")
			return err
		}
		conf.Auth = append(conf.Auth, ssh.PublicKeys(signer))
		logger.Info(fmt.Sprintf("[%s] Using private key %s ", spy.Service.Host, spy.Service.PrivateKeyPath))
	}
	//If krb5 conf path is provided, use it
	if spy.Service.Krb5ConfPath != "" {
		c, _ := config.Load(spy.Service.Krb5ConfPath)

		//FIXME Error handling is shit for wrong password or unreachable auth server

		sshGSSAPIClient, err := NewKrb5InitiatorClient(spy.Service.User, c)
		if err != nil {
			logger.Warning("Unable to create sshGSSAPIClient")
			return err
		}
		conf.Auth = append(conf.Auth, ssh.GSSAPIWithMICAuthMethod(&sshGSSAPIClient, spy.Service.Host))
		logger.Info(fmt.Sprintf(" [%s] Using krb5 conf %s", spy.Service.Host, spy.Service.Krb5ConfPath))
	}

	spy.Client, err = ssh.Dial("tcp", fmt.Sprintf("%s:%d", spy.Service.Host, spy.Service.Port), conf)
	if err != nil {
		logger.Warning(fmt.Sprintf("[%s] unable to connect: %s", spy.Service.Host, err))
		return err
	}

	return nil
}

func (spy *Spy) TailFiles() error {
	var wg sync.WaitGroup

	for index, file := range spy.Service.Files {
		wg.Add(1)
		sess, err := spy.Client.NewSession()
		if err != nil {
			return err
		}
		spy.Sessions = append(spy.Sessions, sess)

		sessStdOut, err := sess.StdoutPipe()
		if err != nil {
			return err
		}
		go func() {
			defer wg.Done()
			copyWithAppend(os.Stdout, sessStdOut, fmt.Sprintf("[%s] ", spy.Service.Files[index].Alias))
		}()

		sessStderr, err := sess.StderrPipe()
		if err != nil {
			return err
		}
		go func() {
			defer wg.Done()
			copyWithAppend(os.Stderr, sessStderr, fmt.Sprintf("[%s] ", spy.Service.Files[index].Alias))
		}()

		logger.Info(fmt.Sprintf("[%s] Tailing %s", spy.Service.Host, file.Path))
		go func(index int, path string) {
			defer wg.Done()
			err := spy.Sessions[index].Run(fmt.Sprintf("tail -f %s", path))
			if err != nil {

			}
		}(index, file.Path)
	}
	wg.Wait()
	return nil

}

func copyWithAppend(dst io.Writer, src io.Reader, appendStr string) {
	scanner := bufio.NewScanner(src)
	for scanner.Scan() {
		line := scanner.Text()
		lineWithAppend := appendStr + line + "\n"
		io.WriteString(dst, lineWithAppend)
	}
}

func (spy *Spy) CloseSessions() {
	logger.Info(fmt.Sprintf("[%s] Closing sessions", spy.Service.Host))
	for _, sess := range spy.Sessions {
		sess.Close()
	}
}

func (spy *Spy) CloseClient() {
	logger.Info(fmt.Sprintf("[%s] Closing client", spy.Service.Host))
	spy.Client.Close()

}
