package protocol

type CoreServerCmd string

const (
	Login                           CoreServerCmd = "Login"
	LoginRet                        CoreServerCmd = "LoginRet"
	Limit                           CoreServerCmd = "Limit"
	LimitRet                        CoreServerCmd = "LimitRet"
	Calc                            CoreServerCmd = "Calc"
	CalcRet                         CoreServerCmd = "CalcRet"
	Symmetric                       CoreServerCmd = "Symmetric"
	SymmetricRet                    CoreServerCmd = "SymmetricRet"
	Monitor                         CoreServerCmd = "Monitor"
	MonitorRet                      CoreServerCmd = "MonitorRet"
	Crypto                          CoreServerCmd = "Crypto"
	CryptoRet                       CoreServerCmd = "CryptoRet"
	Process                         CoreServerCmd = "Process"
	ProcessRet                      CoreServerCmd = "ProcessRet"
	UserPrivateKeyNotFound          CoreServerCmd = "用户私钥文件未找到"
	DecryptCipherSymmetricKey       CoreServerCmd = "解密加密对称密钥失败"
	NotFoundProductLimitInformation CoreServerCmd = "平台找不到数据产品的限制信息"
	ProductDataLimitGOSqliteFailed  CoreServerCmd = "数据产品限制信息入库失败"
	LocalCanUse                     CoreServerCmd = "LocalCanUseSyncReq"
)
const ErrCode = 400
const SuccessCode = 200
