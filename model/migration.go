/**
 * @Author: woshishabii
 * @Description:
 * @File: migration
 * @Version: 0.0.1
 * @Date: 4/15/2023 3:48 PM
 */

package model

// 执行数据迁移

func migration() {
	// 自动迁移模式
	_ = DB.AutoMigrate(&User{}, &School{}, &Session{}, &Exam{})
}
