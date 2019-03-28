package models

import "star-edge-cloud/edge/models"

// RulesCollection - 当前在使用的规则
// todo:加锁
var RulesCollection *models.Rules

// Config - 运行时配置
type Config struct {
	Port string
}
