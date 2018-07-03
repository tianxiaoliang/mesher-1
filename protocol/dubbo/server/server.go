package server

import (
	"net"
	"sync"
	"time"

	"github.com/ServiceComb/go-chassis/core/lager"
	"github.com/ServiceComb/go-chassis/core/server"
	"github.com/go-chassis/mesher/protocol/dubbo/proxy"
	"github.com/go-chassis/mesher/protocol/dubbo/utils"
)

const (
	NAME = "dubbo"
)

//ConnectionMgr -------连接管理
type ConnectionMgr struct {
	conns map[int]*DubboConnection
	count int
}

//NewConnectMgr is a function which new connection manager and returns it
func NewConnectMgr() *ConnectionMgr {
	tmp := new(ConnectionMgr)
	tmp.count = 0
	tmp.conns = make(map[int]*DubboConnection)
	return tmp
}

//GetConnection is a method to get connection
func (this *ConnectionMgr) GetConnection(conn *net.TCPConn) *DubboConnection {
	dubbConn := NewDubboConnetction(conn, nil)
	key := this.count
	this.conns[key] = dubbConn
	this.count++
	return dubbConn
}

//DeactiveAllConn is a function to close all connection
func (this *ConnectionMgr) DeactiveAllConn() {
	for _, v := range this.conns {
		v.Close()
	}
}

func init() {
	server.InstallPlugin(NAME, newServer)
}

func newServer(opts server.Options) server.ProtocolServer {

	return &DubboServer{
		opts:       opts,
		routineMgr: util.NewRoutineManager(),
	}
}

//-----------------------dubbo server---------------------------

//DubboServer is a struct
type DubboServer struct {
	connMgr    *ConnectionMgr
	opts       server.Options
	mux        sync.RWMutex
	exit       chan chan error
	routineMgr *util.RoutineManager
}

func (d *DubboServer) String() string {
	return NAME
}

//Init is a method to initialize the server
func (d *DubboServer) Init() error {
	d.connMgr = NewConnectMgr()
	lager.Logger.Info("Dubbo server init success.")
	return nil
}

//Register is a method to register the schema to the server
func (d *DubboServer) Register(schema interface{}, options ...server.RegisterOption) (string, error) {
	return "", nil
}

//Stop is a method to disconnect all connection
func (d *DubboServer) Stop() error {
	d.connMgr.DeactiveAllConn()
	d.routineMgr.Done()
	return nil
}

//Start is a method to start server
func (d *DubboServer) Start() error {
	d.Init()
	dubboproxy.DubboListenAddr = d.opts.Address
	host, _, err := net.SplitHostPort(d.opts.Address)
	if err != nil {
		return err
	}
	ip := net.ParseIP(host)
	if ip == nil {
		return &util.BaseError{"Invalid host"}
	}
	tcpAddr, err := net.ResolveTCPAddr("tcp", d.opts.Address)
	if err != nil {
		lager.Logger.Error("ResolveTCPAddr err: ", err)
		return err
	}
	l, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		lager.Logger.Error("listening falied, reason:", err)
		return err
	}
	d.routineMgr.Spawn(d, l, "Acceptloop")
	return nil
}

//Svc is a method
func (d *DubboServer) Svc(arg interface{}) interface{} {
	d.AcceptLoop(arg.(*net.TCPListener))
	return nil
}

//AcceptLoop is a method
func (d *DubboServer) AcceptLoop(l *net.TCPListener) {
	for {
		conn, err := l.AcceptTCP()
		if err != nil {
			select {
			case <-time.After(time.Second * 3):
				lager.Logger.Info("Sleep three second")
			}
		}
		dubbConn := d.connMgr.GetConnection(conn)
		dubbConn.Open()
	}

	defer l.Close()
}
