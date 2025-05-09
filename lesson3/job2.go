package lesson3

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

/*
*
题目2：事务语句
假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和
transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
要求 ：
编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，
需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，
向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
*/
func Job2Method() {

	var err error
	db, err = gorm.Open(mysql.Open("root:20250423qwER@(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		// 配置日志级别（生产环境建议关闭）
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("数据库连接失败:%v", err)
	}

	err = db.Debug().AutoMigrate(&Account{}, &Transaction{})

	if err != nil {
		log.Fatalf("数据自动迁移失败:%v", err)
	}

	fmt.Println("数据库连接成功")

	// 创建测试账户

	/*accountList := []Account{
		{Balance: 200},
		{Balance: 100},
	}
	db.Create(&accountList)
	*/
	fromAccount := Account{ID: 2}
	db.First(&fromAccount)

	toAccount := Account{ID: 1}
	db.First(&toAccount)

	transferAccount := 100.00

	err = transferAccounts(&fromAccount, &toAccount, transferAccount, db)
	fmt.Println("\n尝试超额转账结果:", err)

}

type Account struct {
	ID      uint64  `gorm:"primaryKey"`
	Balance float64 `gorm:"not null;check:balance >= 0"`
}

type Transaction struct {
	ID            uint64    `gorm:"primaryKey"`
	FromAccountId uint64    `gorm:"not null"` //转出账户ID
	ToAccountId   uint64    `gorm:"not null"` //转入账户ID
	Amount        float64   `gorm:"not null;check:amount > 0"`
	CreatedAt     time.Time `gorm:"autoCreateTime"` // 自动记录创建时间
}

func transferAccounts(fromAccount *Account, toAccount *Account, transferAccount float64, db *gorm.DB) error {

	return db.Transaction(func(tx *gorm.DB) error {
		//1.检查账户余额是否充足
		if fromAccount.Balance < transferAccount {
			return errors.New("转出账户余额不足")
		}

		// 2. 扣除转出账户余额
		result := tx.Model(&Account{}).Where("id=? and balance > 0", fromAccount.ID).Update("balance", gorm.Expr("balance - ?", transferAccount))
		//gorm.Expr 是 GORM 提供的工具函数，用于生成 原生 SQL 表达式，将参数安全地嵌入 SQL 片段中，防止 SQL 注入
		if result.Error != nil {
			return fmt.Errorf("扣除余额失败：:%v", result.Error)
		}

		if result.RowsAffected == 0 {
			return errors.New("余额不足，扣除余额失败")
		}

		// 3. 增加转入账户余额
		result = tx.Model(&Account{}).Where("id=?", toAccount.ID).Update("balance", gorm.Expr("balance + ?", transferAccount))
		if result.Error != nil {
			return fmt.Errorf("转入余额失败：:%v", result.Error)
		}
		if result.RowsAffected == 0 {
			return errors.New("转入余额失败，用户不存在")
		}

		// 4. 创建交易记录
		transaction := Transaction{FromAccountId: fromAccount.ID, ToAccountId: toAccount.ID, Amount: transferAccount}
		if err := tx.Create(&transaction).Error; err != nil {
			return fmt.Errorf("创建交易记录失败：:%v", err)
		}

		return nil

	})

}
