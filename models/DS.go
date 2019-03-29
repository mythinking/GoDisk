/***********************

	数据结构

************************/

package models

//分类表
type Category struct {
	Id          int
	Name        string `orm:"size(64)"`
	Key         string `orm:"size(64)"`
	Description string
}

//文件存储表
type Attachment struct {
	Id      int
	Cid     int
	Name    string
	Mark    string `orm:"size(64)"`
	Path    string
	Type    string
	Created string `orm:"size(10)"`
}

// 全局配置表  Option=配置选项  Value=配置内容
type Config struct {
	Id       int
	Option   string `orm:"size(16)"`
	Value    string `orm:"size(32)"`
	Addition string `orm:"size(16)"`
}

//七牛云配置
type QiniuConfigOption struct {
	QnAk     string //Accesskey
	QnSk     string //Secretkey
	QnBucket string //Bucket
	QnZone   string //Zone
}

//又拍云配置
type UpyunConfigOption struct {
	UpBucket   string //Bucket
	UpOperator string //Operator
	UpPassword string //Password
	UpDomain   string //Domain
}

// 阿里云对象存储配置
type OssConfigOption struct {
	OssBucket   string //Bucket
	OssAk       string //Accesskey
	OssSk       string //Secretkey
	OssEndpoint string //地域节点
}

//腾讯云对象存储配置
type CosConfigOption struct {
	CosBucket string //Bucket
	CosAppid  string //APPID
	CosRegion string //Region
	CosSkid   string //SecretID
	CosSk     string //SecretKey
}

// 网站后台提交的表单字段 映射到此结构体 需要持续添加
type SiteConfigOption struct {
	WebTitle    string `form:"WebTitle"`
	Keywords    string `form:"Keywords"`
	Description string `form:"Description"`
	CopyRight   string `form:"CopyRight"`
	LogoUrl     string `form:"LogoUrl"`
}

// 用户配置信息
type UserConfigOption struct {
	Author   string `form:"Author"`
	Password string `form:"Password"`
}
