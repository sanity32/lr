package lr

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func NewMgr(port int) *Mgr {
	return &Mgr{
		Port:               port,
		ExecutableFilename: DEFAULT_LRSRV_FILENAME,
		archive:            LRSRV_STABLE_ZIPPED,
	}
}

type Mgr struct {
	Port               int
	ExecutableFilename string
	archive            EmbedZipFile
	cmnd               *exec.Cmd
	client             *Client
}

func (m *Mgr) Finalize() {

	if c := m.cmnd; c != nil {
		if p := c.Process; p != nil {
			p.Kill()
		}
	}
}

func (m *Mgr) Client() *Client {
	if c := m.client; c != nil {
		return c
	}
	panic(ErrClientNotInit)
}

func (m Mgr) addr() string {
	return fmt.Sprintf("localhost:%v", m.Port)
}

func (m Mgr) filepath() string {
	return m.ExecutableFilename
}

func (m Mgr) executableFileExist() bool {
	_, err := os.Stat(m.filepath())
	return err == nil || os.IsExist(err)
}

func (m Mgr) deploy() error {
	return m.archive.ExtractFirst(m.filepath())
}

func (m Mgr) prepareExecutableFile() error {
	if !m.executableFileExist() {
		return m.deploy()
	}
	return nil
}

func (m *Mgr) runExecutableFile() error {
	arg := fmt.Sprintf("-port=%v", m.Port)
	m.cmnd = exec.Command("./"+m.filepath(), arg)
	return m.cmnd.Start()
}

func (m *Mgr) Init() error {
	if err := m.prepareExecutableFile(); err != nil {
		return err
	}
	m.client = NewClient(m.addr())
	if err := m.client.Connect(); err != nil {
		if err := m.runExecutableFile(); err != nil {
			return err
		}
	}
	return m.client.ConnectN(3, time.Millisecond*500)
}
