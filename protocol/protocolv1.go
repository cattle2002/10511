package protocol

type DetailLimitType struct {
	CanUseNumber  int64 `json:"CanUseNumber"`
	CanUseRecords int64 `json:"CanUseRecords"`
	CanUseBytes   int64 `json:"CanUseBytes"`
	CanUseTime    int64 `json:"CanUseTime"`
}
type DetailLimitValue struct {
	CanUseNumberLocal int64 `json:"CanUseNumberLocal"`
	CanUseTimeLocal   int64 `json:"CanUseTimeLocal"`
}

type LimitResDetail struct {
	CanUseNumberLocal int64 `json:"CanUseNumberLocal"`
	CanUseTimeLocal   int64 `json:"CanUseTimeLocal"`
}
type LimitResPayload struct {
	ID     int64          `json:"ID"`
	Way    int            `json:"Way"`
	Detail LimitResDetail `json:"Detail"`
}

type LimitRes struct {
	Cmd      string          `json:"Cmd"`
	Program  string          `json:"Program"`
	RetCode  int             `json:"RetCode"`
	ErrorMsg string          `json:"ErrorMsg"`
	Payload  LimitResPayload `json:"Payload"`
}

type LimitReqPayload struct {
	ID          int64  `json:"ID"`
	Buyer       string `json:"Buyer"`
	ProductName string `json:"ProductName"`
	ProductID   int64  `json:"ProductID"`
}
type LimitReq struct {
	Cmd     string          `json:"Cmd"`
	Program string          `json:"Program"`
	Payload LimitReqPayload `json:"Payload"`
}

type SymmetricReqPayload struct {
	ID        int64  `json:"ID"`
	ProductID int64  `json:"ProductID"`
	Seller    string `json:"Seller"`
	Buyer     string `json:"Buyer"`
}
type SymmetricReq struct {
	Cmd     string              `json:"Cmd"`
	Program string              `json:"Program"`
	Payload SymmetricReqPayload `json:"Payload"`
}
type SymmetricResPayload struct {
	ID        int64  `json:"ID"`
	ProductID int64  `json:"ProductID"`
	Seller    string `json:"Seller"`
	Buyer     string `json:"Buyer"`
	AlgoType  string `json:"AlgoType"`
	Key       string `json:"Key"`
}

type SymmetricRes struct {
	Cmd      string              `json:"Cmd"`
	Program  string              `json:"Program"`
	RetCode  int64               `json:"RetCode"`
	ErrorMsg string              `json:"ErrorMsg"`
	Payload  SymmetricResPayload `json:"Payload"`
}
type HttpCalcResponsePayload struct {
	ID       int64  `json:"ID"`
	HaveData bool   `json:"HaveData"`
	Data     []byte `json:"Data"`
	Url      string `json:"Url"`
}
type HttpCalcResponse struct {
	Cmd      string                  `json:"Cmd"`
	RetCode  int                     `json:"RetCode"`
	ErrorMsg string                  `json:"ErrorMsg"`
	Payload  HttpCalcResponsePayload `json:"Payload"`
}
type HttpCalcRequestPayload struct {
	ID                                    int64  `json:"ID"`
	ProductID                             int64  `json:"ProductID"`
	ProductName                           string `json:"ProductName"`
	Buyer                                 string `json:"Buyer"`
	Seller                                string `json:"Seller"`
	CertificateAlgorithmTypeAlgorithmType string `json:"CertificateTypeAlgorithmType"`
	SymmetricKeyAlgorithmType             string `json:"SymmetricKeyAlgorithmType"`
	HaveData                              bool   `json:"HaveData"`
	ProductData                           string `json:"ProductData"`
	ProductUrl                            string `json:"ProductPosition"`
	CipherSymmetricKey                    string `json:"CipherSymmetricKey"`
}
type HttpCalcRequest struct {
	Cmd     string                 `json:"Cmd"`
	Payload HttpCalcRequestPayload `json:"Payload"`
}

type HttpListAlgoRequest struct {
	PageNum  int `json:"PageNum"`
	PageSize int `json:"PageSize"`
}

type HttpListAlgoResponse struct {
	Code int              `json:"Code"`
	Msg  string           `json:"Msg"`
	Data *[]HttpAlgoModel `json:"Data"`
}
type HttpAlgoModel struct {
	AlgoID       int64  `json:"AlgoID"`
	AlgoType     string `json:"AlgoType"`
	AlgoPosition string `json:"AlgoPosition"`
	RunningEnv   string `json:"RunningEnv"`
	FileType     string `json:"FileType"`
}

type LoginReqPayload struct {
	ID        int64  `json:"ID"`
	TimeStamp int64  `json:"TimeStamp"`
	User      string `json:"User"`
	Password  string `json:"Password"`
}
type LoginReq struct {
	Cmd     string          `json:"Cmd"`
	Program string          `json:"Program"`
	Payload LoginReqPayload `json:"Payload"`
}

type LoginResPayload struct {
	ID        int64 `json:"ID"`
	TimeStamp int64 `json:"TimeStamp"`
}
type LoginRes struct {
	Cmd      string          `json:"Cmd"`
	Program  string          `json:"Program"`
	RetCode  int             `json:"RetCode"`
	ErrorMsg string          `json:"ErrorMsg"`
	Payload  LoginResPayload `json:"Payload"`
}
type HttpCryptoReqPayload struct {
	ID           int64  `json:"ID"`
	User         string `json:"User"`
	SymmetricKey string `json:"SymmetricKey"`
}
type HttpCryptoReq struct {
	Cmd     string               `json:"Cmd"`
	Payload HttpCryptoReqPayload `json:"Payload"`
}
type HttpCryptoResPayload struct {
	CipherSymmetricKey string `json:"CipherSymmetricKey"`
	AlgoType           string `json:"AlgoType"`
}
type HttpCryptoRes struct {
	Cmd      string               `json:"Cmd"`
	RetCode  int                  `json:"RetCode"`
	ErrorMsg string               `json:"ErrorMsg"`
	Payload  HttpCryptoResPayload `json:"Payload"`
}
type HttpDecryptReqPayload struct {
	ID                 int64  `json:"ID"`
	User               string `json:"User"`
	CipherSymmetricKey string `json:"CipherSymmetricKey"`
}
type HttpDecryptReq struct {
	Cmd     string                `json:"Cmd"`
	Payload HttpDecryptReqPayload `json:"Payload"`
}
type HttpDecryptResPayload struct {
	SymmetricKey string `json:"SymmetricKey"`
}
type HttpDecryptRes struct {
	Cmd      string                `json:"Cmd"`
	RetCode  string                `json:"RetCode"`
	ErrorMsg string                `json:"ErrorMsg"`
	Payload  HttpDecryptResPayload `json:"Payload"`
}

type ProcessAlgoInputs struct {
	Key        string `json:"Key"`
	Type       int    `json:"Type"`
	Value      string `json:"Value"`
	Must       int    `json:"Must"`
	AllowError int    `json:"AllowError"`
	Status     int    `json:"Status"`
	HasValue   bool   `json:"HasValue"`
}
type ProcessAlgo struct {
	Method            string              `json:"Method"`
	ProductID         int64               `json:"ProductID"`
	Data              string              `json:"Data"`
	VisitType         string              `json:"VisitType"`
	Number            int64               `json:"Number"`
	Pwd               string              `json:"Pwd"`
	SymmetricPwdType  string              `json:"SymmetricPwdType"`
	BuyerPublicKey    string              `json:"BuyerPublicKey"`
	AsymmetricPwdType string              `json:"AsymmetricPwdType"`
	Inputs            []ProcessAlgoInputs `json:"Inputs"`
}
type ProcessToMonitorReqPayload struct {
	ID     int64         `json:"ID"`
	Seller string        `json:"Seller"`
	Buyer  string        `json:"Buyer"`
	Algo   []ProcessAlgo `json:"Algo"`
}
type ProcessToMonitorReq struct {
	Cmd     string                     `json:"Cmd"`
	Payload ProcessToMonitorReqPayload `json:"Payload"`
}
type LocalCanUseSyncReqPayload struct {
	ID                int64  `json:"ID"`
	ProductID         int64  `json:"ProductID"`
	UserName          string `json:"UserName"`
	CanUseNumberLocal int64  `json:"CanUseNumberLocal"`
}
type LocalCanUseSyncReq struct {
	Cmd     string                    `json:"Cmd"`
	Program string                    `json:"Program"`
	Payload LocalCanUseSyncReqPayload `json:"Payload"`
}
type LocalCanUseSyncResPayload struct {
	ID int64 `json:"ID"`
}
type LocalCanUseSyncRes struct {
	Cmd      string                    `json:"Cmd"`
	Program  string                    `json:"Program"`
	RetCode  int                       `json:"RetCode"`
	ErrorMsg string                    `json:"ErrorMsg"`
	Payload  LocalCanUseSyncResPayload `json:"Payload"`
}
type HttpStatusReq struct {
	Cmd     string `json:"Cmd"`
	Program string `json:"Program"`
}
type HttpStatusRes struct {
	Code int    `json:"Code"`
	Msg  string `json:"Msg"`
	Data string `json:"Data"`
}
